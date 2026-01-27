package ports

import (
	"context"

	"github.com/gruzdev-dev/codex-documents/core/domain"
)

//go:generate mockgen -source=share.go -destination=share_mocks.go -package=ports ShareService

type ShareService interface {
	Share(ctx context.Context, req domain.ShareRequest) (*domain.ShareResponse, error)
	GetSharedResources(ctx context.Context) (*domain.SharedResourcesResponse, error)
}
