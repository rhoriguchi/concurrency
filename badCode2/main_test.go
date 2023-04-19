package main

import (
	"testing"
)

// Please ignore this file for the moment. It will be covered during the interactive part of our session.

func TestProcess(t *testing.T) {
	/*
		// arrange
		m := &mockMonitor{}

		// act
		go process()
		time.Sleep(5 * time.Second)

		// assert
		if m.callsToWriteString != 2 {
			t.Errorf("Function WriteString() has not been called twice, but %d times", m.callsToWriteString)
		}
	*/
}

type mockMonitor struct {
	callsToWriteString int
}

func (mock *mockMonitor) WriteString(filter int, s string) {
	mock.callsToWriteString++
}

func (mock *mockMonitor) RegisterFilter(filter ...int) {

}

func (mock *mockMonitor) Stop() {

}
