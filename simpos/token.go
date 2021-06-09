package simpos

import (
	"fmt"
	"net/http"
	"net/url"

	"golang.org/x/net/html"
)

func GetToken(cookie []string) (string, error) {
	simPOSURL := &url.URL{
		Scheme:   "https",
		Host:     "tools.uat.tutuka.cloud",
		Path:     "/",
		RawQuery: "target=simpos",
	}

	reqHeader := http.Header{"cookie": cookie}

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
	for {
		tt := z.Next()
		switch {
		case tt == html.ErrorToken:
			// End of the document, it's done
			return "", ErrTokenUnavailable
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
						return jwt, nil
					}
				}
			}
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
