package app

import (
	"github.com/luxordynamics/luxor/pkg/logging"
	"log"
	"testing"
)


func TestFromString_Success(t *testing.T) {
	data :=`
		{ 
			"httpsEnabled": false, 
			"logType": "DEBUG"
		}`

	config, err := fromString(data)

	if err != nil {
		t.Error(err)
	}

	if config.LogType != logging.Debug {
		log.Println("LogLevel should be DEBUG but was " + config.LogType)
		t.Fail()
	}
}
