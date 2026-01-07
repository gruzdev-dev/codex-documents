package models

import (
	"encoding/json"
	"fmt"
)

// This is a specialized resource that defines the characteristics and capabilities of a device.
type DeviceDefinition struct {
	ResourceType              string                                 `json:"resourceType" bson:"resource_type"`                                                   // Type of resource
	Id                        *string                                `json:"id,omitempty" bson:"id,omitempty"`                                                    // Logical id of this artifact
	Meta                      *Meta                                  `json:"meta,omitempty" bson:"meta,omitempty"`                                                // Metadata about the resource
	ImplicitRules             *string                                `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`                             // A set of rules under which this content was created
	Language                  *string                                `json:"language,omitempty" bson:"language,omitempty"`                                        // Language of the resource content
	Text                      *Narrative                             `json:"text,omitempty" bson:"text,omitempty"`                                                // Text summary of the resource, for human interpretation
	Contained                 []json.RawMessage                      `json:"contained,omitempty" bson:"contained,omitempty"`                                      // Contained, inline Resources
	Url                       *string                                `json:"url,omitempty" bson:"url,omitempty"`                                                  // Canonical identifier for this DeviceDefinition, represented as an absolute URI (globally unique)
	Identifier                []Identifier                           `json:"identifier,omitempty" bson:"identifier,omitempty"`                                    // Additional identifier for the DeviceDefinition
	Version                   *string                                `json:"version,omitempty" bson:"version,omitempty"`                                          // Business version of the DeviceDefinition
	VersionAlgorithmString    *string                                `json:"versionAlgorithmString,omitempty" bson:"version_algorithm_string,omitempty"`          // How to compare versions
	VersionAlgorithmCoding    *Coding                                `json:"versionAlgorithmCoding,omitempty" bson:"version_algorithm_coding,omitempty"`          // How to compare versions
	Name                      *string                                `json:"name,omitempty" bson:"name,omitempty"`                                                // Name for this DeviceDefinition (computer friendly)
	Title                     *string                                `json:"title,omitempty" bson:"title,omitempty"`                                              // Name for this DeviceDefinition (human friendly)
	Status                    string                                 `json:"status" bson:"status"`                                                                // draft | active | retired | unknown
	Experimental              *bool                                  `json:"experimental,omitempty" bson:"experimental,omitempty"`                                // For testing only - never for real usage
	Date                      *string                                `json:"date,omitempty" bson:"date,omitempty"`                                                // Date last changed
	Publisher                 *string                                `json:"publisher,omitempty" bson:"publisher,omitempty"`                                      // Name of the publisher/steward (organization or individual)
	Contact                   []ContactDetail                        `json:"contact,omitempty" bson:"contact,omitempty"`                                          // Contact details for the publisher
	Description               *string                                `json:"description,omitempty" bson:"description,omitempty"`                                  // Natural language description of the DeviceDefinition
	UseContext                []UsageContext                         `json:"useContext,omitempty" bson:"use_context,omitempty"`                                   // The context that the content is intended to support
	Jurisdiction              []CodeableConcept                      `json:"jurisdiction,omitempty" bson:"jurisdiction,omitempty"`                                // Jurisdiction of the authority that maintains the DeviceDefinition (if applicable)
	Purpose                   *string                                `json:"purpose,omitempty" bson:"purpose,omitempty"`                                          // Why this DeviceDefinition is defined
	Copyright                 *string                                `json:"copyright,omitempty" bson:"copyright,omitempty"`                                      // Notice about intellectual property ownership, can include restrictions on use
	CopyrightLabel            *string                                `json:"copyrightLabel,omitempty" bson:"copyright_label,omitempty"`                           // Copyright holder and year(s)
	ApprovalDate              *string                                `json:"approvalDate,omitempty" bson:"approval_date,omitempty"`                               // When DeviceDefinition was approved by publisher
	LastReviewDate            *string                                `json:"lastReviewDate,omitempty" bson:"last_review_date,omitempty"`                          // Date on which the asset content was last reviewed by the publisher
	EffectivePeriod           *Period                                `json:"effectivePeriod,omitempty" bson:"effective_period,omitempty"`                         // The effective date range for the DeviceDefinition
	PartNumber                *string                                `json:"partNumber,omitempty" bson:"part_number,omitempty"`                                   // The part number or catalog number of the device
	Manufacturer              *Reference                             `json:"manufacturer,omitempty" bson:"manufacturer,omitempty"`                                // Name of device manufacturer
	ModelNumber               *string                                `json:"modelNumber,omitempty" bson:"model_number,omitempty"`                                 // The catalog or model number for the device for example as defined by the manufacturer
	UdiDeviceIdentifier       []DeviceDefinitionUdiDeviceIdentifier  `json:"udiDeviceIdentifier,omitempty" bson:"udi_device_identifier,omitempty"`                // Unique Device Identifier (UDI) Barcode string
	RegulatoryIdentifier      []DeviceDefinitionRegulatoryIdentifier `json:"regulatoryIdentifier,omitempty" bson:"regulatory_identifier,omitempty"`               // Regulatory identifier(s) associated with this device
	DeviceName                []DeviceDefinitionDeviceName           `json:"deviceName,omitempty" bson:"device_name,omitempty"`                                   // The name or names of the device as given by the manufacturer
	Classification            []DeviceDefinitionClassification       `json:"classification,omitempty" bson:"classification,omitempty"`                            // What kind of device or device system this is
	ConformsTo                []DeviceDefinitionConformsTo           `json:"conformsTo,omitempty" bson:"conforms_to,omitempty"`                                   // Identifies the standards, specifications, or formal guidances for the capabilities supported by the device
	HasPart                   []DeviceDefinitionHasPart              `json:"hasPart,omitempty" bson:"has_part,omitempty"`                                         // A device, part of the current one
	Packaging                 []DeviceDefinitionPackaging            `json:"packaging,omitempty" bson:"packaging,omitempty"`                                      // Information about the packaging of the device, i.e. how the device is packaged
	DeviceVersion             []DeviceDefinitionDeviceVersion        `json:"deviceVersion,omitempty" bson:"device_version,omitempty"`                             // The version of the device or software
	Safety                    []CodeableConcept                      `json:"safety,omitempty" bson:"safety,omitempty"`                                            // Safety characteristics of the device
	ShelfLifeStorage          []ProductShelfLife                     `json:"shelfLifeStorage,omitempty" bson:"shelf_life_storage,omitempty"`                      // Shelf Life and storage information
	OutputLanguage            []string                               `json:"outputLanguage,omitempty" bson:"output_language,omitempty"`                           // Language code for the human-readable text strings produced by the device (all supported)
	Property                  []DeviceDefinitionProperty             `json:"property,omitempty" bson:"property,omitempty"`                                        // Inherent, essentially fixed, characteristics of this kind of device, e.g., time properties, size, etc
	Link                      []DeviceDefinitionLink                 `json:"link,omitempty" bson:"link,omitempty"`                                                // An associated device, attached to, used with, communicating with or linking a previous or new device model to the focal device
	Note                      []Annotation                           `json:"note,omitempty" bson:"note,omitempty"`                                                // Device notes and comments
	Material                  []DeviceDefinitionMaterial             `json:"material,omitempty" bson:"material,omitempty"`                                        // A substance used to create the material(s) of which the device is made
	ProductionIdentifierInUDI []CodeableConcept                      `json:"productionIdentifierInUDI,omitempty" bson:"production_identifier_in_u_d_i,omitempty"` // lot-number | manufactured-date | serial-number | expiration-date | biological-source | software-version
	Guideline                 *DeviceDefinitionGuideline             `json:"guideline,omitempty" bson:"guideline,omitempty"`                                      // Information aimed at providing directions for the usage of this model of device
	CorrectiveAction          *DeviceDefinitionCorrectiveAction      `json:"correctiveAction,omitempty" bson:"corrective_action,omitempty"`                       // Tracking of latest field safety corrective action
	ChargeItem                []DeviceDefinitionChargeItem           `json:"chargeItem,omitempty" bson:"charge_item,omitempty"`                                   // Billing code or reference associated with the device
}

func (r *DeviceDefinition) Validate() error {
	if r.ResourceType != "DeviceDefinition" {
		return fmt.Errorf("invalid resourceType: expected 'DeviceDefinition', got '%s'", r.ResourceType)
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
	if r.VersionAlgorithmCoding != nil {
		if err := r.VersionAlgorithmCoding.Validate(); err != nil {
			return fmt.Errorf("VersionAlgorithmCoding: %w", err)
		}
	}
	var emptyString string
	if r.Status == emptyString {
		return fmt.Errorf("field 'Status' is required")
	}
	for i, item := range r.Contact {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Contact[%d]: %w", i, err)
		}
	}
	for i, item := range r.UseContext {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("UseContext[%d]: %w", i, err)
		}
	}
	for i, item := range r.Jurisdiction {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Jurisdiction[%d]: %w", i, err)
		}
	}
	if r.EffectivePeriod != nil {
		if err := r.EffectivePeriod.Validate(); err != nil {
			return fmt.Errorf("EffectivePeriod: %w", err)
		}
	}
	if r.Manufacturer != nil {
		if err := r.Manufacturer.Validate(); err != nil {
			return fmt.Errorf("Manufacturer: %w", err)
		}
	}
	for i, item := range r.UdiDeviceIdentifier {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("UdiDeviceIdentifier[%d]: %w", i, err)
		}
	}
	for i, item := range r.RegulatoryIdentifier {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("RegulatoryIdentifier[%d]: %w", i, err)
		}
	}
	for i, item := range r.DeviceName {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("DeviceName[%d]: %w", i, err)
		}
	}
	for i, item := range r.Classification {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Classification[%d]: %w", i, err)
		}
	}
	for i, item := range r.ConformsTo {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("ConformsTo[%d]: %w", i, err)
		}
	}
	for i, item := range r.HasPart {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("HasPart[%d]: %w", i, err)
		}
	}
	for i, item := range r.Packaging {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Packaging[%d]: %w", i, err)
		}
	}
	for i, item := range r.DeviceVersion {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("DeviceVersion[%d]: %w", i, err)
		}
	}
	for i, item := range r.Safety {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Safety[%d]: %w", i, err)
		}
	}
	for i, item := range r.ShelfLifeStorage {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("ShelfLifeStorage[%d]: %w", i, err)
		}
	}
	for i, item := range r.Property {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Property[%d]: %w", i, err)
		}
	}
	for i, item := range r.Link {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Link[%d]: %w", i, err)
		}
	}
	for i, item := range r.Note {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Note[%d]: %w", i, err)
		}
	}
	for i, item := range r.Material {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Material[%d]: %w", i, err)
		}
	}
	for i, item := range r.ProductionIdentifierInUDI {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("ProductionIdentifierInUDI[%d]: %w", i, err)
		}
	}
	if r.Guideline != nil {
		if err := r.Guideline.Validate(); err != nil {
			return fmt.Errorf("Guideline: %w", err)
		}
	}
	if r.CorrectiveAction != nil {
		if err := r.CorrectiveAction.Validate(); err != nil {
			return fmt.Errorf("CorrectiveAction: %w", err)
		}
	}
	for i, item := range r.ChargeItem {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("ChargeItem[%d]: %w", i, err)
		}
	}
	return nil
}

type DeviceDefinitionDeviceVersion struct {
	Id        *string          `json:"id,omitempty" bson:"id,omitempty"`               // Unique id for inter-element referencing
	Type      *CodeableConcept `json:"type,omitempty" bson:"type,omitempty"`           // The type of the device version, e.g. manufacturer, approved, internal
	Component *Identifier      `json:"component,omitempty" bson:"component,omitempty"` // The hardware or software module of the device to which the version applies
	Value     string           `json:"value" bson:"value"`                             // The version text
}

func (r *DeviceDefinitionDeviceVersion) Validate() error {
	if r.Type != nil {
		if err := r.Type.Validate(); err != nil {
			return fmt.Errorf("Type: %w", err)
		}
	}
	if r.Component != nil {
		if err := r.Component.Validate(); err != nil {
			return fmt.Errorf("Component: %w", err)
		}
	}
	var emptyString string
	if r.Value == emptyString {
		return fmt.Errorf("field 'Value' is required")
	}
	return nil
}

type DeviceDefinitionUdiDeviceIdentifier struct {
	Id                     *string                                                 `json:"id,omitempty" bson:"id,omitempty"`                                           // Unique id for inter-element referencing
	DeviceIdentifier       string                                                  `json:"deviceIdentifier" bson:"device_identifier"`                                  // The identifier that is to be associated with every Device that references this DeviceDefintiion for the issuer and jurisdiction provided in the DeviceDefinition.udiDeviceIdentifier
	Issuer                 string                                                  `json:"issuer" bson:"issuer"`                                                       // The organization that assigns the identifier algorithm
	Jurisdiction           string                                                  `json:"jurisdiction" bson:"jurisdiction"`                                           // The jurisdiction to which the deviceIdentifier applies
	MarketDistribution     []DeviceDefinitionUdiDeviceIdentifierMarketDistribution `json:"marketDistribution,omitempty" bson:"market_distribution,omitempty"`          // Indicates whether and when the device is available on the market
	DeviceIdentifierSystem *string                                                 `json:"deviceIdentifierSystem,omitempty" bson:"device_identifier_system,omitempty"` // The namespace for the device identifier value
}

func (r *DeviceDefinitionUdiDeviceIdentifier) Validate() error {
	var emptyString string
	if r.DeviceIdentifier == emptyString {
		return fmt.Errorf("field 'DeviceIdentifier' is required")
	}
	if r.Issuer == emptyString {
		return fmt.Errorf("field 'Issuer' is required")
	}
	if r.Jurisdiction == emptyString {
		return fmt.Errorf("field 'Jurisdiction' is required")
	}
	for i, item := range r.MarketDistribution {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("MarketDistribution[%d]: %w", i, err)
		}
	}
	return nil
}

type DeviceDefinitionUdiDeviceIdentifierMarketDistribution struct {
	Id              *string `json:"id,omitempty" bson:"id,omitempty"`        // Unique id for inter-element referencing
	MarketPeriod    *Period `json:"marketPeriod" bson:"market_period"`       // Begin and end dates for the commercial distribution of the device
	SubJurisdiction string  `json:"subJurisdiction" bson:"sub_jurisdiction"` // National state or territory where the device is commercialized
}

func (r *DeviceDefinitionUdiDeviceIdentifierMarketDistribution) Validate() error {
	if r.MarketPeriod == nil {
		return fmt.Errorf("field 'MarketPeriod' is required")
	}
	if r.MarketPeriod != nil {
		if err := r.MarketPeriod.Validate(); err != nil {
			return fmt.Errorf("MarketPeriod: %w", err)
		}
	}
	var emptyString string
	if r.SubJurisdiction == emptyString {
		return fmt.Errorf("field 'SubJurisdiction' is required")
	}
	return nil
}

type DeviceDefinitionRegulatoryIdentifier struct {
	Id               *string `json:"id,omitempty" bson:"id,omitempty"`                              // Unique id for inter-element referencing
	Type             string  `json:"type" bson:"type"`                                              // basic | master | license
	Identifier       string  `json:"identifier" bson:"identifier"`                                  // The identifier itself
	Issuer           string  `json:"issuer" bson:"issuer"`                                          // The organization that issued this identifier
	Jurisdiction     string  `json:"jurisdiction" bson:"jurisdiction"`                              // Relevant jurisdiction governing the identifier
	IdentifierSystem *string `json:"identifierSystem,omitempty" bson:"identifier_system,omitempty"` // The namespace for the device identifier value
}

func (r *DeviceDefinitionRegulatoryIdentifier) Validate() error {
	var emptyString string
	if r.Type == emptyString {
		return fmt.Errorf("field 'Type' is required")
	}
	if r.Identifier == emptyString {
		return fmt.Errorf("field 'Identifier' is required")
	}
	if r.Issuer == emptyString {
		return fmt.Errorf("field 'Issuer' is required")
	}
	if r.Jurisdiction == emptyString {
		return fmt.Errorf("field 'Jurisdiction' is required")
	}
	return nil
}

type DeviceDefinitionConformsTo struct {
	Id            *string           `json:"id,omitempty" bson:"id,omitempty"`             // Unique id for inter-element referencing
	Category      *CodeableConcept  `json:"category,omitempty" bson:"category,omitempty"` // Describes the common type of the standard, specification, or formal guidance
	Specification *CodeableConcept  `json:"specification" bson:"specification"`           // Identifies the standard, specification, or formal guidance that the device adheres to the Device Specification type
	Version       []string          `json:"version,omitempty" bson:"version,omitempty"`   // The specific form or variant of the standard, specification or formal guidance
	Source        []RelatedArtifact `json:"source,omitempty" bson:"source,omitempty"`     // Standard, regulation, certification, or guidance website, document, or other publication, or similar, supporting the conformance
}

func (r *DeviceDefinitionConformsTo) Validate() error {
	if r.Category != nil {
		if err := r.Category.Validate(); err != nil {
			return fmt.Errorf("Category: %w", err)
		}
	}
	if r.Specification == nil {
		return fmt.Errorf("field 'Specification' is required")
	}
	if r.Specification != nil {
		if err := r.Specification.Validate(); err != nil {
			return fmt.Errorf("Specification: %w", err)
		}
	}
	for i, item := range r.Source {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Source[%d]: %w", i, err)
		}
	}
	return nil
}

type DeviceDefinitionPackaging struct {
	Id                  *string                                `json:"id,omitempty" bson:"id,omitempty"`                                     // Unique id for inter-element referencing
	Identifier          *Identifier                            `json:"identifier,omitempty" bson:"identifier,omitempty"`                     // Business identifier of the packaged medication
	Type                *CodeableConcept                       `json:"type,omitempty" bson:"type,omitempty"`                                 // A code that defines the specific type of packaging
	Count               *int                                   `json:"count,omitempty" bson:"count,omitempty"`                               // The number of items contained in the package (devices or sub-packages)
	Distributor         []DeviceDefinitionPackagingDistributor `json:"distributor,omitempty" bson:"distributor,omitempty"`                   // An organization that distributes the packaged device
	UdiDeviceIdentifier []DeviceDefinitionUdiDeviceIdentifier  `json:"udiDeviceIdentifier,omitempty" bson:"udi_device_identifier,omitempty"` // Unique Device Identifier (UDI) Barcode string on the packaging
	Packaging           []DeviceDefinitionPackaging            `json:"packaging,omitempty" bson:"packaging,omitempty"`                       // Allows packages within packages
}

func (r *DeviceDefinitionPackaging) Validate() error {
	if r.Identifier != nil {
		if err := r.Identifier.Validate(); err != nil {
			return fmt.Errorf("Identifier: %w", err)
		}
	}
	if r.Type != nil {
		if err := r.Type.Validate(); err != nil {
			return fmt.Errorf("Type: %w", err)
		}
	}
	for i, item := range r.Distributor {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Distributor[%d]: %w", i, err)
		}
	}
	for i, item := range r.UdiDeviceIdentifier {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("UdiDeviceIdentifier[%d]: %w", i, err)
		}
	}
	for i, item := range r.Packaging {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Packaging[%d]: %w", i, err)
		}
	}
	return nil
}

type DeviceDefinitionMaterial struct {
	Id                  *string          `json:"id,omitempty" bson:"id,omitempty"`                                    // Unique id for inter-element referencing
	Substance           *CodeableConcept `json:"substance" bson:"substance"`                                          // A relevant substance that the device contains, may contain, or is made of
	Alternate           *bool            `json:"alternate,omitempty" bson:"alternate,omitempty"`                      // Indicates an alternative material of the device
	AllergenicIndicator *bool            `json:"allergenicIndicator,omitempty" bson:"allergenic_indicator,omitempty"` // Whether the substance is a known or suspected allergen
}

func (r *DeviceDefinitionMaterial) Validate() error {
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

type DeviceDefinitionGuideline struct {
	Id               *string           `json:"id,omitempty" bson:"id,omitempty"`                              // Unique id for inter-element referencing
	UseContext       []UsageContext    `json:"useContext,omitempty" bson:"use_context,omitempty"`             // The circumstances that form the setting for using the device
	UsageInstruction *string           `json:"usageInstruction,omitempty" bson:"usage_instruction,omitempty"` // Detailed written and visual directions for the user on how to use the device
	RelatedArtifact  []RelatedArtifact `json:"relatedArtifact,omitempty" bson:"related_artifact,omitempty"`   // A source of information or reference for this guideline
	Indication       []CodeableConcept `json:"indication,omitempty" bson:"indication,omitempty"`              // A clinical condition for which the device was designed to be used
	Contraindication []CodeableConcept `json:"contraindication,omitempty" bson:"contraindication,omitempty"`  // A specific situation when a device should not be used because it may cause harm
	Warning          []CodeableConcept `json:"warning,omitempty" bson:"warning,omitempty"`                    // Specific hazard alert information that a user needs to know before using the device
	IntendedUse      *string           `json:"intendedUse,omitempty" bson:"intended_use,omitempty"`           // A description of the general purpose or medical use of the device or its function
}

func (r *DeviceDefinitionGuideline) Validate() error {
	for i, item := range r.UseContext {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("UseContext[%d]: %w", i, err)
		}
	}
	for i, item := range r.RelatedArtifact {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("RelatedArtifact[%d]: %w", i, err)
		}
	}
	for i, item := range r.Indication {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Indication[%d]: %w", i, err)
		}
	}
	for i, item := range r.Contraindication {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Contraindication[%d]: %w", i, err)
		}
	}
	for i, item := range r.Warning {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Warning[%d]: %w", i, err)
		}
	}
	return nil
}

type DeviceDefinitionCorrectiveAction struct {
	Id     *string `json:"id,omitempty" bson:"id,omitempty"`       // Unique id for inter-element referencing
	Recall bool    `json:"recall" bson:"recall"`                   // Whether the corrective action was a recall
	Scope  *string `json:"scope,omitempty" bson:"scope,omitempty"` // model | lot-numbers | serial-numbers
	Period *Period `json:"period" bson:"period"`                   // Start and end dates of the  corrective action
}

func (r *DeviceDefinitionCorrectiveAction) Validate() error {
	if r.Period == nil {
		return fmt.Errorf("field 'Period' is required")
	}
	if r.Period != nil {
		if err := r.Period.Validate(); err != nil {
			return fmt.Errorf("Period: %w", err)
		}
	}
	return nil
}

type DeviceDefinitionClassification struct {
	Id            *string           `json:"id,omitempty" bson:"id,omitempty"`                       // Unique id for inter-element referencing
	Type          *CodeableConcept  `json:"type" bson:"type"`                                       // A classification or risk class of the device model
	Justification []RelatedArtifact `json:"justification,omitempty" bson:"justification,omitempty"` // Further information qualifying this classification of the device model
}

func (r *DeviceDefinitionClassification) Validate() error {
	if r.Type == nil {
		return fmt.Errorf("field 'Type' is required")
	}
	if r.Type != nil {
		if err := r.Type.Validate(); err != nil {
			return fmt.Errorf("Type: %w", err)
		}
	}
	for i, item := range r.Justification {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Justification[%d]: %w", i, err)
		}
	}
	return nil
}

type DeviceDefinitionHasPart struct {
	Id                        *string          `json:"id,omitempty" bson:"id,omitempty"`                             // Unique id for inter-element referencing
	DefinitionCanonical       *string          `json:"definitionCanonical" bson:"definition_canonical"`              // Reference to the part
	DefinitionCodeableConcept *CodeableConcept `json:"definitionCodeableConcept" bson:"definition_codeable_concept"` // Reference to the part
	Count                     *int             `json:"count,omitempty" bson:"count,omitempty"`                       // Number of occurrences of the part
}

func (r *DeviceDefinitionHasPart) Validate() error {
	if r.DefinitionCanonical == nil {
		return fmt.Errorf("field 'DefinitionCanonical' is required")
	}
	if r.DefinitionCodeableConcept == nil {
		return fmt.Errorf("field 'DefinitionCodeableConcept' is required")
	}
	if r.DefinitionCodeableConcept != nil {
		if err := r.DefinitionCodeableConcept.Validate(); err != nil {
			return fmt.Errorf("DefinitionCodeableConcept: %w", err)
		}
	}
	return nil
}

type DeviceDefinitionProperty struct {
	Id                   *string          `json:"id,omitempty" bson:"id,omitempty"`                   // Unique id for inter-element referencing
	Type                 *CodeableConcept `json:"type" bson:"type"`                                   // Code that specifies the property being represented
	ValueQuantity        *Quantity        `json:"valueQuantity" bson:"value_quantity"`                // Value of the property
	ValueCodeableConcept *CodeableConcept `json:"valueCodeableConcept" bson:"value_codeable_concept"` // Value of the property
	ValueString          *string          `json:"valueString" bson:"value_string"`                    // Value of the property
	ValueBoolean         *bool            `json:"valueBoolean" bson:"value_boolean"`                  // Value of the property
	ValueInteger         *int             `json:"valueInteger" bson:"value_integer"`                  // Value of the property
	ValueRange           *Range           `json:"valueRange" bson:"value_range"`                      // Value of the property
	ValueAttachment      *Attachment      `json:"valueAttachment" bson:"value_attachment"`            // Value of the property
}

func (r *DeviceDefinitionProperty) Validate() error {
	if r.Type == nil {
		return fmt.Errorf("field 'Type' is required")
	}
	if r.Type != nil {
		if err := r.Type.Validate(); err != nil {
			return fmt.Errorf("Type: %w", err)
		}
	}
	if r.ValueQuantity == nil {
		return fmt.Errorf("field 'ValueQuantity' is required")
	}
	if r.ValueQuantity != nil {
		if err := r.ValueQuantity.Validate(); err != nil {
			return fmt.Errorf("ValueQuantity: %w", err)
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
	if r.ValueBoolean == nil {
		return fmt.Errorf("field 'ValueBoolean' is required")
	}
	if r.ValueInteger == nil {
		return fmt.Errorf("field 'ValueInteger' is required")
	}
	if r.ValueRange == nil {
		return fmt.Errorf("field 'ValueRange' is required")
	}
	if r.ValueRange != nil {
		if err := r.ValueRange.Validate(); err != nil {
			return fmt.Errorf("ValueRange: %w", err)
		}
	}
	if r.ValueAttachment == nil {
		return fmt.Errorf("field 'ValueAttachment' is required")
	}
	if r.ValueAttachment != nil {
		if err := r.ValueAttachment.Validate(); err != nil {
			return fmt.Errorf("ValueAttachment: %w", err)
		}
	}
	return nil
}

type DeviceDefinitionChargeItem struct {
	Id              *string            `json:"id,omitempty" bson:"id,omitempty"`                            // Unique id for inter-element referencing
	ChargeItemCode  *CodeableReference `json:"chargeItemCode" bson:"charge_item_code"`                      // The code or reference for the charge item
	Count           *Quantity          `json:"count" bson:"count"`                                          // Coefficient applicable to the billing code
	EffectivePeriod *Period            `json:"effectivePeriod,omitempty" bson:"effective_period,omitempty"` // A specific time period in which this charge item applies
	UseContext      []UsageContext     `json:"useContext,omitempty" bson:"use_context,omitempty"`           // The context to which this charge item applies
}

func (r *DeviceDefinitionChargeItem) Validate() error {
	if r.ChargeItemCode == nil {
		return fmt.Errorf("field 'ChargeItemCode' is required")
	}
	if r.ChargeItemCode != nil {
		if err := r.ChargeItemCode.Validate(); err != nil {
			return fmt.Errorf("ChargeItemCode: %w", err)
		}
	}
	if r.Count == nil {
		return fmt.Errorf("field 'Count' is required")
	}
	if r.Count != nil {
		if err := r.Count.Validate(); err != nil {
			return fmt.Errorf("Count: %w", err)
		}
	}
	if r.EffectivePeriod != nil {
		if err := r.EffectivePeriod.Validate(); err != nil {
			return fmt.Errorf("EffectivePeriod: %w", err)
		}
	}
	for i, item := range r.UseContext {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("UseContext[%d]: %w", i, err)
		}
	}
	return nil
}

type DeviceDefinitionLink struct {
	Id                           *string          `json:"id,omitempty" bson:"id,omitempty"`                                    // Unique id for inter-element referencing
	Relation                     *Coding          `json:"relation" bson:"relation"`                                            // The type indicates the relationship of the related device to the device instance
	RelatedDeviceCanonical       *string          `json:"relatedDeviceCanonical" bson:"related_device_canonical"`              // A reference to the linked device
	RelatedDeviceCodeableConcept *CodeableConcept `json:"relatedDeviceCodeableConcept" bson:"related_device_codeable_concept"` // A reference to the linked device
}

func (r *DeviceDefinitionLink) Validate() error {
	if r.Relation == nil {
		return fmt.Errorf("field 'Relation' is required")
	}
	if r.Relation != nil {
		if err := r.Relation.Validate(); err != nil {
			return fmt.Errorf("Relation: %w", err)
		}
	}
	if r.RelatedDeviceCanonical == nil {
		return fmt.Errorf("field 'RelatedDeviceCanonical' is required")
	}
	if r.RelatedDeviceCodeableConcept == nil {
		return fmt.Errorf("field 'RelatedDeviceCodeableConcept' is required")
	}
	if r.RelatedDeviceCodeableConcept != nil {
		if err := r.RelatedDeviceCodeableConcept.Validate(); err != nil {
			return fmt.Errorf("RelatedDeviceCodeableConcept: %w", err)
		}
	}
	return nil
}

type DeviceDefinitionDeviceName struct {
	Id   *string          `json:"id,omitempty" bson:"id,omitempty"` // Unique id for inter-element referencing
	Name string           `json:"name" bson:"name"`                 // A name that is used to refer to the device
	Type *CodeableConcept `json:"type" bson:"type"`                 // registered-name | user-friendly-name | patient-reported-name
}

func (r *DeviceDefinitionDeviceName) Validate() error {
	var emptyString string
	if r.Name == emptyString {
		return fmt.Errorf("field 'Name' is required")
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

type DeviceDefinitionPackagingDistributor struct {
	Id                    *string     `json:"id,omitempty" bson:"id,omitempty"`                                        // Unique id for inter-element referencing
	Name                  *string     `json:"name,omitempty" bson:"name,omitempty"`                                    // Distributor's human-readable name
	OrganizationReference []Reference `json:"organizationReference,omitempty" bson:"organization_reference,omitempty"` // Distributor as an Organization resource
}

func (r *DeviceDefinitionPackagingDistributor) Validate() error {
	for i, item := range r.OrganizationReference {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("OrganizationReference[%d]: %w", i, err)
		}
	}
	return nil
}
