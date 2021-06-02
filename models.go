package main

type Config struct {
	Name             string
	Cookie           []string
	TestCard         TestCard
	SharedTestConfig SharedTestConfig
	TestCases        []TestCase
}

type SharedTestConfig struct {
	AmountMin                     float64
	AmountMax                     float64
	DefaultOriginalCurrencyCode   string
	DefaultOriginalCurrencyPlaces string
	Token                         string
}

type TestCard struct {
	Number     string
	ExpiryDate string
	Cvv        string
	Pin        string
}

type TestCase struct {
	Included                      bool
	Name                          string
	Run                           int
	Mode                          string
	ATM                           bool
	SettleType                    string
	Reversal                      string
	Mcc                           string
	Source                        string
	Foreign                       bool
	OriginalCurrencyCode          string
	OriginalCurrencyDecimalPlaces string
	Acquirer                      string
	Province                      string
	Country                       string
}

type Payload interface {
	GetMethod() string
}

type Auth struct {
	Method string `json:"method"`
	Params struct {
		Amount                        string `json:"amount"`
		CardNumber                    string `json:"cardNumber"`
		Cvv                           string `json:"cvv;omitempty"`
		Expirydate                    string `json:"expirydate"`
		Foreign                       string `json:"foreign"`
		IsPartialReversal             string `json:"isPartialReversal;omitempty"`
		MerchantCategoryCode          string `json:"merchantCategoryCode"`
		Network                       string `json:"network"`
		OriginalAmount                string `json:"originalAmount;omitempty"`
		OriginalCurrencyCode          string `json:"originalCurrencyCode;omitempty"`
		OriginalCurrencyDecimalPlaces string `json:"originalCurrencyDecimalPlaces;omitempty"`
		PartialReversalAmount         string `json:"partialReversalAmount;omitempty"`
		Pin                           string `json:"pin"`
		Source                        string `json:"source"`
		Token                         string `json:"token"`
		Type                          string `json:"type"`
		Acquirer                      string `json:"acquirer"`
		Province                      string `json:"province"`
		Country                       string `json:"country"`
	} `json:"params"`
}

func (a *Auth) GetMethod() string {
	return a.Method
}

type Settle struct {
	Method string `json:"method"`
	Params struct {
		Acquirer                      string `json:"acquirer"`
		Amount                        string `json:"amount"`
		CardNumber                    string `json:"cardNumber"`
		Country                       string `json:"country"`
		Foreign                       string `json:"foreign"`
		MerchantCategoryCode          string `json:"merchantCategoryCode"`
		OriginalAmount                string `json:"originalAmount"`
		OriginalCurrencyCode          string `json:"originalCurrencyCode"`
		OriginalCurrencyDecimalPlaces string `json:"originalCurrencyDecimalPlaces;omitempty"`
		Partial                       bool   `json:"partial"`
		Province                      string `json:"province"`
		SettlementAmount              string `json:"settlementAmount"`
		Token                         string `json:"token"`
		Type                          string `json:"type"`
	} `json:"params"`
}

func (s *Settle) GetMethod() string {
	return s.Method
}
