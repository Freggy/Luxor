package app

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestBuildModel(t *testing.T) {
	data := `
	{
		"id": "TestID", 
    	"files":[
    		{
				"path": "/home/yannic/test/lol"
			},
			{
				"path": "/home/yannic/test/rofl"
			},
			{
				"path": "/home/yannic/test/sub"
			}
		]
	}`

	s, err := NewServiceFromJson([]byte(data))

	if err != nil {
		t.Error("Could not unmarshal JSON")
	}

	m, err := BuildModel(s)

	if err != nil {
		t.Error("Could not build model")
	}

	b, err := json.Marshal(m)

	if err != nil {
		t.Error("Could not marshal JSON")
	}

	fmt.Println(string(b))
}
