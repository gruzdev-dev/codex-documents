package models

import (
	"encoding/json"
	"fmt"
)

// Detailed definition of a medicinal product, typically for uses other than direct patient care (e.g. regulatory use, drug catalogs, to support prescribing, adverse events management etc.).
type MedicinalProductDefinition struct {
	Id                             *string                                    `json:"id,omitempty" bson:"id,omitempty"`                                                            // Logical id of this artifact
	Meta                           *Meta                                      `json:"meta,omitempty" bson:"meta,omitempty"`                                                        // Metadata about the resource
	ImplicitRules                  *string                                    `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`                                     // A set of rules under which this content was created
	Language                       *string                                    `json:"language,omitempty" bson:"language,omitempty"`                                                // Language of the resource content
	Text                           *Narrative                                 `json:"text,omitempty" bson:"text,omitempty"`                                                        // Text summary of the resource, for human interpretation
	Contained                      []json.RawMessage                          `json:"contained,omitempty" bson:"contained,omitempty"`                                              // Contained, inline Resources
	Identifier                     []Identifier                               `json:"identifier,omitempty" bson:"identifier,omitempty"`                                            // Business identifier for this product. Could be an MPID
	Type                           *CodeableConcept                           `json:"type,omitempty" bson:"type,omitempty"`                                                        // Regulatory type, e.g. Investigational or Authorized
	Domain                         *CodeableConcept                           `json:"domain,omitempty" bson:"domain,omitempty"`                                                    // If this medicine applies to human or veterinary uses
	Version                        *string                                    `json:"version,omitempty" bson:"version,omitempty"`                                                  // A business identifier relating to a specific version of the product
	Status                         *CodeableConcept                           `json:"status,omitempty" bson:"status,omitempty"`                                                    // The status within the lifecycle of this product record
	StatusDate                     *string                                    `json:"statusDate,omitempty" bson:"status_date,omitempty"`                                           // The date at which the given status became applicable
	Description                    *string                                    `json:"description,omitempty" bson:"description,omitempty"`                                          // General description of this product
	CombinedPharmaceuticalDoseForm *CodeableConcept                           `json:"combinedPharmaceuticalDoseForm,omitempty" bson:"combined_pharmaceutical_dose_form,omitempty"` // The dose form for a single part product, or combined form of a multiple part product
	Route                          []CodeableConcept                          `json:"route,omitempty" bson:"route,omitempty"`                                                      // The path by which the product is taken into or makes contact with the body
	Indication                     *string                                    `json:"indication,omitempty" bson:"indication,omitempty"`                                            // Description of indication(s) for this product, used when structured indications are not required
	LegalStatusOfSupply            *CodeableConcept                           `json:"legalStatusOfSupply,omitempty" bson:"legal_status_of_supply,omitempty"`                       // The legal status of supply of the medicinal product as classified by the regulator
	AdditionalMonitoringIndicator  *CodeableConcept                           `json:"additionalMonitoringIndicator,omitempty" bson:"additional_monitoring_indicator,omitempty"`    // Whether the Medicinal Product is subject to additional monitoring for regulatory reasons
	SpecialMeasures                []CodeableConcept                          `json:"specialMeasures,omitempty" bson:"special_measures,omitempty"`                                 // Whether the Medicinal Product is subject to special measures for regulatory reasons
	PediatricUseIndicator          *CodeableConcept                           `json:"pediatricUseIndicator,omitempty" bson:"pediatric_use_indicator,omitempty"`                    // If authorised for use in children
	Classification                 []CodeableConcept                          `json:"classification,omitempty" bson:"classification,omitempty"`                                    // Allows the product to be classified by various systems
	MarketingStatus                []MarketingStatus                          `json:"marketingStatus,omitempty" bson:"marketing_status,omitempty"`                                 // Marketing status of the medicinal product, in contrast to marketing authorization
	PackagedMedicinalProduct       []CodeableConcept                          `json:"packagedMedicinalProduct,omitempty" bson:"packaged_medicinal_product,omitempty"`              // Package type for the product
	ComprisedOf                    []Reference                                `json:"comprisedOf,omitempty" bson:"comprised_of,omitempty"`                                         // Types of medicinal manufactured items and/or devices that this product consists of, such as tablets, capsule, or syringes
	Ingredient                     []CodeableConcept                          `json:"ingredient,omitempty" bson:"ingredient,omitempty"`                                            // The ingredients of this medicinal product - when not detailed in other resources
	Impurity                       []CodeableReference                        `json:"impurity,omitempty" bson:"impurity,omitempty"`                                                // Any component of the drug product which is not the chemical entity defined as the drug substance, or an excipient in the drug product
	AttachedDocument               []Reference                                `json:"attachedDocument,omitempty" bson:"attached_document,omitempty"`                               // Additional documentation about the medicinal product
	MasterFile                     []Reference                                `json:"masterFile,omitempty" bson:"master_file,omitempty"`                                           // A master file for the medicinal product (e.g. Pharmacovigilance System Master File)
	Contact                        []MedicinalProductDefinitionContact        `json:"contact,omitempty" bson:"contact,omitempty"`                                                  // A product specific contact, person (in a role), or an organization
	ClinicalTrial                  []Reference                                `json:"clinicalTrial,omitempty" bson:"clinical_trial,omitempty"`                                     // Clinical trials or studies that this product is involved in
	Code                           []Coding                                   `json:"code,omitempty" bson:"code,omitempty"`                                                        // A code that this product is known by, within some formal terminology
	Name                           []MedicinalProductDefinitionName           `json:"name" bson:"name"`                                                                            // The product's name, including full name and possibly coded parts
	CrossReference                 []MedicinalProductDefinitionCrossReference `json:"crossReference,omitempty" bson:"cross_reference,omitempty"`                                   // Reference to another product, e.g. for linking authorised to investigational product
	Operation                      []MedicinalProductDefinitionOperation      `json:"operation,omitempty" bson:"operation,omitempty"`                                              // A manufacturing or administrative process for the medicinal product
	Characteristic                 []MedicinalProductDefinitionCharacteristic `json:"characteristic,omitempty" bson:"characteristic,omitempty"`                                    // Key product features such as "sugar free", "modified release"
}

func (r *MedicinalProductDefinition) Validate() error {
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
	if r.Domain != nil {
		if err := r.Domain.Validate(); err != nil {
			return fmt.Errorf("Domain: %w", err)
		}
	}
	if r.Status != nil {
		if err := r.Status.Validate(); err != nil {
			return fmt.Errorf("Status: %w", err)
		}
	}
	if r.CombinedPharmaceuticalDoseForm != nil {
		if err := r.CombinedPharmaceuticalDoseForm.Validate(); err != nil {
			return fmt.Errorf("CombinedPharmaceuticalDoseForm: %w", err)
		}
	}
	for i, item := range r.Route {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Route[%d]: %w", i, err)
		}
	}
	if r.LegalStatusOfSupply != nil {
		if err := r.LegalStatusOfSupply.Validate(); err != nil {
			return fmt.Errorf("LegalStatusOfSupply: %w", err)
		}
	}
	if r.AdditionalMonitoringIndicator != nil {
		if err := r.AdditionalMonitoringIndicator.Validate(); err != nil {
			return fmt.Errorf("AdditionalMonitoringIndicator: %w", err)
		}
	}
	for i, item := range r.SpecialMeasures {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("SpecialMeasures[%d]: %w", i, err)
		}
	}
	if r.PediatricUseIndicator != nil {
		if err := r.PediatricUseIndicator.Validate(); err != nil {
			return fmt.Errorf("PediatricUseIndicator: %w", err)
		}
	}
	for i, item := range r.Classification {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Classification[%d]: %w", i, err)
		}
	}
	for i, item := range r.MarketingStatus {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("MarketingStatus[%d]: %w", i, err)
		}
	}
	for i, item := range r.PackagedMedicinalProduct {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("PackagedMedicinalProduct[%d]: %w", i, err)
		}
	}
	for i, item := range r.ComprisedOf {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("ComprisedOf[%d]: %w", i, err)
		}
	}
	for i, item := range r.Ingredient {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Ingredient[%d]: %w", i, err)
		}
	}
	for i, item := range r.Impurity {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Impurity[%d]: %w", i, err)
		}
	}
	for i, item := range r.AttachedDocument {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("AttachedDocument[%d]: %w", i, err)
		}
	}
	for i, item := range r.MasterFile {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("MasterFile[%d]: %w", i, err)
		}
	}
	for i, item := range r.Contact {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Contact[%d]: %w", i, err)
		}
	}
	for i, item := range r.ClinicalTrial {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("ClinicalTrial[%d]: %w", i, err)
		}
	}
	for i, item := range r.Code {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Code[%d]: %w", i, err)
		}
	}
	if len(r.Name) < 1 {
		return fmt.Errorf("field 'Name' must have at least 1 elements")
	}
	for i, item := range r.Name {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Name[%d]: %w", i, err)
		}
	}
	for i, item := range r.CrossReference {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("CrossReference[%d]: %w", i, err)
		}
	}
	for i, item := range r.Operation {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Operation[%d]: %w", i, err)
		}
	}
	for i, item := range r.Characteristic {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Characteristic[%d]: %w", i, err)
		}
	}
	return nil
}

type MedicinalProductDefinitionCrossReference struct {
	Id      *string            `json:"id,omitempty" bson:"id,omitempty"`     // Unique id for inter-element referencing
	Product *CodeableReference `json:"product" bson:"product"`               // Reference to another product, e.g. for linking authorised to investigational product
	Type    *CodeableConcept   `json:"type,omitempty" bson:"type,omitempty"` // The type of relationship, for instance branded to generic or virtual to actual product
}

func (r *MedicinalProductDefinitionCrossReference) Validate() error {
	if r.Product == nil {
		return fmt.Errorf("field 'Product' is required")
	}
	if r.Product != nil {
		if err := r.Product.Validate(); err != nil {
			return fmt.Errorf("Product: %w", err)
		}
	}
	if r.Type != nil {
		if err := r.Type.Validate(); err != nil {
			return fmt.Errorf("Type: %w", err)
		}
	}
	return nil
}

type MedicinalProductDefinitionOperation struct {
	Id                       *string            `json:"id,omitempty" bson:"id,omitempty"`                                              // Unique id for inter-element referencing
	Type                     *CodeableReference `json:"type,omitempty" bson:"type,omitempty"`                                          // The type of manufacturing operation e.g. manufacturing itself, re-packaging
	EffectiveDate            *Period            `json:"effectiveDate,omitempty" bson:"effective_date,omitempty"`                       // Date range of applicability
	Organization             []Reference        `json:"organization,omitempty" bson:"organization,omitempty"`                          // The organization responsible for the particular process, e.g. the manufacturer or importer
	ConfidentialityIndicator *CodeableConcept   `json:"confidentialityIndicator,omitempty" bson:"confidentiality_indicator,omitempty"` // Specifies whether this process is considered proprietary or confidential
}

func (r *MedicinalProductDefinitionOperation) Validate() error {
	if r.Type != nil {
		if err := r.Type.Validate(); err != nil {
			return fmt.Errorf("Type: %w", err)
		}
	}
	if r.EffectiveDate != nil {
		if err := r.EffectiveDate.Validate(); err != nil {
			return fmt.Errorf("EffectiveDate: %w", err)
		}
	}
	for i, item := range r.Organization {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Organization[%d]: %w", i, err)
		}
	}
	if r.ConfidentialityIndicator != nil {
		if err := r.ConfidentialityIndicator.Validate(); err != nil {
			return fmt.Errorf("ConfidentialityIndicator: %w", err)
		}
	}
	return nil
}

type MedicinalProductDefinitionCharacteristic struct {
	Id                   *string          `json:"id,omitempty" bson:"id,omitempty"`                                       // Unique id for inter-element referencing
	Type                 *CodeableConcept `json:"type" bson:"type"`                                                       // A code expressing the type of characteristic
	ValueCodeableConcept *CodeableConcept `json:"valueCodeableConcept,omitempty" bson:"value_codeable_concept,omitempty"` // A value for the characteristic
	ValueMarkdown        *string          `json:"valueMarkdown,omitempty" bson:"value_markdown,omitempty"`                // A value for the characteristic
	ValueQuantity        *Quantity        `json:"valueQuantity,omitempty" bson:"value_quantity,omitempty"`                // A value for the characteristic
	ValueRange           *Range           `json:"valueRange,omitempty" bson:"value_range,omitempty"`                      // A value for the characteristic
	ValueInteger         *int             `json:"valueInteger,omitempty" bson:"value_integer,omitempty"`                  // A value for the characteristic
	ValueDate            *string          `json:"valueDate,omitempty" bson:"value_date,omitempty"`                        // A value for the characteristic
	ValueBoolean         *bool            `json:"valueBoolean,omitempty" bson:"value_boolean,omitempty"`                  // A value for the characteristic
	ValueAttachment      *Attachment      `json:"valueAttachment,omitempty" bson:"value_attachment,omitempty"`            // A value for the characteristic
}

func (r *MedicinalProductDefinitionCharacteristic) Validate() error {
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
	return nil
}

type MedicinalProductDefinitionName struct {
	Id          *string                               `json:"id,omitempty" bson:"id,omitempty"`       // Unique id for inter-element referencing
	ProductName string                                `json:"productName" bson:"product_name"`        // The full product name
	Type        *CodeableConcept                      `json:"type,omitempty" bson:"type,omitempty"`   // Type of product name, such as rINN, BAN, Proprietary, Non-Proprietary
	Part        []MedicinalProductDefinitionNamePart  `json:"part,omitempty" bson:"part,omitempty"`   // Coding words or phrases of the name
	Usage       []MedicinalProductDefinitionNameUsage `json:"usage,omitempty" bson:"usage,omitempty"` // Country and jurisdiction where the name applies
}

func (r *MedicinalProductDefinitionName) Validate() error {
	var emptyString string
	if r.ProductName == emptyString {
		return fmt.Errorf("field 'ProductName' is required")
	}
	if r.Type != nil {
		if err := r.Type.Validate(); err != nil {
			return fmt.Errorf("Type: %w", err)
		}
	}
	for i, item := range r.Part {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Part[%d]: %w", i, err)
		}
	}
	for i, item := range r.Usage {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Usage[%d]: %w", i, err)
		}
	}
	return nil
}

type MedicinalProductDefinitionContact struct {
	Id      *string          `json:"id,omitempty" bson:"id,omitempty"`     // Unique id for inter-element referencing
	Type    *CodeableConcept `json:"type,omitempty" bson:"type,omitempty"` // Allows the contact to be classified, for example QPPV, Pharmacovigilance Enquiry Information
	Contact *Reference       `json:"contact" bson:"contact"`               // A product specific contact, person (in a role), or an organization
}

func (r *MedicinalProductDefinitionContact) Validate() error {
	if r.Type != nil {
		if err := r.Type.Validate(); err != nil {
			return fmt.Errorf("Type: %w", err)
		}
	}
	if r.Contact == nil {
		return fmt.Errorf("field 'Contact' is required")
	}
	if r.Contact != nil {
		if err := r.Contact.Validate(); err != nil {
			return fmt.Errorf("Contact: %w", err)
		}
	}
	return nil
}

type MedicinalProductDefinitionNamePart struct {
	Id   *string          `json:"id,omitempty" bson:"id,omitempty"` // Unique id for inter-element referencing
	Part string           `json:"part" bson:"part"`                 // A fragment of a product name
	Type *CodeableConcept `json:"type" bson:"type"`                 // Identifying type for this part of the name (e.g. strength part)
}

func (r *MedicinalProductDefinitionNamePart) Validate() error {
	var emptyString string
	if r.Part == emptyString {
		return fmt.Errorf("field 'Part' is required")
	}
	if r.Type == nil {
		return fmt.Errorf("field 'Type' is required")
	}
	if r.Type != nil {
		if err := r.Type.Validate(); err != nil {
			return fmt.Errorf("Type: %w", err)
		}
	}
	return nil
}

type MedicinalProductDefinitionNameUsage struct {
	Id           *string          `json:"id,omitempty" bson:"id,omitempty"`                     // Unique id for inter-element referencing
	Country      *CodeableConcept `json:"country" bson:"country"`                               // Country code for where this name applies
	Jurisdiction *CodeableConcept `json:"jurisdiction,omitempty" bson:"jurisdiction,omitempty"` // Jurisdiction code for where this name applies
	Language     *CodeableConcept `json:"language" bson:"language"`                             // Language code for this name
}

func (r *MedicinalProductDefinitionNameUsage) Validate() error {
	if r.Country == nil {
		return fmt.Errorf("field 'Country' is required")
	}
	if r.Country != nil {
		if err := r.Country.Validate(); err != nil {
			return fmt.Errorf("Country: %w", err)
		}
	}
	if r.Jurisdiction != nil {
		if err := r.Jurisdiction.Validate(); err != nil {
			return fmt.Errorf("Jurisdiction: %w", err)
		}
	}
	if r.Language == nil {
		return fmt.Errorf("field 'Language' is required")
	}
	if r.Language != nil {
		if err := r.Language.Validate(); err != nil {
			return fmt.Errorf("Language: %w", err)
		}
	}
	return nil
}
