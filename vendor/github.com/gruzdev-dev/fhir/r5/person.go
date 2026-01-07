package models

import (
	"encoding/json"
	"fmt"
)

// Demographics and administrative information about a person independent of a specific health-related context.
type Person struct {
	ResourceType         string                `json:"resourceType" bson:"resource_type"`                                     // Type of resource
	Id                   *string               `json:"id,omitempty" bson:"id,omitempty"`                                      // Logical id of this artifact
	Meta                 *Meta                 `json:"meta,omitempty" bson:"meta,omitempty"`                                  // Metadata about the resource
	ImplicitRules        *string               `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`               // A set of rules under which this content was created
	Language             *string               `json:"language,omitempty" bson:"language,omitempty"`                          // Language of the resource content
	Text                 *Narrative            `json:"text,omitempty" bson:"text,omitempty"`                                  // Text summary of the resource, for human interpretation
	Contained            []json.RawMessage     `json:"contained,omitempty" bson:"contained,omitempty"`                        // Contained, inline Resources
	Identifier           []Identifier          `json:"identifier,omitempty" bson:"identifier,omitempty"`                      // A human identifier for this person
	Active               bool                  `json:"active,omitempty" bson:"active,omitempty"`                              // This person's record is in active use
	Name                 []HumanName           `json:"name,omitempty" bson:"name,omitempty"`                                  // A name associated with the person
	Telecom              []ContactPoint        `json:"telecom,omitempty" bson:"telecom,omitempty"`                            // A contact detail for the person
	Gender               *string               `json:"gender,omitempty" bson:"gender,omitempty"`                              // male | female | other | unknown
	BirthDate            *string               `json:"birthDate,omitempty" bson:"birth_date,omitempty"`                       // The date on which the person was born
	DeceasedBoolean      *bool                 `json:"deceasedBoolean,omitempty" bson:"deceased_boolean,omitempty"`           // Indicates if the individual is deceased or not
	DeceasedDateTime     *string               `json:"deceasedDateTime,omitempty" bson:"deceased_date_time,omitempty"`        // Indicates if the individual is deceased or not
	Address              []Address             `json:"address,omitempty" bson:"address,omitempty"`                            // One or more addresses for the person
	MaritalStatus        *CodeableConcept      `json:"maritalStatus,omitempty" bson:"marital_status,omitempty"`               // Marital (civil) status of a person
	Photo                []Attachment          `json:"photo,omitempty" bson:"photo,omitempty"`                                // Image of the person
	Communication        []PersonCommunication `json:"communication,omitempty" bson:"communication,omitempty"`                // A language which may be used to communicate with the person about his or her health
	ManagingOrganization *Reference            `json:"managingOrganization,omitempty" bson:"managing_organization,omitempty"` // The organization that is the custodian of the person record
	Link                 []PersonLink          `json:"link,omitempty" bson:"link,omitempty"`                                  // Link to a resource that concerns the same actual person
}

func (r *Person) Validate() error {
	if r.ResourceType != "Person" {
		return fmt.Errorf("invalid resourceType: expected 'Person', got '%s'", r.ResourceType)
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
	for i, item := range r.Name {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Name[%d]: %w", i, err)
		}
	}
	for i, item := range r.Telecom {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Telecom[%d]: %w", i, err)
		}
	}
	for i, item := range r.Address {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Address[%d]: %w", i, err)
		}
	}
	if r.MaritalStatus != nil {
		if err := r.MaritalStatus.Validate(); err != nil {
			return fmt.Errorf("MaritalStatus: %w", err)
		}
	}
	for i, item := range r.Photo {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Photo[%d]: %w", i, err)
		}
	}
	for i, item := range r.Communication {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Communication[%d]: %w", i, err)
		}
	}
	if r.ManagingOrganization != nil {
		if err := r.ManagingOrganization.Validate(); err != nil {
			return fmt.Errorf("ManagingOrganization: %w", err)
		}
	}
	for i, item := range r.Link {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Link[%d]: %w", i, err)
		}
	}
	return nil
}

type PersonCommunication struct {
	Id        *string          `json:"id,omitempty" bson:"id,omitempty"`               // Unique id for inter-element referencing
	Language  *CodeableConcept `json:"language" bson:"language"`                       // The language which can be used to communicate with the person about his or her health
	Preferred bool             `json:"preferred,omitempty" bson:"preferred,omitempty"` // Language preference indicator
}

func (r *PersonCommunication) Validate() error {
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

type PersonLink struct {
	Id        *string    `json:"id,omitempty" bson:"id,omitempty"`               // Unique id for inter-element referencing
	Target    *Reference `json:"target" bson:"target"`                           // The resource to which this actual person is associated
	Assurance *string    `json:"assurance,omitempty" bson:"assurance,omitempty"` // level1 | level2 | level3 | level4
}

func (r *PersonLink) Validate() error {
	if r.Target == nil {
		return fmt.Errorf("field 'Target' is required")
	}
	if r.Target != nil {
		if err := r.Target.Validate(); err != nil {
			return fmt.Errorf("Target: %w", err)
		}
	}
	return nil
}
