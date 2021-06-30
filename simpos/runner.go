package simpos

import (
	"fmt"
	"net/http"
	"time"
)

func Run(f string) {
	c, err := ParseConfig(f)
	if err != nil {
		fmt.Println("ERROR occured during config parsing", err)
		return
	}

	token, err := GetToken(c.Cookie)
	if err != nil {
		fmt.Println("ERROR occured during token retrieval", err)
		return
	}
	c.Shared.Token = token
	testcases := c.TestCases

	card := c.TestCard
	shared := c.Shared

	runner := &http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}
	t := time.Now()
	fmt.Println("========== START ==========")
	for _, tc := range testcases {
		if tc.Included {
			for i := 0; i < tc.Runs; i++ {
				fmt.Printf("Running test case %q. Run: %d\n", tc.Name, i+1)
				pl, _ := NewPayload(tc, shared, card)
				req, err := NewRequest(pl)
				if err != nil {
					fmt.Println("ERROR: Unable to create a new request with payload ", i, err)
					continue
				}
				resp, err := runner.Do(req)
				if err != nil {
					fmt.Println("ERROR: The API server failed to respond to one or more instances of payload ", i, err)
					continue
				}
				result := &Result{}
				err = result.FromJSON(resp.Body)
				if err != nil {
					fmt.Println("ERROR: Unable to parse json in response body", err)
				}
				fmt.Println("========== RESULT ==========")
				fmt.Printf("Result Code: %v\n", result.ResultCode)
				fmt.Printf("Result Text: %v\n\n", result.ResultText)
				fmt.Printf("WALLET REQUEST\n%v\n\n", result.WalletRequest)
				fmt.Printf("WALLET RESPONSE\n%v\n\n", result.WalletResponse)

				if result.ReversalWalletRequest != "" {
					fmt.Printf("REVERSAL WALLET REQUEST\n%v\n\n", result.ReversalWalletRequest)
					fmt.Printf("REVERSAL WALLET RESPONSE\n%v\n\n", result.ReversalWalletResponse)
				}

				p, ok := result.IsoResponsePacket.(map[string]interface{})
				if ok {
					fmt.Printf("ISO Response - DE 39: %v\n", p["39"])
				}

				fmt.Println("============================")
				time.Sleep(time.Millisecond)
			}
		}
	}

	fmt.Println("Running adjustment queue to send adjustment payloads...")
	err = runTask(companionTaskUrl, adjustment)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Please try to run it manually.")
	}
	fmt.Println("Adjustment Queue has been processed.")
	fmt.Println("DONE!")
	duration := time.Since(t)
	fmt.Printf("It took %v to run all tests.", duration)
}
