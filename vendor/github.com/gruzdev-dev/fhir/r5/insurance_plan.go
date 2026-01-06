package models

import (
	"encoding/json"
	"fmt"
)

// Details of a Health Insurance plan provided by an organization under an InsuranceProduct.
type InsurancePlan struct {
	Id            *string                     `json:"id,omitempty" bson:"id,omitempty"`                        // Logical id of this artifact
	Meta          *Meta                       `json:"meta,omitempty" bson:"meta,omitempty"`                    // Metadata about the resource
	ImplicitRules *string                     `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"` // A set of rules under which this content was created
	Language      *string                     `json:"language,omitempty" bson:"language,omitempty"`            // Language of the resource content
	Text          *Narrative                  `json:"text,omitempty" bson:"text,omitempty"`                    // Text summary of the resource, for human interpretation
	Contained     []json.RawMessage           `json:"contained,omitempty" bson:"contained,omitempty"`          // Contained, inline Resources
	Identifier    []Identifier                `json:"identifier,omitempty" bson:"identifier,omitempty"`        // Business Identifier for Plan
	Type          *CodeableConcept            `json:"type,omitempty" bson:"type,omitempty"`                    // Classification of Plan
	Product       *Reference                  `json:"product,omitempty" bson:"product,omitempty"`              // The product that this plan is available under
	CoverageArea  []Reference                 `json:"coverageArea,omitempty" bson:"coverage_area,omitempty"`   // Where product-plan applies
	Network       []Reference                 `json:"network,omitempty" bson:"network,omitempty"`              // What networks provide coverage
	GeneralCost   []InsurancePlanGeneralCost  `json:"generalCost,omitempty" bson:"general_cost,omitempty"`     // Overall costs
	SpecificCost  []InsurancePlanSpecificCost `json:"specificCost,omitempty" bson:"specific_cost,omitempty"`   // Individual cost elements
}

func (r *InsurancePlan) Validate() error {
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
	if r.Type != nil {
		if err := r.Type.Validate(); err != nil {
			return fmt.Errorf("Type: %w", err)
		}
	}
	if r.Product != nil {
		if err := r.Product.Validate(); err != nil {
			return fmt.Errorf("Product: %w", err)
		}
	}
	for i, item := range r.CoverageArea {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("CoverageArea[%d]: %w", i, err)
		}
	}
	for i, item := range r.Network {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Network[%d]: %w", i, err)
		}
	}
	for i, item := range r.GeneralCost {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("GeneralCost[%d]: %w", i, err)
		}
	}
	for i, item := range r.SpecificCost {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("SpecificCost[%d]: %w", i, err)
		}
	}
	return nil
}

type InsurancePlanSpecificCost struct {
	Id       *string                            `json:"id,omitempty" bson:"id,omitempty"`           // Unique id for inter-element referencing
	Category *CodeableConcept                   `json:"category" bson:"category"`                   // General category of benefit
	Benefit  []InsurancePlanSpecificCostBenefit `json:"benefit,omitempty" bson:"benefit,omitempty"` // Benefits list
}

func (r *InsurancePlanSpecificCost) Validate() error {
	if r.Category == nil {
		return fmt.Errorf("field 'Category' is required")
	}
	if r.Category != nil {
		if err := r.Category.Validate(); err != nil {
			return fmt.Errorf("Category: %w", err)
		}
	}
	for i, item := range r.Benefit {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Benefit[%d]: %w", i, err)
		}
	}
	return nil
}

type InsurancePlanSpecificCostBenefit struct {
	Id   *string                                `json:"id,omitempty" bson:"id,omitempty"`     // Unique id for inter-element referencing
	Type *CodeableConcept                       `json:"type" bson:"type"`                     // Classification of benefit provided
	Cost []InsurancePlanSpecificCostBenefitCost `json:"cost,omitempty" bson:"cost,omitempty"` // List of the costs
}

func (r *InsurancePlanSpecificCostBenefit) Validate() error {
	if r.Type == nil {
		return fmt.Errorf("field 'Type' is required")
	}
	if r.Type != nil {
		if err := r.Type.Validate(); err != nil {
			return fmt.Errorf("Type: %w", err)
		}
	}
	for i, item := range r.Cost {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Cost[%d]: %w", i, err)
		}
	}
	return nil
}

type InsurancePlanSpecificCostBenefitCost struct {
	Id            *string           `json:"id,omitempty" bson:"id,omitempty"`                       // Unique id for inter-element referencing
	Type          *CodeableConcept  `json:"type" bson:"type"`                                       // Classification of specific cost
	Applicability *CodeableConcept  `json:"applicability,omitempty" bson:"applicability,omitempty"` // in-network | out-of-network | other
	Qualifier     []CodeableConcept `json:"qualifier,omitempty" bson:"qualifier,omitempty"`         // Additional information about the cost
	Value         *Quantity         `json:"value,omitempty" bson:"value,omitempty"`                 // The actual cost value
}

func (r *InsurancePlanSpecificCostBenefitCost) Validate() error {
	if r.Type == nil {
		return fmt.Errorf("field 'Type' is required")
	}
	if r.Type != nil {
		if err := r.Type.Validate(); err != nil {
			return fmt.Errorf("Type: %w", err)
		}
	}
	if r.Applicability != nil {
		if err := r.Applicability.Validate(); err != nil {
			return fmt.Errorf("Applicability: %w", err)
		}
	}
	for i, item := range r.Qualifier {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Qualifier[%d]: %w", i, err)
		}
	}
	if r.Value != nil {
		if err := r.Value.Validate(); err != nil {
			return fmt.Errorf("Value: %w", err)
		}
	}
	return nil
}

type InsurancePlanGeneralCost struct {
	Id        *string          `json:"id,omitempty" bson:"id,omitempty"`                // Unique id for inter-element referencing
	Type      *CodeableConcept `json:"type,omitempty" bson:"type,omitempty"`            // Classification of specific cost
	GroupSize *int             `json:"groupSize,omitempty" bson:"group_size,omitempty"` // Number of enrollees
	Cost      *Money           `json:"cost,omitempty" bson:"cost,omitempty"`            // Cost value
	Comment   *string          `json:"comment,omitempty" bson:"comment,omitempty"`      // Additional cost information
}

func (r *InsurancePlanGeneralCost) Validate() error {
	if r.Type != nil {
		if err := r.Type.Validate(); err != nil {
			return fmt.Errorf("Type: %w", err)
		}
	}
	if r.Cost != nil {
		if err := r.Cost.Validate(); err != nil {
			return fmt.Errorf("Cost: %w", err)
		}
	}
	return nil
}
