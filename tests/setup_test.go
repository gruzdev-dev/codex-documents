package tests

import (
	"context"
	"net"
	"os"
	"testing"

	httpadapter "codex-documents/adapters/http"
	"codex-documents/pkg/container"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go/modules/mongodb"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.uber.org/dig"

	grpcadapter "codex-documents/adapters/grpc"
	"codex-documents/api/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
	"net/http/httptest"
)

type TestEnv struct {
	Container  *dig.Container
	DB         *mongo.Database
	GRPCClient proto.AuthIntegrationClient
	ServerURL  string
	Cleanup    func()
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
	os.Setenv("INTERNAL_SERVICE_SECRET", "test-secret")

	c, err := container.BuildAppContainer()
	require.NoError(t, err)

	var httpHandler *httpadapter.Handler
	var authHandler *grpcadapter.AuthHandler
	var db *mongo.Database
	err = c.Invoke(func(h *httpadapter.Handler, ah *grpcadapter.AuthHandler, database *mongo.Database) {
		httpHandler = h
		authHandler = ah
		db = database
	})
	require.NoError(t, err)

	const bufSize = 1024 * 1024
	lis := bufconn.Listen(bufSize)

	s := grpc.NewServer()
	proto.RegisterAuthIntegrationServer(s, authHandler)

	go func() {
		if err := s.Serve(lis); err != nil {
			return
		}
	}()

	conn, err := grpc.NewClient("passthrough://bufnet",
		grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
			return lis.Dial()
		}),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	require.NoError(t, err)
	grpcClient := proto.NewAuthIntegrationClient(conn)

	router := mux.NewRouter()
	httpHandler.RegisterRoutes(router)

	cleanup := func() {
		s.Stop()
		_ = conn.Close()
		_ = mongoContainer.Terminate(context.Background())
	}

	ts := httptest.NewServer(router)

	return &TestEnv{
		Container:  c,
		ServerURL:  ts.URL,
		DB:         db,
		GRPCClient: grpcClient,
		Cleanup:    cleanup,
	}
}
