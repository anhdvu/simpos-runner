package main

import (
	"os"

	"gopkg.in/yaml.v3"
)

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
	Runs                          int
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

func ParseConfig() (*Config, error) {
	raw, err := os.ReadFile("config.yaml")

	if err != nil {
		return nil, err
	}

	config := &Config{}
	err = yaml.Unmarshal(raw, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
