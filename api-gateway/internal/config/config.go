package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Route struct {
	Path   string `yaml:"path"`
	Method string `yaml:"method"`
	Target string `yaml:"target"`
}

type Service struct {
	BaseURL string  `yaml:"base_url"`
	Routes  []Route `yaml:"routes"`
}

type Config struct {
	Services map[string]Service `yaml:"services"`
}

func LoadConfig(filename string) (*Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
