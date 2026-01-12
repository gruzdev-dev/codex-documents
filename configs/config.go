package configs

import (
	"os"
)

type Config struct {
	HTTP struct {
		Port string
	}
	GRPC struct {
		Port string
	}
	Auth struct {
		JWTSecret      string
		InternalSecret string
	}
	MongoDB struct {
		Host       string
		Port       string
		Username   string
		Password   string
		Database   string
		AuthSource string
	}
}

func NewConfig() (*Config, error) {
	var cfg Config

	if envPort := os.Getenv("HTTP_PORT"); envPort != "" {
		cfg.HTTP.Port = envPort
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
	if envMongoHost := os.Getenv("MONGO_HOST"); envMongoHost != "" {
		cfg.MongoDB.Host = envMongoHost
	}
	if envMongoPort := os.Getenv("MONGO_PORT"); envMongoPort != "" {
		cfg.MongoDB.Port = envMongoPort
	}
	if envMongoUsername := os.Getenv("MONGO_USERNAME"); envMongoUsername != "" {
		cfg.MongoDB.Username = envMongoUsername
	}
	if envMongoPassword := os.Getenv("MONGO_PASSWORD"); envMongoPassword != "" {
		cfg.MongoDB.Password = envMongoPassword
	}
	if envMongoDB := os.Getenv("MONGO_DATABASE"); envMongoDB != "" {
		cfg.MongoDB.Database = envMongoDB
	}
	if envMongoAuthSource := os.Getenv("MONGO_AUTH_SOURCE"); envMongoAuthSource != "" {
		cfg.MongoDB.AuthSource = envMongoAuthSource
	}

	return &cfg, nil
}
