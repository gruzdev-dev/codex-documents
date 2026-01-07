package models

import (
	"encoding/json"
	"fmt"
)

// Details of a Health Insurance product provided by an organization.
type InsuranceProduct struct {
	Id             *string                    `json:"id,omitempty" bson:"id,omitempty"`                          // Logical id of this artifact
	Meta           *Meta                      `json:"meta,omitempty" bson:"meta,omitempty"`                      // Metadata about the resource
	ImplicitRules  *string                    `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`   // A set of rules under which this content was created
	Language       *string                    `json:"language,omitempty" bson:"language,omitempty"`              // Language of the resource content
	Text           *Narrative                 `json:"text,omitempty" bson:"text,omitempty"`                      // Text summary of the resource, for human interpretation
	Contained      []json.RawMessage          `json:"contained,omitempty" bson:"contained,omitempty"`            // Contained, inline Resources
	Identifier     []Identifier               `json:"identifier,omitempty" bson:"identifier,omitempty"`          // Business Identifier for Product
	Status         *string                    `json:"status,omitempty" bson:"status,omitempty"`                  // draft | active | retired | unknown
	Type           []CodeableConcept          `json:"type,omitempty" bson:"type,omitempty"`                      // Kind of product
	Name           *string                    `json:"name,omitempty" bson:"name,omitempty"`                      // Official name
	Alias          []string                   `json:"alias,omitempty" bson:"alias,omitempty"`                    // Alternate names
	Period         *Period                    `json:"period,omitempty" bson:"period,omitempty"`                  // When the product is available
	OwnedBy        *Reference                 `json:"ownedBy,omitempty" bson:"owned_by,omitempty"`               // Product issuer
	AdministeredBy *Reference                 `json:"administeredBy,omitempty" bson:"administered_by,omitempty"` // Product administrator
	CoverageArea   []Reference                `json:"coverageArea,omitempty" bson:"coverage_area,omitempty"`     // Where product applies
	Contact        []ExtendedContactDetail    `json:"contact,omitempty" bson:"contact,omitempty"`                // Official contact details relevant to the health insurance product
	Endpoint       []Reference                `json:"endpoint,omitempty" bson:"endpoint,omitempty"`              // Technical endpoint
	Network        []Reference                `json:"network,omitempty" bson:"network,omitempty"`                // What networks are Included
	Coverage       []InsuranceProductCoverage `json:"coverage,omitempty" bson:"coverage,omitempty"`              // Coverage details
	Related        []InsuranceProductRelated  `json:"related,omitempty" bson:"related,omitempty"`                // Associated insurance product
}

func (r *InsuranceProduct) Validate() error {
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
	for i, item := range r.Type {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Type[%d]: %w", i, err)
		}
	}
	if r.Period != nil {
		if err := r.Period.Validate(); err != nil {
			return fmt.Errorf("Period: %w", err)
		}
	}
	if r.OwnedBy != nil {
		if err := r.OwnedBy.Validate(); err != nil {
			return fmt.Errorf("OwnedBy: %w", err)
		}
	}
	if r.AdministeredBy != nil {
		if err := r.AdministeredBy.Validate(); err != nil {
			return fmt.Errorf("AdministeredBy: %w", err)
		}
	}
	for i, item := range r.CoverageArea {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("CoverageArea[%d]: %w", i, err)
		}
	}
	for i, item := range r.Contact {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Contact[%d]: %w", i, err)
		}
	}
	for i, item := range r.Endpoint {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Endpoint[%d]: %w", i, err)
		}
	}
	for i, item := range r.Network {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Network[%d]: %w", i, err)
		}
	}
	for i, item := range r.Coverage {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Coverage[%d]: %w", i, err)
		}
	}
	for i, item := range r.Related {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Related[%d]: %w", i, err)
		}
	}
	return nil
}

type InsuranceProductCoverageBenefit struct {
	Id          *string                                `json:"id,omitempty" bson:"id,omitempty"`                   // Unique id for inter-element referencing
	Type        *CodeableConcept                       `json:"type" bson:"type"`                                   // Classification of benefit provided
	Requirement *string                                `json:"requirement,omitempty" bson:"requirement,omitempty"` // Referral requirements
	Limit       []InsuranceProductCoverageBenefitLimit `json:"limit,omitempty" bson:"limit,omitempty"`             // Limits on the provided benefits
}

func (r *InsuranceProductCoverageBenefit) Validate() error {
	if r.Type == nil {
		return fmt.Errorf("field 'Type' is required")
	}
	if r.Type != nil {
		if err := r.Type.Validate(); err != nil {
			return fmt.Errorf("Type: %w", err)
		}
	}
	for i, item := range r.Limit {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Limit[%d]: %w", i, err)
		}
	}
	return nil
}

type InsuranceProductCoverageBenefitLimit struct {
	Id    *string          `json:"id,omitempty" bson:"id,omitempty"`       // Unique id for inter-element referencing
	Value *Quantity        `json:"value,omitempty" bson:"value,omitempty"` // Maximum value allowed
	Code  *CodeableConcept `json:"code,omitempty" bson:"code,omitempty"`   // Benefit limit details
}

func (r *InsuranceProductCoverageBenefitLimit) Validate() error {
	if r.Value != nil {
		if err := r.Value.Validate(); err != nil {
			return fmt.Errorf("Value: %w", err)
		}
	}
	if r.Code != nil {
		if err := r.Code.Validate(); err != nil {
			return fmt.Errorf("Code: %w", err)
		}
	}
	return nil
}

type InsuranceProductRelated struct {
	Id           *string          `json:"id,omitempty" bson:"id,omitempty"`                     // Unique id for inter-element referencing
	Product      *Reference       `json:"product,omitempty" bson:"product,omitempty"`           // Related Product reference
	Relationship *CodeableConcept `json:"relationship,omitempty" bson:"relationship,omitempty"` // Relationship of this product to the related product
	Period       *Period          `json:"period,omitempty" bson:"period,omitempty"`             // Period that this Relationship is valid
}

func (r *InsuranceProductRelated) Validate() error {
	if r.Product != nil {
		if err := r.Product.Validate(); err != nil {
			return fmt.Errorf("Product: %w", err)
		}
	}
	if r.Relationship != nil {
		if err := r.Relationship.Validate(); err != nil {
			return fmt.Errorf("Relationship: %w", err)
		}
	}
	if r.Period != nil {
		if err := r.Period.Validate(); err != nil {
			return fmt.Errorf("Period: %w", err)
		}
	}
	return nil
}

type InsuranceProductCoverage struct {
	Id      *string                           `json:"id,omitempty" bson:"id,omitempty"`           // Unique id for inter-element referencing
	Type    *CodeableConcept                  `json:"type" bson:"type"`                           // Classification of Coverage
	Network []Reference                       `json:"network,omitempty" bson:"network,omitempty"` // What networks provide coverage
	Benefit []InsuranceProductCoverageBenefit `json:"benefit" bson:"benefit"`                     // List of benefits
}

func (r *InsuranceProductCoverage) Validate() error {
	if r.Type == nil {
		return fmt.Errorf("field 'Type' is required")
	}
	if r.Type != nil {
		if err := r.Type.Validate(); err != nil {
			return fmt.Errorf("Type: %w", err)
		}
	}
	for i, item := range r.Network {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Network[%d]: %w", i, err)
		}
	}
	if len(r.Benefit) < 1 {
		return fmt.Errorf("field 'Benefit' must have at least 1 elements")
	}
	for i, item := range r.Benefit {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Benefit[%d]: %w", i, err)
		}
	}
	return nil
}
