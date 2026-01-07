package ports

import (
	"context"
	models "github.com/gruzdev-dev/fhir/r5"
)

type PatientService interface {
	Create(ctx context.Context, patient *models.Patient) error
	Get(ctx context.Context, id string) (*models.Patient, error)
	Update(ctx context.Context, patient *models.Patient) error
}

type PatientRepository interface {
	Create(ctx context.Context, patient *models.Patient) error
	GetByID(ctx context.Context, id string) (*models.Patient, error)
	Update(ctx context.Context, patient *models.Patient) error
}
