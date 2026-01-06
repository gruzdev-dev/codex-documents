package models

import (
	"encoding/json"
	"fmt"
)

// The definition and characteristics of a medicinal manufactured item, such as a tablet or capsule, as contained in a packaged medicinal product.
type ManufacturedItemDefinition struct {
	Id                   *string                               `json:"id,omitempty" bson:"id,omitempty"`                                   // Logical id of this artifact
	Meta                 *Meta                                 `json:"meta,omitempty" bson:"meta,omitempty"`                               // Metadata about the resource
	ImplicitRules        *string                               `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`            // A set of rules under which this content was created
	Language             *string                               `json:"language,omitempty" bson:"language,omitempty"`                       // Language of the resource content
	Text                 *Narrative                            `json:"text,omitempty" bson:"text,omitempty"`                               // Text summary of the resource, for human interpretation
	Contained            []json.RawMessage                     `json:"contained,omitempty" bson:"contained,omitempty"`                     // Contained, inline Resources
	Identifier           []Identifier                          `json:"identifier,omitempty" bson:"identifier,omitempty"`                   // Unique identifier
	Status               string                                `json:"status" bson:"status"`                                               // draft | active | retired | unknown
	Name                 *string                               `json:"name,omitempty" bson:"name,omitempty"`                               // A descriptive name applied to this item
	ManufacturedDoseForm *CodeableConcept                      `json:"manufacturedDoseForm" bson:"manufactured_dose_form"`                 // Dose form as manufactured (before any necessary transformation)
	UnitOfPresentation   *CodeableConcept                      `json:"unitOfPresentation,omitempty" bson:"unit_of_presentation,omitempty"` // The “real-world” units in which the quantity of the item is described
	Manufacturer         []Reference                           `json:"manufacturer,omitempty" bson:"manufacturer,omitempty"`               // Manufacturer of the item, one of several possible
	MarketingStatus      []MarketingStatus                     `json:"marketingStatus,omitempty" bson:"marketing_status,omitempty"`        // Allows specifying that an item is on the market for sale, or that it is not available, and the dates and locations associated
	Ingredient           []CodeableConcept                     `json:"ingredient,omitempty" bson:"ingredient,omitempty"`                   // The ingredients of this manufactured item. Only needed if these are not specified by incoming references from the Ingredient resource
	Property             []ManufacturedItemDefinitionProperty  `json:"property,omitempty" bson:"property,omitempty"`                       // General characteristics of this item
	Component            []ManufacturedItemDefinitionComponent `json:"component,omitempty" bson:"component,omitempty"`                     // Physical parts of the manufactured item, that it is intrinsically made from. This is distinct from the ingredients that are part of its chemical makeup
}

func (r *ManufacturedItemDefinition) Validate() error {
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
	if r.ManufacturedDoseForm == nil {
		return fmt.Errorf("field 'ManufacturedDoseForm' is required")
	}
	if r.ManufacturedDoseForm != nil {
		if err := r.ManufacturedDoseForm.Validate(); err != nil {
			return fmt.Errorf("ManufacturedDoseForm: %w", err)
		}
	}
	if r.UnitOfPresentation != nil {
		if err := r.UnitOfPresentation.Validate(); err != nil {
			return fmt.Errorf("UnitOfPresentation: %w", err)
		}
	}
	for i, item := range r.Manufacturer {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Manufacturer[%d]: %w", i, err)
		}
	}
	for i, item := range r.MarketingStatus {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("MarketingStatus[%d]: %w", i, err)
		}
	}
	for i, item := range r.Ingredient {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Ingredient[%d]: %w", i, err)
		}
	}
	for i, item := range r.Property {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Property[%d]: %w", i, err)
		}
	}
	for i, item := range r.Component {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Component[%d]: %w", i, err)
		}
	}
	return nil
}

type ManufacturedItemDefinitionProperty struct {
	Id                   *string          `json:"id,omitempty" bson:"id,omitempty"`                                       // Unique id for inter-element referencing
	Type                 *CodeableConcept `json:"type" bson:"type"`                                                       // A code expressing the type of characteristic
	ValueCodeableConcept *CodeableConcept `json:"valueCodeableConcept,omitempty" bson:"value_codeable_concept,omitempty"` // A value for the characteristic
	ValueQuantity        *Quantity        `json:"valueQuantity,omitempty" bson:"value_quantity,omitempty"`                // A value for the characteristic
	ValueRange           *Range           `json:"valueRange,omitempty" bson:"value_range,omitempty"`                      // A value for the characteristic
	ValueDate            *string          `json:"valueDate,omitempty" bson:"value_date,omitempty"`                        // A value for the characteristic
	ValueBoolean         *bool            `json:"valueBoolean,omitempty" bson:"value_boolean,omitempty"`                  // A value for the characteristic
	ValueMarkdown        *string          `json:"valueMarkdown,omitempty" bson:"value_markdown,omitempty"`                // A value for the characteristic
	ValueAttachment      *Attachment      `json:"valueAttachment,omitempty" bson:"value_attachment,omitempty"`            // A value for the characteristic
	ValueReference       *Reference       `json:"valueReference,omitempty" bson:"value_reference,omitempty"`              // A value for the characteristic
}

func (r *ManufacturedItemDefinitionProperty) Validate() error {
	if r.Type == nil {
		return fmt.Errorf("field 'Type' is required")
	}
	if r.Type != nil {
		if err := r.Type.Validate(); err != nil {
			return fmt.Errorf("Type: %w", err)
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
	if r.ValueAttachment != nil {
		if err := r.ValueAttachment.Validate(); err != nil {
			return fmt.Errorf("ValueAttachment: %w", err)
		}
	}
	if r.ValueReference != nil {
		if err := r.ValueReference.Validate(); err != nil {
			return fmt.Errorf("ValueReference: %w", err)
		}
	}
	return nil
}

type ManufacturedItemDefinitionComponent struct {
	Id          *string                                          `json:"id,omitempty" bson:"id,omitempty"`                   // Unique id for inter-element referencing
	Type        *CodeableConcept                                 `json:"type" bson:"type"`                                   // Defining type of the component e.g. shell, layer, ink
	Function    []CodeableConcept                                `json:"function,omitempty" bson:"function,omitempty"`       // The function of this component within the item e.g. delivers active ingredient, masks taste
	Amount      []Quantity                                       `json:"amount,omitempty" bson:"amount,omitempty"`           // The measurable amount of total quantity of all substances in the component, expressible in different ways (e.g. by mass or volume)
	Constituent []ManufacturedItemDefinitionComponentConstituent `json:"constituent,omitempty" bson:"constituent,omitempty"` // A reference to a constituent of the manufactured item as a whole, linked here so that its component location within the item can be indicated. This not where the item's ingredient are primarily stated (for which see Ingredient.for or ManufacturedItemDefinition.ingredient)
	Property    []ManufacturedItemDefinitionProperty             `json:"property,omitempty" bson:"property,omitempty"`       // General characteristics of this component
	Component   []ManufacturedItemDefinitionComponent            `json:"component,omitempty" bson:"component,omitempty"`     // A component that this component contains or is made from
}

func (r *ManufacturedItemDefinitionComponent) Validate() error {
	if r.Type == nil {
		return fmt.Errorf("field 'Type' is required")
	}
	if r.Type != nil {
		if err := r.Type.Validate(); err != nil {
			return fmt.Errorf("Type: %w", err)
		}
	}
	for i, item := range r.Function {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Function[%d]: %w", i, err)
		}
	}
	for i, item := range r.Amount {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Amount[%d]: %w", i, err)
		}
	}
	for i, item := range r.Constituent {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Constituent[%d]: %w", i, err)
		}
	}
	for i, item := range r.Property {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Property[%d]: %w", i, err)
		}
	}
	for i, item := range r.Component {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Component[%d]: %w", i, err)
		}
	}
	return nil
}

type ManufacturedItemDefinitionComponentConstituent struct {
	Id            *string             `json:"id,omitempty" bson:"id,omitempty"`                        // Unique id for inter-element referencing
	Amount        []Quantity          `json:"amount,omitempty" bson:"amount,omitempty"`                // The measurable amount of the substance, expressible in different ways (e.g. by mass or volume)
	Location      []CodeableConcept   `json:"location,omitempty" bson:"location,omitempty"`            // The physical location of the constituent/ingredient within the component
	Function      []CodeableConcept   `json:"function,omitempty" bson:"function,omitempty"`            // The function of this constituent within the component e.g. binder
	HasIngredient []CodeableReference `json:"hasIngredient,omitempty" bson:"has_ingredient,omitempty"` // The ingredient that is the constituent of the given component
}

func (r *ManufacturedItemDefinitionComponentConstituent) Validate() error {
	for i, item := range r.Amount {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Amount[%d]: %w", i, err)
		}
	}
	for i, item := range r.Location {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Location[%d]: %w", i, err)
		}
	}
	for i, item := range r.Function {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Function[%d]: %w", i, err)
		}
	}
	for i, item := range r.HasIngredient {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("HasIngredient[%d]: %w", i, err)
		}
	}
	return nil
}
