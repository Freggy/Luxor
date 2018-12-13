package app

import (
	"encoding/json"
	"github.com/luxordynamics/luxor/pkg/logging"
	"io/ioutil"
)

type Config struct {
	HttpsEnabled bool            `json:"httpsEnabled"`
	LogLevel     logging.LogType `json:"logType"`
}

func FromFile(path string) (*Config, error) {
	content, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, err
	}

	conf, err := fromString(string(content))

	if err != nil {
		return nil, err
	}

	return conf, nil
}

func fromString(content string) (*Config, error) {
	conf := Config{}

	if err := json.Unmarshal([]byte(content), &conf); err != nil {
		return nil, err
	}
	return &conf, nil
}
