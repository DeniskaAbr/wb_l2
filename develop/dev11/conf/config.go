package conf

import (
	"encoding/json"
	"net/http"
	"os"
)

type Configuration struct {
	Host   string       `json:"host"`
	Port   string       `json:"port"`
	Server *http.Server `json:"-"`
}

func (c *Configuration) ReadFromFile(file string) (Configuration, error) {
	cfg := Configuration{}

	jsonFile, err := os.ReadFile(file)
	if err != nil {
		return cfg, err
	}

	err = json.Unmarshal(jsonFile, &cfg)
	if err != nil {
		return cfg, err
	}

	return cfg, nil
}
