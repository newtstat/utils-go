package utils

import (
	"io"
	"strings"
	"testing"
)

var testJSON = JSON

type testJSONT struct {
	Test bool
}

func TestJsonT_Unmarshal(t *testing.T) {
	t.Helper()

	t.Run("normal/testJSON.Unmarshal", func(t *testing.T) {
		testStruct := testJSONT{}
		r := strings.NewReader(`{"test":true}`)
		if err := testJSON.Unmarshal(r, &testStruct); err != nil {
			t.Errorf("TestJsonT_Unmarshal(): testJSON.Unmarshal(): err != nil: %v", err)
		}
	})

	t.Run("non-normal/ioutil.ReadAll", func(t *testing.T) {
		testJSON.ioutilReadAllFn = func(r io.Reader) (bytes []byte, e error) {
			return nil, ErrorDummyErrorForTest
		}
		testStruct := testJSONT{}
		r := strings.NewReader(`{"test":true}`)
		if err := testJSON.Unmarshal(r, &testStruct); err == nil {
			t.Errorf("TestJsonT_Unmarshal(): testJSON.Unmarshal(): err == nil: %v", err)
		}
	})

	t.Run("non-normal/testJSON.Unmarshal", func(t *testing.T) {
		testStruct := testJSONT{}
		r := strings.NewReader(`{"test":true}` + `UnmarshalError`)
		if err := testJSON.Unmarshal(r, &testStruct); err == nil {
			t.Errorf("TestJsonT_Unmarshal(): testJSON.Unmarshal(): err == nil: %v", err)
		}
	})
}
