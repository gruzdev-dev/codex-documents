package models

import (
	"encoding/json"
	"fmt"
)

// A container for a collection of resources.
type Bundle struct {
	ResourceType  string          `json:"resourceType" bson:"resource_type"`                       // Type of resource
	Id            *string         `json:"id,omitempty" bson:"id,omitempty"`                        // Logical id of this artifact
	Meta          *Meta           `json:"meta,omitempty" bson:"meta,omitempty"`                    // Metadata about the resource
	ImplicitRules *string         `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"` // A set of rules under which this content was created
	Language      *string         `json:"language,omitempty" bson:"language,omitempty"`            // Language of the resource content
	Identifier    *Identifier     `json:"identifier,omitempty" bson:"identifier,omitempty"`        // Persistent identifier for the bundle
	Type          string          `json:"type" bson:"type"`                                        // document | message | transaction | transaction-response | batch | batch-response | history | searchset | collection | subscription-notification
	Timestamp     *string         `json:"timestamp,omitempty" bson:"timestamp,omitempty"`          // When the bundle was assembled
	Total         *int            `json:"total,omitempty" bson:"total,omitempty"`                  // Total matches across all pages
	Link          []BundleLink    `json:"link,omitempty" bson:"link,omitempty"`                    // Links related to this Bundle
	Entry         []BundleEntry   `json:"entry,omitempty" bson:"entry,omitempty"`                  // Entry in the bundle - will have a resource or information
	Signature     *Signature      `json:"signature,omitempty" bson:"signature,omitempty"`          // Digital Signature (deprecated: use Provenance Signatures)
	Issues        json.RawMessage `json:"issues,omitempty" bson:"issues,omitempty"`                // OperationOutcome with issues about the Bundle
}

func (r *Bundle) Validate() error {
	if r.ResourceType != "Bundle" {
		return fmt.Errorf("invalid resourceType: expected 'Bundle', got '%s'", r.ResourceType)
	}
	if r.Meta != nil {
		if err := r.Meta.Validate(); err != nil {
			return fmt.Errorf("Meta: %w", err)
		}
	}
	if r.Identifier != nil {
		if err := r.Identifier.Validate(); err != nil {
			return fmt.Errorf("Identifier: %w", err)
		}
	}
	var emptyString string
	if r.Type == emptyString {
		return fmt.Errorf("field 'Type' is required")
	}
	for i, item := range r.Link {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Link[%d]: %w", i, err)
		}
	}
	for i, item := range r.Entry {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Entry[%d]: %w", i, err)
		}
	}
	if r.Signature != nil {
		if err := r.Signature.Validate(); err != nil {
			return fmt.Errorf("Signature: %w", err)
		}
	}
	return nil
}

type BundleLink struct {
	Id       *string `json:"id,omitempty" bson:"id,omitempty"` // Unique id for inter-element referencing
	Relation string  `json:"relation" bson:"relation"`         // See http://www.iana.org/assignments/link-relations/link-relations.xhtml#link-relations-1
	Url      string  `json:"url" bson:"url"`                   // Reference details for the link
}

func (r *BundleLink) Validate() error {
	var emptyString string
	if r.Relation == emptyString {
		return fmt.Errorf("field 'Relation' is required")
	}
	if r.Url == emptyString {
		return fmt.Errorf("field 'Url' is required")
	}
	return nil
}

type BundleEntry struct {
	Id       *string              `json:"id,omitempty" bson:"id,omitempty"`             // Unique id for inter-element referencing
	Link     []BundleLink         `json:"link,omitempty" bson:"link,omitempty"`         // Links related to this entry
	FullUrl  *string              `json:"fullUrl,omitempty" bson:"full_url,omitempty"`  // URI for resource (e.g. the absolute URL server address, URI for UUID/OID, etc.)
	Resource json.RawMessage      `json:"resource,omitempty" bson:"resource,omitempty"` // A resource in the bundle
	Search   *BundleEntrySearch   `json:"search,omitempty" bson:"search,omitempty"`     // Search related information
	Request  *BundleEntryRequest  `json:"request,omitempty" bson:"request,omitempty"`   // Additional execution information (transaction/batch/history)
	Response *BundleEntryResponse `json:"response,omitempty" bson:"response,omitempty"` // Results of execution (transaction/batch/history)
}

func (r *BundleEntry) Validate() error {
	for i, item := range r.Link {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Link[%d]: %w", i, err)
		}
	}
	if r.Search != nil {
		if err := r.Search.Validate(); err != nil {
			return fmt.Errorf("Search: %w", err)
		}
	}
	if r.Request != nil {
		if err := r.Request.Validate(); err != nil {
			return fmt.Errorf("Request: %w", err)
		}
	}
	if r.Response != nil {
		if err := r.Response.Validate(); err != nil {
			return fmt.Errorf("Response: %w", err)
		}
	}
	return nil
}

type BundleEntrySearch struct {
	Id    *string  `json:"id,omitempty" bson:"id,omitempty"`       // Unique id for inter-element referencing
	Mode  *string  `json:"mode,omitempty" bson:"mode,omitempty"`   // match | include - why this is in the result set
	Score *float64 `json:"score,omitempty" bson:"score,omitempty"` // Search ranking (between 0 and 1)
}

func (r *BundleEntrySearch) Validate() error {
	return nil
}

type BundleEntryRequest struct {
	Id              *string `json:"id,omitempty" bson:"id,omitempty"`                             // Unique id for inter-element referencing
	Method          string  `json:"method" bson:"method"`                                         // GET | HEAD | POST | PUT | DELETE | PATCH
	Url             string  `json:"url" bson:"url"`                                               // URL for HTTP equivalent of this entry
	IfNoneMatch     *string `json:"ifNoneMatch,omitempty" bson:"if_none_match,omitempty"`         // For managing cache validation
	IfModifiedSince *string `json:"ifModifiedSince,omitempty" bson:"if_modified_since,omitempty"` // For managing cache currency
	IfMatch         *string `json:"ifMatch,omitempty" bson:"if_match,omitempty"`                  // For managing update contention
	IfNoneExist     *string `json:"ifNoneExist,omitempty" bson:"if_none_exist,omitempty"`         // For conditional creates
}

func (r *BundleEntryRequest) Validate() error {
	var emptyString string
	if r.Method == emptyString {
		return fmt.Errorf("field 'Method' is required")
	}
	if r.Url == emptyString {
		return fmt.Errorf("field 'Url' is required")
	}
	return nil
}

type BundleEntryResponse struct {
	Id           *string         `json:"id,omitempty" bson:"id,omitempty"`                      // Unique id for inter-element referencing
	Status       string          `json:"status" bson:"status"`                                  // Status response code (text optional)
	Location     *string         `json:"location,omitempty" bson:"location,omitempty"`          // The location (if the operation returns a location)
	Etag         *string         `json:"etag,omitempty" bson:"etag,omitempty"`                  // The Etag for the resource (if relevant)
	LastModified *string         `json:"lastModified,omitempty" bson:"last_modified,omitempty"` // Server's date time modified
	Outcome      json.RawMessage `json:"outcome,omitempty" bson:"outcome,omitempty"`            // OperationOutcome with hints and warnings (for batch/transaction)
}

func (r *BundleEntryResponse) Validate() error {
	var emptyString string
	if r.Status == emptyString {
		return fmt.Errorf("field 'Status' is required")
	}
	return nil
}
