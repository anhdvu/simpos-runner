package simpos

import (
	"encoding/base64"
	"strings"
	"testing"
)

func TestGetToken(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		cookie := []string{"CFID=11258", "CFTOKEN=e15ae0e6d40a7cf8-BB2533BF-AE61-662B-BFEA8BA789C9CB2E"}

		token, err := GetToken(cookie)

		if err != nil {
			t.Errorf("Expected err == nil but got err != nil. %v", err)
		}

		parts := strings.Split(token, ".")

		want := `{"typ":"JWT","alg":"HS256"}`

		got, _ := base64.StdEncoding.DecodeString(parts[0])

		if string(got) != want {
			t.Errorf("Wanted %q but got %q", want, string(got))
		}

	})

	t.Run("invalid cookie", func(t *testing.T) {
		cookie := []string{"a=b", "x=y"}

		_, err := GetToken(cookie)

		if err == nil {
			t.Errorf("Expected error != nil but got error == nil")
		}

		if err != ErrTokenUnavailable {
			t.Errorf("Expected %v but got a different error.", ErrTokenUnavailable)
		}
	})
}
