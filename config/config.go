package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	APIKey string `json:"GH_TOKEN"`
}

func ReadAPIKey() (Config, error) {
	file, err := os.ReadFile("config.json")
	if err != nil {
		return Config{}, fmt.Errorf("error reading config.json")
	}
	c := Config{}
	err = json.Unmarshal(file, &c)
	if err != nil {
		return Config{}, fmt.Errorf("error unmarshalling the file")
	}
	return c, nil
}
