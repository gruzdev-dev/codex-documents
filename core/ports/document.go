package ports

import (
	"codex-documents/core/domain"
	"context"
	models "github.com/gruzdev-dev/fhir/r5"
)

type DocumentRepository interface {
	Create(ctx context.Context, doc *models.DocumentReference) error
	GetByID(ctx context.Context, id string) (*models.DocumentReference, error)
	SearchByPatient(ctx context.Context, patientID string) ([]*models.DocumentReference, error)

	CreateObservation(ctx context.Context, obs *models.Observation) error
	SearchObservations(ctx context.Context, patientID string, code string) ([]*models.Observation, error)
}

type AccessRepository interface {
	SaveShare(ctx context.Context, share *domain.SharedAccess) error
	GetShareByToken(ctx context.Context, token string) (*domain.SharedAccess, error)
}
