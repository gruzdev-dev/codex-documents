package main

import (
	"context"
	"golang.org/x/sync/errgroup"
	"log"
	"os/signal"
	"syscall"

	"codex-documents/adapters/grpc"
	"codex-documents/api/proto"
	grpcServer "codex-documents/servers/grpc"
	httpServer "codex-documents/servers/http"
)

func main() {
	container, err := BuildContainer()
	if err != nil {
		log.Fatalf("Fatal error building container: %v", err)
	}

	err = container.Invoke(func(
		httpSrv *httpServer.Server,
		grpcSrv *grpcServer.Server,
		authHandler *grpc.AuthHandler,
	) error {
		proto.RegisterAuthIntegrationServer(grpcSrv.GetGRPCServer(), authHandler)

		ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
		defer stop()

		g, ctx := errgroup.WithContext(ctx)

		g.Go(func() error {
			return httpSrv.Start()
		})

		g.Go(func() error {
			return grpcSrv.Start(ctx)
		})

		return g.Wait()
	})

	if err != nil {
		log.Fatalf("Application stopped with error: %v", err)
	}
}
