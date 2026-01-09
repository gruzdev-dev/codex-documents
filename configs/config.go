package configs

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		Port string `yaml:"port"`
	} `yaml:"server"`
	GRPC struct {
		Port string `yaml:"port"`
	} `yaml:"grpc"`
	Auth struct {
		JWTSecret      string `yaml:"jwt_secret"`
		InternalSecret string `yaml:"internal_secret"`
	} `yaml:"auth"`
	MongoDB struct {
		URI      string `yaml:"uri"`
		Database string `yaml:"database"`
	} `yaml:"mongodb"`
}

func NewConfig() (*Config, error) {
	var cfg Config

	configFile, err := os.ReadFile("config.yaml")
	if err == nil {
		if err := yaml.Unmarshal(configFile, &cfg); err != nil {
			return nil, err
		}
	}

	if envPort := os.Getenv("SERVER_PORT"); envPort != "" {
		cfg.Server.Port = envPort
	}
	if envGRPCPort := os.Getenv("GRPC_PORT"); envGRPCPort != "" {
		cfg.GRPC.Port = envGRPCPort
	}
	if envSecret := os.Getenv("JWT_SECRET"); envSecret != "" {
		cfg.Auth.JWTSecret = envSecret
	}
	if envInternalSecret := os.Getenv("INTERNAL_SERVICE_SECRET"); envInternalSecret != "" {
		cfg.Auth.InternalSecret = envInternalSecret
	}
	if envMongoURI := os.Getenv("MONGO_URI"); envMongoURI != "" {
		cfg.MongoDB.URI = envMongoURI
	}
	if envMongoDB := os.Getenv("MONGO_DATABASE"); envMongoDB != "" {
		cfg.MongoDB.Database = envMongoDB
	}

	return &cfg, nil
}
