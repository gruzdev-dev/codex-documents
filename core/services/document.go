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

type DocumentService struct {
	repo         ports.DocumentRepository
	fileProvider ports.FileProvider
	validator    *validator.DocumentValidator
}

func NewDocumentService(
	repo ports.DocumentRepository,
	fileProvider ports.FileProvider,
	v *validator.DocumentValidator,
) *DocumentService {
	return &DocumentService{
		repo:         repo,
		fileProvider: fileProvider,
		validator:    v,
	}
}

func (s *DocumentService) CreateDocument(ctx context.Context, doc *models.DocumentReference) (*domain.CreateDocumentResult, error) {
	if err := s.validator.Validate(doc); err != nil {
		return nil, fmt.Errorf("%w: %v", domain.ErrInvalidInput, err)
	}
	user, ok := identity.FromCtx(ctx)
	if !ok {
		return nil, domain.ErrAccessDenied
	}

	if doc.Id != nil && *doc.Id != "" {
		return nil, fmt.Errorf("%w: document ID must not be provided during creation", domain.ErrInvalidInput)
	}
	id := uuid.New().String()
	doc.Id = &id

	patientRef := fmt.Sprintf("Patient/%s", user.PatientID)
	doc.Subject = &models.Reference{
		Reference: &patientRef,
	}

	uploadUrls := make(map[string]string)

	if len(doc.Content) > 0 {
		for i := range doc.Content {
			attachment := doc.Content[i].Attachment
			if attachment == nil {
				continue
			}

			if attachment.ContentType == nil || attachment.Size == nil {
				continue
			}

			if attachment.Url != nil || attachment.Data != nil {
				continue
			}

			presignedUrls, err := s.fileProvider.GetPresignedUrls(ctx, domain.GetPresignedUrlsRequest{
				UserId:      user.PatientID,
				ContentType: *attachment.ContentType,
				Size:        *attachment.Size,
			})

			if err != nil {
				return nil, fmt.Errorf("%w: %v", domain.ErrInternal, err)
			}

			doc.Content[i].Attachment.Id = &presignedUrls.FileId
			doc.Content[i].Attachment.Url = &presignedUrls.DownloadUrl
			uploadUrls[presignedUrls.FileId] = presignedUrls.UploadUrl
		}
	}

	created, err := s.repo.Create(ctx, doc)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", domain.ErrInternal, err)
	}

	return &domain.CreateDocumentResult{
		Document:   created,
		UploadUrls: uploadUrls,
	}, nil
}

func (s *DocumentService) GetDocument(ctx context.Context, id string) (*models.DocumentReference, error) {
	user, ok := identity.FromCtx(ctx)
	if !ok || !user.HasScope("patient/*.read") {
		return nil, domain.ErrAccessDenied
	}

	if id == "" {
		return nil, domain.ErrDocumentIDRequired
	}

	doc, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", domain.ErrInternal, err)
	}
	if doc == nil {
		return nil, domain.ErrDocumentNotFound
	}

	if !s.isOwner(user, doc) {
		return nil, domain.ErrAccessDenied
	}

	return doc, nil
}

func (s *DocumentService) DeleteDocument(ctx context.Context, id string) error {
	user, ok := identity.FromCtx(ctx)
	if !ok || !user.HasScope("patient/*.write") {
		return domain.ErrAccessDenied
	}

	existing, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return fmt.Errorf("%w: %v", domain.ErrInternal, err)
	}
	if existing == nil {
		return domain.ErrDocumentNotFound
	}

	if !s.isOwner(user, existing) {
		return domain.ErrAccessDenied
	}

	if existing.Content != nil {
		for _, content := range existing.Content {
			if content.Attachment == nil || content.Attachment.Id == nil {
				continue
			}
			_ = s.fileProvider.DeleteFile(ctx, *content.Attachment.Id)
		}
	}

	if err := s.repo.Delete(ctx, id); err != nil {
		return fmt.Errorf("%w: %v", domain.ErrInternal, err)
	}

	return nil
}

func (s *DocumentService) ListDocuments(ctx context.Context, patientID string, limit, offset int) (*domain.ListResponse[models.DocumentReference], error) {
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

	return &domain.ListResponse[models.DocumentReference]{
		Items: items,
		Total: total,
	}, nil
}

func (s *DocumentService) isOwner(user domain.Identity, doc *models.DocumentReference) bool {
	if doc.Subject == nil || doc.Subject.Reference == nil {
		return false
	}
	return strings.HasSuffix(*doc.Subject.Reference, user.PatientID)
}
