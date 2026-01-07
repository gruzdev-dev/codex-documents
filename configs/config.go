package configs

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		Port string `yaml:"port"`
	} `yaml:"server"`
	Auth struct {
		JWTSecret string `yaml:"jwt_secret"`
	} `yaml:"auth"`
	MongoDB struct {
		URI        string `yaml:"uri"`
		Database   string `yaml:"database"`
	} `yaml:"mongodb"`
}

func NewConfig() (*Config, error) {
	configFile, err := os.ReadFile("config.yaml")
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal(configFile, &cfg); err != nil {
		return nil, err
	}

	if envPort := os.Getenv("SERVER_PORT"); envPort != "" {
		cfg.Server.Port = envPort
	}
	if envSecret := os.Getenv("JWT_SECRET"); envSecret != "" {
		cfg.Auth.JWTSecret = envSecret
	}
	if envMongoURI := os.Getenv("MONGO_URI"); envMongoURI != "" {
		cfg.MongoDB.URI = envMongoURI
	}
	if envMongoDB := os.Getenv("MONGO_DATABASE"); envMongoDB != "" {
		cfg.MongoDB.Database = envMongoDB
	}

	return &cfg, nil
}
