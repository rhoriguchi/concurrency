package monitor

import (
	"bytes"
	"io/ioutil"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func readFromChanOrEmpty(c chan string) string {
	select {
	case s := <-c:
		return s
	default:
		return ""
	}
}

func TestRegisterFilter(t *testing.T) {
	tests := []struct {
		name    string
		filters []int
		err     string
	}{
		{"Correct", []int{123}, ""},
		{"Correct multiple", []int{0xa0, 0x0a}, ""},
		{"filter=2", []int{2}, ""},
		{"filter=1", []int{1}, ""},
		{"Negative filter", []int{-100}, "Filter '-100' has to be a positive integer greater than 0"},
		{"0 filter", []int{0}, "Filter '0' has to be a positive integer greater than 0"},
		{"Binary level dependence", []int{101, 201}, "Filter '201' and '101' are not independent (on binary level)"},
		{"Duplicates", []int{123, 123}, "Filter '123' already exists"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resetGlobals()
			RegisterFilter(tt.filters...)
			actualError := readFromChanOrEmpty(errmsg)
			assert.Equal(t, tt.err, actualError)
		})
	}
}

func TestRegisterWriter(t *testing.T) {
	var tests = []struct {
		name string
		f    func(t *testing.T)
	}{{"Wrong registered filter", func(t *testing.T) {
		RegisterFilter(0x02)
		RegisterWriter(0x20, ioutil.Discard)
		actualError := readFromChanOrEmpty(errmsg)
		assert.Equal(t, "'32' is not a valid filter", actualError)
	}}, {"No registered filters", func(t *testing.T) {
		RegisterWriter(0x02, ioutil.Discard)
		actualError := readFromChanOrEmpty(errmsg)
		assert.Equal(t, "'2' is not a valid filter", actualError)
	}}, {"Register on correct filter", func(t *testing.T) {
		RegisterFilter(0x02)
		RegisterWriter(0x02, ioutil.Discard)
		actualError := readFromChanOrEmpty(errmsg)
		assert.Equal(t, "", actualError)
	}}, {"Register on filter with binary overlap", func(t *testing.T) {
		RegisterFilter(0x11)
		RegisterWriter(0x10, ioutil.Discard)
		actualError := readFromChanOrEmpty(errmsg)
		assert.Equal(t, "'16' is not a valid filter", actualError)
	}},
	}
	for _, tt := range tests {
		resetGlobals()
		t.Run(tt.name, tt.f)
	}
}

func TestWriteString(t *testing.T) {
	type args struct {
		filter int
		s      string
	}
	type output struct {
		writer1 string
		writer2 string
	}
	tests := []struct {
		name   string
		args   args
		output output
	}{
		{"Write to one writer", args{1 << 1, "aaaa"}, output{"aaaa\n", ""}},
		{"Write to multiple writers", args{1<<1 | 1<<2, "aaaa"}, output{"aaaa\n", "aaaa\n"}},
		{"Write to filter with no registered writers", args{1 << 3, "aaaa"}, output{"", ""}},
		{"Write to not registered filter", args{1 << 4, "aaaa"}, output{"", ""}},
	}
	for _, tt := range tests {
		resetGlobals()
		t.Run(tt.name, func(t *testing.T) {
			RegisterFilter(1<<1, 1<<2, 1<<3)
			writer1 := &bytes.Buffer{}
			writer2 := &bytes.Buffer{}
			RegisterWriter(1<<1, writer1)
			RegisterWriter(1<<2, writer2)
			WriteString(tt.args.filter, tt.args.s)

			// TODO: WAN-2333 try removing the sleep, this might also cause flaky tests
			time.Sleep(time.Second)

			assert.Equal(t, "", readFromChanOrEmpty(errmsg))

			assert.Equal(t, tt.output.writer1, writer1.String())
			assert.Equal(t, tt.output.writer2, writer2.String())
		})
	}
}
