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
	GoodTillCanceled  TimeInForce = "gtc"
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
	AccountURL             string `json:"account",url:"account,omitifempty"`
	InstrumentURL          string `json:"instrument",url:"instrument,omitifempty"`
	Symbol                 string
	Type                   OrderType
	TimeInForce            TimeInForce `json:"time_in_force",url:"time_in_force,omitifempty"`
	Trigger                Trigger
	Price                  *float64
	StopPrice              *float64 `json:"stop_price",url:"stop_price,omitifempty"`
	Quantity               float64  `json:"quantity,string",url:"quantity,omitifempty"`
	Side                   OrderSide
	ClientID               *string `json:"client_id",url:"client_id,omitifempty"`
	ExtendedHours          bool    `json:"extended_hours",url:"extended_hours,omitifempty"`
	OverrideDayTradeChecks bool    `json:"override_day_trade_checks",url:"override_day_trade_checks,omitifempty"`
	OverrideDTBPChecks     bool    `json:"override_dtbp_checks",url:"override_dtbp_checks,omitifempty"`
}

type OrderTicket struct {
	*Order
	ID                 string
	UpdatedAt          time.Time `json:"updated_at"`
	Executions         []string
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
	Since         time.Time `url:"since,omitifempty"`
	InstrumentURL string    `url:"instrument,omitifempty"`
	Cursor        string    `url:"cursor,omitifempty"`
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
	for {
		var resp struct {
			Results []*OrderTicket
			Next    string
		}

		err := c.getJSON(url, req, &resp)
		if err != nil {
			return nil, err
		}

		result = append(result, resp.Results...)

		if resp.Next == "" {
			break
		}

		req.Cursor = resp.Next
	}

	return result, nil
}

func (c *Client) CancelOrder(orderId string) (*OrderTicket, error) {
	url := Endpoint + "/orders/" + orderId + "/cancel"
	resp := &OrderTicket{}
	err := c.postForm(url, nil, resp)
	return resp, err
}
