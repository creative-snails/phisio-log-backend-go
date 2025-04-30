package config

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type ServerConfig struct {
	Port int	 `yaml:"port"`
	Host string  `yaml:"host"`
}

type DatabaseConfig struct {
	Port int        `yaml:"port"`
	Host string     `yaml:"host"`
	User string     `yaml:"user"`
	Password string `yaml:"password"`
	Dbname string   `yaml:"dbname"`
	Sslmode string  `yaml:"sslmode"`
}

type Config struct {
	Server ServerConfig `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
}

func LoadConfig(configPath string) (*Config, error) {
	err := godotenv.Load()
	if err != nil {
	return nil, fmt.Errorf("error loading .env file: %w", err)
	}

	viper.SetConfigFile(configPath)
	viper.AutomaticEnv()

	viper.BindEnv("server.port", "SERVER_PORT")
	viper.BindEnv("server.host", "SERVER_HOST")
	viper.BindEnv("database.port", "DB_PORT")
	viper.BindEnv("database.host", "DB_HOST")
	viper.BindEnv("database.user", "DB_USER")
	viper.BindEnv("database.password", "DB_PASSWORD")
	viper.BindEnv("database.dbname", "DB_NAME")
	viper.BindEnv("database.sslmode", "DB_SSLMODE")

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	config := &Config{}
	if err := viper.Unmarshal(config); err != nil {
		return nil, fmt.Errorf("unable to decode into struct: %w", err)
	}

	return config, nil
}