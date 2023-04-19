package monitor // Package monitor used by dns-inspection

import (
	"context"
	"fmt"
	"io"
	"net"
	"sync"
	"syscall"
)

type monitor struct {
	sink       map[int][]io.Writer   // Receiving ends for each filter
	filterChan map[int](chan string) // Buffered Channel for each filter
	filterList []int                 // List of filterChan
	ctx        context.Context
	sync.Mutex
}

var distributor monitor
var stop context.CancelFunc
var errmsg chan string

// Messages to filter 'Ignore' will not be written.
const Ignore = -1

func resetGlobals() {
	distributor = monitor{}
	distributor.sink = make(map[int][]io.Writer)
	distributor.filterChan = make(map[int](chan string))
	distributor.ctx, stop = context.WithCancel(context.Background())
	// Create channel for error handling
	// This is a buffered channel in case, the logging can't handle that much and fast stuff (should not happen...)
	errmsg = make(chan string, 512)
}

func init() {
	resetGlobals()
}

func handleConn(conn *net.TCPConn, filter int) {
	defer func() {
		distributor.remove(filter, conn)
		conn.Close()
	}()
	connCtrl := make(chan bool)
	distributor.register(filter, conn)
	conn.SetLinger(0) // If connection is closed, discard unset data
	buf := make([]byte, 1)
	go func() {
		for {
			n, err := conn.Read(buf)
			if err != nil || n == 0 {
				connCtrl <- true

				return
			}
		}
	}()

	select {
	case <-connCtrl:
		break
	case <-distributor.ctx.Done():
		break
	}
}

// RegisterFilter adds new filter(s).
func RegisterFilter(filter ...int) {
	var skip bool
	for _, f := range filter {
		skip = false
		if f < 1 {
			errmsg <- fmt.Sprintf("Filter '%d' has to be a positive integer greater than 0", f)

			continue
		}
		for _, x := range distributor.filterList {
			if f == x {
				errmsg <- fmt.Sprintf("Filter '%d' already exists", f)
				skip = true

				break
			}
			if f&x != 0 {
				errmsg <- fmt.Sprintf("Filter '%d' and '%d' are not independent (on binary level)", f, x)
				skip = true

				break
			}
		}
		if !skip {
			distributor.newFilter(f)
		}
	}
}

func (m *monitor) newFilter(filter int) {
	filterChan := make(chan string, 512)
	go func() {
		for {
			select {
			case msgString := <-filterChan:
				distributor.writestring(filter, msgString)
			case <-m.ctx.Done():
				safeChanClose(filterChan)

				break
			}
		}
	}()

	m.Lock()
	defer m.Unlock()
	m.filterChan[filter] = filterChan
	m.filterList = append(m.filterList, filter)
}

// closes a channel if it hasn't been closed yet
// needed to avoid panic when trying to close an already closed channel.
func safeChanClose(ch chan string) {
	select {
	case <-ch:
		// channel is closed
	default:
		// channel is not closed
		close(ch)
	}
}

// SetLogging provides a interface, to set a central logging.
func SetLogging(fn func(string)) {
	go func() {
		for {
			select {
			case msg := <-errmsg:
				fn(msg)
			case <-distributor.ctx.Done():
				break
			}
		}
	}()
}

// Stop stops all monitoring.
func Stop() {
	stop()
}

// Serve a TCP socket, so clients can connect and get logged dns messages.
func Serve(socketAddress string, filter int) error {
	addr, err := net.ResolveTCPAddr("tcp", socketAddress)
	if err != nil {
		return err
	}
	servant, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return err
	}
	go func() {
		<-distributor.ctx.Done()
		servant.Close()
	}()
	go func() {
		for {
			conn, err := servant.AcceptTCP()
			if err != nil {
				if err == syscall.EINVAL {
					return
				}
				errmsg <- fmt.Sprintf("Error while waiting for TCP connections: %v", err)

				continue
			}
			go handleConn(conn, filter)
		}
	}()

	return nil
}

func (m *monitor) remove(filter int, writer io.Writer) {
	m.Lock()
	defer m.Unlock()
	for i := len(m.sink[filter]) - 1; i >= 0; i-- {
		if m.sink[filter][i] == writer {
			m.sink[filter] = append(m.sink[filter][:i], m.sink[filter][i+1:]...)

			break
		}
	}
}

func RegisterWriter(filter int, writer io.Writer) {
	if distributor.filterChan[filter] == nil {
		errmsg <- fmt.Sprintf("'%d' is not a valid filter", filter)

		return
	}
	distributor.register(filter, writer)
}

func (m *monitor) register(filter int, writer io.Writer) {
	m.Lock()
	defer m.Unlock()
	m.sink[filter] = append(m.sink[filter], writer)
}

// WriteString writes string s to registered writers.
func WriteString(filter int, s string) {
	if filter&Ignore == Ignore {
		return
	}
	for _, x := range distributor.filterList {
		if filter&x == x {
			distributor.filterChan[x] <- s
		}
	}
}

func (m *monitor) writestring(filter int, s string) {
	s += "\n"
	for _, w := range m.sink[filter] {
		n, err := io.WriteString(w, s)
		if err != nil {
			errmsg <- fmt.Sprintf("Could not write string: %v", err)
		}
		if n != len(s) {
			errmsg <- fmt.Sprintf("Could not write complete string: %v", s)
		}
	}
}
