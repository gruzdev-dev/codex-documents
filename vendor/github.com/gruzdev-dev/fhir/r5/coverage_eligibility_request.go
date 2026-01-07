package models

import (
	"encoding/json"
	"fmt"
)

// The CoverageEligibilityRequest provides patient and insurance coverage information to an insurer for them to respond, in the form of an CoverageEligibilityResponse, with information regarding whether the stated coverage is valid and in-force and optionally to provide the insurance details of the policy.
type CoverageEligibilityRequest struct {
	ResourceType   string                                     `json:"resourceType" bson:"resource_type"`                         // Type of resource
	Id             *string                                    `json:"id,omitempty" bson:"id,omitempty"`                          // Logical id of this artifact
	Meta           *Meta                                      `json:"meta,omitempty" bson:"meta,omitempty"`                      // Metadata about the resource
	ImplicitRules  *string                                    `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`   // A set of rules under which this content was created
	Language       *string                                    `json:"language,omitempty" bson:"language,omitempty"`              // Language of the resource content
	Text           *Narrative                                 `json:"text,omitempty" bson:"text,omitempty"`                      // Text summary of the resource, for human interpretation
	Contained      []json.RawMessage                          `json:"contained,omitempty" bson:"contained,omitempty"`            // Contained, inline Resources
	Identifier     []Identifier                               `json:"identifier,omitempty" bson:"identifier,omitempty"`          // Business Identifier for coverage eligiblity request
	Status         string                                     `json:"status" bson:"status"`                                      // active | cancelled | draft | entered-in-error
	StatusReason   *string                                    `json:"statusReason,omitempty" bson:"status_reason,omitempty"`     // Reason for status change
	Priority       *CodeableConcept                           `json:"priority,omitempty" bson:"priority,omitempty"`              // Desired processing priority
	Purpose        []string                                   `json:"purpose" bson:"purpose"`                                    // auth-requirements | benefits | discovery | validation
	Patient        *Reference                                 `json:"patient" bson:"patient"`                                    // Intended recipient of products and services
	Event          []CoverageEligibilityRequestEvent          `json:"event,omitempty" bson:"event,omitempty"`                    // Event information
	ServicedDate   *string                                    `json:"servicedDate,omitempty" bson:"serviced_date,omitempty"`     // Estimated date or dates of service
	ServicedPeriod *Period                                    `json:"servicedPeriod,omitempty" bson:"serviced_period,omitempty"` // Estimated date or dates of service
	Created        string                                     `json:"created" bson:"created"`                                    // Creation date
	Enterer        *Reference                                 `json:"enterer,omitempty" bson:"enterer,omitempty"`                // Author
	Provider       *Reference                                 `json:"provider,omitempty" bson:"provider,omitempty"`              // Party responsible for the request
	Insurer        *Reference                                 `json:"insurer" bson:"insurer"`                                    // Coverage issuer
	Facility       *Reference                                 `json:"facility,omitempty" bson:"facility,omitempty"`              // Servicing facility
	SupportingInfo []CoverageEligibilityRequestSupportingInfo `json:"supportingInfo,omitempty" bson:"supporting_info,omitempty"` // Supporting information
	Insurance      []CoverageEligibilityRequestInsurance      `json:"insurance,omitempty" bson:"insurance,omitempty"`            // Patient insurance information
	Item           []CoverageEligibilityRequestItem           `json:"item,omitempty" bson:"item,omitempty"`                      // Item to be evaluated for eligibiity
}

func (r *CoverageEligibilityRequest) Validate() error {
	if r.ResourceType != "CoverageEligibilityRequest" {
		return fmt.Errorf("invalid resourceType: expected 'CoverageEligibilityRequest', got '%s'", r.ResourceType)
	}
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
	var emptyString string
	if r.Status == emptyString {
		return fmt.Errorf("field 'Status' is required")
	}
	if r.Priority != nil {
		if err := r.Priority.Validate(); err != nil {
			return fmt.Errorf("Priority: %w", err)
		}
	}
	if len(r.Purpose) < 1 {
		return fmt.Errorf("field 'Purpose' must have at least 1 elements")
	}
	if r.Patient == nil {
		return fmt.Errorf("field 'Patient' is required")
	}
	if r.Patient != nil {
		if err := r.Patient.Validate(); err != nil {
			return fmt.Errorf("Patient: %w", err)
		}
	}
	for i, item := range r.Event {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Event[%d]: %w", i, err)
		}
	}
	if r.ServicedPeriod != nil {
		if err := r.ServicedPeriod.Validate(); err != nil {
			return fmt.Errorf("ServicedPeriod: %w", err)
		}
	}
	if r.Created == emptyString {
		return fmt.Errorf("field 'Created' is required")
	}
	if r.Enterer != nil {
		if err := r.Enterer.Validate(); err != nil {
			return fmt.Errorf("Enterer: %w", err)
		}
	}
	if r.Provider != nil {
		if err := r.Provider.Validate(); err != nil {
			return fmt.Errorf("Provider: %w", err)
		}
	}
	if r.Insurer == nil {
		return fmt.Errorf("field 'Insurer' is required")
	}
	if r.Insurer != nil {
		if err := r.Insurer.Validate(); err != nil {
			return fmt.Errorf("Insurer: %w", err)
		}
	}
	if r.Facility != nil {
		if err := r.Facility.Validate(); err != nil {
			return fmt.Errorf("Facility: %w", err)
		}
	}
	for i, item := range r.SupportingInfo {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("SupportingInfo[%d]: %w", i, err)
		}
	}
	for i, item := range r.Insurance {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Insurance[%d]: %w", i, err)
		}
	}
	for i, item := range r.Item {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Item[%d]: %w", i, err)
		}
	}
	return nil
}

type CoverageEligibilityRequestInsurance struct {
	Id                  *string    `json:"id,omitempty" bson:"id,omitempty"`                                    // Unique id for inter-element referencing
	Focal               *bool      `json:"focal,omitempty" bson:"focal,omitempty"`                              // Applicable coverage
	Coverage            *Reference `json:"coverage" bson:"coverage"`                                            // Insurance information
	BusinessArrangement *string    `json:"businessArrangement,omitempty" bson:"business_arrangement,omitempty"` // Additional provider contract number
}

func (r *CoverageEligibilityRequestInsurance) Validate() error {
	if r.Coverage == nil {
		return fmt.Errorf("field 'Coverage' is required")
	}
	if r.Coverage != nil {
		if err := r.Coverage.Validate(); err != nil {
			return fmt.Errorf("Coverage: %w", err)
		}
	}
	return nil
}

type CoverageEligibilityRequestItem struct {
	Id                     *string                                   `json:"id,omitempty" bson:"id,omitempty"`                                           // Unique id for inter-element referencing
	SupportingInfoSequence []int                                     `json:"supportingInfoSequence,omitempty" bson:"supporting_info_sequence,omitempty"` // Applicable exception or supporting information
	Category               *CodeableConcept                          `json:"category,omitempty" bson:"category,omitempty"`                               // Benefit classification
	ProductOrService       *CodeableConcept                          `json:"productOrService,omitempty" bson:"product_or_service,omitempty"`             // Billing, service, product, or drug code
	Modifier               []CodeableConcept                         `json:"modifier,omitempty" bson:"modifier,omitempty"`                               // Product or service billing modifiers
	Provider               *Reference                                `json:"provider,omitempty" bson:"provider,omitempty"`                               // Perfoming practitioner
	Quantity               *Quantity                                 `json:"quantity,omitempty" bson:"quantity,omitempty"`                               // Count of products or services
	UnitPrice              *Money                                    `json:"unitPrice,omitempty" bson:"unit_price,omitempty"`                            // Fee, charge or cost per item
	Facility               *Reference                                `json:"facility,omitempty" bson:"facility,omitempty"`                               // Servicing facility
	Diagnosis              []CoverageEligibilityRequestItemDiagnosis `json:"diagnosis,omitempty" bson:"diagnosis,omitempty"`                             // Applicable diagnosis
	Detail                 []Reference                               `json:"detail,omitempty" bson:"detail,omitempty"`                                   // Product or service details
}

func (r *CoverageEligibilityRequestItem) Validate() error {
	if r.Category != nil {
		if err := r.Category.Validate(); err != nil {
			return fmt.Errorf("Category: %w", err)
		}
	}
	if r.ProductOrService != nil {
		if err := r.ProductOrService.Validate(); err != nil {
			return fmt.Errorf("ProductOrService: %w", err)
		}
	}
	for i, item := range r.Modifier {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Modifier[%d]: %w", i, err)
		}
	}
	if r.Provider != nil {
		if err := r.Provider.Validate(); err != nil {
			return fmt.Errorf("Provider: %w", err)
		}
	}
	if r.Quantity != nil {
		if err := r.Quantity.Validate(); err != nil {
			return fmt.Errorf("Quantity: %w", err)
		}
	}
	if r.UnitPrice != nil {
		if err := r.UnitPrice.Validate(); err != nil {
			return fmt.Errorf("UnitPrice: %w", err)
		}
	}
	if r.Facility != nil {
		if err := r.Facility.Validate(); err != nil {
			return fmt.Errorf("Facility: %w", err)
		}
	}
	for i, item := range r.Diagnosis {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Diagnosis[%d]: %w", i, err)
		}
	}
	for i, item := range r.Detail {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Detail[%d]: %w", i, err)
		}
	}
	return nil
}

type CoverageEligibilityRequestItemDiagnosis struct {
	Id                       *string          `json:"id,omitempty" bson:"id,omitempty"`                                               // Unique id for inter-element referencing
	DiagnosisCodeableConcept *CodeableConcept `json:"diagnosisCodeableConcept,omitempty" bson:"diagnosis_codeable_concept,omitempty"` // Nature of illness or problem
	DiagnosisReference       *Reference       `json:"diagnosisReference,omitempty" bson:"diagnosis_reference,omitempty"`              // Nature of illness or problem
}

func (r *CoverageEligibilityRequestItemDiagnosis) Validate() error {
	if r.DiagnosisCodeableConcept != nil {
		if err := r.DiagnosisCodeableConcept.Validate(); err != nil {
			return fmt.Errorf("DiagnosisCodeableConcept: %w", err)
		}
	}
	if r.DiagnosisReference != nil {
		if err := r.DiagnosisReference.Validate(); err != nil {
			return fmt.Errorf("DiagnosisReference: %w", err)
		}
	}
	return nil
}

type CoverageEligibilityRequestEvent struct {
	Id           *string          `json:"id,omitempty" bson:"id,omitempty"`   // Unique id for inter-element referencing
	Type         *CodeableConcept `json:"type" bson:"type"`                   // Specific event
	WhenDateTime *string          `json:"whenDateTime" bson:"when_date_time"` // Occurance date or period
	WhenPeriod   *Period          `json:"whenPeriod" bson:"when_period"`      // Occurance date or period
}

func (r *CoverageEligibilityRequestEvent) Validate() error {
	if r.Type == nil {
		return fmt.Errorf("field 'Type' is required")
	}
	if r.Type != nil {
		if err := r.Type.Validate(); err != nil {
			return fmt.Errorf("Type: %w", err)
		}
	}
	if r.WhenDateTime == nil {
		return fmt.Errorf("field 'WhenDateTime' is required")
	}
	if r.WhenPeriod == nil {
		return fmt.Errorf("field 'WhenPeriod' is required")
	}
	if r.WhenPeriod != nil {
		if err := r.WhenPeriod.Validate(); err != nil {
			return fmt.Errorf("WhenPeriod: %w", err)
		}
	}
	return nil
}

type CoverageEligibilityRequestSupportingInfo struct {
	Id           *string    `json:"id,omitempty" bson:"id,omitempty"`                       // Unique id for inter-element referencing
	Sequence     int        `json:"sequence" bson:"sequence"`                               // Information instance identifier
	Information  *Reference `json:"information" bson:"information"`                         // Data to be provided
	AppliesToAll *bool      `json:"appliesToAll,omitempty" bson:"applies_to_all,omitempty"` // Applies to all items
}

func (r *CoverageEligibilityRequestSupportingInfo) Validate() error {
	if r.Sequence == 0 {
		return fmt.Errorf("field 'Sequence' is required")
	}
	if r.Information == nil {
		return fmt.Errorf("field 'Information' is required")
	}
	if r.Information != nil {
		if err := r.Information.Validate(); err != nil {
			return fmt.Errorf("Information: %w", err)
		}
	}
	return nil
}
