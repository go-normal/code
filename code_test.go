package code

import (
	"testing"
)

func TestAdd(t *testing.T) {
	Add(12345, "this is the error message for error 12345")
}

func TestAdd2ST(t *testing.T) {
	code := 123456
	_, err := Add(code, string(code))
	if err != nil {
		t.Error(err)
	}
	_, err = Add(code, string(code))

	if err == nil {
		t.Error("duplicate error code")
	}
}

func TestMessage(t *testing.T) {
	paramError := MustAdd(1234567, "param error: %s >> %s")
	if paramError.Message("mobile", "invalid length") != "param error: mobile >> invalid length" {
		t.Fail()
	}
}
