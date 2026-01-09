package ports

import (
	"context"
)

type FileProvider interface {
	GetUploadURL(ctx context.Context, fileName string, contentType string) (string, error)
}
