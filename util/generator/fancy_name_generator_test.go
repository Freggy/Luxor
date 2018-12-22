package generator

import (
	"encoding/json"
	"testing"
)

func TestMarshallAssets(t *testing.T) {
	var (
		nouns      NounList
		adjectives AdjectiveList
	)
	if err := json.Unmarshal([]byte(Nouns), &nouns); err != nil {
		t.Fail()
	}

	if err := json.Unmarshal([]byte(Adjectives), &adjectives); err != nil {
		t.Fail()
	}
}

func TestGenerateName(t *testing.T) {
	name := GenerateName("TestMessage")

	if name != "der_provisorische_Programmcode" {
		t.Fail()
	}

	name = GenerateName("TestMessage")

	// Test twice because the output should be always the same
	if name != "der_provisorische_Programmcode" {
		t.Fail()
	}
}
