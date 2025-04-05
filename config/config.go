package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		Port int	 `yaml:"port"`
		Host string  `yaml:"host"`
	} `yaml:"server"`

	Database struct {
		Port int        `yaml:"port"`
		Host string     `yaml:"host"`
		User string     `yaml:"user"`
		Password string `yaml:"password"`
		Dbname string   `yaml:"dbname"`
		Sslmode string  `yaml:"sslmode"`
	} `yaml:"database"`
}

func LoadConfig(configPath string) (*Config, error) {
	config := &Config{}

	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(config); err != nil {
		return nil, err
	}

	return config, nil
}