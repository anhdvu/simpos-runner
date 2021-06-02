package main

import (
	"fmt"
	"net/http"
	"net/url"

	"golang.org/x/net/html"
)

func GetToken() {
	simPOSURL := &url.URL{
		Scheme:   "https",
		Host:     "tools.uat.tutuka.cloud",
		Path:     "/",
		RawQuery: "target=simpos",
	}

	reqHeader := http.Header{"cookie": []string{"CFID=20557", "CFTOKEN=3b79a1e3c773c94-E7F38875-D295-8B7C-22E05E8FB40599AD", "AWSALB=tugAHNMcNux96I1rB8PSGBdahIZcYro+F/J3FXiAMgYvpsE+9Z7nGMmQgVNbx9M/S5t5WBSHseSjVPZnvphy6yk9ilb5Q4/ABOHnqNMOswe9TU7nzLeLe5q7chMA", "AWSALBCORS=tugAHNMcNux96I1rB8PSGBdahIZcYro+F/J3FXiAMgYvpsE+9Z7nGMmQgVNbx9M/S5t5WBSHseSjVPZnvphy6yk9ilb5Q4/ABOHnqNMOswe9TU7nzLeLe5q7chMA", "CFGLOBALS=urltoken=CFID#=20557&CFTOKEN#=3b79a1e3c773c94-E7F38875-D295-8B7C-22E05E8FB40599AD#lastvisit={ts '2021-04-29 06:46:58'}#hitcount=683#timecreated={ts '2020-04-20 10:46:05'}#cftoken=3b79a1e3c773c94-E7F38875-D295-8B7C-22E05E8FB40599AD#cfid=20557#"}}

	req := &http.Request{
		Method: http.MethodGet,
		URL:    simPOSURL,
		Header: reqHeader,
	}

	response, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	z := html.NewTokenizer(response.Body)
	fmt.Println(z.TagName())
	for {
		tt := z.Next()
		switch {
		case tt == html.ErrorToken:
			// End of the document, it's done
			return
		case tt == html.StartTagToken:
			t := z.Token()

			// Check if the token is an <input> tag
			if t.Data == "input" {
				for _, a := range t.Attr {
					if a.Key == "id" && a.Val == "token" {
						ok, jwt := getJWToken(t)
						if !ok {
							continue
						}
						fmt.Println(jwt)
						break
					}
				}
			}
			break
		}
	}
}

func getJWToken(t html.Token) (ok bool, jwt string) {
	// Iterate over token attributes until we find an "href"
	for _, attr := range t.Attr {
		if attr.Key == "value" {
			jwt = attr.Val
			ok = true
		}
	}

	// "bare" return will return the variables (ok, href) as
	// defined in the function definition
	return
}
