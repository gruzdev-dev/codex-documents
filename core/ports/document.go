package ports

import (
	"context"

	"github.com/gruzdev-dev/codex-documents/core/domain"
	models "github.com/gruzdev-dev/fhir/r5"
)

type DocumentRepository interface {
	Create(ctx context.Context, doc *models.DocumentReference) (*models.DocumentReference, error)
	GetByID(ctx context.Context, id string) (*models.DocumentReference, error)
	Update(ctx context.Context, doc *models.DocumentReference) (*models.DocumentReference, error)
	Delete(ctx context.Context, id string) error
	Search(ctx context.Context, patientID string, limit, offset int) ([]models.DocumentReference, int64, error)
}

type DocumentService interface {
	CreateDocument(ctx context.Context, doc *models.DocumentReference) (*domain.CreateDocumentResult, error)
	GetDocument(ctx context.Context, id string) (*models.DocumentReference, error)
	DeleteDocument(ctx context.Context, id string) error
	ListDocuments(ctx context.Context, patientID string, limit, offset int) (*domain.ListResponse[models.DocumentReference], error)
}
