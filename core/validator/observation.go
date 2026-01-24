package validator

import (
	"errors"

	models "github.com/gruzdev-dev/fhir/r5"
)

type ObservationValidator struct{}

func NewObservationValidator() *ObservationValidator {
	return &ObservationValidator{}
}

func (v *ObservationValidator) Validate(obs *models.Observation) error {
	if obs == nil {
		return errors.New("observation resource is nil")
	}
	return nil
}
