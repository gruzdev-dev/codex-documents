package models

import (
	"encoding/json"
	"fmt"
)

// This resource is primarily used for the identification and definition of a medication, including ingredients, for the purposes of prescribing, dispensing, and administering a medication as well as for making statements about medication use.
type Medication struct {
	ResourceType                 string                 `json:"resourceType" bson:"resource_type"`                                                      // Type of resource
	Id                           *string                `json:"id,omitempty" bson:"id,omitempty"`                                                       // Logical id of this artifact
	Meta                         *Meta                  `json:"meta,omitempty" bson:"meta,omitempty"`                                                   // Metadata about the resource
	ImplicitRules                *string                `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`                                // A set of rules under which this content was created
	Language                     *string                `json:"language,omitempty" bson:"language,omitempty"`                                           // Language of the resource content
	Text                         *Narrative             `json:"text,omitempty" bson:"text,omitempty"`                                                   // Text summary of the resource, for human interpretation
	Contained                    []json.RawMessage      `json:"contained,omitempty" bson:"contained,omitempty"`                                         // Contained, inline Resources
	Identifier                   []Identifier           `json:"identifier,omitempty" bson:"identifier,omitempty"`                                       // Business identifier for this medication
	Code                         *CodeableConcept       `json:"code,omitempty" bson:"code,omitempty"`                                                   // Codes that identify this medication
	Status                       *string                `json:"status,omitempty" bson:"status,omitempty"`                                               // active | inactive | entered-in-error
	MarketingAuthorizationHolder *Reference             `json:"marketingAuthorizationHolder,omitempty" bson:"marketing_authorization_holder,omitempty"` // Organization that has authorization to market medication
	DoseForm                     *CodeableConcept       `json:"doseForm,omitempty" bson:"dose_form,omitempty"`                                          // powder | tablets | capsule +
	PackageSize                  *MedicationPackageSize `json:"packageSize,omitempty" bson:"package_size,omitempty"`                                    // When the code does not specify the package size, this backbone element can be used to specify the overall amount of medication in the package
	Ingredient                   []MedicationIngredient `json:"ingredient,omitempty" bson:"ingredient,omitempty"`                                       // Components of a medication. These can be ingredient substances or other medications, in the case of combination packaged medications
	Instance                     *MedicationInstance    `json:"instance,omitempty" bson:"instance,omitempty"`                                           // Details about packaged medications
	Definition                   *Reference             `json:"definition,omitempty" bson:"definition,omitempty"`                                       // Formal definition of the medication
}

func (r *Medication) Validate() error {
	if r.ResourceType != "Medication" {
		return fmt.Errorf("invalid resourceType: expected 'Medication', got '%s'", r.ResourceType)
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
	if r.Code != nil {
		if err := r.Code.Validate(); err != nil {
			return fmt.Errorf("Code: %w", err)
		}
	}
	if r.MarketingAuthorizationHolder != nil {
		if err := r.MarketingAuthorizationHolder.Validate(); err != nil {
			return fmt.Errorf("MarketingAuthorizationHolder: %w", err)
		}
	}
	if r.DoseForm != nil {
		if err := r.DoseForm.Validate(); err != nil {
			return fmt.Errorf("DoseForm: %w", err)
		}
	}
	if r.PackageSize != nil {
		if err := r.PackageSize.Validate(); err != nil {
			return fmt.Errorf("PackageSize: %w", err)
		}
	}
	for i, item := range r.Ingredient {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Ingredient[%d]: %w", i, err)
		}
	}
	if r.Instance != nil {
		if err := r.Instance.Validate(); err != nil {
			return fmt.Errorf("Instance: %w", err)
		}
	}
	if r.Definition != nil {
		if err := r.Definition.Validate(); err != nil {
			return fmt.Errorf("Definition: %w", err)
		}
	}
	return nil
}

type MedicationPackageSize struct {
	Id            *string   `json:"id,omitempty" bson:"id,omitempty"`                         // Unique id for inter-element referencing
	NumberOfItems *Quantity `json:"numberOfItems,omitempty" bson:"number_of_items,omitempty"` // The number of items in the package that are represented by the code
	AmountPerItem *Ratio    `json:"amountPerItem,omitempty" bson:"amount_per_item,omitempty"` // The amount per each item
}

func (r *MedicationPackageSize) Validate() error {
	if r.NumberOfItems != nil {
		if err := r.NumberOfItems.Validate(); err != nil {
			return fmt.Errorf("NumberOfItems: %w", err)
		}
	}
	if r.AmountPerItem != nil {
		if err := r.AmountPerItem.Validate(); err != nil {
			return fmt.Errorf("AmountPerItem: %w", err)
		}
	}
	return nil
}

type MedicationIngredient struct {
	Id                      *string            `json:"id,omitempty" bson:"id,omitempty"`                                             // Unique id for inter-element referencing
	Item                    *CodeableReference `json:"item" bson:"item"`                                                             // The specific substance or medication that is the ingredient
	IsActive                *bool              `json:"isActive,omitempty" bson:"is_active,omitempty"`                                // Active ingredient indicator
	Role                    *CodeableConcept   `json:"role,omitempty" bson:"role,omitempty"`                                         // A code that defines the type of ingredient, active, base, etc.
	StrengthRatio           *Ratio             `json:"strengthRatio,omitempty" bson:"strength_ratio,omitempty"`                      // Quantity of ingredient present
	StrengthCodeableConcept *CodeableConcept   `json:"strengthCodeableConcept,omitempty" bson:"strength_codeable_concept,omitempty"` // Quantity of ingredient present
	StrengthQuantity        *Quantity          `json:"strengthQuantity,omitempty" bson:"strength_quantity,omitempty"`                // Quantity of ingredient present
}

func (r *MedicationIngredient) Validate() error {
	if r.Item == nil {
		return fmt.Errorf("field 'Item' is required")
	}
	if r.Item != nil {
		if err := r.Item.Validate(); err != nil {
			return fmt.Errorf("Item: %w", err)
		}
	}
	if r.Role != nil {
		if err := r.Role.Validate(); err != nil {
			return fmt.Errorf("Role: %w", err)
		}
	}
	if r.StrengthRatio != nil {
		if err := r.StrengthRatio.Validate(); err != nil {
			return fmt.Errorf("StrengthRatio: %w", err)
		}
	}
	if r.StrengthCodeableConcept != nil {
		if err := r.StrengthCodeableConcept.Validate(); err != nil {
			return fmt.Errorf("StrengthCodeableConcept: %w", err)
		}
	}
	if r.StrengthQuantity != nil {
		if err := r.StrengthQuantity.Validate(); err != nil {
			return fmt.Errorf("StrengthQuantity: %w", err)
		}
	}
	return nil
}

type MedicationInstance struct {
	Id             *string      `json:"id,omitempty" bson:"id,omitempty"`                          // Unique id for inter-element referencing
	Identifier     []Identifier `json:"identifier,omitempty" bson:"identifier,omitempty"`          // Identifier for the physical instance, typically a serial number
	LotNumber      *string      `json:"lotNumber,omitempty" bson:"lot_number,omitempty"`           // Identifier assigned to batch
	ExpirationDate *string      `json:"expirationDate,omitempty" bson:"expiration_date,omitempty"` // When instance will expire
}

func (r *MedicationInstance) Validate() error {
	for i, item := range r.Identifier {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Identifier[%d]: %w", i, err)
		}
	}
	return nil
}
