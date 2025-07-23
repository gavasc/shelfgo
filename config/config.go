package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	AllowedFormats []string `json:"allowed_formats"`
}

// opens the config file and decodes it, returning the config if no error occurs
func Load() (Config, error) {
	log.Print("Loading config file")
	confFile, err := os.Open("config/config.json")
	if err != nil {
		return Config{}, err
	}

	config := Config{}

	decoder := json.NewDecoder(confFile)
	if err := decoder.Decode(&config); err != nil {
		return Config{}, err
	}

	return config, nil
}
