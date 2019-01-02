package robinhood

import (
	"reflect"
	"testing"
	"time"
)

func TestListPositions(t *testing.T) {
	// One page of 5 positions, then one page of 1 position.
	responses, err := loadResponses(
		"testdata/responses/positions.0.json",
		"testdata/responses/positions.1.json",
	)
	if err != nil {
		t.Fatal(err)
	}

	client := NewTestClient(responses)
	positions, err := client.ListPositions(false)
	if err != nil {
		t.Errorf(err.Error())
	}

	if len(positions) != 6 {
		t.Errorf("got %d positions, expected %d", len(positions), 6)
	}

	expected := &Position{
		AccountURL:              "https://api.robinhood.com/accounts/8UD09348/",
		IntradayAverageBuyPrice: 2.34,
		URL:                     "https://api.robinhood.com/accounts/8UD09348/positions/a44552fb-9f59-4168-86f1-c93998fa019d/",
		CreatedAt:               time.Date(2018, 1, 11, 17, 48, 47, 128378000, time.UTC),
		UpdatedAt:               time.Date(2018, 1, 11, 18, 11, 42, 883624000, time.UTC),
		AverageBuyPrice:         2.34,
		InstrumentURL:           "https://api.robinhood.com/instruments/a44552fb-9f59-4168-86f1-c93998fa019d/",
		Quantity:                3.0,
	}

	if !reflect.DeepEqual(positions[0], expected) {
		t.Errorf("did not parse expected position: got %+v, wanted %+v", positions[0], expected)
	}
}
