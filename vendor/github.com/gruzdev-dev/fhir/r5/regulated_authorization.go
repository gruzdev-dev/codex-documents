package models

import (
	"encoding/json"
	"fmt"
)

// Regulatory approval, clearance or licensing related to a regulated product, treatment, facility or activity that is cited in a guidance, regulation, rule or legislative act. An example is Market Authorization relating to a Medicinal Product.
type RegulatedAuthorization struct {
	Id               *string                     `json:"id,omitempty" bson:"id,omitempty"`                              // Logical id of this artifact
	Meta             *Meta                       `json:"meta,omitempty" bson:"meta,omitempty"`                          // Metadata about the resource
	ImplicitRules    *string                     `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`       // A set of rules under which this content was created
	Language         *string                     `json:"language,omitempty" bson:"language,omitempty"`                  // Language of the resource content
	Text             *Narrative                  `json:"text,omitempty" bson:"text,omitempty"`                          // Text summary of the resource, for human interpretation
	Contained        []json.RawMessage           `json:"contained,omitempty" bson:"contained,omitempty"`                // Contained, inline Resources
	Identifier       []Identifier                `json:"identifier,omitempty" bson:"identifier,omitempty"`              // Business identifier for the authorization, typically assigned by the authorizing body
	Subject          []Reference                 `json:"subject,omitempty" bson:"subject,omitempty"`                    // The product type, treatment, facility or activity that is being authorized
	Type             *CodeableConcept            `json:"type,omitempty" bson:"type,omitempty"`                          // Overall type of this authorization, for example drug marketing approval, orphan drug designation
	Description      *string                     `json:"description,omitempty" bson:"description,omitempty"`            // General textual supporting information
	Region           []CodeableConcept           `json:"region,omitempty" bson:"region,omitempty"`                      // The territory in which the authorization has been granted
	Status           *CodeableConcept            `json:"status,omitempty" bson:"status,omitempty"`                      // The status that is authorised e.g. approved. Intermediate states can be tracked with cases and applications
	StatusDate       *string                     `json:"statusDate,omitempty" bson:"status_date,omitempty"`             // The date at which the current status was assigned
	ValidityPeriod   *Period                     `json:"validityPeriod,omitempty" bson:"validity_period,omitempty"`     // The time period in which the regulatory approval etc. is in effect, e.g. a Marketing Authorization includes the date of authorization and/or expiration date
	Indication       []CodeableReference         `json:"indication,omitempty" bson:"indication,omitempty"`              // Condition for which the use of the regulated product applies
	IntendedUse      *CodeableConcept            `json:"intendedUse,omitempty" bson:"intended_use,omitempty"`           // The intended use of the product, e.g. prevention, treatment
	Basis            []CodeableConcept           `json:"basis,omitempty" bson:"basis,omitempty"`                        // The legal/regulatory framework or reasons under which this authorization is granted
	Holder           *Reference                  `json:"holder,omitempty" bson:"holder,omitempty"`                      // The organization that has been granted this authorization, by the regulator
	Regulator        *Reference                  `json:"regulator,omitempty" bson:"regulator,omitempty"`                // The regulatory authority or authorizing body granting the authorization
	AttachedDocument []Reference                 `json:"attachedDocument,omitempty" bson:"attached_document,omitempty"` // Additional information or supporting documentation about the authorization
	Case             *RegulatedAuthorizationCase `json:"case,omitempty" bson:"case,omitempty"`                          // The case or regulatory procedure for granting or amending a regulated authorization. Note: This area is subject to ongoing review and the workgroup is seeking implementer feedback on its use (see link at bottom of page)
}

func (r *RegulatedAuthorization) Validate() error {
	if r.Meta != nil {
		if err := r.Meta.Validate(); err != nil {
			return fmt.Errorf("Meta: %w", err)
		}
	}
	if r.Text != nil {
		if err := r.Text.Validate(); err != nil {
			return fmt.Errorf("Text: %w", err)
		}
	}
	for i, item := range r.Identifier {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Identifier[%d]: %w", i, err)
		}
	}
	for i, item := range r.Subject {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Subject[%d]: %w", i, err)
		}
	}
	if r.Type != nil {
		if err := r.Type.Validate(); err != nil {
			return fmt.Errorf("Type: %w", err)
		}
	}
	for i, item := range r.Region {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Region[%d]: %w", i, err)
		}
	}
	if r.Status != nil {
		if err := r.Status.Validate(); err != nil {
			return fmt.Errorf("Status: %w", err)
		}
	}
	if r.ValidityPeriod != nil {
		if err := r.ValidityPeriod.Validate(); err != nil {
			return fmt.Errorf("ValidityPeriod: %w", err)
		}
	}
	for i, item := range r.Indication {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Indication[%d]: %w", i, err)
		}
	}
	if r.IntendedUse != nil {
		if err := r.IntendedUse.Validate(); err != nil {
			return fmt.Errorf("IntendedUse: %w", err)
		}
	}
	for i, item := range r.Basis {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Basis[%d]: %w", i, err)
		}
	}
	if r.Holder != nil {
		if err := r.Holder.Validate(); err != nil {
			return fmt.Errorf("Holder: %w", err)
		}
	}
	if r.Regulator != nil {
		if err := r.Regulator.Validate(); err != nil {
			return fmt.Errorf("Regulator: %w", err)
		}
	}
	for i, item := range r.AttachedDocument {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("AttachedDocument[%d]: %w", i, err)
		}
	}
	if r.Case != nil {
		if err := r.Case.Validate(); err != nil {
			return fmt.Errorf("Case: %w", err)
		}
	}
	return nil
}

type RegulatedAuthorizationCase struct {
	Id           *string                      `json:"id,omitempty" bson:"id,omitempty"`                       // Unique id for inter-element referencing
	Identifier   *Identifier                  `json:"identifier,omitempty" bson:"identifier,omitempty"`       // Identifier by which this case can be referenced
	Type         *CodeableConcept             `json:"type,omitempty" bson:"type,omitempty"`                   // The defining type of case
	Status       *CodeableConcept             `json:"status,omitempty" bson:"status,omitempty"`               // The status associated with the case
	DatePeriod   *Period                      `json:"datePeriod,omitempty" bson:"date_period,omitempty"`      // Relevant date for this case
	DateDateTime *string                      `json:"dateDateTime,omitempty" bson:"date_date_time,omitempty"` // Relevant date for this case
	Application  []RegulatedAuthorizationCase `json:"application,omitempty" bson:"application,omitempty"`     // Applications submitted to obtain a regulated authorization. Steps within the longer running case or procedure
}

func (r *RegulatedAuthorizationCase) Validate() error {
	if r.Identifier != nil {
		if err := r.Identifier.Validate(); err != nil {
			return fmt.Errorf("Identifier: %w", err)
		}
	}
	if r.Type != nil {
		if err := r.Type.Validate(); err != nil {
			return fmt.Errorf("Type: %w", err)
		}
	}
	if r.Status != nil {
		if err := r.Status.Validate(); err != nil {
			return fmt.Errorf("Status: %w", err)
		}
	}
	if r.DatePeriod != nil {
		if err := r.DatePeriod.Validate(); err != nil {
			return fmt.Errorf("DatePeriod: %w", err)
		}
	}
	for i, item := range r.Application {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Application[%d]: %w", i, err)
		}
	}
	return nil
}
