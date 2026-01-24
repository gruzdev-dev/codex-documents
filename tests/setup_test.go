//go:build integration

package tests

import (
	"context"
	"log"
	"net"
	"testing"
	"time"

	httpadapter "github.com/gruzdev-dev/codex-documents/adapters/http"
	mongostorage "github.com/gruzdev-dev/codex-documents/adapters/storage/mongodb"
	"github.com/gruzdev-dev/codex-documents/configs"
	"github.com/gruzdev-dev/codex-documents/core/ports"
	"github.com/gruzdev-dev/codex-documents/core/services"
	"github.com/gruzdev-dev/codex-documents/core/validator"
	"github.com/gruzdev-dev/codex-documents/pkg/database"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go/modules/mongodb"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.uber.org/dig"
	"go.uber.org/mock/gomock"

	"net/http/httptest"

	grpcadapter "github.com/gruzdev-dev/codex-documents/adapters/grpc"
	"github.com/gruzdev-dev/codex-documents/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

type TestEnv struct {
	Container        *dig.Container
	DB               *mongo.Database
	GRPCClient       proto.AuthIntegrationClient
	ServerURL        string
	MockFileProvider *ports.MockFileProvider
	Cleanup          func()
}

func SetupTestEnv(t *testing.T) *TestEnv {
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	mockFileProvider := ports.NewMockFileProvider(ctrl)

	mongoContainer, err := mongodb.Run(ctx, "mongo:7.0",
		mongodb.WithUsername("testusername"),
		mongodb.WithPassword("testpassword"))
	require.NoError(t, err)

	cfg := initConfig(t, ctx, mongoContainer)

	c, err := buildTestContainer(cfg, mockFileProvider)
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
			log.Printf("Server exited with error: %v", err)
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
	ts := httptest.NewServer(router)

	return &TestEnv{
		Container:        c,
		ServerURL:        ts.URL,
		DB:               db,
		GRPCClient:       grpcClient,
		MockFileProvider: mockFileProvider,
		Cleanup: func() {
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			ctrl.Finish()
			s.GracefulStop()
			_ = conn.Close()
			ts.Close()
			_ = mongoContainer.Terminate(ctx)
		},
	}
}

func buildTestContainer(cfg *configs.Config, mockFileProvider *ports.MockFileProvider) (*dig.Container, error) {
	c := dig.New()

	if err := c.Provide(func() *configs.Config {
		return cfg
	}); err != nil {
		return nil, err
	}

	if err := c.Provide(database.NewMongoDB); err != nil {
		return nil, err
	}

	if err := c.Provide(mongostorage.NewPatientRepo, dig.As(new(ports.PatientRepository))); err != nil {
		return nil, err
	}

	if err := c.Provide(validator.NewPatientValidator); err != nil {
		return nil, err
	}

	if err := c.Provide(services.NewPatientService, dig.As(new(ports.PatientService))); err != nil {
		return nil, err
	}

	if err := c.Provide(mongostorage.NewDocumentRepo, dig.As(new(ports.DocumentRepository))); err != nil {
		return nil, err
	}

	if err := c.Provide(validator.NewDocumentValidator); err != nil {
		return nil, err
	}

	if err := c.Provide(func() ports.FileProvider {
		return mockFileProvider
	}); err != nil {
		return nil, err
	}

	if err := c.Provide(services.NewDocumentService, dig.As(new(ports.DocumentService))); err != nil {
		return nil, err
	}

	if err := c.Provide(mongostorage.NewObservationRepo, dig.As(new(ports.ObservationRepository))); err != nil {
		return nil, err
	}

	if err := c.Provide(validator.NewObservationValidator); err != nil {
		return nil, err
	}

	if err := c.Provide(services.NewObservationService, dig.As(new(ports.ObservationService))); err != nil {
		return nil, err
	}

	if err := c.Provide(httpadapter.NewHandler); err != nil {
		return nil, err
	}

	if err := c.Provide(grpcadapter.NewAuthHandler); err != nil {
		return nil, err
	}

	return c, nil
}

func initConfig(t *testing.T, ctx context.Context, mongoContainer *mongodb.MongoDBContainer) *configs.Config {
	host, err := mongoContainer.Host(ctx)
	if err != nil {
		t.Fatalf("failed to get host: %v", err)
	}

	port, err := mongoContainer.MappedPort(ctx, "27017")
	if err != nil {
		t.Fatalf("failed to get mapped port: %v", err)
	}

	cfg := &configs.Config{}
	cfg.HTTP.Port = "8080"
	cfg.GRPC.Port = "8081"
	cfg.Auth.JWTSecret = "secret-key"
	cfg.Auth.InternalSecret = "test-secret"
	cfg.MongoDB.Host = host
	cfg.MongoDB.Port = port.Port()
	cfg.MongoDB.Username = "testusername"
	cfg.MongoDB.Password = "testpassword"
	cfg.MongoDB.Database = "test_db"
	cfg.MongoDB.AuthSource = "admin"
	cfg.FileService.Addr = ""

	return cfg
}
