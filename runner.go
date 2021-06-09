package main

import (
	"fmt"
	"net/http"
	"time"
)

func Run() {

	c, _ := ParseConfig("config.yaml")
	token, _ := GetToken(c.Cookie)
	c.Shared.Token = token
	testcases := c.TestCases

	// fmt.Printf("%+v\n", testcases)
	card := c.TestCard
	shared := c.Shared

	runner := &http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}

	var pls []Payload

	for _, tc := range testcases {
		if tc.Included {
			for i := 0; i < tc.Runs; i++ {
				pl, _ := NewPayload(tc, shared, card)
				pls = append(pls, pl)
				time.Sleep(time.Millisecond)
			}
		}
	}

	for _, pl := range pls {
		req, err := NewRequest(pl)
		if err != nil {
			fmt.Println(err)
			continue
		}
		resp, err := runner.Do(req)
		if err != nil {
			fmt.Println(err)
			continue
		}
		data := &Result{}
		err = data.FromJSON(resp.Body)
		if err != nil {
			fmt.Println("Could not parse json in response body", err)
		}

		fmt.Printf("%+v\n", data)
	}
}
