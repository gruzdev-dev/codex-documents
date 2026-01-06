package models

import (
	"encoding/json"
	"fmt"
)

// A medically related item or items, in a container or package.
type PackagedProductDefinition struct {
	Id                    *string                                        `json:"id,omitempty" bson:"id,omitempty"`                                         // Logical id of this artifact
	Meta                  *Meta                                          `json:"meta,omitempty" bson:"meta,omitempty"`                                     // Metadata about the resource
	ImplicitRules         *string                                        `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`                  // A set of rules under which this content was created
	Language              *string                                        `json:"language,omitempty" bson:"language,omitempty"`                             // Language of the resource content
	Text                  *Narrative                                     `json:"text,omitempty" bson:"text,omitempty"`                                     // Text summary of the resource, for human interpretation
	Contained             []json.RawMessage                              `json:"contained,omitempty" bson:"contained,omitempty"`                           // Contained, inline Resources
	Identifier            []Identifier                                   `json:"identifier,omitempty" bson:"identifier,omitempty"`                         // A unique identifier for this package as whole - not for the content of the package
	Name                  *string                                        `json:"name,omitempty" bson:"name,omitempty"`                                     // A name for this package. Typically as listed in a drug formulary, catalogue, inventory etc
	Type                  *CodeableConcept                               `json:"type,omitempty" bson:"type,omitempty"`                                     // A high level category e.g. medicinal product, raw material, shipping container etc
	PackageFor            []Reference                                    `json:"packageFor,omitempty" bson:"package_for,omitempty"`                        // The product that this is a pack for
	Status                *CodeableConcept                               `json:"status,omitempty" bson:"status,omitempty"`                                 // The status within the lifecycle of this item. High level - not intended to duplicate details elsewhere e.g. legal status, or authorization/marketing status
	StatusDate            *string                                        `json:"statusDate,omitempty" bson:"status_date,omitempty"`                        // The date at which the given status became applicable
	ContainedItemQuantity []Quantity                                     `json:"containedItemQuantity,omitempty" bson:"contained_item_quantity,omitempty"` // A total of the complete count of contained items of a particular type/form, independent of sub-packaging or organization. This can be considered as the pack size. See also packaging.containedItem.amount (especially the long definition)
	Description           *string                                        `json:"description,omitempty" bson:"description,omitempty"`                       // Textual description. Note that this is not the name of the package or product
	LegalStatusOfSupply   []PackagedProductDefinitionLegalStatusOfSupply `json:"legalStatusOfSupply,omitempty" bson:"legal_status_of_supply,omitempty"`    // The legal status of supply of the packaged item as classified by the regulator
	MarketingStatus       []MarketingStatus                              `json:"marketingStatus,omitempty" bson:"marketing_status,omitempty"`              // Allows specifying that an item is on the market for sale, or that it is not available, and the dates and locations associated
	CopackagedIndicator   bool                                           `json:"copackagedIndicator,omitempty" bson:"copackaged_indicator,omitempty"`      // Identifies if the drug product is supplied with another item such as a diluent or adjuvant
	Manufacturer          []Reference                                    `json:"manufacturer,omitempty" bson:"manufacturer,omitempty"`                     // Manufacturer of this package type (multiple means these are all possible manufacturers)
	AttachedDocument      []Reference                                    `json:"attachedDocument,omitempty" bson:"attached_document,omitempty"`            // Additional information or supporting documentation about the packaged product
	Packaging             *PackagedProductDefinitionPackaging            `json:"packaging,omitempty" bson:"packaging,omitempty"`                           // A packaging item, as a container for medically related items, possibly with other packaging items within, or a packaging component, such as bottle cap
	Characteristic        []PackagedProductDefinitionPackagingProperty   `json:"characteristic,omitempty" bson:"characteristic,omitempty"`                 // Allows the key features to be recorded, such as "hospital pack", "nurse prescribable"
}

func (r *PackagedProductDefinition) Validate() error {
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
	for i, item := range r.PackageFor {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("PackageFor[%d]: %w", i, err)
		}
	}
	if r.Status != nil {
		if err := r.Status.Validate(); err != nil {
			return fmt.Errorf("Status: %w", err)
		}
	}
	for i, item := range r.ContainedItemQuantity {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("ContainedItemQuantity[%d]: %w", i, err)
		}
	}
	for i, item := range r.LegalStatusOfSupply {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("LegalStatusOfSupply[%d]: %w", i, err)
		}
	}
	for i, item := range r.MarketingStatus {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("MarketingStatus[%d]: %w", i, err)
		}
	}
	for i, item := range r.Manufacturer {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Manufacturer[%d]: %w", i, err)
		}
	}
	for i, item := range r.AttachedDocument {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("AttachedDocument[%d]: %w", i, err)
		}
	}
	if r.Packaging != nil {
		if err := r.Packaging.Validate(); err != nil {
			return fmt.Errorf("Packaging: %w", err)
		}
	}
	for i, item := range r.Characteristic {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Characteristic[%d]: %w", i, err)
		}
	}
	return nil
}

type PackagedProductDefinitionLegalStatusOfSupply struct {
	Id           *string          `json:"id,omitempty" bson:"id,omitempty"`                     // Unique id for inter-element referencing
	Code         *CodeableConcept `json:"code,omitempty" bson:"code,omitempty"`                 // The actual status of supply. In what situation this package type may be supplied for use
	Jurisdiction *CodeableConcept `json:"jurisdiction,omitempty" bson:"jurisdiction,omitempty"` // The place where the legal status of supply applies
}

func (r *PackagedProductDefinitionLegalStatusOfSupply) Validate() error {
	if r.Code != nil {
		if err := r.Code.Validate(); err != nil {
			return fmt.Errorf("Code: %w", err)
		}
	}
	if r.Jurisdiction != nil {
		if err := r.Jurisdiction.Validate(); err != nil {
			return fmt.Errorf("Jurisdiction: %w", err)
		}
	}
	return nil
}

type PackagedProductDefinitionPackaging struct {
	Id                *string                                           `json:"id,omitempty" bson:"id,omitempty"`                                // Unique id for inter-element referencing
	Identifier        []Identifier                                      `json:"identifier,omitempty" bson:"identifier,omitempty"`                // An identifier that is specific to this particular part of the packaging. Including possibly a Data Carrier Identifier
	Type              *CodeableConcept                                  `json:"type,omitempty" bson:"type,omitempty"`                            // The physical type of the container of the items
	ComponentPart     bool                                              `json:"componentPart,omitempty" bson:"component_part,omitempty"`         // Is this a part of the packaging (e.g. a cap or bottle stopper), rather than the packaging itself (e.g. a bottle or vial)
	Quantity          *int                                              `json:"quantity,omitempty" bson:"quantity,omitempty"`                    // The quantity of this level of packaging in the package that contains it (with the outermost level being 1)
	Material          []CodeableConcept                                 `json:"material,omitempty" bson:"material,omitempty"`                    // Material type of the package item
	AlternateMaterial []CodeableConcept                                 `json:"alternateMaterial,omitempty" bson:"alternate_material,omitempty"` // A possible alternate material for this part of the packaging, that is allowed to be used instead of the usual material
	ShelfLifeStorage  []ProductShelfLife                                `json:"shelfLifeStorage,omitempty" bson:"shelf_life_storage,omitempty"`  // Shelf Life and storage information
	Manufacturer      []Reference                                       `json:"manufacturer,omitempty" bson:"manufacturer,omitempty"`            // Manufacturer of this packaging item (multiple means these are all potential manufacturers)
	Property          []PackagedProductDefinitionPackagingProperty      `json:"property,omitempty" bson:"property,omitempty"`                    // General characteristics of this item
	ContainedItem     []PackagedProductDefinitionPackagingContainedItem `json:"containedItem,omitempty" bson:"contained_item,omitempty"`         // The item(s) within the packaging
	Packaging         []PackagedProductDefinitionPackaging              `json:"packaging,omitempty" bson:"packaging,omitempty"`                  // Allows containers (and parts of containers) within containers, still as a part of single packaged product
}

func (r *PackagedProductDefinitionPackaging) Validate() error {
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
	for i, item := range r.Material {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Material[%d]: %w", i, err)
		}
	}
	for i, item := range r.AlternateMaterial {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("AlternateMaterial[%d]: %w", i, err)
		}
	}
	for i, item := range r.ShelfLifeStorage {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("ShelfLifeStorage[%d]: %w", i, err)
		}
	}
	for i, item := range r.Manufacturer {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Manufacturer[%d]: %w", i, err)
		}
	}
	for i, item := range r.Property {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Property[%d]: %w", i, err)
		}
	}
	for i, item := range r.ContainedItem {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("ContainedItem[%d]: %w", i, err)
		}
	}
	for i, item := range r.Packaging {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Packaging[%d]: %w", i, err)
		}
	}
	return nil
}

type PackagedProductDefinitionPackagingProperty struct {
	Id                   *string          `json:"id,omitempty" bson:"id,omitempty"`                                       // Unique id for inter-element referencing
	Type                 *CodeableConcept `json:"type" bson:"type"`                                                       // A code expressing the type of characteristic
	ValueCodeableConcept *CodeableConcept `json:"valueCodeableConcept,omitempty" bson:"value_codeable_concept,omitempty"` // A value for the characteristic
	ValueQuantity        *Quantity        `json:"valueQuantity,omitempty" bson:"value_quantity,omitempty"`                // A value for the characteristic
	ValueDate            *string          `json:"valueDate,omitempty" bson:"value_date,omitempty"`                        // A value for the characteristic
	ValueBoolean         *bool            `json:"valueBoolean,omitempty" bson:"value_boolean,omitempty"`                  // A value for the characteristic
	ValueAttachment      *Attachment      `json:"valueAttachment,omitempty" bson:"value_attachment,omitempty"`            // A value for the characteristic
}

func (r *PackagedProductDefinitionPackagingProperty) Validate() error {
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
	if r.ValueAttachment != nil {
		if err := r.ValueAttachment.Validate(); err != nil {
			return fmt.Errorf("ValueAttachment: %w", err)
		}
	}
	return nil
}

type PackagedProductDefinitionPackagingContainedItem struct {
	Id     *string            `json:"id,omitempty" bson:"id,omitempty"`         // Unique id for inter-element referencing
	Item   *CodeableReference `json:"item" bson:"item"`                         // The actual item(s) of medication, as manufactured, or a device, or other medically related item (food, biologicals, raw materials, medical fluids, gases etc.), as contained in the package
	Amount *Quantity          `json:"amount,omitempty" bson:"amount,omitempty"` // The number of this type of item within this packaging or for continuous items such as liquids it is the quantity (for example 25ml). See also PackagedProductDefinition.containedItemQuantity (especially the long definition)
}

func (r *PackagedProductDefinitionPackagingContainedItem) Validate() error {
	if r.Item == nil {
		return fmt.Errorf("field 'Item' is required")
	}
	if r.Item != nil {
		if err := r.Item.Validate(); err != nil {
			return fmt.Errorf("Item: %w", err)
		}
	}
	if r.Amount != nil {
		if err := r.Amount.Validate(); err != nil {
			return fmt.Errorf("Amount: %w", err)
		}
	}
	return nil
}
