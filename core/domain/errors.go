package domain

import "errors"

var (
	ErrPatientNotFound   = errors.New("patient not found")
	ErrPatientIDRequired = errors.New("patient id is required")
	ErrAccessDenied      = errors.New("access denied: identity mismatch or insufficient scopes")
	ErrInvalidInput      = errors.New("invalid input data")
	ErrInternal          = errors.New("internal server error")
)
