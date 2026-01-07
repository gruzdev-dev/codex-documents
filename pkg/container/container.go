package container

import (
	"codex-documents/adapters/http"
	"codex-documents/adapters/storage/mongodb"
	"codex-documents/configs"
	"codex-documents/core/ports"
	"codex-documents/core/services"
	"codex-documents/core/validator"
	"codex-documents/pkg/database"

	"go.uber.org/dig"
)

func BuildAppContainer() (*dig.Container, error) {
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

	if err := c.Provide(http.NewHandler); err != nil {
		return nil, err
	}

	return c, nil
}
