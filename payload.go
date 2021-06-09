package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Payload interface {
	GetMethod() string
}

type Auth struct {
	Method string `json:"method"`
	Params struct {
		Amount                        string `json:"amount"`
		CardNumber                    string `json:"cardNumber"`
		Cvv                           string `json:"cvv,omitempty"`
		Expirydate                    string `json:"expirydate"`
		Foreign                       string `json:"foreign"`
		IsPartialReversal             string `json:"isPartialReversal,omitempty"`
		MerchantCategoryCode          string `json:"merchantCategoryCode"`
		Network                       string `json:"network"`
		OriginalAmount                string `json:"originalAmount,omitempty"`
		OriginalCurrencyCode          string `json:"originalCurrencyCode,omitempty"`
		OriginalCurrencyDecimalPlaces string `json:"originalCurrencyDecimalPlaces,omitempty"`
		PartialReversalAmount         string `json:"partialReversalAmount,omitempty"`
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
		OriginalAmount                string `json:"originalAmount,omitempty"`
		OriginalCurrencyCode          string `json:"originalCurrencyCode"`
		OriginalCurrencyDecimalPlaces string `json:"originalCurrencyDecimalPlaces,omitempty"`
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

func NewPayload(tc TestCase, shared SharedConfig, card TestCard) (Payload, error) {
	mode := tc.Mode

	switch mode {
	case pos, web:
		return makeAuth(tc, shared, card)
	case settlement:
		return makeSettle(tc, shared, card)

	}
	return nil, ErrUnsupportedMode
}

func makeAuth(tc TestCase, shared SharedConfig, card TestCard) (Payload, error) {
	amount := randomizeAmount(shared)

	pl := &Auth{}

	pl.Params.Amount = fmt.Sprintf("%.2f", amount)
	pl.Params.CardNumber = card.Number
	pl.Params.Expirydate = card.ExpiryDate

	mode := tc.Mode

	// Handle Method and its related fields
	switch mode {
	case pos:
		switch tc.Reversal {
		case "full":
			pl.Method = posDeductRev
		case "partial":
			pl.Method = posDeductRev
			pl.Params.IsPartialReversal = "1"
			pl.Params.PartialReversalAmount = fmt.Sprintf("%.2f", makePartialAmount(amount))
		default:
			pl.Method = posDeduct
		}

		if tc.ATM {
			pl.Params.Type = "01"
		} else {
			pl.Params.Type = "00"
		}

		pl.Params.Network = "0"
		pl.Params.Pin = card.Pin
		if tc.Source == mag || tc.Source == nfc {
			pl.Params.Source = strings.ToUpper(tc.Source)
		} else {
			tc.Source = strings.ToUpper(emv)
		}

	case web:
		switch tc.Reversal {
		case "full":
			pl.Method = webDeductReverse
		case "partial":
			pl.Method = webDeductReverse
			pl.Params.IsPartialReversal = "1"
			pl.Params.PartialReversalAmount = fmt.Sprintf("%.2f", randomizeAmount(shared))
		default:
			pl.Method = webDeduct
		}

		pl.Params.Cvv = card.Cvv
		pl.Params.Pin = ""
		pl.Params.Source = ""
		pl.Params.Type = ""
		pl.Params.Network = ""

	}

	// Handle fields related to foreign transactions
	if tc.Foreign {
		pl.Params.Foreign = "1"
		pl.Params.OriginalAmount = fmt.Sprintf("%.2f", randomizeAmount(shared))
		if tc.OriginalCurrencyCode != "" {
			pl.Params.OriginalCurrencyCode = tc.OriginalCurrencyCode
			pl.Params.OriginalCurrencyDecimalPlaces = tc.OriginalCurrencyDecimalPlaces
		} else {
			pl.Params.OriginalCurrencyCode = shared.DefaultOriginalCurrencyCode
			pl.Params.OriginalCurrencyDecimalPlaces = shared.DefaultOriginalCurrencyDecimalPlaces
		}
	} else {
		pl.Params.Foreign = "0"
		pl.Params.OriginalAmount = ""
		pl.Params.OriginalCurrencyCode = ""
		pl.Params.OriginalCurrencyDecimalPlaces = ""
	}

	pl.Params.MerchantCategoryCode = tc.Mcc
	pl.Params.Token = shared.Token
	pl.Params.Acquirer = formatAcquirer(tc.Acquirer, acquirerLength)
	pl.Params.Province = formatAcquirer(tc.Province, provinceLength)
	pl.Params.Country = formatAcquirer(tc.Country, countryLength)

	return pl, nil
}

func makeSettle(tc TestCase, shared SharedConfig, card TestCard) (Payload, error) {
	amount := randomizeAmount(shared)

	pl := &Settle{}
	pl.Params.Amount = fmt.Sprintf("%.2f", amount)
	pl.Params.CardNumber = card.Number

	switch tc.SettleType {
	case refund:
		pl.Method = refundAdj
		pl.Params.SettlementAmount = ""
		pl.Params.Type = "20"
	case fxdeduct:
		pl.Method = forexAdj
		pl.Params.SettlementAmount = fmt.Sprintf("%.2f", amount+makePartialAmount(amount))
		pl.Params.Type = "21"
	case fxload:
		pl.Method = forexAdj
		pl.Params.SettlementAmount = fmt.Sprintf("%.2f", amount-makePartialAmount(amount))
		pl.Params.Type = "20"
	default:
		return nil, ErrSettleTypeNotSet
	}

	if tc.OriginalCurrencyCode != "" {
		pl.Params.OriginalCurrencyCode = tc.OriginalCurrencyCode
	} else {
		pl.Params.OriginalCurrencyCode = shared.DefaultOriginalCurrencyCode
	}

	if tc.Foreign {
		pl.Params.Foreign = "1"
		pl.Params.OriginalAmount = fmt.Sprintf("%.2f", randomizeAmount(shared))
		if tc.OriginalCurrencyDecimalPlaces != "" {
			pl.Params.OriginalCurrencyDecimalPlaces = tc.OriginalCurrencyDecimalPlaces
		} else {
			pl.Params.OriginalCurrencyDecimalPlaces = shared.DefaultOriginalCurrencyDecimalPlaces
		}
	} else {
		pl.Params.Foreign = "0"
		pl.Params.OriginalAmount = ""
		pl.Params.OriginalCurrencyDecimalPlaces = ""
	}

	pl.Params.Partial = false
	pl.Params.MerchantCategoryCode = tc.Mcc
	pl.Params.Token = shared.Token
	pl.Params.Acquirer = formatAcquirer(tc.Acquirer, acquirerLength)
	pl.Params.Province = formatAcquirer(tc.Province, provinceLength)
	pl.Params.Country = formatAcquirer(tc.Country, countryLength)

	return pl, nil
}

func NewRequest(p Payload) (*http.Request, error) {
	buf := &bytes.Buffer{}
	e := json.NewEncoder(buf)
	err := e.Encode(p)
	if err != nil {
		fmt.Println("Error occured at function NewRequest.")
		return nil, err
	}

	return http.NewRequest(http.MethodPost, baseUrl, buf)
}

type Result struct {
	IsoRequest             string
	IsoResponse            string
	IsoResponsePacket      map[string]string
	ResultCode             int
	ResultText             string
	ReversalIsoRequest     string
	ReversalIsoResponse    string
	ReversalWalletRequest  string
	ReversalWalletResponse string
	WalletRequest          string
	WalletResponse         string
}

func (r *Result) FromJSON(b io.Reader) error {
	d := json.NewDecoder(b)
	return d.Decode(r)
}
