package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Port           string `json:"port"`
	Driver         string `json:"database_driver"`
	DatabaseURL    string `json:"database_url"`
	Migrations     string `json:"database_migrations"`
	ChatServiceURL string `json:"chatservice_url"`
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) ReadConfig(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, c)
	if err != nil {
		return err
	}

	return nil
}
