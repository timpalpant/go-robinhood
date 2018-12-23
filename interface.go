package robinhood

type Instrument struct {
	MinTickSize        *float64 `json:"min_tick_size,string"`
	Splits             string
	MarginInitialRatio float64
	URL                string
	Quote              string
	Symbol             string
	BloombergUnique    string
	ListDate           string `json:"list_date"`
	Fundamentals       string
	State              string
	Country            string
	DayTradeRatio      float64 `json:"day_trade_ratio,string"`
	Tradeable          bool
	MaintenanceRatio   float64 `json:"maintenance_ratio,string"`
	ID                 string
	Market             string
	Name               string
}
