package services

import (
	"context"
	"fmt"

	"github.com/gruzdev-dev/codex-documents/core/domain"
	"github.com/gruzdev-dev/codex-documents/core/ports"
	"github.com/gruzdev-dev/codex-documents/core/validator"
	"github.com/gruzdev-dev/codex-documents/pkg/identity"

	"github.com/google/uuid"
	models "github.com/gruzdev-dev/fhir/r5"
)

type PatientService struct {
	repo      ports.PatientRepository
	validator *validator.PatientValidator
}

func NewPatientService(repo ports.PatientRepository, v *validator.PatientValidator) *PatientService {
	return &PatientService{
		repo:      repo,
		validator: v,
	}
}

func (s *PatientService) Create(ctx context.Context, patient *models.Patient) (*models.Patient, error) {
	if err := s.validator.Validate(patient); err != nil {
		return nil, fmt.Errorf("%w: %v", domain.ErrInvalidInput, err)
	}

	if patient.Id != nil && *patient.Id != "" {
		return nil, fmt.Errorf("%w: patient ID must not be provided during creation", domain.ErrInvalidInput)
	}

	if patient.Id == nil {
		id := uuid.New().String()
		patient.Id = &id
	}

	created, err := s.repo.Create(ctx, patient)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", domain.ErrInternal, err)
	}

	return created, nil
}

func (s *PatientService) Get(ctx context.Context, id string) (*models.Patient, error) {
	user, ok := identity.FromCtx(ctx)
	if !ok {
		return nil, domain.ErrAccessDenied
	}

	if !user.IsPatient(id) && !user.HasScope("patient/*.read") {
		return nil, domain.ErrAccessDenied
	}

	patient, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", domain.ErrInternal, err)
	}

	if patient == nil {
		return nil, domain.ErrPatientNotFound
	}

	return patient, nil
}

func (s *PatientService) Update(ctx context.Context, patient *models.Patient) (*models.Patient, error) {
	user, ok := identity.FromCtx(ctx)
	if !ok {
		return nil, domain.ErrAccessDenied
	}

	if patient.Id == nil {
		return nil, domain.ErrPatientIDRequired
	}

	if !user.IsPatient(*patient.Id) && !user.HasScope("patient/*.write") {
		return nil, domain.ErrAccessDenied
	}

	if err := s.validator.Validate(patient); err != nil {
		return nil, fmt.Errorf("%w: %v", domain.ErrInvalidInput, err)
	}

	existing, err := s.repo.GetByID(ctx, *patient.Id)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", domain.ErrInternal, err)
	}

	if existing == nil {
		return nil, domain.ErrPatientNotFound
	}

	updated, err := s.repo.Update(ctx, patient)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", domain.ErrInternal, err)
	}

	return updated, nil
}
