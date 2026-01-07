package models

import (
	"encoding/json"
	"fmt"
)

// A Capability Statement documents a set of capabilities (behaviors) of a FHIR Server or Client for a particular version of FHIR that may be used as a statement of actual server functionality or a statement of required or desired server implementation.
type CapabilityStatement struct {
	Id                     *string                            `json:"id,omitempty" bson:"id,omitempty"`                                           // Logical id of this artifact
	Meta                   *Meta                              `json:"meta,omitempty" bson:"meta,omitempty"`                                       // Metadata about the resource
	ImplicitRules          *string                            `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`                    // A set of rules under which this content was created
	Language               *string                            `json:"language,omitempty" bson:"language,omitempty"`                               // Language of the resource content
	Text                   *Narrative                         `json:"text,omitempty" bson:"text,omitempty"`                                       // Text summary of the resource, for human interpretation
	Contained              []json.RawMessage                  `json:"contained,omitempty" bson:"contained,omitempty"`                             // Contained, inline Resources
	Url                    *string                            `json:"url,omitempty" bson:"url,omitempty"`                                         // Canonical identifier for this capability statement, represented as a URI (globally unique)
	Identifier             []Identifier                       `json:"identifier,omitempty" bson:"identifier,omitempty"`                           // Additional identifier for the CapabilityStatement (business identifier)
	Version                *string                            `json:"version,omitempty" bson:"version,omitempty"`                                 // Business version of the capability statement
	VersionAlgorithmString *string                            `json:"versionAlgorithmString,omitempty" bson:"version_algorithm_string,omitempty"` // How to compare versions
	VersionAlgorithmCoding *Coding                            `json:"versionAlgorithmCoding,omitempty" bson:"version_algorithm_coding,omitempty"` // How to compare versions
	Name                   *string                            `json:"name,omitempty" bson:"name,omitempty"`                                       // Name for this capability statement (computer friendly)
	Title                  *string                            `json:"title,omitempty" bson:"title,omitempty"`                                     // Name for this capability statement (human friendly)
	Status                 string                             `json:"status" bson:"status"`                                                       // draft | active | retired | unknown
	Experimental           bool                               `json:"experimental,omitempty" bson:"experimental,omitempty"`                       // For testing only - never for real usage
	Date                   string                             `json:"date" bson:"date"`                                                           // Date last changed
	Publisher              *string                            `json:"publisher,omitempty" bson:"publisher,omitempty"`                             // Name of the publisher/steward (organization or individual)
	Contact                []ContactDetail                    `json:"contact,omitempty" bson:"contact,omitempty"`                                 // Contact details for the publisher
	Description            *string                            `json:"description,omitempty" bson:"description,omitempty"`                         // Natural language description of the capability statement
	UseContext             []UsageContext                     `json:"useContext,omitempty" bson:"use_context,omitempty"`                          // The context that the content is intended to support
	ActorDefinition        []string                           `json:"actorDefinition,omitempty" bson:"actor_definition,omitempty"`                // ActorDefinitions the CapabilityStatement supports
	Jurisdiction           []CodeableConcept                  `json:"jurisdiction,omitempty" bson:"jurisdiction,omitempty"`                       // Jurisdiction of the authority that maintains the capability statement (if applicable)
	Purpose                *string                            `json:"purpose,omitempty" bson:"purpose,omitempty"`                                 // Why this capability statement is defined
	Copyright              *string                            `json:"copyright,omitempty" bson:"copyright,omitempty"`                             // Notice about intellectual property ownership, can include restrictions on use
	CopyrightLabel         *string                            `json:"copyrightLabel,omitempty" bson:"copyright_label,omitempty"`                  // Copyright holder and year(s)
	Kind                   string                             `json:"kind" bson:"kind"`                                                           // instance | capability | requirements
	Instantiates           []string                           `json:"instantiates,omitempty" bson:"instantiates,omitempty"`                       // Canonical URL of another capability statement this implements
	Imports                []string                           `json:"imports,omitempty" bson:"imports,omitempty"`                                 // Canonical URL of another capability statement this adds to
	Software               *CapabilityStatementSoftware       `json:"software,omitempty" bson:"software,omitempty"`                               // Software that is covered by this capability statement
	Implementation         *CapabilityStatementImplementation `json:"implementation,omitempty" bson:"implementation,omitempty"`                   // If this describes a specific instance
	FhirVersion            string                             `json:"fhirVersion" bson:"fhir_version"`                                            // FHIR Version the system supports
	Format                 []string                           `json:"format" bson:"format"`                                                       // formats supported (xml | json | ttl | mime type)
	PatchFormat            []string                           `json:"patchFormat,omitempty" bson:"patch_format,omitempty"`                        // Patch formats supported (Mime types for FHIR and JSON And XML Patch)
	AcceptLanguage         []string                           `json:"acceptLanguage,omitempty" bson:"accept_language,omitempty"`                  // Languages supported
	ImplementationGuide    []string                           `json:"implementationGuide,omitempty" bson:"implementation_guide,omitempty"`        // Implementation guides supported
	Rest                   []CapabilityStatementRest          `json:"rest,omitempty" bson:"rest,omitempty"`                                       // If the endpoint is a RESTful one
	Messaging              []CapabilityStatementMessaging     `json:"messaging,omitempty" bson:"messaging,omitempty"`                             // If messaging is supported
	Document               []CapabilityStatementDocument      `json:"document,omitempty" bson:"document,omitempty"`                               // Document definition
}

func (r *CapabilityStatement) Validate() error {
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
	if r.FhirVersion == emptyString {
		return fmt.Errorf("field 'FhirVersion' is required")
	}
	if len(r.Format) < 1 {
		return fmt.Errorf("field 'Format' must have at least 1 elements")
	}
	for i, item := range r.Rest {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Rest[%d]: %w", i, err)
		}
	}
	for i, item := range r.Messaging {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Messaging[%d]: %w", i, err)
		}
	}
	for i, item := range r.Document {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Document[%d]: %w", i, err)
		}
	}
	return nil
}

type CapabilityStatementSoftware struct {
	Id          *string `json:"id,omitempty" bson:"id,omitempty"`                    // Unique id for inter-element referencing
	Name        string  `json:"name" bson:"name"`                                    // A name the software is known by
	Version     *string `json:"version,omitempty" bson:"version,omitempty"`          // Version covered by this statement
	ReleaseDate *string `json:"releaseDate,omitempty" bson:"release_date,omitempty"` // Date this version was released
}

func (r *CapabilityStatementSoftware) Validate() error {
	var emptyString string
	if r.Name == emptyString {
		return fmt.Errorf("field 'Name' is required")
	}
	return nil
}

type CapabilityStatementImplementation struct {
	Id          *string    `json:"id,omitempty" bson:"id,omitempty"`               // Unique id for inter-element referencing
	Description string     `json:"description" bson:"description"`                 // Describes this specific instance
	Url         *string    `json:"url,omitempty" bson:"url,omitempty"`             // Base URL for the installation
	Custodian   *Reference `json:"custodian,omitempty" bson:"custodian,omitempty"` // Organization that manages the data
}

func (r *CapabilityStatementImplementation) Validate() error {
	var emptyString string
	if r.Description == emptyString {
		return fmt.Errorf("field 'Description' is required")
	}
	if r.Custodian != nil {
		if err := r.Custodian.Validate(); err != nil {
			return fmt.Errorf("Custodian: %w", err)
		}
	}
	return nil
}

type CapabilityStatementRestSecurity struct {
	Id          *string           `json:"id,omitempty" bson:"id,omitempty"`                   // Unique id for inter-element referencing
	Cors        bool              `json:"cors,omitempty" bson:"cors,omitempty"`               // Adds CORS Headers (http://enable-cors.org/)
	Service     []CodeableConcept `json:"service,omitempty" bson:"service,omitempty"`         // OAuth | SMART-on-FHIR | NTLM | Basic | Kerberos | Certificates
	Description *string           `json:"description,omitempty" bson:"description,omitempty"` // General description of how security works
}

func (r *CapabilityStatementRestSecurity) Validate() error {
	for i, item := range r.Service {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Service[%d]: %w", i, err)
		}
	}
	return nil
}

type CapabilityStatementRestResourceInteraction struct {
	Id            *string `json:"id,omitempty" bson:"id,omitempty"`                       // Unique id for inter-element referencing
	Code          string  `json:"code" bson:"code"`                                       // read | vread | update | update-conditional | patch | patch-conditional | delete | delete-conditional-single | delete-conditional-multiple | delete-history | delete-history-version | history-instance | history-type | create | create-conditional | search-type
	Documentation *string `json:"documentation,omitempty" bson:"documentation,omitempty"` // Anything special about interaction behavior
}

func (r *CapabilityStatementRestResourceInteraction) Validate() error {
	var emptyString string
	if r.Code == emptyString {
		return fmt.Errorf("field 'Code' is required")
	}
	return nil
}

type CapabilityStatementRestResourceSearchParam struct {
	Id            *string `json:"id,omitempty" bson:"id,omitempty"`                       // Unique id for inter-element referencing
	Name          string  `json:"name" bson:"name"`                                       // Name for parameter in search url
	Definition    *string `json:"definition,omitempty" bson:"definition,omitempty"`       // Source of definition for parameter
	Type          string  `json:"type" bson:"type"`                                       // number | date | string | token | reference | composite | quantity | uri | special | resource
	Documentation *string `json:"documentation,omitempty" bson:"documentation,omitempty"` // Server-specific usage
}

func (r *CapabilityStatementRestResourceSearchParam) Validate() error {
	var emptyString string
	if r.Name == emptyString {
		return fmt.Errorf("field 'Name' is required")
	}
	if r.Type == emptyString {
		return fmt.Errorf("field 'Type' is required")
	}
	return nil
}

type CapabilityStatementRestResourceOperation struct {
	Id            *string `json:"id,omitempty" bson:"id,omitempty"`                       // Unique id for inter-element referencing
	Name          string  `json:"name" bson:"name"`                                       // Name by which the operation/query is invoked
	Definition    string  `json:"definition" bson:"definition"`                           // The defined operation/query
	Documentation *string `json:"documentation,omitempty" bson:"documentation,omitempty"` // Specific details about operation behavior
}

func (r *CapabilityStatementRestResourceOperation) Validate() error {
	var emptyString string
	if r.Name == emptyString {
		return fmt.Errorf("field 'Name' is required")
	}
	if r.Definition == emptyString {
		return fmt.Errorf("field 'Definition' is required")
	}
	return nil
}

type CapabilityStatementRestInteraction struct {
	Id            *string `json:"id,omitempty" bson:"id,omitempty"`                       // Unique id for inter-element referencing
	Code          string  `json:"code" bson:"code"`                                       // transaction | batch | search-system | history-system
	Documentation *string `json:"documentation,omitempty" bson:"documentation,omitempty"` // Anything special about interaction behavior
}

func (r *CapabilityStatementRestInteraction) Validate() error {
	var emptyString string
	if r.Code == emptyString {
		return fmt.Errorf("field 'Code' is required")
	}
	return nil
}

type CapabilityStatementRest struct {
	Id            *string                                      `json:"id,omitempty" bson:"id,omitempty"`                       // Unique id for inter-element referencing
	Mode          string                                       `json:"mode" bson:"mode"`                                       // client | server
	Documentation *string                                      `json:"documentation,omitempty" bson:"documentation,omitempty"` // General description of implementation
	Security      *CapabilityStatementRestSecurity             `json:"security,omitempty" bson:"security,omitempty"`           // Information about security of implementation
	Resource      []CapabilityStatementRestResource            `json:"resource,omitempty" bson:"resource,omitempty"`           // Resource served on the REST interface
	Interaction   []CapabilityStatementRestInteraction         `json:"interaction,omitempty" bson:"interaction,omitempty"`     // What interactions are supported?
	SearchParam   []CapabilityStatementRestResourceSearchParam `json:"searchParam,omitempty" bson:"search_param,omitempty"`    // Search parameters for searching all resources
	Operation     []CapabilityStatementRestResourceOperation   `json:"operation,omitempty" bson:"operation,omitempty"`         // Definition of a system level operation
	Compartment   []string                                     `json:"compartment,omitempty" bson:"compartment,omitempty"`     // Compartments served/used by system
}

func (r *CapabilityStatementRest) Validate() error {
	var emptyString string
	if r.Mode == emptyString {
		return fmt.Errorf("field 'Mode' is required")
	}
	if r.Security != nil {
		if err := r.Security.Validate(); err != nil {
			return fmt.Errorf("Security: %w", err)
		}
	}
	for i, item := range r.Resource {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Resource[%d]: %w", i, err)
		}
	}
	for i, item := range r.Interaction {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Interaction[%d]: %w", i, err)
		}
	}
	for i, item := range r.SearchParam {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("SearchParam[%d]: %w", i, err)
		}
	}
	for i, item := range r.Operation {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Operation[%d]: %w", i, err)
		}
	}
	return nil
}

type CapabilityStatementRestResource struct {
	Id                *string                                      `json:"id,omitempty" bson:"id,omitempty"`                                // Unique id for inter-element referencing
	Type              string                                       `json:"type" bson:"type"`                                                // A resource type that is supported
	Definition        *string                                      `json:"definition,omitempty" bson:"definition,omitempty"`                // The definition for an additional resource
	Profile           *string                                      `json:"profile,omitempty" bson:"profile,omitempty"`                      // System-wide profile
	SupportedProfile  []string                                     `json:"supportedProfile,omitempty" bson:"supported_profile,omitempty"`   // Use-case specific profiles
	Documentation     *string                                      `json:"documentation,omitempty" bson:"documentation,omitempty"`          // Additional information about the use of the resource type
	Interaction       []CapabilityStatementRestResourceInteraction `json:"interaction,omitempty" bson:"interaction,omitempty"`              // What interactions are supported?
	Versioning        *string                                      `json:"versioning,omitempty" bson:"versioning,omitempty"`                // no-version | versioned | versioned-update
	ReadHistory       bool                                         `json:"readHistory,omitempty" bson:"read_history,omitempty"`             // Whether vRead can return past versions
	UpdateCreate      bool                                         `json:"updateCreate,omitempty" bson:"update_create,omitempty"`           // If update can commit to a new identity
	ConditionalCreate bool                                         `json:"conditionalCreate,omitempty" bson:"conditional_create,omitempty"` // If allows/uses conditional create
	ConditionalRead   *string                                      `json:"conditionalRead,omitempty" bson:"conditional_read,omitempty"`     // not-supported | modified-since | not-match | full-support
	ConditionalUpdate bool                                         `json:"conditionalUpdate,omitempty" bson:"conditional_update,omitempty"` // If allows/uses conditional update
	ConditionalPatch  bool                                         `json:"conditionalPatch,omitempty" bson:"conditional_patch,omitempty"`   // If allows/uses conditional patch
	ConditionalDelete *string                                      `json:"conditionalDelete,omitempty" bson:"conditional_delete,omitempty"` // not-supported | single | multiple - how conditional delete is supported
	ReferencePolicy   []string                                     `json:"referencePolicy,omitempty" bson:"reference_policy,omitempty"`     // literal | logical | resolves | enforced | local
	SearchInclude     []string                                     `json:"searchInclude,omitempty" bson:"search_include,omitempty"`         // _include values supported by the server
	SearchRevInclude  []string                                     `json:"searchRevInclude,omitempty" bson:"search_rev_include,omitempty"`  // _revinclude values supported by the server
	SearchParam       []CapabilityStatementRestResourceSearchParam `json:"searchParam,omitempty" bson:"search_param,omitempty"`             // Search parameters supported by implementation
	Operation         []CapabilityStatementRestResourceOperation   `json:"operation,omitempty" bson:"operation,omitempty"`                  // Definition of a resource operation
}

func (r *CapabilityStatementRestResource) Validate() error {
	var emptyString string
	if r.Type == emptyString {
		return fmt.Errorf("field 'Type' is required")
	}
	for i, item := range r.Interaction {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Interaction[%d]: %w", i, err)
		}
	}
	for i, item := range r.SearchParam {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("SearchParam[%d]: %w", i, err)
		}
	}
	for i, item := range r.Operation {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Operation[%d]: %w", i, err)
		}
	}
	return nil
}

type CapabilityStatementMessaging struct {
	Id               *string                                        `json:"id,omitempty" bson:"id,omitempty"`                              // Unique id for inter-element referencing
	Endpoint         []CapabilityStatementMessagingEndpoint         `json:"endpoint,omitempty" bson:"endpoint,omitempty"`                  // Where messages should be sent
	ReliableCache    *int                                           `json:"reliableCache,omitempty" bson:"reliable_cache,omitempty"`       // Reliable Message Cache Length (min)
	Documentation    *string                                        `json:"documentation,omitempty" bson:"documentation,omitempty"`        // Messaging interface behavior details
	SupportedMessage []CapabilityStatementMessagingSupportedMessage `json:"supportedMessage,omitempty" bson:"supported_message,omitempty"` // Messages supported by this system
}

func (r *CapabilityStatementMessaging) Validate() error {
	for i, item := range r.Endpoint {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Endpoint[%d]: %w", i, err)
		}
	}
	for i, item := range r.SupportedMessage {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("SupportedMessage[%d]: %w", i, err)
		}
	}
	return nil
}

type CapabilityStatementMessagingEndpoint struct {
	Id       *string `json:"id,omitempty" bson:"id,omitempty"` // Unique id for inter-element referencing
	Protocol *Coding `json:"protocol" bson:"protocol"`         // http | ftp | mllp +
	Address  string  `json:"address" bson:"address"`           // Network address or identifier of the end-point
}

func (r *CapabilityStatementMessagingEndpoint) Validate() error {
	if r.Protocol == nil {
		return fmt.Errorf("field 'Protocol' is required")
	}
	if r.Protocol != nil {
		if err := r.Protocol.Validate(); err != nil {
			return fmt.Errorf("Protocol: %w", err)
		}
	}
	var emptyString string
	if r.Address == emptyString {
		return fmt.Errorf("field 'Address' is required")
	}
	return nil
}

type CapabilityStatementMessagingSupportedMessage struct {
	Id         *string `json:"id,omitempty" bson:"id,omitempty"` // Unique id for inter-element referencing
	Mode       string  `json:"mode" bson:"mode"`                 // sender | receiver
	Definition string  `json:"definition" bson:"definition"`     // Message supported by this system
}

func (r *CapabilityStatementMessagingSupportedMessage) Validate() error {
	var emptyString string
	if r.Mode == emptyString {
		return fmt.Errorf("field 'Mode' is required")
	}
	if r.Definition == emptyString {
		return fmt.Errorf("field 'Definition' is required")
	}
	return nil
}

type CapabilityStatementDocument struct {
	Id            *string `json:"id,omitempty" bson:"id,omitempty"`                       // Unique id for inter-element referencing
	Mode          string  `json:"mode" bson:"mode"`                                       // producer | consumer
	Documentation *string `json:"documentation,omitempty" bson:"documentation,omitempty"` // Description of document support
	Profile       string  `json:"profile" bson:"profile"`                                 // Constraint on the resources used in the document
}

func (r *CapabilityStatementDocument) Validate() error {
	var emptyString string
	if r.Mode == emptyString {
		return fmt.Errorf("field 'Mode' is required")
	}
	if r.Profile == emptyString {
		return fmt.Errorf("field 'Profile' is required")
	}
	return nil
}
