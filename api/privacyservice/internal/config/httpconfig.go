package config

import (
	"encoding/json"
	"net"
	"os"
)

type HTTPConfig interface {
	Address() string
	ReadConfig(string) error
}

type httpConfig struct {
	Port string `json:"port"`
	Host string `json:"host"`
}

func NewHttpConfig() *httpConfig {
	return &httpConfig{}
}

func (c *httpConfig) ReadConfig(path string) error {
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

func (h *httpConfig) Address() string {
	return net.JoinHostPort(h.Host, h.Port)
}
