package main

import (
	httpAdapter "codex-documents/adapters/http"
	storageAdapter "codex-documents/adapters/storage"
	"codex-documents/configs"
	"codex-documents/core/ports"
	"codex-documents/core/service"
	httpServer "codex-documents/servers/http"

	"go.uber.org/dig"
)

func BuildContainer() (*dig.Container, error) {
	container := dig.New()

	if err := container.Provide(configs.NewConfig); err != nil {
		return nil, err
	}
	if err := container.Provide(storageAdapter.NewInMemoryRepo); err != nil {
		return nil, err
	}
	if err := container.Provide(service.NewUserService, dig.As(new(ports.UserService))); err != nil {
		return nil, err
	}
	if err := container.Provide(httpAdapter.NewHandler); err != nil {
		return nil, err
	}
	if err := container.Provide(httpServer.NewServer); err != nil {
		return nil, err
	}

	return container, nil
}
