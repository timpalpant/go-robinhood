package robinhood

import (
	"testing"
)

func TestListAllInstruments(t *testing.T) {
	// Three pages of 100 instruments each.
	responses, err := loadResponses(
		"testdata/responses/instruments.0.json",
		"testdata/responses/instruments.1.json",
		"testdata/responses/instruments.2.json",
	)
	if err != nil {
		t.Fatal(err)
	}

	client := NewTestClient(responses)
	instruments, err := client.ListAllInstruments()
	if err != nil {
		t.Errorf(err.Error())
	}

	if len(instruments) != 300 {
		t.Errorf("got %d instruments, expected %d", len(instruments), 300)
	}
}
