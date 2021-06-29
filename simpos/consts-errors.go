package simpos

import (
	"errors"
)

const (
	baseUrl          string = "https://tools.uat.tutuka.cloud/api/json.cfm"
	companionTaskUrl string = "https://ve.uat.tutuka.cloud/index.cfm?FuseAction=tasks.processCompanion"

	pos            string = "pos"
	web            string = "web"
	settlement     string = "settlement"
	payment        string = "payment"
	acquirerLength int    = 22
	provinceLength int    = 13
	countryLength  int    = 3

	posDeduct        string = "POSDeduct"
	webDeduct        string = "WEBDeduct"
	posDeductRev     string = "POSDeductReverse"
	webDeductReverse string = "WEBDeductReverse"
	refundAdj        string = "RefundAdjustment"
	forexAdj         string = "ForexAdjustment"
	refundAuth       string = "RefundAuth"
	loadAdj          string = "LoadAdjustment"

	refund     string = "refund"
	fxload     string = "fxload"
	fxdeduct   string = "fxdeduct"
	noauth     string = "noauth"
	chargeback string = "chargeback"

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
	ErrQueueRun         = errors.New("Unable to run the queue")
	ErrNoQueueSpecified = errors.New("No proper queue specified. Please check typo.")
)

var taskUrls = map[string]string{
	"reversal":   "https://ve.uat.tutuka.cloud/index.cfm?FuseAction=tasks.processCompanionReversals",
	"adjustment": "https://ve.uat.tutuka.cloud/index.cfm?FuseAction=tasks.processCompanionAdjustments",
}
