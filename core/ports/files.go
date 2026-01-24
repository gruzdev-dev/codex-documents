package ports

import (
	"context"

	"github.com/gruzdev-dev/codex-documents/core/domain"
)

//go:generate mockgen -source=files.go -destination=files_mocks.go -package=ports FileProvider

type FileProvider interface {
	GetPresignedUrls(ctx context.Context, data domain.GetPresignedUrlsRequest) (*domain.PresignedUrlsResponse, error)
	DeleteFile(ctx context.Context, fileId string) error
}
