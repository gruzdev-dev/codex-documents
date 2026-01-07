package validator

import (
	"errors"
	"strings"

	models "github.com/gruzdev-dev/fhir/r5"
)

type PatientValidator struct{}

func NewPatientValidator() *PatientValidator {
	return &PatientValidator{}
}

func (v *PatientValidator) Validate(p *models.Patient) error {
	if p == nil {
		return errors.New("patient resource is nil")
	}

	if err := v.checkEmailPresence(p); err != nil {
		return err
	}

	return nil
}

func (v *PatientValidator) checkEmailPresence(p *models.Patient) error {
	hasEmail := false
	for _, telecom := range p.Telecom {
		if telecom.System != nil && *telecom.System == "email" {
			if telecom.Value != nil && strings.Contains(*telecom.Value, "@") {
				hasEmail = true
				break
			}
		}
	}

	if !hasEmail {
		return errors.New("patient must have a valid email in telecom contacts")
	}

	return nil
}