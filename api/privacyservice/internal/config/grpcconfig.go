package config

import (
	"encoding/json"
	"net"
	"os"
)

type GRPCConfig interface {
	Address() string
	ReadConfig(string) error
}

type grpcConfig struct {
	Port string `json:"port"`
	Host string `json:"host"`
}

func NewGRPCConfig() *grpcConfig {
	return &grpcConfig{}
}

func (c *grpcConfig) ReadConfig(path string) error {
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

func (g *grpcConfig) Address() string {
	return net.JoinHostPort(g.Host, g.Port)
}
