package config

import (
	"encoding/json"
	"os"
)

type GeneralConfig struct {
	Driver            string `json:"database_driver"`
	DatabaseURL       string `json:"database_url"`
	Migrations        string `json:"database_migrations"`
	PrivacyClientGRPC string `json:"privacyservice_grpc"`
	PrivacyClientHTTP string `json:"privacyservice_http"`
}

func NewConfig() *GeneralConfig {
	return &GeneralConfig{}
}

func (c *GeneralConfig) ReadConfig(path string) error {
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
