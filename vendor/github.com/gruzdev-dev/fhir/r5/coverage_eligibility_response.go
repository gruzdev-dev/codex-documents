package models

import (
	"encoding/json"
	"fmt"
)

// This resource provides eligibility and plan details from the processing of an CoverageEligibilityRequest resource.
type CoverageEligibilityResponse struct {
	Id             *string                                `json:"id,omitempty" bson:"id,omitempty"`                          // Logical id of this artifact
	Meta           *Meta                                  `json:"meta,omitempty" bson:"meta,omitempty"`                      // Metadata about the resource
	ImplicitRules  *string                                `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`   // A set of rules under which this content was created
	Language       *string                                `json:"language,omitempty" bson:"language,omitempty"`              // Language of the resource content
	Text           *Narrative                             `json:"text,omitempty" bson:"text,omitempty"`                      // Text summary of the resource, for human interpretation
	Contained      []json.RawMessage                      `json:"contained,omitempty" bson:"contained,omitempty"`            // Contained, inline Resources
	Identifier     []Identifier                           `json:"identifier,omitempty" bson:"identifier,omitempty"`          // Business Identifier for coverage eligiblity request
	Status         string                                 `json:"status" bson:"status"`                                      // active | cancelled | draft | entered-in-error
	StatusReason   *string                                `json:"statusReason,omitempty" bson:"status_reason,omitempty"`     // Reason for status change
	Purpose        []string                               `json:"purpose" bson:"purpose"`                                    // auth-requirements | benefits | discovery | validation
	Patient        *Reference                             `json:"patient" bson:"patient"`                                    // Intended recipient of products and services
	Event          []CoverageEligibilityResponseEvent     `json:"event,omitempty" bson:"event,omitempty"`                    // Event information
	ServicedDate   *string                                `json:"servicedDate,omitempty" bson:"serviced_date,omitempty"`     // Estimated date or dates of service
	ServicedPeriod *Period                                `json:"servicedPeriod,omitempty" bson:"serviced_period,omitempty"` // Estimated date or dates of service
	Created        string                                 `json:"created" bson:"created"`                                    // Response creation date
	Requestor      *Reference                             `json:"requestor,omitempty" bson:"requestor,omitempty"`            // Party responsible for the request
	Request        *Reference                             `json:"request,omitempty" bson:"request,omitempty"`                // Eligibility request reference
	Outcome        string                                 `json:"outcome" bson:"outcome"`                                    // queued | complete | error | partial
	Disposition    *string                                `json:"disposition,omitempty" bson:"disposition,omitempty"`        // Disposition Message
	Insurer        *Reference                             `json:"insurer" bson:"insurer"`                                    // Coverage issuer
	Insurance      []CoverageEligibilityResponseInsurance `json:"insurance,omitempty" bson:"insurance,omitempty"`            // Patient insurance information
	PreAuthRef     *string                                `json:"preAuthRef,omitempty" bson:"pre_auth_ref,omitempty"`        // Preauthorization reference
	Form           *CodeableConcept                       `json:"form,omitempty" bson:"form,omitempty"`                      // Printed form identifier
	Error          []CoverageEligibilityResponseError     `json:"error,omitempty" bson:"error,omitempty"`                    // Processing errors
}

func (r *CoverageEligibilityResponse) Validate() error {
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
	if r.Requestor != nil {
		if err := r.Requestor.Validate(); err != nil {
			return fmt.Errorf("Requestor: %w", err)
		}
	}
	if r.Request != nil {
		if err := r.Request.Validate(); err != nil {
			return fmt.Errorf("Request: %w", err)
		}
	}
	if r.Outcome == emptyString {
		return fmt.Errorf("field 'Outcome' is required")
	}
	if r.Insurer == nil {
		return fmt.Errorf("field 'Insurer' is required")
	}
	if r.Insurer != nil {
		if err := r.Insurer.Validate(); err != nil {
			return fmt.Errorf("Insurer: %w", err)
		}
	}
	for i, item := range r.Insurance {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Insurance[%d]: %w", i, err)
		}
	}
	if r.Form != nil {
		if err := r.Form.Validate(); err != nil {
			return fmt.Errorf("Form: %w", err)
		}
	}
	for i, item := range r.Error {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Error[%d]: %w", i, err)
		}
	}
	return nil
}

type CoverageEligibilityResponseError struct {
	Id         *string          `json:"id,omitempty" bson:"id,omitempty"`                 // Unique id for inter-element referencing
	Code       *CodeableConcept `json:"code" bson:"code"`                                 // Error code detailing processing issues
	Expression []string         `json:"expression,omitempty" bson:"expression,omitempty"` // FHIRPath of element(s) related to issue
}

func (r *CoverageEligibilityResponseError) Validate() error {
	if r.Code == nil {
		return fmt.Errorf("field 'Code' is required")
	}
	if r.Code != nil {
		if err := r.Code.Validate(); err != nil {
			return fmt.Errorf("Code: %w", err)
		}
	}
	return nil
}

type CoverageEligibilityResponseEvent struct {
	Id           *string          `json:"id,omitempty" bson:"id,omitempty"`   // Unique id for inter-element referencing
	Type         *CodeableConcept `json:"type" bson:"type"`                   // Specific event
	WhenDateTime *string          `json:"whenDateTime" bson:"when_date_time"` // Occurance date or period
	WhenPeriod   *Period          `json:"whenPeriod" bson:"when_period"`      // Occurance date or period
}

func (r *CoverageEligibilityResponseEvent) Validate() error {
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

type CoverageEligibilityResponseInsurance struct {
	Id            *string                                    `json:"id,omitempty" bson:"id,omitempty"`                        // Unique id for inter-element referencing
	Coverage      *Reference                                 `json:"coverage" bson:"coverage"`                                // Insurance information
	Inforce       bool                                       `json:"inforce,omitempty" bson:"inforce,omitempty"`              // Coverage inforce indicator
	BenefitPeriod *Period                                    `json:"benefitPeriod,omitempty" bson:"benefit_period,omitempty"` // When the benefits are applicable
	Item          []CoverageEligibilityResponseInsuranceItem `json:"item,omitempty" bson:"item,omitempty"`                    // Benefits and authorization details
}

func (r *CoverageEligibilityResponseInsurance) Validate() error {
	if r.Coverage == nil {
		return fmt.Errorf("field 'Coverage' is required")
	}
	if r.Coverage != nil {
		if err := r.Coverage.Validate(); err != nil {
			return fmt.Errorf("Coverage: %w", err)
		}
	}
	if r.BenefitPeriod != nil {
		if err := r.BenefitPeriod.Validate(); err != nil {
			return fmt.Errorf("BenefitPeriod: %w", err)
		}
	}
	for i, item := range r.Item {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Item[%d]: %w", i, err)
		}
	}
	return nil
}

type CoverageEligibilityResponseInsuranceItem struct {
	Id                      *string                                           `json:"id,omitempty" bson:"id,omitempty"`                                            // Unique id for inter-element referencing
	Category                *CodeableConcept                                  `json:"category,omitempty" bson:"category,omitempty"`                                // Benefit classification
	ProductOrService        *CodeableConcept                                  `json:"productOrService,omitempty" bson:"product_or_service,omitempty"`              // Billing, service, product, or drug code
	Modifier                []CodeableConcept                                 `json:"modifier,omitempty" bson:"modifier,omitempty"`                                // Product or service billing modifiers
	Provider                *Reference                                        `json:"provider,omitempty" bson:"provider,omitempty"`                                // Performing practitioner
	Excluded                bool                                              `json:"excluded,omitempty" bson:"excluded,omitempty"`                                // Excluded from the plan
	Name                    *string                                           `json:"name,omitempty" bson:"name,omitempty"`                                        // Short name for the benefit
	Description             *string                                           `json:"description,omitempty" bson:"description,omitempty"`                          // Description of the benefit or services covered
	Network                 *CodeableConcept                                  `json:"network,omitempty" bson:"network,omitempty"`                                  // In or out of network
	Unit                    *CodeableConcept                                  `json:"unit,omitempty" bson:"unit,omitempty"`                                        // Individual or family
	Term                    *CodeableConcept                                  `json:"term,omitempty" bson:"term,omitempty"`                                        // Annual or lifetime
	Benefit                 []CoverageEligibilityResponseInsuranceItemBenefit `json:"benefit,omitempty" bson:"benefit,omitempty"`                                  // Benefit Summary
	AuthorizationRequired   bool                                              `json:"authorizationRequired,omitempty" bson:"authorization_required,omitempty"`     // Authorization required flag
	AuthorizationSupporting []CodeableConcept                                 `json:"authorizationSupporting,omitempty" bson:"authorization_supporting,omitempty"` // Type of required supporting materials
	AuthorizationUrl        *string                                           `json:"authorizationUrl,omitempty" bson:"authorization_url,omitempty"`               // Preauthorization requirements endpoint
}

func (r *CoverageEligibilityResponseInsuranceItem) Validate() error {
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
	if r.Network != nil {
		if err := r.Network.Validate(); err != nil {
			return fmt.Errorf("Network: %w", err)
		}
	}
	if r.Unit != nil {
		if err := r.Unit.Validate(); err != nil {
			return fmt.Errorf("Unit: %w", err)
		}
	}
	if r.Term != nil {
		if err := r.Term.Validate(); err != nil {
			return fmt.Errorf("Term: %w", err)
		}
	}
	for i, item := range r.Benefit {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Benefit[%d]: %w", i, err)
		}
	}
	for i, item := range r.AuthorizationSupporting {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("AuthorizationSupporting[%d]: %w", i, err)
		}
	}
	return nil
}

type CoverageEligibilityResponseInsuranceItemBenefit struct {
	Id                 *string          `json:"id,omitempty" bson:"id,omitempty"`                                   // Unique id for inter-element referencing
	Type               *CodeableConcept `json:"type" bson:"type"`                                                   // Benefit classification
	AllowedUnsignedInt *int             `json:"allowedUnsignedInt,omitempty" bson:"allowed_unsigned_int,omitempty"` // Benefits allowed
	AllowedString      *string          `json:"allowedString,omitempty" bson:"allowed_string,omitempty"`            // Benefits allowed
	AllowedMoney       *Money           `json:"allowedMoney,omitempty" bson:"allowed_money,omitempty"`              // Benefits allowed
	UsedUnsignedInt    *int             `json:"usedUnsignedInt,omitempty" bson:"used_unsigned_int,omitempty"`       // Benefits used
	UsedString         *string          `json:"usedString,omitempty" bson:"used_string,omitempty"`                  // Benefits used
	UsedMoney          *Money           `json:"usedMoney,omitempty" bson:"used_money,omitempty"`                    // Benefits used
}

func (r *CoverageEligibilityResponseInsuranceItemBenefit) Validate() error {
	if r.Type == nil {
		return fmt.Errorf("field 'Type' is required")
	}
	if r.Type != nil {
		if err := r.Type.Validate(); err != nil {
			return fmt.Errorf("Type: %w", err)
		}
	}
	if r.AllowedMoney != nil {
		if err := r.AllowedMoney.Validate(); err != nil {
			return fmt.Errorf("AllowedMoney: %w", err)
		}
	}
	if r.UsedMoney != nil {
		if err := r.UsedMoney.Validate(); err != nil {
			return fmt.Errorf("UsedMoney: %w", err)
		}
	}
	return nil
}
