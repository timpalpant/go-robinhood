package robinhood

import (
	"time"
)

type OrderState string

const (
	Queued          OrderState = "queued"
	Unconfirmed                = "unconfirmed"
	Confirmed                  = "confirmed"
	PartiallyFilled            = "partially_filled"
	Filled                     = "filled"
	Rejected                   = "rejected"
	Canceled                   = "cancelled"
	Failed                     = "failed"
)

type OrderType string

const (
	Market OrderType = "market"
	Limit  OrderType = "limit"
)

type TimeInForce string

const (
	GoodForDay        TimeInForce = "gfd"
	GoodTilCanceled   TimeInForce = "gtc"
	ImmediateOrCancel TimeInForce = "ioc"
	OPG               TimeInForce = "opg"
)

type Trigger string

const (
	Immediate Trigger = "immediate"
	Stop      Trigger = "stop"
)

type OrderSide string

const (
	Buy  OrderSide = "buy"
	Sell OrderSide = "sell"
)

type Order struct {
	AccountURL             string      `json:"account" url:"account,omitempty"`
	InstrumentURL          string      `json:"instrument" url:"instrument,omitempty"`
	Symbol                 string      `url:"symbol,omitempty"`
	Type                   OrderType   `url:"type,omitempty"`
	TimeInForce            TimeInForce `json:"time_in_force" url:"time_in_force,omitempty"`
	Trigger                Trigger     `url:"trigger,omitempty"`
	Price                  *float64    `json:",string" url:"price,omitempty"`
	StopPrice              *float64    `json:"stop_price,string" url:"stop_price,omitempty"`
	Quantity               float64     `json:"quantity,string" url:"quantity,omitempty"`
	Side                   OrderSide   `url:"side,omitempty"`
	ClientID               *string     `json:"client_id" url:"client_id,omitempty"`
	ExtendedHours          bool        `json:"extended_hours" url:"extended_hours,omitempty"`
	OverrideDayTradeChecks bool        `json:"override_day_trade_checks" url:"override_day_trade_checks,omitempty"`
	OverrideDTBPChecks     bool        `json:"override_dtbp_checks" url:"override_dtbp_checks,omitempty"`
}

type OrderExecution struct {
}

type OrderTicket struct {
	*Order
	ID                 string
	UpdatedAt          time.Time `json:"updated_at"`
	Executions         []OrderExecution
	Fees               float64 `json:"fees,string"`
	CancelURL          string  `json:"cancel"`
	CumulativeQuantity float64 `json:"cumulative_quantity,string"`
	RejectReason       string  `json:"reject_reason"`
	State              OrderState
	LastTransactionAt  time.Time `json:"last_transaction_at"`
	URL                string
	CreatedAt          time.Time `json:"created_at"`
	PositionURL        string    `json:"position"`
	AveragePrice       *float64  `json:"average_price,string"`
}

type ListOrdersRequest struct {
	Since         time.Time `url:"since,omitempty"`
	InstrumentURL string    `url:"instrument,omitempty"`
}

func (c *Client) PlaceOrder(o *Order) (*OrderTicket, error) {
	url := Endpoint + "/orders/"
	resp := &OrderTicket{Order: o}
	err := c.postForm(url, o, resp)
	return resp, err
}

func (c *Client) GetOrder(orderId string) (*OrderTicket, error) {
	url := Endpoint + "/orders/" + orderId
	resp := &OrderTicket{}
	err := c.getJSON(url, nil, resp)
	return resp, err
}

func (c *Client) ListOrders(req *ListOrdersRequest) ([]*OrderTicket, error) {
	url := Endpoint + "/orders/"
	var result []*OrderTicket
	for url != "" {
		var resp struct {
			Results []*OrderTicket
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

func (c *Client) CancelOrder(orderId string) (*OrderTicket, error) {
	url := Endpoint + "/orders/" + orderId + "/cancel"
	resp := &OrderTicket{}
	err := c.postForm(url, nil, resp)
	return resp, err
}
