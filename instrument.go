package robinhood

type Instrument struct {
	MarginInitialRatio float64 `json:"margin_initial_ratio,string"`
	RHSTradability     string  `json:"rhs_tradability"`
	ID                 string
	MarketURL          string   `json:"market"`
	SimpleName         string   `json:"simple_name"`
	MinTickSize        *float64 `json:"min_tick_size,string"`
	MaintenanceRatio   float64  `json:"maintenance_ratio,string"`
	Tradability        string
	State              string
	Type               string
	Tradeable          bool
	FundamentalsURL    string `json:"fundamentals"`
	QuoteURL           string `json:"quote"`
	Symbol             string
	DayTradeRatio      float64 `json:"day_trade_ratio,string"`
	Name               string
	TradableChainID    *string `json:"tradable_chain_id"`
	SplitsURL          string  `json:"splits"`
	URL                string
	BloombergUnique    string `json:"bloomberg_unique"`
	ListDate           string `json:"list_date"`
	Country            string
}

// List all known instruments.
func (c *Client) ListAllInstruments() ([]*Instrument, error) {
	url := Endpoint + "/instruments/"
	var result []*Instrument
	for url != "" {
		var resp struct {
			Results []*Instrument
			Next    string
		}

		if err := c.getJSON(url, &resp); err != nil {
			return nil, err
		}

		result = append(result, resp.Results...)
		url = resp.Next
		if len(resp.Results) == 0 {
			break
		}
	}

	return result, nil
}
