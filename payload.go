package main

import "fmt"

func NewPayload(tc *TestCase, shared *SharedTestConfig, card *TestCard) (Payload, error) {
	mode := tc.Mode

	switch mode {
	case pos, web:
		return makeAuth(tc, shared, card)
	case settlement:
		return makeSettle(tc, shared, card)

	}
	return nil, ErrUnsupportedMode
}

func makeAuth(tc *TestCase, shared *SharedTestConfig, card *TestCard) (Payload, error) {
	amount := randomizeAmount(shared)

	pl := &Auth{}

	mode := tc.Mode

	switch mode {
	case pos:
	case web:
		switch tc.Reversal {
		case "":
			pl.Method = webDeduct
		}
	}

	pl.Params.Amount = fmt.Sprintf("%.2f", amount)
	pl.Params.CardNumber = card.Number
	pl.Params.Expirydate = card.ExpiryDate

	if tc.Foreign {
		pl.Params.Foreign = "1"
		pl.Params.OriginalAmount = fmt.Sprintf("%.2f", randomizeAmount(shared))

		if tc.OriginalCurrencyCode != "" {
			pl.Params.OriginalCurrencyCode = tc.OriginalCurrencyCode
			pl.Params.OriginalCurrencyDecimalPlaces = tc.OriginalCurrencyDecimalPlaces
		} else {
			pl.Params.OriginalCurrencyCode = shared.DefaultOriginalCurrencyCode
			pl.Params.OriginalCurrencyDecimalPlaces = shared.DefaultOriginalCurrencyPlaces
		}
	} else {
		pl.Params.Foreign = "0"
		pl.Params.OriginalAmount = ""
		pl.Params.OriginalCurrencyCode = ""
		pl.Params.OriginalCurrencyDecimalPlaces = ""
	}

	pl.Params.Network = "0"
	pl.Params.Pin = card.Pin

	pl.Params.Source = tc.Source
	pl.Params.Token = shared.Token

	pl.Params.Acquirer = formatAcquirer(tc.Acquirer, acquirerLength)
	pl.Params.Province = formatAcquirer(tc.Province, provinceLength)
	pl.Params.Country = formatAcquirer(tc.Country, countryLength)

	return pl, nil
}

func makeSettle(tc *TestCase, shared *SharedTestConfig, card *TestCard) (Payload, error) {
	if tc.SettleType == "none" {
		return nil, ErrSettleTypeNotSet
	}
	pl := &Settle{}

	return pl, nil
}
