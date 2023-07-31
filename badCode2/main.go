package main

import (
	"badCode2/monitor"
	"fmt"
	"time"
)

const port = "2023"
const (
	filterRegId = 0x02
	writerRegId = 0x02
)

func main() {
	// accept incoming connections in the background
	err := monitor.Serve("localhost:"+port, writerRegId)
	if err != nil {
		panic(err)
	}

	fmt.Println("Catch me on TCP port 2023...")

	process()
}

func process() {
	data := []struct {
		delay int
		msg   string
	}{
		{2, "early bird!"},
		{4, "welcome, the party has started"},
		{9, "bye bye, time to go"},
	}

	monitor.RegisterFilter(filterRegId)

	fmt.Println("Catch me on TCP port 2023...")

	// send timed messages
	for _, d := range data {
		go func(delay int, msg string) {
			time.Sleep(time.Duration(delay) * time.Second)
			monitor.WriteString(writerRegId, msg)
		}(d.delay, d.msg)
	}

	// auto-shutdown after 10 seconds
	<-time.After(10 * time.Second)
	fmt.Println("Time to leave. Good bye!")
	monitor.Stop()

	// allow the shutdown code to complete in the hope to discover more race conditions or other issues
	time.Sleep(10 * time.Millisecond)
}
