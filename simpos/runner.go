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
				fmt.Println("WALLET REQUEST\n", result.WalletRequest)
				fmt.Println("WALLET RESPONSE\n", result.WalletResponse)

				if result.ReversalWalletRequest != "" {
					fmt.Println("REVERSAL WALLET REQUEST\n", result.ReversalWalletRequest)
					fmt.Println("REVERSAL WALLET RESPONSE\n", result.ReversalWalletResponse)
				}
				fmt.Println("DE 39: ", result.IsoResponsePacket["39"])
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
