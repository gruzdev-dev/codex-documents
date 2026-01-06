package models

import (
	"encoding/json"
	"fmt"
)

// The details of a healthcare service available at a location or in a catalog.  In the case where there is a hierarchy of services (for example, Lab -> Pathology -> Wound Cultures), this can be represented using a set of linked HealthcareServices.
type HealthcareService struct {
	Id                   *string                        `json:"id,omitempty" bson:"id,omitempty"`                                       // Logical id of this artifact
	Meta                 *Meta                          `json:"meta,omitempty" bson:"meta,omitempty"`                                   // Metadata about the resource
	ImplicitRules        *string                        `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`                // A set of rules under which this content was created
	Language             *string                        `json:"language,omitempty" bson:"language,omitempty"`                           // Language of the resource content
	Text                 *Narrative                     `json:"text,omitempty" bson:"text,omitempty"`                                   // Text summary of the resource, for human interpretation
	Contained            []json.RawMessage              `json:"contained,omitempty" bson:"contained,omitempty"`                         // Contained, inline Resources
	Identifier           []Identifier                   `json:"identifier,omitempty" bson:"identifier,omitempty"`                       // External identifiers for this item
	Active               bool                           `json:"active,omitempty" bson:"active,omitempty"`                               // Whether this HealthcareService record is in active use
	ProvidedBy           *Reference                     `json:"providedBy,omitempty" bson:"provided_by,omitempty"`                      // Organization that provides this service
	OfferedIn            []Reference                    `json:"offeredIn,omitempty" bson:"offered_in,omitempty"`                        // The service within which this service is offered
	Category             []CodeableConcept              `json:"category,omitempty" bson:"category,omitempty"`                           // Broad category of service being performed or delivered
	Type                 []CodeableConcept              `json:"type,omitempty" bson:"type,omitempty"`                                   // Type of service that may be delivered or performed
	Specialty            []CodeableConcept              `json:"specialty,omitempty" bson:"specialty,omitempty"`                         // Specialties handled by the HealthcareService
	Location             []Reference                    `json:"location,omitempty" bson:"location,omitempty"`                           // Location(s) where service may be provided
	Name                 *string                        `json:"name,omitempty" bson:"name,omitempty"`                                   // Description of service as presented to a consumer while searching
	Comment              *string                        `json:"comment,omitempty" bson:"comment,omitempty"`                             // Additional description and/or any specific issues not covered elsewhere
	ExtraDetails         *string                        `json:"extraDetails,omitempty" bson:"extra_details,omitempty"`                  // Extra details about the service that can't be placed in the other fields
	Photo                *Attachment                    `json:"photo,omitempty" bson:"photo,omitempty"`                                 // Facilitates quick identification of the service
	Contact              []ExtendedContactDetail        `json:"contact,omitempty" bson:"contact,omitempty"`                             // Official contact details for the HealthcareService
	CoverageArea         []Reference                    `json:"coverageArea,omitempty" bson:"coverage_area,omitempty"`                  // Location(s) service is intended for/available to
	ServiceProvisionCode []CodeableConcept              `json:"serviceProvisionCode,omitempty" bson:"service_provision_code,omitempty"` // Conditions under which service is available/offered
	Eligibility          []HealthcareServiceEligibility `json:"eligibility,omitempty" bson:"eligibility,omitempty"`                     // Specific eligibility requirements required to use the service
	Program              []CodeableConcept              `json:"program,omitempty" bson:"program,omitempty"`                             // Programs that this service is applicable to
	Characteristic       []CodeableConcept              `json:"characteristic,omitempty" bson:"characteristic,omitempty"`               // Collection of characteristics (attributes)
	Communication        []CodeableConcept              `json:"communication,omitempty" bson:"communication,omitempty"`                 // The language that this service is offered in
	ReferralMethod       []CodeableConcept              `json:"referralMethod,omitempty" bson:"referral_method,omitempty"`              // Ways that the service accepts referrals
	ReferralRequired     bool                           `json:"referralRequired,omitempty" bson:"referral_required,omitempty"`          // A referral is required for access to this service
	AppointmentRequired  bool                           `json:"appointmentRequired,omitempty" bson:"appointment_required,omitempty"`    // An appointment is required for access to this service
	Availability         *Availability                  `json:"availability,omitempty" bson:"availability,omitempty"`                   // Times the healthcare service is available (including exceptions)
	Endpoint             []Reference                    `json:"endpoint,omitempty" bson:"endpoint,omitempty"`                           // Technical endpoints providing access to electronic services operated for the healthcare service
}

func (r *HealthcareService) Validate() error {
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
	if r.ProvidedBy != nil {
		if err := r.ProvidedBy.Validate(); err != nil {
			return fmt.Errorf("ProvidedBy: %w", err)
		}
	}
	for i, item := range r.OfferedIn {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("OfferedIn[%d]: %w", i, err)
		}
	}
	for i, item := range r.Category {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Category[%d]: %w", i, err)
		}
	}
	for i, item := range r.Type {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Type[%d]: %w", i, err)
		}
	}
	for i, item := range r.Specialty {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Specialty[%d]: %w", i, err)
		}
	}
	for i, item := range r.Location {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Location[%d]: %w", i, err)
		}
	}
	if r.Photo != nil {
		if err := r.Photo.Validate(); err != nil {
			return fmt.Errorf("Photo: %w", err)
		}
	}
	for i, item := range r.Contact {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Contact[%d]: %w", i, err)
		}
	}
	for i, item := range r.CoverageArea {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("CoverageArea[%d]: %w", i, err)
		}
	}
	for i, item := range r.ServiceProvisionCode {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("ServiceProvisionCode[%d]: %w", i, err)
		}
	}
	for i, item := range r.Eligibility {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Eligibility[%d]: %w", i, err)
		}
	}
	for i, item := range r.Program {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Program[%d]: %w", i, err)
		}
	}
	for i, item := range r.Characteristic {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Characteristic[%d]: %w", i, err)
		}
	}
	for i, item := range r.Communication {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Communication[%d]: %w", i, err)
		}
	}
	for i, item := range r.ReferralMethod {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("ReferralMethod[%d]: %w", i, err)
		}
	}
	if r.Availability != nil {
		if err := r.Availability.Validate(); err != nil {
			return fmt.Errorf("Availability: %w", err)
		}
	}
	for i, item := range r.Endpoint {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Endpoint[%d]: %w", i, err)
		}
	}
	return nil
}

type HealthcareServiceEligibility struct {
	Id                   *string          `json:"id,omitempty" bson:"id,omitempty"`                                       // Unique id for inter-element referencing
	Code                 *CodeableConcept `json:"code,omitempty" bson:"code,omitempty"`                                   // Coded value for the eligibility
	ValueCodeableConcept *CodeableConcept `json:"valueCodeableConcept,omitempty" bson:"value_codeable_concept,omitempty"` // Value associated with the eligibility code
	ValueBoolean         *bool            `json:"valueBoolean,omitempty" bson:"value_boolean,omitempty"`                  // Value associated with the eligibility code
	ValueQuantity        *Quantity        `json:"valueQuantity,omitempty" bson:"value_quantity,omitempty"`                // Value associated with the eligibility code
	ValueRange           *Range           `json:"valueRange,omitempty" bson:"value_range,omitempty"`                      // Value associated with the eligibility code
	ValueReference       *Reference       `json:"valueReference,omitempty" bson:"value_reference,omitempty"`              // Value associated with the eligibility code
	Comment              *string          `json:"comment,omitempty" bson:"comment,omitempty"`                             // Describes the eligibility conditions for the service
	Period               *Period          `json:"period,omitempty" bson:"period,omitempty"`                               // The period this eligibility rule applies
}

func (r *HealthcareServiceEligibility) Validate() error {
	if r.Code != nil {
		if err := r.Code.Validate(); err != nil {
			return fmt.Errorf("Code: %w", err)
		}
	}
	if r.ValueCodeableConcept != nil {
		if err := r.ValueCodeableConcept.Validate(); err != nil {
			return fmt.Errorf("ValueCodeableConcept: %w", err)
		}
	}
	if r.ValueQuantity != nil {
		if err := r.ValueQuantity.Validate(); err != nil {
			return fmt.Errorf("ValueQuantity: %w", err)
		}
	}
	if r.ValueRange != nil {
		if err := r.ValueRange.Validate(); err != nil {
			return fmt.Errorf("ValueRange: %w", err)
		}
	}
	if r.ValueReference != nil {
		if err := r.ValueReference.Validate(); err != nil {
			return fmt.Errorf("ValueReference: %w", err)
		}
	}
	if r.Period != nil {
		if err := r.Period.Validate(); err != nil {
			return fmt.Errorf("Period: %w", err)
		}
	}
	return nil
}
