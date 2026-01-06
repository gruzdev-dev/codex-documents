package models

import (
	"fmt"
)

// Signature Type: A signature along with supporting context. The signature may be a digital signature that is cryptographic in nature, or some other signature acceptable to the domain. This other signature may be as simple as a graphical image representing a hand-written signature, or a signature ceremony Different signature approaches have different utilities.
type Signature struct {
	Id           *string    `json:"id,omitempty" bson:"id,omitempty"`                      // Unique id for inter-element referencing
	Type         []Coding   `json:"type,omitempty" bson:"type,omitempty"`                  // Indication of the reason the entity signed the object(s)
	When         *string    `json:"when,omitempty" bson:"when,omitempty"`                  // When the signature was created
	Who          *Reference `json:"who,omitempty" bson:"who,omitempty"`                    // Who signed
	OnBehalfOf   *Reference `json:"onBehalfOf,omitempty" bson:"on_behalf_of,omitempty"`    // The party represented
	TargetFormat *string    `json:"targetFormat,omitempty" bson:"target_format,omitempty"` // The technical format of the signed resources
	SigFormat    *string    `json:"sigFormat,omitempty" bson:"sig_format,omitempty"`       // The technical format of the signature
	Data         *string    `json:"data,omitempty" bson:"data,omitempty"`                  // The actual signature content (XML Signature, JSON Jose, picture, etc.)
}

func (r *Signature) Validate() error {
	for i, item := range r.Type {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Type[%d]: %w", i, err)
		}
	}
	if r.Who != nil {
		if err := r.Who.Validate(); err != nil {
			return fmt.Errorf("Who: %w", err)
		}
	}
	if r.OnBehalfOf != nil {
		if err := r.OnBehalfOf.Validate(); err != nil {
			return fmt.Errorf("OnBehalfOf: %w", err)
		}
	}
	return nil
}
