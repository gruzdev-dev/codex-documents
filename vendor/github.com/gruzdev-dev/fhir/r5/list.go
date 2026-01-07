package models

import (
	"encoding/json"
	"fmt"
)

// A List is a curated collection of resources, for things such as problem lists, allergy lists, facility list, organization list, etc.
type List struct {
	ResourceType  string            `json:"resourceType" bson:"resource_type"`                       // Type of resource
	Id            *string           `json:"id,omitempty" bson:"id,omitempty"`                        // Logical id of this artifact
	Meta          *Meta             `json:"meta,omitempty" bson:"meta,omitempty"`                    // Metadata about the resource
	ImplicitRules *string           `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"` // A set of rules under which this content was created
	Language      *string           `json:"language,omitempty" bson:"language,omitempty"`            // Language of the resource content
	Text          *Narrative        `json:"text,omitempty" bson:"text,omitempty"`                    // Text summary of the resource, for human interpretation
	Contained     []json.RawMessage `json:"contained,omitempty" bson:"contained,omitempty"`          // Contained, inline Resources
	Identifier    []Identifier      `json:"identifier,omitempty" bson:"identifier,omitempty"`        // Business identifier
	Status        string            `json:"status" bson:"status"`                                    // current | retired | entered-in-error
	Mode          string            `json:"mode" bson:"mode"`                                        // working | snapshot | changes
	Title         *string           `json:"title,omitempty" bson:"title,omitempty"`                  // Descriptive name for the list
	Code          *CodeableConcept  `json:"code,omitempty" bson:"code,omitempty"`                    // What the purpose of this list is
	Subject       []Reference       `json:"subject,omitempty" bson:"subject,omitempty"`              // If all resources have the same subject(s)
	Encounter     *Reference        `json:"encounter,omitempty" bson:"encounter,omitempty"`          // Context in which list created
	Date          *string           `json:"date,omitempty" bson:"date,omitempty"`                    // When the list was prepared
	Source        *Reference        `json:"source,omitempty" bson:"source,omitempty"`                // Who and/or what defined the list contents (aka Author)
	OrderedBy     *CodeableConcept  `json:"orderedBy,omitempty" bson:"ordered_by,omitempty"`         // What order the list has
	Note          []Annotation      `json:"note,omitempty" bson:"note,omitempty"`                    // Comments about the list
	Entry         []ListEntry       `json:"entry,omitempty" bson:"entry,omitempty"`                  // Entries in the list
	EmptyReason   *CodeableConcept  `json:"emptyReason,omitempty" bson:"empty_reason,omitempty"`     // Why list is empty
}

func (r *List) Validate() error {
	if r.ResourceType != "List" {
		return fmt.Errorf("invalid resourceType: expected 'List', got '%s'", r.ResourceType)
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
	if r.Mode == emptyString {
		return fmt.Errorf("field 'Mode' is required")
	}
	if r.Code != nil {
		if err := r.Code.Validate(); err != nil {
			return fmt.Errorf("Code: %w", err)
		}
	}
	for i, item := range r.Subject {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Subject[%d]: %w", i, err)
		}
	}
	if r.Encounter != nil {
		if err := r.Encounter.Validate(); err != nil {
			return fmt.Errorf("Encounter: %w", err)
		}
	}
	if r.Source != nil {
		if err := r.Source.Validate(); err != nil {
			return fmt.Errorf("Source: %w", err)
		}
	}
	if r.OrderedBy != nil {
		if err := r.OrderedBy.Validate(); err != nil {
			return fmt.Errorf("OrderedBy: %w", err)
		}
	}
	for i, item := range r.Note {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Note[%d]: %w", i, err)
		}
	}
	for i, item := range r.Entry {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Entry[%d]: %w", i, err)
		}
	}
	if r.EmptyReason != nil {
		if err := r.EmptyReason.Validate(); err != nil {
			return fmt.Errorf("EmptyReason: %w", err)
		}
	}
	return nil
}

type ListEntry struct {
	Id      *string          `json:"id,omitempty" bson:"id,omitempty"`           // Unique id for inter-element referencing
	Flag    *CodeableConcept `json:"flag,omitempty" bson:"flag,omitempty"`       // Status/Workflow information about this item
	Deleted bool             `json:"deleted,omitempty" bson:"deleted,omitempty"` // If this item is actually marked as deleted
	Date    *string          `json:"date,omitempty" bson:"date,omitempty"`       // When item added to list
	Item    *Reference       `json:"item" bson:"item"`                           // Actual entry
}

func (r *ListEntry) Validate() error {
	if r.Flag != nil {
		if err := r.Flag.Validate(); err != nil {
			return fmt.Errorf("Flag: %w", err)
		}
	}
	if r.Item == nil {
		return fmt.Errorf("field 'Item' is required")
	}
	if r.Item != nil {
		if err := r.Item.Validate(); err != nil {
			return fmt.Errorf("Item: %w", err)
		}
	}
	return nil
}
