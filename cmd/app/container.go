package main

import (
	"codex-documents/adapters/http"
	"codex-documents/adapters/storage/mongodb"
	"codex-documents/configs"
	"codex-documents/pkg/database"
	httpServer "codex-documents/servers/http"

	"go.uber.org/dig"
)

func BuildContainer() (*dig.Container, error) {
	container := dig.New()

	if err := container.Provide(configs.NewConfig); err != nil {
		return nil, err
	}

	if err := container.Provide(database.NewMongoDB); err != nil {
		return nil, err
	}

	if err := container.Provide(mongodb.NewPatientRepo); err != nil {
		return nil, err
	}

	if err := container.Provide(http.NewHandler); err != nil {
		return nil, err
	}

	if err := container.Provide(httpServer.NewServer); err != nil {
		return nil, err
	}

	return container, nil
}
