package app

import (
	"github.com/luxordynamics/luxor/pkg/logging"
	"log"
	"strconv"
	"testing"
)


func TestFromString_Success(t *testing.T) {
	data :=`
		{ 
			"httpsEnabled": false, 
			"logType": 1
		}`

	config, err := fromString(data)

	if err != nil {
		t.Error(err)
	}

	if config.LogLevel != logging.Debug {
		log.Println("LogLevel should be DEBUG(0) but was " + strconv.Itoa(int(config.LogLevel)))
		t.Fail()
	}
}
