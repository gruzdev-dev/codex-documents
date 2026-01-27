package ports

import (
	"context"

	"github.com/gruzdev-dev/codex-documents/core/domain"
	models "github.com/gruzdev-dev/fhir/r5"
)

//go:generate mockgen -source=observation.go -destination=observation_mocks.go -package=ports ObservationRepository,ObservationService

type ObservationRepository interface {
	Create(ctx context.Context, obs *models.Observation) (*models.Observation, error)
	GetByID(ctx context.Context, id string) (*models.Observation, error)
	GetByIDs(ctx context.Context, ids []string) ([]models.Observation, error)
	Update(ctx context.Context, obs *models.Observation) (*models.Observation, error)
	Delete(ctx context.Context, id string) error
	Search(ctx context.Context, patientID string, limit, offset int) ([]models.Observation, int64, error)
}

type ObservationService interface {
	Create(ctx context.Context, obs *models.Observation) (*models.Observation, error)
	Get(ctx context.Context, id string) (*models.Observation, error)
	Update(ctx context.Context, obs *models.Observation) (*models.Observation, error)
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, patientID string, limit, offset int) (*domain.ListResponse[models.Observation], error)
}
