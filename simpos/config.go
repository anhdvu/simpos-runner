package simpos

import (
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Name      string       `yaml:"name"`
	Cookie    []string     `yaml:"cookie"`
	TestCard  TestCard     `yaml:"testcard"`
	Shared    SharedConfig `yaml:"shared"`
	TestCases []TestCase   `yaml:"testcases"`
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
	Included                      bool   `yaml:"included"`
	Name                          string `yaml:"name"`
	Runs                          int    `yaml:"runs"`
	Mode                          string `yaml:"mode"`
	Function                      string `yaml:"function"`
	ATM                           bool   `yaml:"atm"`
	Source                        string `yaml:"source"`
	Foreign                       bool   `yaml:"foreign"`
	OriginalCurrencyCode          string `yaml:"originalCurrencyCode"`
	OriginalCurrencyDecimalPlaces string `yaml:"originalCurrencyDecimalPlaces"`
	Acquirer                      string `yaml:"acquirer"`
	Province                      string `yaml:"province"`
	Country                       string `yaml:"country"`
	Mcc                           string `yaml:"mcc"`
	Reversal                      string `yaml:"reversal"`
	Advice                        bool   `yaml:"advice"`
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

func (c *Config) fromJSON(r io.Reader) error {
	d := yaml.NewDecoder(r)
	err := d.Decode(c)
	if err != nil {
		return err
	}
	return nil
}
