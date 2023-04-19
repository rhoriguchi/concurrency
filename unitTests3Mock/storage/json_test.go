package storage

import "testing"

func TestLoadStream(t *testing.T) {
	// arrange

	// act
	out, err := loadStream()

	// assert
	if err != nil {
		t.Errorf("loading data failed: %v", err)
	}

	// no data expected, because data file does not exist in package storage directory
	if len(out) > 0 {
		t.Error("channel should be closed")
	}
}
