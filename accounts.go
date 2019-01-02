package robinhood

import (
	"time"
)

type AccountType string

const (
	CashAccount   AccountType = "cash"
	MarginAccount AccountType = "margin"
)

type CashBalances struct {
	CashHeldForOrders          float64   `json:"cash_held_for_orders,string"`
	CreatedAt                  time.Time `json:"created_at"`
	UpdatedAt                  time.Time `json:"updated_at"`
	Cash                       float64   `json:",string"`
	BuyingPower                float64   `json:"buying_power,string"`
	CashAvailableForWithdrawal float64   `json:"cash_available_for_withdrawal,string"`
	UnclearedDeposits          float64   `json:"uncleared_deposits,string"`
	UnsettledFunds             float64   `json:"unsettled_funds,string"`
}

type MarginBalances struct {
	DayTradeBuyingPower               float64   `json:"day_trade_buying_power,string"`
	CreatedAt                         time.Time `json:"created_at"`
	UpdatedAt                         time.Time `json:"updated_at"`
	OvernightBuyingPowerHeldForOrders float64   `json:"overnight_buying_power_held_for_orders,string"`
	CashHeldForOrders                 float64   `json:"cash_held_for_orders,string"`
	DayTradeBuyingPowerHeldForOrders  float64   `json:"day_trade_buying_power_held_for_orders,string"`
	MarkedPatternDayTraderDate        *time.Time
	Cash                              float64 `json:",string"`
	UnallocatedMarginCash             float64 `json:"unallocated_margin_cash,string"`
	CashAvailableForWithdrawal        float64 `json:"cash_available_for_withdrawal,string"`
	MarginLimit                       float64 `json:"margin_limit,string"`
	OvernightBuyingPower              float64 `json:"overnight_buying_power,string"`
	UnclearedDeposits                 float64 `json:"uncleared_deposits,string"`
	UnsettledFunds                    float64 `json:"unsettled_funds,string"`
	DayTradeRatio                     float64 `json:"day_trade_ratio,string"`
	OvernightRatio                    float64 `json:"overnight_ratio,string"`
}

type Account struct {
	Deactivated                bool
	URL                        string
	CreatedAt                  time.Time       `json:"created_at"`
	UpdatedAt                  time.Time       `json:"updated_at"`
	CashBalances               *CashBalances   `json:"cash_balances"`
	MarginBalances             *MarginBalances `json:"margin_balances"`
	PortfolioURL               string          `json:"portfolio"`
	WithdrawalHalted           bool            `json:"withdrawal_halted"`
	CashAvailableForWithdrawal float64         `json:"cash_available_for_withdrawal,string"`
	Type                       AccountType
	SMA                        float64 `json:",string"`
	SweepEnabled               bool    `json:"sweep_enabled"`
	DepositHalted              bool    `json:"deposit_halted"`
	BuyingPower                float64 `json:"buying_power,string"`
	UserURL                    string  `json:"user"`
	MaxACHEarlyAccessAmount    float64 `json:"max_ach_early_access_amount,string"`
	CashHeldForOrders          float64 `json:"cash_held_for_orders,string"`
	OnlyPositionClosingTrades  bool    `json:"only_position_closing_trades"`
	PositionsURL               string  `json:"positions"`
	Cash                       float64 `json:",string"`
	SMAHeldForOrders           float64 `json:"sma_held_for_orders,string"`
	AccountNumber              string  `json:"account_number"`
	UnclearedDeposits          float64 `json:"uncleared_deposits,string"`
	UnsettledFunds             float64 `json:"unsettled_funds,string"`
}

func (c *Client) ListAccounts() ([]*Account, error) {
	url := Endpoint + "/accounts/"
	var result []*Account
	for url != "" {
		var resp struct {
			Results []*Account
			Next    string
		}

		err := c.getJSON(url, nil, &resp)
		if err != nil {
			return nil, err
		}

		result = append(result, resp.Results...)
		url = resp.Next
	}

	return result, nil
}
