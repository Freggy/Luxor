package rest

import (
	"encoding/json"
	"log"
	"testing"
)

func TestErrorToJsonString(t *testing.T) {
	myErr := NewError(400, "My Error Message", "MyErrorType")
	jsonErr, err := myErr.ToJson()

	if err != nil {
		log.Println("Error while marshalling")
		t.Fail()
	}

	if string(jsonErr) != `{"code":400,"message":"My Error Message","type":"MyErrorType"}` {
		t.Fail()
	}
}

func TestErrorUnmarshalling(t *testing.T) {
	myErr := NewError(400, "My Error Message", "MyErrorType")
	data, err := myErr.ToJson()

	if err != nil {
		log.Println("Error while marshalling")
		t.Fail()
	}

	var otherErr Error

	if err := json.Unmarshal(data, &otherErr); err != nil {
		log.Println("Error while unmarshalling")
		t.Fail()
	}

	if otherErr.Code != myErr.Code {
		t.Fail()
	} else if otherErr.ErrorType != myErr.ErrorType {
		t.Fail()
	} else if otherErr.Message != myErr.Message {
		t.Fail()
	}
}
