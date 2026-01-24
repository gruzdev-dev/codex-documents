package domain

import "errors"

var (
	ErrPatientNotFound   = errors.New("patient not found")
	ErrPatientIDRequired = errors.New("patient id is required")

	ErrDocumentNotFound   = errors.New("document not found")
	ErrDocumentIDRequired = errors.New("document id is required")

	ErrObservationNotFound    = errors.New("observation not found")
	ErrObservationIDRequired  = errors.New("observation id is required")
	ErrInvalidDerivedFromRef  = errors.New("derivedFrom must reference DocumentReference resources")
	ErrDerivedFromDocNotFound = errors.New("referenced document not found")

	ErrAccessDenied = errors.New("access denied: identity mismatch or insufficient scopes")
	ErrInvalidInput = errors.New("invalid input data")
	ErrInternal     = errors.New("internal server error")
)
