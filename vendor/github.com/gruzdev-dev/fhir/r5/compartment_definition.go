package models

import (
	"encoding/json"
	"fmt"
)

// A compartment definition that defines how resources are accessed on a server.
type CompartmentDefinition struct {
	ResourceType           string                          `json:"resourceType" bson:"resource_type"`                                          // Type of resource
	Id                     *string                         `json:"id,omitempty" bson:"id,omitempty"`                                           // Logical id of this artifact
	Meta                   *Meta                           `json:"meta,omitempty" bson:"meta,omitempty"`                                       // Metadata about the resource
	ImplicitRules          *string                         `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`                    // A set of rules under which this content was created
	Language               *string                         `json:"language,omitempty" bson:"language,omitempty"`                               // Language of the resource content
	Text                   *Narrative                      `json:"text,omitempty" bson:"text,omitempty"`                                       // Text summary of the resource, for human interpretation
	Contained              []json.RawMessage               `json:"contained,omitempty" bson:"contained,omitempty"`                             // Contained, inline Resources
	Url                    string                          `json:"url" bson:"url"`                                                             // Canonical identifier for this compartment definition, represented as a URI (globally unique)
	Version                *string                         `json:"version,omitempty" bson:"version,omitempty"`                                 // Business version of the compartment definition
	VersionAlgorithmString *string                         `json:"versionAlgorithmString,omitempty" bson:"version_algorithm_string,omitempty"` // How to compare versions
	VersionAlgorithmCoding *Coding                         `json:"versionAlgorithmCoding,omitempty" bson:"version_algorithm_coding,omitempty"` // How to compare versions
	Name                   string                          `json:"name" bson:"name"`                                                           // Name for this compartment definition (computer friendly)
	Title                  *string                         `json:"title,omitempty" bson:"title,omitempty"`                                     // Name for this compartment definition (human friendly)
	Status                 string                          `json:"status" bson:"status"`                                                       // draft | active | retired | unknown
	Experimental           *bool                           `json:"experimental,omitempty" bson:"experimental,omitempty"`                       // For testing only - never for real usage
	Date                   *string                         `json:"date,omitempty" bson:"date,omitempty"`                                       // Date last changed
	Publisher              *string                         `json:"publisher,omitempty" bson:"publisher,omitempty"`                             // Name of the publisher/steward (organization or individual)
	Contact                []ContactDetail                 `json:"contact,omitempty" bson:"contact,omitempty"`                                 // Contact details for the publisher
	Description            *string                         `json:"description,omitempty" bson:"description,omitempty"`                         // Natural language description of the compartment definition
	UseContext             []UsageContext                  `json:"useContext,omitempty" bson:"use_context,omitempty"`                          // The context that the content is intended to support
	Purpose                *string                         `json:"purpose,omitempty" bson:"purpose,omitempty"`                                 // Why this compartment definition is defined
	Code                   string                          `json:"code" bson:"code"`                                                           // Patient | Encounter | RelatedPerson | Practitioner | Device | EpisodeOfCare
	Search                 bool                            `json:"search" bson:"search"`                                                       // Whether the search syntax is supported
	Resource               []CompartmentDefinitionResource `json:"resource,omitempty" bson:"resource,omitempty"`                               // How a resource is related to the compartment
}

func (r *CompartmentDefinition) Validate() error {
	if r.ResourceType != "CompartmentDefinition" {
		return fmt.Errorf("invalid resourceType: expected 'CompartmentDefinition', got '%s'", r.ResourceType)
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
	var emptyString string
	if r.Url == emptyString {
		return fmt.Errorf("field 'Url' is required")
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
	if r.Code == emptyString {
		return fmt.Errorf("field 'Code' is required")
	}
	for i, item := range r.Resource {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Resource[%d]: %w", i, err)
		}
	}
	return nil
}

type CompartmentDefinitionResource struct {
	Id         *string  `json:"id,omitempty" bson:"id,omitempty"`                  // Unique id for inter-element referencing
	Code       string   `json:"code" bson:"code"`                                  // Name of resource type
	Param      []string `json:"param,omitempty" bson:"param,omitempty"`            // Search Parameter Name, or chained parameters
	StartParam *string  `json:"startParam,omitempty" bson:"start_param,omitempty"` // Search Param for interpreting $everything.start
	EndParam   *string  `json:"endParam,omitempty" bson:"end_param,omitempty"`     // Search Param for interpreting $everything.end
}

func (r *CompartmentDefinitionResource) Validate() error {
	var emptyString string
	if r.Code == emptyString {
		return fmt.Errorf("field 'Code' is required")
	}
	return nil
}
