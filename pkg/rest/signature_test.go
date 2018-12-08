package rest

import (
	"testing"
	"time"
)

func TestCreateHmacBase64Signature(t *testing.T) {

	sig := &HmacSignature{
		"GET",
		"/api/test",
		"application/json",
		time.UTC.String(),
		[]byte(`{"test":"test"}`),
		"",
	}

	stringSig, err := sig.Create("superdupersecret")

	if err != nil {
		t.Fail()
	} else if stringSig != "zKw8juIJCzv9GZ4DPh5b+DZl0MHxmpeSBpcM838+46Q=" {
		t.Fail()
	}
}

func TestSuccessfulCompareHmacBase64Signature(t *testing.T) {

	sig := &HmacSignature{
		"GET",
		"/api/test",
		"application/json",
		time.UTC.String(),
		[]byte(`{"test":"test"}`),
		"",
	}

	if _, err := sig.Create("superdupersecret"); err != nil {
		t.Fail()
	}

	succ, err := sig.Verify("zKw8juIJCzv9GZ4DPh5b+DZl0MHxmpeSBpcM838+46Q=")

	if err != nil {
		t.Fail()
	}

	if !succ {
		t.Fail()
	}
}

func TestFailingCompareHmacBase64Signature(t *testing.T) {

	sig := &HmacSignature{
		"GET",
		"/api/test",
		"application/json",
		time.UTC.String(),
		[]byte(`{"test":"test"}`),
		"",
	}

	if _, err := sig.Create("superdupersecret"); err != nil {
		t.Fail()
	}

	succ, err := sig.Verify("SGFsbG9JY2hCaW5FaW5LcmFzcw==")

	if err != nil {
		t.Fail()
	}

	if succ {
		t.Fail()
	}
}
