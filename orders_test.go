package robinhood

import (
	"reflect"
	"testing"
	"time"
)

func TestParseOrderResponse(t *testing.T) {
	responses, err := loadResponses("testdata/responses/place_order.json")
	if err != nil {
		t.Fatal(err)
	}

	order := &Order{}

	client := NewTestClient(responses)
	ticket, err := client.PlaceOrder(order)
	if err != nil {
		t.Fatal(err)
	}

	expected := &OrderTicket{
		Executions:         []string{},
		UpdatedAt:          time.Date(2016, 4, 1, 21, 24, 13, 698563000, time.UTC),
		Fees:               0.0,
		CancelURL:          "https://api.robinhood.com/orders/15390ade-face-caca-0987-9fdac5824701/cancel/",
		ID:                 "15390ade-face-caca-0987-9fdac5824701",
		CumulativeQuantity: 0.0,
		State:              Queued,
		LastTransactionAt:  time.Date(2016, 4, 1, 23, 34, 54, 237390000, time.UTC),
		URL:                "https://api.robinhood.com/orders/15390ade-face-caca-0987-9fdac5824701/",
		CreatedAt:          time.Date(2016, 4, 1, 22, 12, 14, 890283000, time.UTC),
		PositionURL:        "https://api.robinhood.com/positions/8UD09348/50810c35-d215-4866-9758-0ada4ac79ffa/",
		Order: &Order{
			InstrumentURL: "https://api.robinhood.com/instruments/50810c35-d215-4866-9758-0ada4ac79ffa/",
			Side:          Sell,
			TimeInForce:   GoodTilCanceled,
			Trigger:       "immediate",
			Type:          Market,
			AccountURL:    "https://api.robinhood.com/accounts/8UD09348/",
			Quantity:      1.0,
		},
	}

	if !reflect.DeepEqual(ticket, expected) {
		t.Errorf("did not parse expected order ticket: got %+v, wanted %+v", ticket, expected)
	}
}
