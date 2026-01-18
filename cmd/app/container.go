package main

import (
	grpcServer "github.com/gruzdev-dev/codex-documents/servers/grpc"
	httpServer "github.com/gruzdev-dev/codex-documents/servers/http"
	"go.uber.org/dig"

	"github.com/gruzdev-dev/codex-documents/adapters/clients/files"
	"github.com/gruzdev-dev/codex-documents/adapters/grpc"
	"github.com/gruzdev-dev/codex-documents/adapters/http"
	"github.com/gruzdev-dev/codex-documents/adapters/storage/mongodb"
	"github.com/gruzdev-dev/codex-documents/configs"
	"github.com/gruzdev-dev/codex-documents/core/ports"
	"github.com/gruzdev-dev/codex-documents/core/services"
	"github.com/gruzdev-dev/codex-documents/core/validator"
	"github.com/gruzdev-dev/codex-documents/pkg/database"
)

func BuildContainer() (*dig.Container, error) {
	c := dig.New()

	if err := c.Provide(configs.NewConfig); err != nil {
		return nil, err
	}

	if err := c.Provide(database.NewMongoDB); err != nil {
		return nil, err
	}

	if err := c.Provide(mongodb.NewPatientRepo, dig.As(new(ports.PatientRepository))); err != nil {
		return nil, err
	}

	if err := c.Provide(validator.NewPatientValidator); err != nil {
		return nil, err
	}

	if err := c.Provide(services.NewPatientService, dig.As(new(ports.PatientService))); err != nil {
		return nil, err
	}

	if err := c.Provide(mongodb.NewDocumentRepo, dig.As(new(ports.DocumentRepository))); err != nil {
		return nil, err
	}

	if err := c.Provide(validator.NewDocumentValidator); err != nil {
		return nil, err
	}

	if err := c.Provide(files.NewClient, dig.As(new(ports.FileProvider))); err != nil {
		return nil, err
	}

	if err := c.Provide(services.NewDocumentService, dig.As(new(ports.DocumentService))); err != nil {
		return nil, err
	}

	if err := c.Provide(http.NewHandler); err != nil {
		return nil, err
	}

	if err := c.Provide(grpc.NewAuthHandler); err != nil {
		return nil, err
	}

	if err := c.Provide(httpServer.NewServer); err != nil {
		return nil, err
	}

	if err := c.Provide(grpcServer.NewServer); err != nil {
		return nil, err
	}

	return c, nil
}
