package models

import (
	"encoding/json"
	"fmt"
)

// A set of rules of how a particular interoperability or standards problem is solved - typically through the use of FHIR resources. This resource is used to gather all the parts of an implementation guide into a logical whole and to publish a computable definition of all the parts.
type ImplementationGuide struct {
	Id                     *string                        `json:"id,omitempty" bson:"id,omitempty"`                                           // Logical id of this artifact
	Meta                   *Meta                          `json:"meta,omitempty" bson:"meta,omitempty"`                                       // Metadata about the resource
	ImplicitRules          *string                        `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`                    // A set of rules under which this content was created
	Language               *string                        `json:"language,omitempty" bson:"language,omitempty"`                               // Language of the resource content
	Text                   *Narrative                     `json:"text,omitempty" bson:"text,omitempty"`                                       // Text summary of the resource, for human interpretation
	Contained              []json.RawMessage              `json:"contained,omitempty" bson:"contained,omitempty"`                             // Contained, inline Resources
	Url                    string                         `json:"url" bson:"url"`                                                             // Canonical identifier for this implementation guide, represented as a URI (globally unique)
	Identifier             []Identifier                   `json:"identifier,omitempty" bson:"identifier,omitempty"`                           // Additional identifier for the implementation guide (business identifier)
	Version                *string                        `json:"version,omitempty" bson:"version,omitempty"`                                 // Business version of the implementation guide
	VersionAlgorithmString *string                        `json:"versionAlgorithmString,omitempty" bson:"version_algorithm_string,omitempty"` // How to compare versions
	VersionAlgorithmCoding *Coding                        `json:"versionAlgorithmCoding,omitempty" bson:"version_algorithm_coding,omitempty"` // How to compare versions
	Name                   string                         `json:"name" bson:"name"`                                                           // Name for this implementation guide (computer friendly)
	Title                  *string                        `json:"title,omitempty" bson:"title,omitempty"`                                     // Name for this implementation guide (human friendly)
	Status                 string                         `json:"status" bson:"status"`                                                       // draft | active | retired | unknown
	Experimental           bool                           `json:"experimental,omitempty" bson:"experimental,omitempty"`                       // For testing only - never for real usage
	Date                   *string                        `json:"date,omitempty" bson:"date,omitempty"`                                       // Date last changed
	Publisher              *string                        `json:"publisher,omitempty" bson:"publisher,omitempty"`                             // Name of the publisher/steward (organization or individual)
	Contact                []ContactDetail                `json:"contact,omitempty" bson:"contact,omitempty"`                                 // Contact details for the publisher
	Description            *string                        `json:"description,omitempty" bson:"description,omitempty"`                         // Natural language description of the implementation guide
	UseContext             []UsageContext                 `json:"useContext,omitempty" bson:"use_context,omitempty"`                          // The context that the content is intended to support
	Jurisdiction           []CodeableConcept              `json:"jurisdiction,omitempty" bson:"jurisdiction,omitempty"`                       // Jurisdiction of the authority that maintains the implementation guide (if applicable)
	Purpose                *string                        `json:"purpose,omitempty" bson:"purpose,omitempty"`                                 // Why this implementation guide is defined
	Copyright              *string                        `json:"copyright,omitempty" bson:"copyright,omitempty"`                             // Notice about intellectual property ownership, can include restrictions on use
	CopyrightLabel         *string                        `json:"copyrightLabel,omitempty" bson:"copyright_label,omitempty"`                  // Copyright holder and year(s)
	PackageId              string                         `json:"packageId" bson:"package_id"`                                                // NPM Package name for IG
	License                *string                        `json:"license,omitempty" bson:"license,omitempty"`                                 // SPDX license code for this IG (or not-open-source)
	FhirVersion            []string                       `json:"fhirVersion" bson:"fhir_version"`                                            // FHIR Version(s) this Implementation Guide targets
	DependsOn              []ImplementationGuideDependsOn `json:"dependsOn,omitempty" bson:"depends_on,omitempty"`                            // Another Implementation guide this depends on
	Global                 []ImplementationGuideGlobal    `json:"global,omitempty" bson:"global,omitempty"`                                   // Profiles that apply globally
	Definition             *ImplementationGuideDefinition `json:"definition,omitempty" bson:"definition,omitempty"`                           // Information needed to build the IG
	Manifest               *ImplementationGuideManifest   `json:"manifest,omitempty" bson:"manifest,omitempty"`                               // Information about an assembled IG
}

func (r *ImplementationGuide) Validate() error {
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
	var emptyString string
	if r.Url == emptyString {
		return fmt.Errorf("field 'Url' is required")
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
	if r.Name == emptyString {
		return fmt.Errorf("field 'Name' is required")
	}
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
	if r.PackageId == emptyString {
		return fmt.Errorf("field 'PackageId' is required")
	}
	if len(r.FhirVersion) < 1 {
		return fmt.Errorf("field 'FhirVersion' must have at least 1 elements")
	}
	for i, item := range r.DependsOn {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("DependsOn[%d]: %w", i, err)
		}
	}
	for i, item := range r.Global {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Global[%d]: %w", i, err)
		}
	}
	if r.Definition != nil {
		if err := r.Definition.Validate(); err != nil {
			return fmt.Errorf("Definition: %w", err)
		}
	}
	if r.Manifest != nil {
		if err := r.Manifest.Validate(); err != nil {
			return fmt.Errorf("Manifest: %w", err)
		}
	}
	return nil
}

type ImplementationGuideGlobal struct {
	Id      *string `json:"id,omitempty" bson:"id,omitempty"` // Unique id for inter-element referencing
	Type    string  `json:"type" bson:"type"`                 // Type this profile applies to
	Profile string  `json:"profile" bson:"profile"`           // Profile that all resources must conform to
}

func (r *ImplementationGuideGlobal) Validate() error {
	var emptyString string
	if r.Type == emptyString {
		return fmt.Errorf("field 'Type' is required")
	}
	if r.Profile == emptyString {
		return fmt.Errorf("field 'Profile' is required")
	}
	return nil
}

type ImplementationGuideDefinition struct {
	Id        *string                                  `json:"id,omitempty" bson:"id,omitempty"`               // Unique id for inter-element referencing
	Grouping  []ImplementationGuideDefinitionGrouping  `json:"grouping,omitempty" bson:"grouping,omitempty"`   // Grouping used to present related resources in the IG
	Resource  []ImplementationGuideDefinitionResource  `json:"resource,omitempty" bson:"resource,omitempty"`   // Resource in the implementation guide
	Page      *ImplementationGuideDefinitionPage       `json:"page,omitempty" bson:"page,omitempty"`           // Page/Section in the Guide
	Parameter []ImplementationGuideDefinitionParameter `json:"parameter,omitempty" bson:"parameter,omitempty"` // Defines how IG is built by tools
	Template  []ImplementationGuideDefinitionTemplate  `json:"template,omitempty" bson:"template,omitempty"`   // A template for building resources
}

func (r *ImplementationGuideDefinition) Validate() error {
	for i, item := range r.Grouping {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Grouping[%d]: %w", i, err)
		}
	}
	for i, item := range r.Resource {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Resource[%d]: %w", i, err)
		}
	}
	if r.Page != nil {
		if err := r.Page.Validate(); err != nil {
			return fmt.Errorf("Page: %w", err)
		}
	}
	for i, item := range r.Parameter {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Parameter[%d]: %w", i, err)
		}
	}
	for i, item := range r.Template {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Template[%d]: %w", i, err)
		}
	}
	return nil
}

type ImplementationGuideDefinitionResource struct {
	Id          *string    `json:"id,omitempty" bson:"id,omitempty"`                    // Unique id for inter-element referencing
	Reference   *Reference `json:"reference" bson:"reference"`                          // Location of the resource
	FhirVersion []string   `json:"fhirVersion,omitempty" bson:"fhir_version,omitempty"` // Versions this applies to (if different to IG)
	Name        *string    `json:"name,omitempty" bson:"name,omitempty"`                // Human readable name for the resource
	Description *string    `json:"description,omitempty" bson:"description,omitempty"`  // Reason why included in guide
	IsExample   bool       `json:"isExample,omitempty" bson:"is_example,omitempty"`     // Is this an example
	Profile     []string   `json:"profile,omitempty" bson:"profile,omitempty"`          // Profile(s) this resource is valid against
	GroupingId  *string    `json:"groupingId,omitempty" bson:"grouping_id,omitempty"`   // Grouping this is part of
}

func (r *ImplementationGuideDefinitionResource) Validate() error {
	if r.Reference == nil {
		return fmt.Errorf("field 'Reference' is required")
	}
	if r.Reference != nil {
		if err := r.Reference.Validate(); err != nil {
			return fmt.Errorf("Reference: %w", err)
		}
	}
	return nil
}

type ImplementationGuideManifestResource struct {
	Id           *string    `json:"id,omitempty" bson:"id,omitempty"`                      // Unique id for inter-element referencing
	Reference    *Reference `json:"reference" bson:"reference"`                            // Location of the resource
	IsExample    bool       `json:"isExample,omitempty" bson:"is_example,omitempty"`       // Is this an example
	Profile      []string   `json:"profile,omitempty" bson:"profile,omitempty"`            // Profile(s) this is an example of
	RelativePath *string    `json:"relativePath,omitempty" bson:"relative_path,omitempty"` // Relative path for page in IG
}

func (r *ImplementationGuideManifestResource) Validate() error {
	if r.Reference == nil {
		return fmt.Errorf("field 'Reference' is required")
	}
	if r.Reference != nil {
		if err := r.Reference.Validate(); err != nil {
			return fmt.Errorf("Reference: %w", err)
		}
	}
	return nil
}

type ImplementationGuideDependsOn struct {
	Id        *string `json:"id,omitempty" bson:"id,omitempty"`                // Unique id for inter-element referencing
	Uri       string  `json:"uri" bson:"uri"`                                  // Identity of the IG that this depends on
	PackageId *string `json:"packageId,omitempty" bson:"package_id,omitempty"` // NPM Package name for IG this depends on
	Version   *string `json:"version,omitempty" bson:"version,omitempty"`      // Version of the IG
	Reason    *string `json:"reason,omitempty" bson:"reason,omitempty"`        // Why dependency exists
}

func (r *ImplementationGuideDependsOn) Validate() error {
	var emptyString string
	if r.Uri == emptyString {
		return fmt.Errorf("field 'Uri' is required")
	}
	return nil
}

type ImplementationGuideDefinitionGrouping struct {
	Id          *string `json:"id,omitempty" bson:"id,omitempty"`                   // Unique id for inter-element referencing
	Name        string  `json:"name" bson:"name"`                                   // Descriptive name for the package
	Description *string `json:"description,omitempty" bson:"description,omitempty"` // Human readable text describing the package
}

func (r *ImplementationGuideDefinitionGrouping) Validate() error {
	var emptyString string
	if r.Name == emptyString {
		return fmt.Errorf("field 'Name' is required")
	}
	return nil
}

type ImplementationGuideDefinitionPage struct {
	Id             *string                             `json:"id,omitempty" bson:"id,omitempty"`                          // Unique id for inter-element referencing
	SourceUrl      *string                             `json:"sourceUrl,omitempty" bson:"source_url,omitempty"`           // Source for page
	SourceString   *string                             `json:"sourceString,omitempty" bson:"source_string,omitempty"`     // Source for page
	SourceMarkdown *string                             `json:"sourceMarkdown,omitempty" bson:"source_markdown,omitempty"` // Source for page
	Name           string                              `json:"name" bson:"name"`                                          // Name of the page when published
	Title          string                              `json:"title" bson:"title"`                                        // Short title shown for navigational assistance
	Generation     string                              `json:"generation" bson:"generation"`                              // html | markdown | xml | generated
	Page           []ImplementationGuideDefinitionPage `json:"page,omitempty" bson:"page,omitempty"`                      // Nested Pages / Sections
}

func (r *ImplementationGuideDefinitionPage) Validate() error {
	var emptyString string
	if r.Name == emptyString {
		return fmt.Errorf("field 'Name' is required")
	}
	if r.Title == emptyString {
		return fmt.Errorf("field 'Title' is required")
	}
	if r.Generation == emptyString {
		return fmt.Errorf("field 'Generation' is required")
	}
	for i, item := range r.Page {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Page[%d]: %w", i, err)
		}
	}
	return nil
}

type ImplementationGuideDefinitionParameter struct {
	Id    *string `json:"id,omitempty" bson:"id,omitempty"` // Unique id for inter-element referencing
	Code  *Coding `json:"code" bson:"code"`                 // Code that identifies parameter
	Value string  `json:"value" bson:"value"`               // Value for named type
}

func (r *ImplementationGuideDefinitionParameter) Validate() error {
	if r.Code == nil {
		return fmt.Errorf("field 'Code' is required")
	}
	if r.Code != nil {
		if err := r.Code.Validate(); err != nil {
			return fmt.Errorf("Code: %w", err)
		}
	}
	var emptyString string
	if r.Value == emptyString {
		return fmt.Errorf("field 'Value' is required")
	}
	return nil
}

type ImplementationGuideDefinitionTemplate struct {
	Id     *string `json:"id,omitempty" bson:"id,omitempty"`       // Unique id for inter-element referencing
	Code   string  `json:"code" bson:"code"`                       // Type of template specified
	Source string  `json:"source" bson:"source"`                   // The source location for the template
	Scope  *string `json:"scope,omitempty" bson:"scope,omitempty"` // The scope in which the template applies
}

func (r *ImplementationGuideDefinitionTemplate) Validate() error {
	var emptyString string
	if r.Code == emptyString {
		return fmt.Errorf("field 'Code' is required")
	}
	if r.Source == emptyString {
		return fmt.Errorf("field 'Source' is required")
	}
	return nil
}

type ImplementationGuideManifest struct {
	Id        *string                               `json:"id,omitempty" bson:"id,omitempty"`               // Unique id for inter-element referencing
	Rendering *string                               `json:"rendering,omitempty" bson:"rendering,omitempty"` // Location of rendered implementation guide
	Resource  []ImplementationGuideManifestResource `json:"resource" bson:"resource"`                       // Resource in the implementation guide
	Page      []ImplementationGuideManifestPage     `json:"page,omitempty" bson:"page,omitempty"`           // HTML page within the parent IG
	Image     []string                              `json:"image,omitempty" bson:"image,omitempty"`         // Image within the IG
	Other     []string                              `json:"other,omitempty" bson:"other,omitempty"`         // Additional linkable file in IG
}

func (r *ImplementationGuideManifest) Validate() error {
	if len(r.Resource) < 1 {
		return fmt.Errorf("field 'Resource' must have at least 1 elements")
	}
	for i, item := range r.Resource {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Resource[%d]: %w", i, err)
		}
	}
	for i, item := range r.Page {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Page[%d]: %w", i, err)
		}
	}
	return nil
}

type ImplementationGuideManifestPage struct {
	Id     *string  `json:"id,omitempty" bson:"id,omitempty"`         // Unique id for inter-element referencing
	Name   string   `json:"name" bson:"name"`                         // HTML page name
	Title  *string  `json:"title,omitempty" bson:"title,omitempty"`   // Title of the page, for references
	Anchor []string `json:"anchor,omitempty" bson:"anchor,omitempty"` // Anchor available on the page
}

func (r *ImplementationGuideManifestPage) Validate() error {
	var emptyString string
	if r.Name == emptyString {
		return fmt.Errorf("field 'Name' is required")
	}
	return nil
}
