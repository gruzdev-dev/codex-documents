package database

import (
	"context"
	"fmt"
	"time"

	"codex-documents/configs"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const (
	connectTimeout = 10 * time.Second
)

func NewMongoDB(cfg *configs.Config) (*mongo.Database, error) {
	uri := buildMongoURI(cfg)

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().
		ApplyURI(uri).
		SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(opts)
	if err != nil {
		return nil, fmt.Errorf("mongo connect error: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout)
	defer cancel()

	if err := client.Database("admin").RunCommand(ctx, bson.D{{Key: "ping", Value: 1}}).Err(); err != nil {
		return nil, fmt.Errorf("mongo ping error: %w", err)
	}

	return client.Database(cfg.MongoDB.Database), nil
}

func buildMongoURI(cfg *configs.Config) string {
	if cfg.MongoDB.Username != "" && cfg.MongoDB.Password != "" {
		authSource := cfg.MongoDB.AuthSource
		if authSource == "" {
			authSource = cfg.MongoDB.Database
		}
		return fmt.Sprintf("mongodb://%s:%s@%s:%s/%s?authSource=%s",
			cfg.MongoDB.Username,
			cfg.MongoDB.Password,
			cfg.MongoDB.Host,
			cfg.MongoDB.Port,
			cfg.MongoDB.Database,
			authSource,
		)
	}
	return fmt.Sprintf("mongodb://%s:%s/%s",
		cfg.MongoDB.Host,
		cfg.MongoDB.Port,
		cfg.MongoDB.Database,
	)
}
