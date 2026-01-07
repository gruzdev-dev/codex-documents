package tests

import (
	"context"
	nethttp "net/http"
	"os"
	"testing"
	"time"

	httpadapter "codex-documents/adapters/http"
	"codex-documents/pkg/container"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go/modules/mongodb"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.uber.org/dig"
)

type TestEnv struct {
	Container *dig.Container
	Router    *mux.Router
	Handler   nethttp.Handler
	DB        *mongo.Database
	Cleanup   func()
}

func SetupTestEnv(t *testing.T) *TestEnv {
	ctx := context.Background()

	mongoContainer, err := mongodb.Run(ctx, "mongo:7.0")
	require.NoError(t, err)

	uri, err := mongoContainer.ConnectionString(ctx)
	require.NoError(t, err)

	os.Setenv("MONGO_URI", uri)
	os.Setenv("MONGO_DATABASE", "test_db")
	os.Setenv("JWT_SECRET", "secret-key")

	c, err := container.BuildAppContainer()
	require.NoError(t, err)

	var handler *httpadapter.Handler
	var db *mongo.Database
	err = c.Invoke(func(h *httpadapter.Handler, database *mongo.Database) {
		handler = h
		db = database
	})
	require.NoError(t, err)

	router := mux.NewRouter()
	handler.RegisterRoutes(router)

	cleanup := func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if db != nil {
			if client := db.Client(); client != nil {
				_ = client.Disconnect(ctx)
			}
		}

		_ = mongoContainer.Terminate(ctx)
	}

	return &TestEnv{
		Container: c,
		Router:    router,
		Handler:   router,
		DB:        db,
		Cleanup:   cleanup,
	}
}
