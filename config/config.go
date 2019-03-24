package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Token string `json:"token"`
	Host  string `json:"host"`
}

func NewConfig(path string) (*Config, error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	cfg := new(Config)
	err = json.Unmarshal(file, cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
