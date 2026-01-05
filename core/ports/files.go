package ports

import (
	"context"
)

type FilesServicePort interface {
	GetUploadURL(ctx context.Context, fileID string) (string, error)
	GetDownloadURL(ctx context.Context, fileID string) (string, error)
}
