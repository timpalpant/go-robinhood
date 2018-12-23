package robinhood

import (
	"testing"
)

func TestInstruments(t *testing.T) {
	client := NewClient("", "")
	instruments, err := client.GetInstruments()
	if err != nil {
		t.Errorf(err.Error())
	}

	t.Logf("got %d instruments", len(instruments))
}
