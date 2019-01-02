package app

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io/ioutil"
)

type Config struct {
	HttpsEnabled bool `json:"httpsEnabled"`
	logger       *logrus.Logger
}

// FromFile reads the the file at the given path. The content needs to be in JSON format.
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
