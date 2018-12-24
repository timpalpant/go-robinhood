package robinhood

import (
	"testing"

	"gopkg.in/d4l3k/messagediff.v1"
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

	expected := &Instrument{
		MarginInitialRatio: 1.0,
		RHSTradability:     "untradable",
		ID:                 "5c0f7dce-716b-4896-a52d-4b7431b4abcb",
		MarketURL:          "https://api.robinhood.com/markets/XASE/",
		SimpleName:         "EVI Industries",
		MinTickSize:        nil,
		MaintenanceRatio:   1.0,
		Tradability:        "untradable",
		State:              "inactive",
		Type:               "stock",
		Tradeable:          false,
		FundamentalsURL:    "https://api.robinhood.com/fundamentals/EVI/",
		QuoteURL:           "https://api.robinhood.com/quotes/EVI/",
		Symbol:             "EVI",
		DayTradeRatio:      1.0,
		Name:               "EVI Industries, Inc.",
		TradableChainID:    nil,
		SplitsURL:          "https://api.robinhood.com/instruments/5c0f7dce-716b-4896-a52d-4b7431b4abcb/splits/",
		URL:                "https://api.robinhood.com/instruments/5c0f7dce-716b-4896-a52d-4b7431b4abcb/",
		Country:            "US",
		BloombergUnique:    "EQ0010554700001000",
		ListDate:           "2009-12-01",
	}

	diff, equal := messagediff.PrettyDiff(instruments[0], expected)
	if !equal {
		t.Errorf(diff)
	}
}
