package simpos

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Name      string
	Cookie    []string
	TestCard  TestCard
	Shared    SharedConfig
	TestCases []TestCase
}

type SharedConfig struct {
	AmountMin                            float64 `yaml:"amountMin"`
	AmountMax                            float64 `yaml:"amountMax"`
	DefaultOriginalCurrencyCode          string  `yaml:"defaultOriginalCurrencyCode"`
	DefaultOriginalCurrencyDecimalPlaces string  `yaml:"defaultOriginalCurrencyDecimalPlaces"`
	Token                                string  `yaml:"token"`
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
	SettleType                    string `yaml:"settleType"`
	Reversal                      string
	Mcc                           string
	Source                        string
	Foreign                       bool
	OriginalCurrencyCode          string
	OriginalCurrencyDecimalPlaces string
	Acquirer                      string
	Province                      string
	Country                       string
	Advice                        bool
}

func ParseConfig(f string) (*Config, error) {
	raw, err := os.ReadFile(f)

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

func (c *Config) ParseConfig(f string) error {
	file, err := os.Open(f)
	if err != nil {
		return err
	}
	defer file.Close()
	d := yaml.NewDecoder(file)
	err = d.Decode(c)
	if err != nil {
		return err
	}
	return nil
}
