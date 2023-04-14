package model

type CalRsp struct {
	SecurityId   string `json:"securityId"`
	Price        string `json:"price"`
	Yield        string `json:"yield"`
	StrikeYield  string `json:"strikeYield"`
	TaxesAccrued string `json:"taxesAccrued"`
	SettleDay    string `json:"settleDay"`
	Success      bool   `json:"success"`
	Message      string `json:"message"`
}

type CalReq struct {
	ClearSpeed       string  `json:"clearSpeed" v:"bail|required|in:1,2,3"`
	CurrencyRate     string  `json:"currencyRate"`
	Price            string  `json:"price"`
	SecurityId       string  `json:"securityId"`
	SettleCurrency   string  `json:"settleCurrency"`
	SettleDay        string  `json:"settleDay"`
	StrikeYield      string  `json:"strikeYield"`
	TaxesAccrued     string  `json:"taxesAccrued"`
	TradeDate        string  `json:"tradeDate"`
	Yield            float64 `json:"yield"`
	TransportSession string  `json:"transportSession"`
}
