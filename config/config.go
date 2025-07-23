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
	confFile, err := os.Open("./config.json")
	log.Print("Loading config file")
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
