package simpos

import (
	"bytes"
	"encoding/json"
	"io"
	"reflect"
	"testing"
)

type MockPayload struct {
	Name string `json:"name"`
}

func (mp *MockPayload) JSON(w io.Writer) error {
	return json.NewEncoder(w).Encode(mp)
}
func TestPayload_JSON(t *testing.T) {
	var mock Payload = &MockPayload{Name: "mock"}
	buf := &bytes.Buffer{}
	want := `{"name":"mock"}
`
	err := mock.JSON(buf)
	got := buf.String()
	if err != nil {
		t.Errorf("Expected err == nil but got err == %v", err)
	}
	if want != got {
		t.Errorf("Wanted %q but got %q", want, got)
	}
}

func TestResult_FromJSON(t *testing.T) {
	t.Run("Happy path", func(t *testing.T) {
		data := `{
			"isoRequest": "xxx",
			"isoResponse": "xxx",
			"isoResponsePacket": {
				"39": "R0000"
			},
			"resultCode": 1,
			"resultText": "success",
			"walletRequest": "zzzz",
			"walletResponse": "zzzz"
			}`

		dataReader := bytes.NewBufferString(data)

		got := &Result{}
		err := got.FromJSON(dataReader)

		if err != nil {
			t.Errorf("Expected error == nil but got error != nil.")
		}

		want := &Result{
			IsoRequest:  "xxx",
			IsoResponse: "xxx",
			IsoResponsePacket: map[string]string{
				"39": "R0000",
			},
			ResultCode:     1,
			ResultText:     "success",
			WalletRequest:  "zzzz",
			WalletResponse: "zzzz",
		}

		if !reflect.DeepEqual(want, got) {
			t.Errorf("Wanted %+v, but got %+v", want, got)
		}
	})

	t.Run("Error case", func(t *testing.T) {
		data := `{
			"isoRequest": 123,
			"isoRes": "xxx",
			"isoResponsePackage": {
				"39": "R0000"
			},
			"resultCode": 1,
			"resultText": "success",
			"walletRequest": "zzzz",
			"walletResponse": "zzzz"
			}`

		dataReader := bytes.NewBufferString(data)

		got := &Result{}
		err := got.FromJSON(dataReader)

		if err == nil {
			t.Errorf("Expected error == nil but got error != nil.")
		}
	})
}
