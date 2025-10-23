package ports

import "codex-documents/core/domain"

type UserService interface {
	HealthCheck() string
}

type UserRepository interface {
	GetUser(id string) (domain.User, error)
}
