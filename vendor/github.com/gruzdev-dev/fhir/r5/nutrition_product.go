package models

import (
	"encoding/json"
	"fmt"
)

// A food or supplement that is consumed by patients.
type NutritionProduct struct {
	Id                *string                          `json:"id,omitempty" bson:"id,omitempty"`                                // Logical id of this artifact
	Meta              *Meta                            `json:"meta,omitempty" bson:"meta,omitempty"`                            // Metadata about the resource
	ImplicitRules     *string                          `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`         // A set of rules under which this content was created
	Language          *string                          `json:"language,omitempty" bson:"language,omitempty"`                    // Language of the resource content
	Text              *Narrative                       `json:"text,omitempty" bson:"text,omitempty"`                            // Text summary of the resource, for human interpretation
	Contained         []json.RawMessage                `json:"contained,omitempty" bson:"contained,omitempty"`                  // Contained, inline Resources
	Code              *CodeableConcept                 `json:"code,omitempty" bson:"code,omitempty"`                            // A code that can identify the product
	Status            string                           `json:"status" bson:"status"`                                            // active | inactive | entered-in-error
	Category          []CodeableConcept                `json:"category,omitempty" bson:"category,omitempty"`                    // Broad product groups, like Fruit, Grain, Beverages, or Vegetables Products
	Manufacturer      []Reference                      `json:"manufacturer,omitempty" bson:"manufacturer,omitempty"`            // Manufacturer, representative or person officially responsible for the product
	Nutrient          []NutritionProductNutrient       `json:"nutrient,omitempty" bson:"nutrient,omitempty"`                    // The product's nutritional information expressed by the nutrients
	IngredientSummary *string                          `json:"ingredientSummary,omitempty" bson:"ingredient_summary,omitempty"` // Textual description of product ingredients
	Ingredient        []NutritionProductIngredient     `json:"ingredient,omitempty" bson:"ingredient,omitempty"`                // Ingredients contained in this product
	Energy            *Quantity                        `json:"energy,omitempty" bson:"energy,omitempty"`                        // The amount of energy present in the product expressed in kilocalories or kilojoules
	Characteristic    []NutritionProductCharacteristic `json:"characteristic,omitempty" bson:"characteristic,omitempty"`        // Specifies descriptive properties of the nutrition product
	Instance          []NutritionProductInstance       `json:"instance,omitempty" bson:"instance,omitempty"`                    // One or several physical instances or occurrences of the nutrition product
	Note              []Annotation                     `json:"note,omitempty" bson:"note,omitempty"`                            // Comments made about the product
}

func (r *NutritionProduct) Validate() error {
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
	if r.Code != nil {
		if err := r.Code.Validate(); err != nil {
			return fmt.Errorf("Code: %w", err)
		}
	}
	var emptyString string
	if r.Status == emptyString {
		return fmt.Errorf("field 'Status' is required")
	}
	for i, item := range r.Category {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Category[%d]: %w", i, err)
		}
	}
	for i, item := range r.Manufacturer {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Manufacturer[%d]: %w", i, err)
		}
	}
	for i, item := range r.Nutrient {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Nutrient[%d]: %w", i, err)
		}
	}
	for i, item := range r.Ingredient {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Ingredient[%d]: %w", i, err)
		}
	}
	if r.Energy != nil {
		if err := r.Energy.Validate(); err != nil {
			return fmt.Errorf("Energy: %w", err)
		}
	}
	for i, item := range r.Characteristic {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Characteristic[%d]: %w", i, err)
		}
	}
	for i, item := range r.Instance {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Instance[%d]: %w", i, err)
		}
	}
	for i, item := range r.Note {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Note[%d]: %w", i, err)
		}
	}
	return nil
}

type NutritionProductIngredient struct {
	Id             *string            `json:"id,omitempty" bson:"id,omitempty"`                          // Unique id for inter-element referencing
	Item           *CodeableReference `json:"item" bson:"item"`                                          // The ingredient contained in the product
	AmountRatio    *Ratio             `json:"amountRatio,omitempty" bson:"amount_ratio,omitempty"`       // The amount of ingredient that is in the product
	AmountQuantity *Quantity          `json:"amountQuantity,omitempty" bson:"amount_quantity,omitempty"` // The amount of ingredient that is in the product
	Allergen       bool               `json:"allergen,omitempty" bson:"allergen,omitempty"`              // A known or suspected allergenic and/or substance that is associated with an intolerance
}

func (r *NutritionProductIngredient) Validate() error {
	if r.Item == nil {
		return fmt.Errorf("field 'Item' is required")
	}
	if r.Item != nil {
		if err := r.Item.Validate(); err != nil {
			return fmt.Errorf("Item: %w", err)
		}
	}
	if r.AmountRatio != nil {
		if err := r.AmountRatio.Validate(); err != nil {
			return fmt.Errorf("AmountRatio: %w", err)
		}
	}
	if r.AmountQuantity != nil {
		if err := r.AmountQuantity.Validate(); err != nil {
			return fmt.Errorf("AmountQuantity: %w", err)
		}
	}
	return nil
}

type NutritionProductCharacteristic struct {
	Id                   *string          `json:"id,omitempty" bson:"id,omitempty"`                   // Unique id for inter-element referencing
	Type                 *CodeableConcept `json:"type" bson:"type"`                                   // Code specifying the type of characteristic
	ValueCodeableConcept *CodeableConcept `json:"valueCodeableConcept" bson:"value_codeable_concept"` // The value of the characteristic
	ValueString          *string          `json:"valueString" bson:"value_string"`                    // The value of the characteristic
	ValueQuantity        *Quantity        `json:"valueQuantity" bson:"value_quantity"`                // The value of the characteristic
	ValueBase64Binary    *string          `json:"valueBase64Binary" bson:"value_base64_binary"`       // The value of the characteristic
	ValueAttachment      *Attachment      `json:"valueAttachment" bson:"value_attachment"`            // The value of the characteristic
	ValueBoolean         *bool            `json:"valueBoolean" bson:"value_boolean"`                  // The value of the characteristic
}

func (r *NutritionProductCharacteristic) Validate() error {
	if r.Type == nil {
		return fmt.Errorf("field 'Type' is required")
	}
	if r.Type != nil {
		if err := r.Type.Validate(); err != nil {
			return fmt.Errorf("Type: %w", err)
		}
	}
	if r.ValueCodeableConcept == nil {
		return fmt.Errorf("field 'ValueCodeableConcept' is required")
	}
	if r.ValueCodeableConcept != nil {
		if err := r.ValueCodeableConcept.Validate(); err != nil {
			return fmt.Errorf("ValueCodeableConcept: %w", err)
		}
	}
	if r.ValueString == nil {
		return fmt.Errorf("field 'ValueString' is required")
	}
	if r.ValueQuantity == nil {
		return fmt.Errorf("field 'ValueQuantity' is required")
	}
	if r.ValueQuantity != nil {
		if err := r.ValueQuantity.Validate(); err != nil {
			return fmt.Errorf("ValueQuantity: %w", err)
		}
	}
	if r.ValueBase64Binary == nil {
		return fmt.Errorf("field 'ValueBase64Binary' is required")
	}
	if r.ValueAttachment == nil {
		return fmt.Errorf("field 'ValueAttachment' is required")
	}
	if r.ValueAttachment != nil {
		if err := r.ValueAttachment.Validate(); err != nil {
			return fmt.Errorf("ValueAttachment: %w", err)
		}
	}
	if r.ValueBoolean == nil {
		return fmt.Errorf("field 'ValueBoolean' is required")
	}
	return nil
}

type NutritionProductInstance struct {
	Id                    *string      `json:"id,omitempty" bson:"id,omitempty"`                                         // Unique id for inter-element referencing
	Quantity              *Quantity    `json:"quantity,omitempty" bson:"quantity,omitempty"`                             // The amount of items or instances
	Identifier            []Identifier `json:"identifier,omitempty" bson:"identifier,omitempty"`                         // The identifier for the physical instance, typically a serial number or manufacturer number
	Name                  *string      `json:"name,omitempty" bson:"name,omitempty"`                                     // The name or brand for the specific product
	LotNumber             *string      `json:"lotNumber,omitempty" bson:"lot_number,omitempty"`                          // The identification of the batch or lot of the product
	Expiry                *string      `json:"expiry,omitempty" bson:"expiry,omitempty"`                                 // The expiry date or date and time for the product
	UseBy                 *string      `json:"useBy,omitempty" bson:"use_by,omitempty"`                                  // The date until which the product is expected to be good for consumption
	BiologicalSourceEvent *Identifier  `json:"biologicalSourceEvent,omitempty" bson:"biological_source_event,omitempty"` // An identifier of the donation, collection, or pooling event from which biological material in this nutrition product was derived
}

func (r *NutritionProductInstance) Validate() error {
	if r.Quantity != nil {
		if err := r.Quantity.Validate(); err != nil {
			return fmt.Errorf("Quantity: %w", err)
		}
	}
	for i, item := range r.Identifier {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Identifier[%d]: %w", i, err)
		}
	}
	if r.BiologicalSourceEvent != nil {
		if err := r.BiologicalSourceEvent.Validate(); err != nil {
			return fmt.Errorf("BiologicalSourceEvent: %w", err)
		}
	}
	return nil
}

type NutritionProductNutrient struct {
	Id             *string            `json:"id,omitempty" bson:"id,omitempty"`                          // Unique id for inter-element referencing
	Item           *CodeableReference `json:"item" bson:"item"`                                          // The (relevant) nutrients in the product
	AmountRatio    *Ratio             `json:"amountRatio,omitempty" bson:"amount_ratio,omitempty"`       // The amount of nutrient present in the product
	AmountQuantity *Quantity          `json:"amountQuantity,omitempty" bson:"amount_quantity,omitempty"` // The amount of nutrient present in the product
}

func (r *NutritionProductNutrient) Validate() error {
	if r.Item == nil {
		return fmt.Errorf("field 'Item' is required")
	}
	if r.Item != nil {
		if err := r.Item.Validate(); err != nil {
			return fmt.Errorf("Item: %w", err)
		}
	}
	if r.AmountRatio != nil {
		if err := r.AmountRatio.Validate(); err != nil {
			return fmt.Errorf("AmountRatio: %w", err)
		}
	}
	if r.AmountQuantity != nil {
		if err := r.AmountQuantity.Validate(); err != nil {
			return fmt.Errorf("AmountQuantity: %w", err)
		}
	}
	return nil
}
