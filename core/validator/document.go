package validator

import (
	"errors"

	models "github.com/gruzdev-dev/fhir/r5"
)

type DocumentValidator struct{}

func NewDocumentValidator() *DocumentValidator {
	return &DocumentValidator{}
}

func (v *DocumentValidator) Validate(doc *models.DocumentReference) error {
	if doc == nil {
		return errors.New("document resource is nil")
	}

	if len(doc.Content) == 0 {
		return errors.New("document must have at least one content entry")
	}

	for _, c := range doc.Content {
		if c.Attachment == nil {
			return errors.New("each content entry must have an attachment")
		}
	}

	return nil
}
