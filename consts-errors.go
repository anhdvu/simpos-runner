package main

import "errors"

const (
	pos            string = "pos"
	web            string = "web"
	settlement     string = "settlement"
	acquirerLength int    = 22
	provinceLength int    = 13
	countryLength  int    = 3
	baseUrl        string = "https://tools.uat.tutuka.cloud/api/json.cfm"
)

var (
	ErrUnsupportedMode  = errors.New("The specified mode is not supported. Please check for typos or specify a supported mode.")
	ErrSettleTypeNotSet = errors.New("In Settlement mode but SettleType was not set. Please check config file again.")
)
