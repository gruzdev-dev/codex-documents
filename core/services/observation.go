package services

import (
	"context"
	"fmt"
	"strings"

	"github.com/gruzdev-dev/codex-documents/core/domain"
	"github.com/gruzdev-dev/codex-documents/core/ports"
	"github.com/gruzdev-dev/codex-documents/core/validator"
	"github.com/gruzdev-dev/codex-documents/pkg/identity"

	"github.com/google/uuid"
	models "github.com/gruzdev-dev/fhir/r5"
)

type ObservationService struct {
	repo      ports.ObservationRepository
	docRepo   ports.DocumentRepository
	validator *validator.ObservationValidator
}

func NewObservationService(
	repo ports.ObservationRepository,
	docRepo ports.DocumentRepository,
	v *validator.ObservationValidator,
) *ObservationService {
	return &ObservationService{
		repo:      repo,
		docRepo:   docRepo,
		validator: v,
	}
}

func (s *ObservationService) Create(ctx context.Context, obs *models.Observation) (*models.Observation, error) {
	if err := s.validator.Validate(obs); err != nil {
		return nil, fmt.Errorf("%w: %v", domain.ErrInvalidInput, err)
	}

	user, ok := identity.FromCtx(ctx)
	if !ok {
		return nil, domain.ErrAccessDenied
	}

	if obs.Id != nil && *obs.Id != "" {
		return nil, fmt.Errorf("%w: observation ID must not be provided during creation", domain.ErrInvalidInput)
	}

	id := uuid.New().String()
	obs.Id = &id

	patientRef := fmt.Sprintf("Patient/%s", user.PatientID)
	obs.Subject = &models.Reference{
		Reference: &patientRef,
	}

	if err := s.validateDerivedFrom(ctx, obs.DerivedFrom, user.PatientID); err != nil {
		return nil, err
	}

	created, err := s.repo.Create(ctx, obs)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", domain.ErrInternal, err)
	}

	return created, nil
}

func (s *ObservationService) Get(ctx context.Context, id string) (*models.Observation, error) {
	user, ok := identity.FromCtx(ctx)
	if !ok || !user.HasScope("patient/*.read") {
		return nil, domain.ErrAccessDenied
	}

	if id == "" {
		return nil, domain.ErrObservationIDRequired
	}

	obs, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", domain.ErrInternal, err)
	}
	if obs == nil {
		return nil, domain.ErrObservationNotFound
	}

	if !s.isOwner(user, obs) {
		return nil, domain.ErrAccessDenied
	}

	return obs, nil
}

func (s *ObservationService) Update(ctx context.Context, obs *models.Observation) (*models.Observation, error) {
	user, ok := identity.FromCtx(ctx)
	if !ok || !user.HasScope("patient/*.write") {
		return nil, domain.ErrAccessDenied
	}

	if obs.Id == nil {
		return nil, domain.ErrObservationIDRequired
	}

	existing, err := s.repo.GetByID(ctx, *obs.Id)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", domain.ErrInternal, err)
	}
	if existing == nil {
		return nil, domain.ErrObservationNotFound
	}

	if !s.isOwner(user, existing) {
		return nil, domain.ErrAccessDenied
	}

	if err := s.validator.Validate(obs); err != nil {
		return nil, fmt.Errorf("%w: %v", domain.ErrInvalidInput, err)
	}

	if derivedFromChanged(existing.DerivedFrom, obs.DerivedFrom) {
		if err := s.validateDerivedFrom(ctx, obs.DerivedFrom, user.PatientID); err != nil {
			return nil, err
		}
	}

	updated, err := s.repo.Update(ctx, obs)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", domain.ErrInternal, err)
	}

	return updated, nil
}

func (s *ObservationService) Delete(ctx context.Context, id string) error {
	user, ok := identity.FromCtx(ctx)
	if !ok || !user.HasScope("patient/*.write") {
		return domain.ErrAccessDenied
	}

	existing, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return fmt.Errorf("%w: %v", domain.ErrInternal, err)
	}
	if existing == nil {
		return domain.ErrObservationNotFound
	}

	if !s.isOwner(user, existing) {
		return domain.ErrAccessDenied
	}

	if err := s.repo.Delete(ctx, id); err != nil {
		return fmt.Errorf("%w: %v", domain.ErrInternal, err)
	}

	return nil
}

func (s *ObservationService) List(ctx context.Context, patientID string, limit, offset int) (*domain.ListResponse[models.Observation], error) {
	user, ok := identity.FromCtx(ctx)
	if !ok || !user.HasScope("patient/*.read") {
		return nil, domain.ErrAccessDenied
	}

	if user.PatientID != patientID {
		return nil, domain.ErrAccessDenied
	}

	items, total, err := s.repo.Search(ctx, patientID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", domain.ErrInternal, err)
	}

	return &domain.ListResponse[models.Observation]{
		Items: items,
		Total: total,
	}, nil
}

func (s *ObservationService) isOwner(user domain.Identity, obs *models.Observation) bool {
	if obs.Subject == nil || obs.Subject.Reference == nil {
		return false
	}
	return strings.HasSuffix(*obs.Subject.Reference, user.PatientID)
}

func (s *ObservationService) validateDerivedFrom(ctx context.Context, derivedFrom []models.Reference, patientID string) error {
	if len(derivedFrom) == 0 {
		return nil
	}

	for _, ref := range derivedFrom {
		if ref.Reference == nil {
			continue
		}

		refStr := *ref.Reference
		if !strings.HasPrefix(refStr, "DocumentReference/") {
			return domain.ErrInvalidDerivedFromRef
		}

		docID := strings.TrimPrefix(refStr, "DocumentReference/")
		if docID == "" {
			return domain.ErrInvalidDerivedFromRef
		}

		doc, err := s.docRepo.GetByID(ctx, docID)
		if err != nil {
			return fmt.Errorf("%w: %v", domain.ErrInternal, err)
		}
		if doc == nil {
			return domain.ErrDerivedFromDocNotFound
		}

		if doc.Subject == nil || doc.Subject.Reference == nil {
			return domain.ErrDerivedFromDocNotFound
		}

		expectedPatientRef := fmt.Sprintf("Patient/%s", patientID)
		if *doc.Subject.Reference != expectedPatientRef {
			return domain.ErrAccessDenied
		}
	}

	return nil
}

func derivedFromChanged(old, new []models.Reference) bool {
	if len(old) != len(new) {
		return true
	}

	oldMap := make(map[string]bool)
	for _, ref := range old {
		if ref.Reference != nil {
			oldMap[*ref.Reference] = true
		}
	}

	for _, ref := range new {
		if ref.Reference == nil {
			return true
		}
		if !oldMap[*ref.Reference] {
			return true
		}
	}

	return false
}
