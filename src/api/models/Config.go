package models

import (
	"encoding/json"
	"log"
	"os"
)

// Config structure for the configuration file
type Config struct {
	Database Database `json:"db"`
	Service  Service  `json:"service"`
}

// Database structure configuration file format
type Database struct {
	Type         string `json:"type"`
	Host         string `json:"host"`
	Port         int    `json:"port"`
	User         string `json:"user"`
	DatabaseName string `json:"database"`
	Password     string `json:"password"`
}

// Service structure configuration file format
type Service struct {
	Port     int    `json:"port"`
	LogLevel string `json:"loglevel"`
}

// Load config of service
func (config *Config) Load(path string) (*Config, error) {
	// config := Config{}
	file, err := os.Open(path)
	defer file.Close()

	if err != nil {
		log.Fatalf("Error getting opening configuration file. Error:%v", err)
	} else {
		decoder := json.NewDecoder(file)
		err := decoder.Decode(&config)
		if err != nil {
			log.Fatalf("Error loading opening configuration file. Error:%v", err)
		}
	}

	return config, nil
}
