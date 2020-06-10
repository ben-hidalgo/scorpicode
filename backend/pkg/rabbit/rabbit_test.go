package rabbit_test

import (
	"backend/pkg/rabbit"
	"testing"
)

const (
	NOT_NIL   = "not nil"
	NOT_EMPTY = "not empty"
	GOT       = "got '%v' %s '%v'"
	WANTED    = "but wanted"
)

func TestHatDao(t *testing.T) {

	if testing.Short() {
		t.Skip()
	}

	// connect rabbit
	rabbitConn, err := rabbit.Connect()
	if err != nil {
		t.Fatalf(GOT, err, WANTED, nil)
	}
	defer rabbitConn.Close()

}
