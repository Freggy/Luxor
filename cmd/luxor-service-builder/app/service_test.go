package app

import (
	"testing"
)

func Test_NewServiceFromJson(t *testing.T) {
	data := `
	{
		"id": "TestID", 
    	"files":[
    		{
				"path": "/path/to/file"
			},
			{
				"path": "/path/to/dir",
				"files": [
					{
						"path": "/path/to/file"
					},
					{
						"path": "/path/to/file"
					}
				]
			}
		]
	}`

	s, err := NewServiceFromJson([]byte(data))

	if err != nil {
		t.Error("Could not unmarshal JSON")
	}

	if s.Id != "TestID" {
		t.Errorf("Id should have been %s but was %s", "TestID", s.Id)
	}

	if len(s.Files) < 2 {
		t.Errorf("Exptected %d entries, but found %d", 2, len(s.Files))
	}

	if len(s.Files[1].Files) < 2 {
		t.Errorf("Exptected %d entries, but found %d", 2, len(s.Files))
	}
}
