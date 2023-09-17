package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type DatabaseConfig struct {
	Driver   string `yaml:"driver"`
	Host     string `yaml:"host"`
	Database string `yaml:"database"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type Config struct {
	Database DatabaseConfig `yaml:"database"`
}

func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	conf := &Config{}
	if err := yaml.Unmarshal(data, conf); err != nil {
		return nil, err
	}

	return conf, nil
}
