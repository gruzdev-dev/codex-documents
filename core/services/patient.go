package services

import (
	"context"
	"fmt"

	"codex-documents/core/domain"
	"codex-documents/core/ports"
	"codex-documents/core/validator"
	"codex-documents/pkg/identity"
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

func (s *PatientService) Create(ctx context.Context, patient *models.Patient) error {
	if err := s.validator.Validate(patient); err != nil {
		return fmt.Errorf("%w: %v", domain.ErrInvalidInput, err)
	}

	if err := s.repo.Create(ctx, patient); err != nil {
		return fmt.Errorf("%w: %v", domain.ErrInternal, err)
	}

	return nil
}

func (s *PatientService) GetPatient(ctx context.Context, id string) (*models.Patient, error) {
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

func (s *PatientService) Update(ctx context.Context, patient *models.Patient) error {
	user, ok := identity.FromCtx(ctx)
	if !ok {
		return domain.ErrAccessDenied
	}

	if patient.Id == nil {
		return domain.ErrPatientIDRequired
	}

	if !user.IsPatient(*patient.Id) && !user.HasScope("patient/*.write") {
		return domain.ErrAccessDenied
	}

	if err := s.validator.Validate(patient); err != nil {
		return fmt.Errorf("%w: %v", domain.ErrInvalidInput, err)
	}

	existing, err := s.repo.GetByID(ctx, *patient.Id)
	if err != nil {
		return fmt.Errorf("%w: %v", domain.ErrInternal, err)
	}

	if existing == nil {
		return domain.ErrPatientNotFound
	}

	if err := s.repo.Update(ctx, patient); err != nil {
		return fmt.Errorf("%w: %v", domain.ErrInternal, err)
	}

	return nil
}
