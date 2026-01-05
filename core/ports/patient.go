package ports

import (
	"codex-documents/fhir/models"
	"context"
)

type PatientRepository interface {
	Create(ctx context.Context, patient *models.Patient) error
	GetByID(ctx context.Context, id string) (*models.Patient, error)
	Update(ctx context.Context, patient *models.Patient) error
}
