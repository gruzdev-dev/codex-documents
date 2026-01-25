package ports

import (
	"context"

	"github.com/gruzdev-dev/codex-documents/core/domain"
)

//go:generate mockgen -source=access.go -destination=access_mocks.go -package=ports TmpAccessClient

type TmpAccessClient interface {
	GenerateTmpToken(ctx context.Context, data domain.GenerateTmpTokenRequest) (*domain.GenerateTmpTokenResponse, error)
}
