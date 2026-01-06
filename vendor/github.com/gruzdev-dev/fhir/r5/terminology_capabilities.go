package models

import (
	"encoding/json"
	"fmt"
)

// A TerminologyCapabilities resource documents a set of capabilities (behaviors) of a FHIR Terminology Server that may be used as a statement of actual server functionality or a statement of required or desired server implementation.
type TerminologyCapabilities struct {
	Id                     *string                                `json:"id,omitempty" bson:"id,omitempty"`                                           // Logical id of this artifact
	Meta                   *Meta                                  `json:"meta,omitempty" bson:"meta,omitempty"`                                       // Metadata about the resource
	ImplicitRules          *string                                `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`                    // A set of rules under which this content was created
	Language               *string                                `json:"language,omitempty" bson:"language,omitempty"`                               // Language of the resource content
	Text                   *Narrative                             `json:"text,omitempty" bson:"text,omitempty"`                                       // Text summary of the resource, for human interpretation
	Contained              []json.RawMessage                      `json:"contained,omitempty" bson:"contained,omitempty"`                             // Contained, inline Resources
	Url                    *string                                `json:"url,omitempty" bson:"url,omitempty"`                                         // Canonical identifier for this terminology capabilities, represented as a URI (globally unique)
	Identifier             []Identifier                           `json:"identifier,omitempty" bson:"identifier,omitempty"`                           // Additional identifier for the terminology capabilities
	Version                *string                                `json:"version,omitempty" bson:"version,omitempty"`                                 // Business version of the terminology capabilities
	VersionAlgorithmString *string                                `json:"versionAlgorithmString,omitempty" bson:"version_algorithm_string,omitempty"` // How to compare versions
	VersionAlgorithmCoding *Coding                                `json:"versionAlgorithmCoding,omitempty" bson:"version_algorithm_coding,omitempty"` // How to compare versions
	Name                   *string                                `json:"name,omitempty" bson:"name,omitempty"`                                       // Name for this terminology capabilities (computer friendly)
	Title                  *string                                `json:"title,omitempty" bson:"title,omitempty"`                                     // Name for this terminology capabilities (human friendly)
	Status                 string                                 `json:"status" bson:"status"`                                                       // draft | active | retired | unknown
	Experimental           bool                                   `json:"experimental,omitempty" bson:"experimental,omitempty"`                       // For testing only - never for real usage
	Date                   string                                 `json:"date" bson:"date"`                                                           // Date last changed
	Publisher              *string                                `json:"publisher,omitempty" bson:"publisher,omitempty"`                             // Name of the publisher/steward (organization or individual)
	Contact                []ContactDetail                        `json:"contact,omitempty" bson:"contact,omitempty"`                                 // Contact details for the publisher
	Description            *string                                `json:"description,omitempty" bson:"description,omitempty"`                         // Natural language description of the terminology capabilities
	UseContext             []UsageContext                         `json:"useContext,omitempty" bson:"use_context,omitempty"`                          // The context that the content is intended to support
	Jurisdiction           []CodeableConcept                      `json:"jurisdiction,omitempty" bson:"jurisdiction,omitempty"`                       // Jurisdiction of the authority that maintains the terminology capabilities (if applicable)
	Purpose                *string                                `json:"purpose,omitempty" bson:"purpose,omitempty"`                                 // Why this terminology capabilities is defined
	Copyright              *string                                `json:"copyright,omitempty" bson:"copyright,omitempty"`                             // Notice about intellectual property ownership, can include restrictions on use
	CopyrightLabel         *string                                `json:"copyrightLabel,omitempty" bson:"copyright_label,omitempty"`                  // Copyright holder and year(s)
	Kind                   string                                 `json:"kind" bson:"kind"`                                                           // instance | capability | requirements
	Software               *TerminologyCapabilitiesSoftware       `json:"software,omitempty" bson:"software,omitempty"`                               // Software that is covered by this terminology capability statement
	Implementation         *TerminologyCapabilitiesImplementation `json:"implementation,omitempty" bson:"implementation,omitempty"`                   // If this describes a specific instance
	LockedDate             bool                                   `json:"lockedDate,omitempty" bson:"locked_date,omitempty"`                          // Whether lockedDate is supported
	CodeSystem             []TerminologyCapabilitiesCodeSystem    `json:"codeSystem,omitempty" bson:"code_system,omitempty"`                          // A code system supported by the server
	Supplements            *TerminologyCapabilitiesSupplements    `json:"supplements,omitempty" bson:"supplements,omitempty"`                         // Supplement Support Information
	Expansion              *TerminologyCapabilitiesExpansion      `json:"expansion,omitempty" bson:"expansion,omitempty"`                             // Information about the [ValueSet/$expand](valueset-operation-expand.html) operation
	CodeSearch             *string                                `json:"codeSearch,omitempty" bson:"code_search,omitempty"`                          // in-compose | in-expansion | in-compose-or-expansion
	ValidateCode           *TerminologyCapabilitiesValidateCode   `json:"validateCode,omitempty" bson:"validate_code,omitempty"`                      // Information about the [ValueSet/$validate-code](valueset-operation-validate-code.html) operation
	Translation            *TerminologyCapabilitiesTranslation    `json:"translation,omitempty" bson:"translation,omitempty"`                         // Information about the [ConceptMap/$translate](conceptmap-operation-translate.html) operation
}

func (r *TerminologyCapabilities) Validate() error {
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
	if r.Date == emptyString {
		return fmt.Errorf("field 'Date' is required")
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
	if r.Kind == emptyString {
		return fmt.Errorf("field 'Kind' is required")
	}
	if r.Software != nil {
		if err := r.Software.Validate(); err != nil {
			return fmt.Errorf("Software: %w", err)
		}
	}
	if r.Implementation != nil {
		if err := r.Implementation.Validate(); err != nil {
			return fmt.Errorf("Implementation: %w", err)
		}
	}
	for i, item := range r.CodeSystem {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("CodeSystem[%d]: %w", i, err)
		}
	}
	if r.Supplements != nil {
		if err := r.Supplements.Validate(); err != nil {
			return fmt.Errorf("Supplements: %w", err)
		}
	}
	if r.Expansion != nil {
		if err := r.Expansion.Validate(); err != nil {
			return fmt.Errorf("Expansion: %w", err)
		}
	}
	if r.ValidateCode != nil {
		if err := r.ValidateCode.Validate(); err != nil {
			return fmt.Errorf("ValidateCode: %w", err)
		}
	}
	if r.Translation != nil {
		if err := r.Translation.Validate(); err != nil {
			return fmt.Errorf("Translation: %w", err)
		}
	}
	return nil
}

type TerminologyCapabilitiesExpansion struct {
	Id           *string                                     `json:"id,omitempty" bson:"id,omitempty"`                     // Unique id for inter-element referencing
	Hierarchical bool                                        `json:"hierarchical,omitempty" bson:"hierarchical,omitempty"` // Whether the server can return nested value sets
	Paging       bool                                        `json:"paging,omitempty" bson:"paging,omitempty"`             // Whether the server supports paging on expansion
	Incomplete   bool                                        `json:"incomplete,omitempty" bson:"incomplete,omitempty"`     // Allow request for incomplete expansions?
	Parameter    []TerminologyCapabilitiesExpansionParameter `json:"parameter,omitempty" bson:"parameter,omitempty"`       // Supported expansion parameter
	TextFilter   *string                                     `json:"textFilter,omitempty" bson:"text_filter,omitempty"`    // Documentation about text searching works
}

func (r *TerminologyCapabilitiesExpansion) Validate() error {
	for i, item := range r.Parameter {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Parameter[%d]: %w", i, err)
		}
	}
	return nil
}

type TerminologyCapabilitiesExpansionParameter struct {
	Id            *string `json:"id,omitempty" bson:"id,omitempty"`                       // Unique id for inter-element referencing
	Name          string  `json:"name" bson:"name"`                                       // Name of the supported expansion parameter
	Documentation *string `json:"documentation,omitempty" bson:"documentation,omitempty"` // Description of support for parameter
}

func (r *TerminologyCapabilitiesExpansionParameter) Validate() error {
	var emptyString string
	if r.Name == emptyString {
		return fmt.Errorf("field 'Name' is required")
	}
	return nil
}

type TerminologyCapabilitiesSupplements struct {
	Id      *string `json:"id,omitempty" bson:"id,omitempty"`           // Unique id for inter-element referencing
	Globals *string `json:"globals,omitempty" bson:"globals,omitempty"` // not-supported | explicit | implicit
}

func (r *TerminologyCapabilitiesSupplements) Validate() error {
	return nil
}

type TerminologyCapabilitiesValidateCode struct {
	Id           *string `json:"id,omitempty" bson:"id,omitempty"` // Unique id for inter-element referencing
	Translations bool    `json:"translations" bson:"translations"` // Whether translations are validated
}

func (r *TerminologyCapabilitiesValidateCode) Validate() error {
	return nil
}

type TerminologyCapabilitiesTranslation struct {
	Id       *string `json:"id,omitempty" bson:"id,omitempty"` // Unique id for inter-element referencing
	NeedsMap bool    `json:"needsMap" bson:"needs_map"`        // Whether the client must identify the map
}

func (r *TerminologyCapabilitiesTranslation) Validate() error {
	return nil
}

type TerminologyCapabilitiesSoftware struct {
	Id      *string `json:"id,omitempty" bson:"id,omitempty"`           // Unique id for inter-element referencing
	Name    string  `json:"name" bson:"name"`                           // A name the software is known by
	Version *string `json:"version,omitempty" bson:"version,omitempty"` // Version covered by this statement
}

func (r *TerminologyCapabilitiesSoftware) Validate() error {
	var emptyString string
	if r.Name == emptyString {
		return fmt.Errorf("field 'Name' is required")
	}
	return nil
}

type TerminologyCapabilitiesImplementation struct {
	Id          *string `json:"id,omitempty" bson:"id,omitempty"`   // Unique id for inter-element referencing
	Description string  `json:"description" bson:"description"`     // Describes this specific instance
	Url         *string `json:"url,omitempty" bson:"url,omitempty"` // Base URL for the implementation
}

func (r *TerminologyCapabilitiesImplementation) Validate() error {
	var emptyString string
	if r.Description == emptyString {
		return fmt.Errorf("field 'Description' is required")
	}
	return nil
}

type TerminologyCapabilitiesCodeSystem struct {
	Id          *string                                    `json:"id,omitempty" bson:"id,omitempty"`                   // Unique id for inter-element referencing
	Uri         *string                                    `json:"uri,omitempty" bson:"uri,omitempty"`                 // Canonical identifier for the code system, represented as a URI (no version portion)
	Supplement  []string                                   `json:"supplement,omitempty" bson:"supplement,omitempty"`   // Canonical identifier for a supported supplement to this code system (including supplement version)
	Version     []TerminologyCapabilitiesCodeSystemVersion `json:"version,omitempty" bson:"version,omitempty"`         // Version of Code System supported
	Content     string                                     `json:"content" bson:"content"`                             // not-present | example | fragment | complete | supplement
	Subsumption bool                                       `json:"subsumption,omitempty" bson:"subsumption,omitempty"` // Whether subsumption is supported
}

func (r *TerminologyCapabilitiesCodeSystem) Validate() error {
	for i, item := range r.Version {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Version[%d]: %w", i, err)
		}
	}
	var emptyString string
	if r.Content == emptyString {
		return fmt.Errorf("field 'Content' is required")
	}
	return nil
}

type TerminologyCapabilitiesCodeSystemVersion struct {
	Id            *string                                          `json:"id,omitempty" bson:"id,omitempty"`                       // Unique id for inter-element referencing
	Code          *string                                          `json:"code,omitempty" bson:"code,omitempty"`                   // Version identifier for this version
	IsDefault     bool                                             `json:"isDefault,omitempty" bson:"is_default,omitempty"`        // If this is the default version for this code system
	Supplement    []string                                         `json:"supplement,omitempty" bson:"supplement,omitempty"`       // Canonical identifier for a supported supplement to this code system version (including supplement version)
	Compositional bool                                             `json:"compositional,omitempty" bson:"compositional,omitempty"` // If compositional grammar is supported
	Language      []string                                         `json:"language,omitempty" bson:"language,omitempty"`           // Language Displays supported
	Filter        []TerminologyCapabilitiesCodeSystemVersionFilter `json:"filter,omitempty" bson:"filter,omitempty"`               // Filter Properties supported
	Property      []string                                         `json:"property,omitempty" bson:"property,omitempty"`           // Properties supported for $lookup
}

func (r *TerminologyCapabilitiesCodeSystemVersion) Validate() error {
	for i, item := range r.Filter {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Filter[%d]: %w", i, err)
		}
	}
	return nil
}

type TerminologyCapabilitiesCodeSystemVersionFilter struct {
	Id   *string  `json:"id,omitempty" bson:"id,omitempty"` // Unique id for inter-element referencing
	Code string   `json:"code" bson:"code"`                 // Code of the property supported
	Op   []string `json:"op" bson:"op"`                     // Operations supported for the property
}

func (r *TerminologyCapabilitiesCodeSystemVersionFilter) Validate() error {
	var emptyString string
	if r.Code == emptyString {
		return fmt.Errorf("field 'Code' is required")
	}
	if len(r.Op) < 1 {
		return fmt.Errorf("field 'Op' must have at least 1 elements")
	}
	return nil
}
