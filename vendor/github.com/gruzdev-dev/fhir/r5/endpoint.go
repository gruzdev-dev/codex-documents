package models

import (
	"encoding/json"
	"fmt"
)

// The technical details of an endpoint that can be used for electronic services, such as for web services providing XDS.b, a REST endpoint for another FHIR server, or a s/Mime email address. This may include any security context information.
type Endpoint struct {
	ResourceType         string            `json:"resourceType" bson:"resource_type"`                                     // Type of resource
	Id                   *string           `json:"id,omitempty" bson:"id,omitempty"`                                      // Logical id of this artifact
	Meta                 *Meta             `json:"meta,omitempty" bson:"meta,omitempty"`                                  // Metadata about the resource
	ImplicitRules        *string           `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`               // A set of rules under which this content was created
	Language             *string           `json:"language,omitempty" bson:"language,omitempty"`                          // Language of the resource content
	Text                 *Narrative        `json:"text,omitempty" bson:"text,omitempty"`                                  // Text summary of the resource, for human interpretation
	Contained            []json.RawMessage `json:"contained,omitempty" bson:"contained,omitempty"`                        // Contained, inline Resources
	Identifier           []Identifier      `json:"identifier,omitempty" bson:"identifier,omitempty"`                      // Identifies this endpoint across multiple systems
	Status               string            `json:"status" bson:"status"`                                                  // active | limited | suspended | error | off | entered-in-error
	ConnectionType       []CodeableConcept `json:"connectionType" bson:"connection_type"`                                 // Protocol/Profile/Standard to be used with this endpoint connection
	Name                 *string           `json:"name,omitempty" bson:"name,omitempty"`                                  // A name that this endpoint can be identified by
	Description          *string           `json:"description,omitempty" bson:"description,omitempty"`                    // Additional details about the endpoint that could be displayed as further information to identify the description beyond its name
	EnvironmentType      []CodeableConcept `json:"environmentType,omitempty" bson:"environment_type,omitempty"`           // The type of environment(s) exposed at this endpoint
	ManagingOrganization *Reference        `json:"managingOrganization,omitempty" bson:"managing_organization,omitempty"` // Organization that manages this endpoint (might not be the organization that exposes the endpoint)
	Contact              []ContactPoint    `json:"contact,omitempty" bson:"contact,omitempty"`                            // Contact details for source (e.g. troubleshooting)
	Period               *Period           `json:"period,omitempty" bson:"period,omitempty"`                              // Interval the endpoint is expected to be operational
	Availability         *Availability     `json:"availability,omitempty" bson:"availability,omitempty"`                  // Times the endpoint is expected to be available (including exceptions)
	Payload              []EndpointPayload `json:"payload,omitempty" bson:"payload,omitempty"`                            // Set of payloads that are provided by this endpoint
	Address              string            `json:"address" bson:"address"`                                                // The technical base address for connecting to this endpoint
	Header               []string          `json:"header,omitempty" bson:"header,omitempty"`                              // Usage depends on the channel type
}

func (r *Endpoint) Validate() error {
	if r.ResourceType != "Endpoint" {
		return fmt.Errorf("invalid resourceType: expected 'Endpoint', got '%s'", r.ResourceType)
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
	var emptyString string
	if r.Status == emptyString {
		return fmt.Errorf("field 'Status' is required")
	}
	if len(r.ConnectionType) < 1 {
		return fmt.Errorf("field 'ConnectionType' must have at least 1 elements")
	}
	for i, item := range r.ConnectionType {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("ConnectionType[%d]: %w", i, err)
		}
	}
	for i, item := range r.EnvironmentType {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("EnvironmentType[%d]: %w", i, err)
		}
	}
	if r.ManagingOrganization != nil {
		if err := r.ManagingOrganization.Validate(); err != nil {
			return fmt.Errorf("ManagingOrganization: %w", err)
		}
	}
	for i, item := range r.Contact {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Contact[%d]: %w", i, err)
		}
	}
	if r.Period != nil {
		if err := r.Period.Validate(); err != nil {
			return fmt.Errorf("Period: %w", err)
		}
	}
	if r.Availability != nil {
		if err := r.Availability.Validate(); err != nil {
			return fmt.Errorf("Availability: %w", err)
		}
	}
	for i, item := range r.Payload {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Payload[%d]: %w", i, err)
		}
	}
	if r.Address == emptyString {
		return fmt.Errorf("field 'Address' is required")
	}
	return nil
}

type EndpointPayload struct {
	Id               *string           `json:"id,omitempty" bson:"id,omitempty"`                              // Unique id for inter-element referencing
	Type             []CodeableConcept `json:"type,omitempty" bson:"type,omitempty"`                          // The type of content that may be used at this endpoint (e.g. XDS Discharge summaries)
	MimeType         []string          `json:"mimeType,omitempty" bson:"mime_type,omitempty"`                 // Mimetype to send. If not specified, the content could be anything (including no payload, if the connectionType defined this)
	ProfileCanonical []string          `json:"profileCanonical,omitempty" bson:"profile_canonical,omitempty"` // The profile that is expected at this endpoint
	ProfileUri       []string          `json:"profileUri,omitempty" bson:"profile_uri,omitempty"`             // The non-fhir based profile that is expected at this endpoint
}

func (r *EndpointPayload) Validate() error {
	for i, item := range r.Type {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Type[%d]: %w", i, err)
		}
	}
	return nil
}
