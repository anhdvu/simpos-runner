package simpos

import (
	"errors"
)

const (
	baseUrl         string = "https://tools.uat.tutuka.cloud/api/json.cfm"
	reversalQueue   string = "https://ve.uat.tutuka.cloud/index.cfm?FuseAction=tasks.processCompanionReversals"
	adjustmentQueue string = "https://ve.uat.tutuka.cloud/index.cfm?FuseAction=tasks.processCompanionAdjustments"

	pos            string = "pos"
	web            string = "web"
	settlement     string = "settlement"
	acquirerLength int    = 22
	provinceLength int    = 13
	countryLength  int    = 3

	posDeduct        string = "POSDeduct"
	webDeduct        string = "WEBDeduct"
	posDeductRev     string = "POSDeductReverse"
	webDeductReverse string = "WEBDeductReverse"
	refundAdj        string = "RefundAdjustment"
	forexAdj         string = "ForexAdjustment"

	refund   string = "refund"
	fxload   string = "fxload"
	fxdeduct string = "fxdeduct"

	reversal   string = "reversal"
	adjustment string = "adjustment"
	both       string = "both"

	mag string = "mag"
	emv string = "emv"
	nfc string = "nfc"

	timeout int = 300
)

var (
	ErrUnsupportedMode  = errors.New("The specified mode is not supported. Please check for typos or specify a supported mode.")
	ErrSettleTypeNotSet = errors.New("In Settlement mode but SettleType was not set. Please check config file again.")
	ErrTokenUnavailable = errors.New("Couldn't retrieve token. Please check provided cookie header in config file.")
)
