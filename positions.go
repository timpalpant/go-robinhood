package robinhood

import (
	"time"
)

type Position struct {
	AccountURL               string  `json:"account"`
	SharesHeldForStockGrants float64 `json:"shares_held_for_stock_grants,string"`
	IntradayQuantity         float64 `json:"intraday_quantity,string"`
	IntradayAverageBuyPrice  float64 `json:"intraday_average_buy_price,string"`
	URL                      string
	CreatedAt                time.Time `json:"created_at"`
	UpdatedAt                time.Time `json:"updated_at"`
	SharesHeldForBuys        float64   `json:"shares_held_for_buys,string"`
	AverageBuyPrice          float64   `json:"average_buy_price,string"`
	InstrumentURL            string    `json:"instrument"`
	SharesHeldForSells       float64   `json:"shares_held_for_sells,string"`
	Quantity                 float64   `json:",string"`
}

type getPositionsRequest struct {
	Nonzero bool `url:"nonzero,omitempty"`
}

func (c *Client) ListPositions(nonzero bool) ([]*Position, error) {
	url := Endpoint + "/positions/"
	return c.doGetPositions(url, nonzero)
}

func (c *Client) ListPositionsForAccount(accountNumber string, nonzero bool) ([]*Position, error) {
	url := Endpoint + "/accounts/" + accountNumber + "/positions/"
	return c.doGetPositions(url, nonzero)
}

func (c *Client) doGetPositions(url string, nonzero bool) ([]*Position, error) {
	req := &getPositionsRequest{
		Nonzero: nonzero,
	}

	var result []*Position
	for url != "" {
		var resp struct {
			Results []*Position
			Next    string
		}

		err := c.getJSON(url, req, &resp)
		if err != nil {
			return nil, err
		}

		result = append(result, resp.Results...)
		url = resp.Next
	}

	return result, nil
}
