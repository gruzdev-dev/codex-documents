package models

import (
	"encoding/json"
	"fmt"
)

// An ingredient of a manufactured item or pharmaceutical product.
type Ingredient struct {
	ResourceType        string                   `json:"resourceType" bson:"resource_type"`                                   // Type of resource
	Id                  *string                  `json:"id,omitempty" bson:"id,omitempty"`                                    // Logical id of this artifact
	Meta                *Meta                    `json:"meta,omitempty" bson:"meta,omitempty"`                                // Metadata about the resource
	ImplicitRules       *string                  `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`             // A set of rules under which this content was created
	Language            *string                  `json:"language,omitempty" bson:"language,omitempty"`                        // Language of the resource content
	Text                *Narrative               `json:"text,omitempty" bson:"text,omitempty"`                                // Text summary of the resource, for human interpretation
	Contained           []json.RawMessage        `json:"contained,omitempty" bson:"contained,omitempty"`                      // Contained, inline Resources
	Identifier          *Identifier              `json:"identifier,omitempty" bson:"identifier,omitempty"`                    // An identifier or code by which the ingredient can be referenced
	Status              string                   `json:"status" bson:"status"`                                                // draft | active | retired | unknown
	For                 []Reference              `json:"for,omitempty" bson:"for,omitempty"`                                  // The product which this ingredient is a constituent part of
	Role                *CodeableConcept         `json:"role" bson:"role"`                                                    // Purpose of the ingredient within the product, e.g. active, inactive
	Function            []CodeableConcept        `json:"function,omitempty" bson:"function,omitempty"`                        // Precise action within the drug product, e.g. antioxidant, alkalizing agent
	Group               *CodeableConcept         `json:"group,omitempty" bson:"group,omitempty"`                              // A classification of the ingredient according to where in the physical item it tends to be used, such the outer shell of a tablet, inner body or ink
	AllergenicIndicator *bool                    `json:"allergenicIndicator,omitempty" bson:"allergenic_indicator,omitempty"` // If the ingredient is a known or suspected allergen
	Comment             *string                  `json:"comment,omitempty" bson:"comment,omitempty"`                          // A place for providing any notes that are relevant to the component, e.g. removed during process, adjusted for loss on drying
	Manufacturer        []IngredientManufacturer `json:"manufacturer,omitempty" bson:"manufacturer,omitempty"`                // An organization that manufactures this ingredient
	Substance           *IngredientSubstance     `json:"substance" bson:"substance"`                                          // The substance that comprises this ingredient
}

func (r *Ingredient) Validate() error {
	if r.ResourceType != "Ingredient" {
		return fmt.Errorf("invalid resourceType: expected 'Ingredient', got '%s'", r.ResourceType)
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
	if r.Identifier != nil {
		if err := r.Identifier.Validate(); err != nil {
			return fmt.Errorf("Identifier: %w", err)
		}
	}
	var emptyString string
	if r.Status == emptyString {
		return fmt.Errorf("field 'Status' is required")
	}
	for i, item := range r.For {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("For[%d]: %w", i, err)
		}
	}
	if r.Role == nil {
		return fmt.Errorf("field 'Role' is required")
	}
	if r.Role != nil {
		if err := r.Role.Validate(); err != nil {
			return fmt.Errorf("Role: %w", err)
		}
	}
	for i, item := range r.Function {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Function[%d]: %w", i, err)
		}
	}
	if r.Group != nil {
		if err := r.Group.Validate(); err != nil {
			return fmt.Errorf("Group: %w", err)
		}
	}
	for i, item := range r.Manufacturer {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Manufacturer[%d]: %w", i, err)
		}
	}
	if r.Substance == nil {
		return fmt.Errorf("field 'Substance' is required")
	}
	if r.Substance != nil {
		if err := r.Substance.Validate(); err != nil {
			return fmt.Errorf("Substance: %w", err)
		}
	}
	return nil
}

type IngredientManufacturer struct {
	Id           *string    `json:"id,omitempty" bson:"id,omitempty"`     // Unique id for inter-element referencing
	Role         *string    `json:"role,omitempty" bson:"role,omitempty"` // allowed | possible | actual
	Manufacturer *Reference `json:"manufacturer" bson:"manufacturer"`     // An organization that manufactures this ingredient
}

func (r *IngredientManufacturer) Validate() error {
	if r.Manufacturer == nil {
		return fmt.Errorf("field 'Manufacturer' is required")
	}
	if r.Manufacturer != nil {
		if err := r.Manufacturer.Validate(); err != nil {
			return fmt.Errorf("Manufacturer: %w", err)
		}
	}
	return nil
}

type IngredientSubstance struct {
	Id       *string                       `json:"id,omitempty" bson:"id,omitempty"`             // Unique id for inter-element referencing
	Code     *CodeableReference            `json:"code" bson:"code"`                             // A code or full resource that represents the ingredient substance
	Strength []IngredientSubstanceStrength `json:"strength,omitempty" bson:"strength,omitempty"` // The quantity of substance, per presentation, or per volume or mass, and type of quantity
}

func (r *IngredientSubstance) Validate() error {
	if r.Code == nil {
		return fmt.Errorf("field 'Code' is required")
	}
	if r.Code != nil {
		if err := r.Code.Validate(); err != nil {
			return fmt.Errorf("Code: %w", err)
		}
	}
	for i, item := range r.Strength {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Strength[%d]: %w", i, err)
		}
	}
	return nil
}

type IngredientSubstanceStrength struct {
	Id                           *string                                        `json:"id,omitempty" bson:"id,omitempty"`                                                       // Unique id for inter-element referencing
	PresentationRatio            *Ratio                                         `json:"presentationRatio,omitempty" bson:"presentation_ratio,omitempty"`                        // The quantity of substance in the unit of presentation
	PresentationRatioRange       *RatioRange                                    `json:"presentationRatioRange,omitempty" bson:"presentation_ratio_range,omitempty"`             // The quantity of substance in the unit of presentation
	PresentationCodeableConcept  *CodeableConcept                               `json:"presentationCodeableConcept,omitempty" bson:"presentation_codeable_concept,omitempty"`   // The quantity of substance in the unit of presentation
	PresentationQuantity         *Quantity                                      `json:"presentationQuantity,omitempty" bson:"presentation_quantity,omitempty"`                  // The quantity of substance in the unit of presentation
	TextPresentation             *string                                        `json:"textPresentation,omitempty" bson:"text_presentation,omitempty"`                          // Text of either the whole presentation strength or a part of it (rest being in Strength.presentation as a ratio)
	ConcentrationRatio           *Ratio                                         `json:"concentrationRatio,omitempty" bson:"concentration_ratio,omitempty"`                      // The strength per unitary volume (or mass)
	ConcentrationRatioRange      *RatioRange                                    `json:"concentrationRatioRange,omitempty" bson:"concentration_ratio_range,omitempty"`           // The strength per unitary volume (or mass)
	ConcentrationCodeableConcept *CodeableConcept                               `json:"concentrationCodeableConcept,omitempty" bson:"concentration_codeable_concept,omitempty"` // The strength per unitary volume (or mass)
	ConcentrationQuantity        *Quantity                                      `json:"concentrationQuantity,omitempty" bson:"concentration_quantity,omitempty"`                // The strength per unitary volume (or mass)
	TextConcentration            *string                                        `json:"textConcentration,omitempty" bson:"text_concentration,omitempty"`                        // Text of either the whole concentration strength or a part of it (rest being in Strength.concentration as a ratio)
	Basis                        *CodeableConcept                               `json:"basis,omitempty" bson:"basis,omitempty"`                                                 // A code that indicates if the strength is, for example, based on the ingredient substance as stated or on the substance base (when the ingredient is a salt)
	MeasurementPoint             *string                                        `json:"measurementPoint,omitempty" bson:"measurement_point,omitempty"`                          // When strength is measured at a particular point or distance
	Country                      []CodeableConcept                              `json:"country,omitempty" bson:"country,omitempty"`                                             // Where the strength range applies
	ReferenceStrength            []IngredientSubstanceStrengthReferenceStrength `json:"referenceStrength,omitempty" bson:"reference_strength,omitempty"`                        // Strength expressed in terms of a reference substance
}

func (r *IngredientSubstanceStrength) Validate() error {
	if r.PresentationRatio != nil {
		if err := r.PresentationRatio.Validate(); err != nil {
			return fmt.Errorf("PresentationRatio: %w", err)
		}
	}
	if r.PresentationRatioRange != nil {
		if err := r.PresentationRatioRange.Validate(); err != nil {
			return fmt.Errorf("PresentationRatioRange: %w", err)
		}
	}
	if r.PresentationCodeableConcept != nil {
		if err := r.PresentationCodeableConcept.Validate(); err != nil {
			return fmt.Errorf("PresentationCodeableConcept: %w", err)
		}
	}
	if r.PresentationQuantity != nil {
		if err := r.PresentationQuantity.Validate(); err != nil {
			return fmt.Errorf("PresentationQuantity: %w", err)
		}
	}
	if r.ConcentrationRatio != nil {
		if err := r.ConcentrationRatio.Validate(); err != nil {
			return fmt.Errorf("ConcentrationRatio: %w", err)
		}
	}
	if r.ConcentrationRatioRange != nil {
		if err := r.ConcentrationRatioRange.Validate(); err != nil {
			return fmt.Errorf("ConcentrationRatioRange: %w", err)
		}
	}
	if r.ConcentrationCodeableConcept != nil {
		if err := r.ConcentrationCodeableConcept.Validate(); err != nil {
			return fmt.Errorf("ConcentrationCodeableConcept: %w", err)
		}
	}
	if r.ConcentrationQuantity != nil {
		if err := r.ConcentrationQuantity.Validate(); err != nil {
			return fmt.Errorf("ConcentrationQuantity: %w", err)
		}
	}
	if r.Basis != nil {
		if err := r.Basis.Validate(); err != nil {
			return fmt.Errorf("Basis: %w", err)
		}
	}
	for i, item := range r.Country {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Country[%d]: %w", i, err)
		}
	}
	for i, item := range r.ReferenceStrength {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("ReferenceStrength[%d]: %w", i, err)
		}
	}
	return nil
}

type IngredientSubstanceStrengthReferenceStrength struct {
	Id                 *string            `json:"id,omitempty" bson:"id,omitempty"`                              // Unique id for inter-element referencing
	Substance          *CodeableReference `json:"substance" bson:"substance"`                                    // Relevant reference substance
	StrengthRatio      *Ratio             `json:"strengthRatio" bson:"strength_ratio"`                           // Strength expressed in terms of a reference substance
	StrengthRatioRange *RatioRange        `json:"strengthRatioRange" bson:"strength_ratio_range"`                // Strength expressed in terms of a reference substance
	StrengthQuantity   *Quantity          `json:"strengthQuantity" bson:"strength_quantity"`                     // Strength expressed in terms of a reference substance
	MeasurementPoint   *string            `json:"measurementPoint,omitempty" bson:"measurement_point,omitempty"` // When strength is measured at a particular point or distance
	Country            []CodeableConcept  `json:"country,omitempty" bson:"country,omitempty"`                    // Where the strength range applies
}

func (r *IngredientSubstanceStrengthReferenceStrength) Validate() error {
	if r.Substance == nil {
		return fmt.Errorf("field 'Substance' is required")
	}
	if r.Substance != nil {
		if err := r.Substance.Validate(); err != nil {
			return fmt.Errorf("Substance: %w", err)
		}
	}
	if r.StrengthRatio == nil {
		return fmt.Errorf("field 'StrengthRatio' is required")
	}
	if r.StrengthRatio != nil {
		if err := r.StrengthRatio.Validate(); err != nil {
			return fmt.Errorf("StrengthRatio: %w", err)
		}
	}
	if r.StrengthRatioRange == nil {
		return fmt.Errorf("field 'StrengthRatioRange' is required")
	}
	if r.StrengthRatioRange != nil {
		if err := r.StrengthRatioRange.Validate(); err != nil {
			return fmt.Errorf("StrengthRatioRange: %w", err)
		}
	}
	if r.StrengthQuantity == nil {
		return fmt.Errorf("field 'StrengthQuantity' is required")
	}
	if r.StrengthQuantity != nil {
		if err := r.StrengthQuantity.Validate(); err != nil {
			return fmt.Errorf("StrengthQuantity: %w", err)
		}
	}
	for i, item := range r.Country {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Country[%d]: %w", i, err)
		}
	}
	return nil
}
