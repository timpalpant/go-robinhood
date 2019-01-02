package robinhood

import (
	"reflect"
	"testing"
	"time"
)

func TestListAccounts_Cash(t *testing.T) {
	responses, err := loadResponses("testdata/responses/accounts.cash.json")
	if err != nil {
		t.Fatal(err)
	}

	client := NewTestClient(responses)
	accts, err := client.ListAccounts()
	if err != nil {
		t.Errorf(err.Error())
	}

	if len(accts) != 1 {
		t.Errorf("got %d accounts, expected %d", len(accts), 6)
	}

	expected := &Account{
		UpdatedAt:    time.Date(2015, 9, 25, 18, 43, 10, 879108000, time.UTC),
		PortfolioURL: "https://api.robinhood.com/accounts/8UD09348/portfolio/",
		CashBalances: &CashBalances{
			CreatedAt:                  time.Date(2016, 3, 12, 1, 8, 27, 672943000, time.UTC),
			Cash:                       214.89,
			BuyingPower:                114.89,
			UpdatedAt:                  time.Date(2016, 3, 18, 9, 3, 59, 954927000, time.UTC),
			CashAvailableForWithdrawal: 114.89,
			UnsettledFunds:             100.0,
		},
		CashAvailableForWithdrawal: 114.89,
		Type:                       CashAccount,
		BuyingPower:                114.89,
		UserURL:                    "https://api.robinhood.com/user/",
		URL:                        "https://api.robinhood.com/accounts/8UD09348/",
		PositionsURL:               "https://api.robinhood.com/accounts/8UD09348/positions/",
		CreatedAt:                  time.Date(2016, 3, 12, 1, 8, 27, 672943000, time.UTC),
		Cash:                       114.89,
		AccountNumber:              "8UD09348",
		UnsettledFunds:             100.0,
	}

	if !reflect.DeepEqual(accts[0], expected) {
		t.Errorf("did not parse expected account: got %+v, wanted %+v", accts[0], expected)
	}
}

func TestListAccounts_Margin(t *testing.T) {
	responses, err := loadResponses("testdata/responses/accounts.margin.json")
	if err != nil {
		t.Fatal(err)
	}

	client := NewTestClient(responses)
	accts, err := client.ListAccounts()
	if err != nil {
		t.Errorf(err.Error())
	}

	if len(accts) != 1 {
		t.Errorf("got %d accounts, expected %d", len(accts), 6)
	}

	expected := &Account{
		UpdatedAt: time.Date(2016, 4, 13, 17, 4, 30, 664674000, time.UTC),
		MarginBalances: &MarginBalances{
			DayTradeBuyingPower:        1644.8050,
			CreatedAt:                  time.Date(2016, 4, 13, 17, 4, 30, 653404000, time.UTC),
			Cash:                       421.41,
			UnallocatedMarginCash:      612.39,
			UpdatedAt:                  time.Date(2016, 6, 30, 17, 25, 44, 637401000, time.UTC),
			CashAvailableForWithdrawal: 421.41,
			OvernightBuyingPower:       612.39,
			UnsettledFunds:             190.98,
			DayTradeRatio:              0.25,
			OvernightRatio:             1.0,
		},
		PortfolioURL:               "https://api.robinhood.com/accounts/8UD09348/portfolio/",
		CashAvailableForWithdrawal: 421.41,
		Type:                       MarginAccount,
		SMA:                        1629.26,
		SweepEnabled:               true,
		BuyingPower:                1629.26,
		UserURL:                    "https://api.robinhood.com/user/",
		MaxACHEarlyAccessAmount:    1000.0,
		URL:                        "https://api.robinhood.com/accounts/8UD09348/",
		PositionsURL:               "https://api.robinhood.com/accounts/8UD09348/positions/",
		CreatedAt:                  time.Date(2016, 3, 12, 1, 8, 27, 672943000, time.UTC),
		Cash:                       421.41,
		AccountNumber:              "8UD09348",
		UnsettledFunds:             190.98,
	}

	if !reflect.DeepEqual(accts[0], expected) {
		t.Errorf("did not parse expected account: got %+v, wanted %+v", accts[0], expected)
	}
}

func TestParseAccountNumber(t *testing.T) {
	accountURL := "https://api.robinhood.com/accounts/8UD09348/"
	id, err := ParseAccountNumber(accountURL)
	if err != nil {
		t.Error(err)
	}

	expected := "8UD09348"
	if id != expected {
		t.Errorf("got %v, expected %v", id, expected)
	}
}

func TestGetPortfolio(t *testing.T) {
	responses, err := loadResponses("testdata/responses/portfolio.json")
	if err != nil {
		t.Fatal(err)
	}

	client := NewTestClient(responses)
	portfolio, err := client.GetPortfolio("8UD09348")
	if err != nil {
		t.Errorf(err.Error())
	}

	expected := &Portfolio{
		URL:                         "https://api.robinhood.com/portfolios/8UD09348/",
		AdjustedEquityPreviousClose: 500.17,
		AccountURL:                  "https://api.robinhood.com/accounts/8UD09348/",
		LastCoreMarketValue:         34.07,
		LastCoreEquity:              499.66,
		Equity:                      499.66,
		MarketValue:                 34.07,
		EquityPreviousClose:         0.17,
	}

	if !reflect.DeepEqual(portfolio, expected) {
		t.Errorf("did not parse expected portfolio: got %+v, wanted %+v", portfolio, expected)
	}
}
