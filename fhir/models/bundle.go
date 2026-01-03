package fhir

import "encoding/json"

// A container for a collection of resources.
type Bundle struct {
	Id            *string         `json:"id,omitempty"`            // Logical id of this artifact
	Meta          *Meta           `json:"meta,omitempty"`          // Metadata about the resource
	ImplicitRules *string         `json:"implicitRules,omitempty"` // A set of rules under which this content was created
	Language      *string         `json:"language,omitempty"`      // Language of the resource content
	Identifier    *Identifier     `json:"identifier,omitempty"`    // Persistent identifier for the bundle
	Type          string          `json:"type,omitempty"`          // document | message | transaction | transaction-response | batch | batch-response | history | searchset | collection | subscription-notification
	Timestamp     *string         `json:"timestamp,omitempty"`     // When the bundle was assembled
	Total         *int            `json:"total,omitempty"`         // Total matches across all pages
	Link          []BundleLink    `json:"link,omitempty"`          // Links related to this Bundle
	Entry         []BundleEntry   `json:"entry,omitempty"`         // Entry in the bundle - will have a resource or information
	Signature     *Signature      `json:"signature,omitempty"`     // Digital Signature (deprecated: use Provenance Signatures)
	Issues        json.RawMessage `json:"issues,omitempty"`        // OperationOutcome with issues about the Bundle
}

type BundleLink struct {
	Id       *string `json:"id,omitempty"`       // Unique id for inter-element referencing
	Relation string  `json:"relation,omitempty"` // See http://www.iana.org/assignments/link-relations/link-relations.xhtml#link-relations-1
	Url      string  `json:"url,omitempty"`      // Reference details for the link
}

type BundleEntry struct {
	Id       *string              `json:"id,omitempty"`       // Unique id for inter-element referencing
	Link     []BundleEntryLink    `json:"link,omitempty"`     // Links related to this entry
	FullUrl  *string              `json:"fullUrl,omitempty"`  // URI for resource (e.g. the absolute URL server address, URI for UUID/OID, etc.)
	Resource json.RawMessage      `json:"resource,omitempty"` // A resource in the bundle
	Search   *BundleEntrySearch   `json:"search,omitempty"`   // Search related information
	Request  *BundleEntryRequest  `json:"request,omitempty"`  // Additional execution information (transaction/batch/history)
	Response *BundleEntryResponse `json:"response,omitempty"` // Results of execution (transaction/batch/history)
}

type BundleEntrySearch struct {
	Id    *string  `json:"id,omitempty"`    // Unique id for inter-element referencing
	Mode  *string  `json:"mode,omitempty"`  // match | include - why this is in the result set
	Score *float64 `json:"score,omitempty"` // Search ranking (between 0 and 1)
}

type BundleEntryRequest struct {
	Id              *string `json:"id,omitempty"`              // Unique id for inter-element referencing
	Method          string  `json:"method,omitempty"`          // GET | HEAD | POST | PUT | DELETE | PATCH
	Url             string  `json:"url,omitempty"`             // URL for HTTP equivalent of this entry
	IfNoneMatch     *string `json:"ifNoneMatch,omitempty"`     // For managing cache validation
	IfModifiedSince *string `json:"ifModifiedSince,omitempty"` // For managing cache currency
	IfMatch         *string `json:"ifMatch,omitempty"`         // For managing update contention
	IfNoneExist     *string `json:"ifNoneExist,omitempty"`     // For conditional creates
}

type BundleEntryResponse struct {
	Id           *string         `json:"id,omitempty"`           // Unique id for inter-element referencing
	Status       string          `json:"status,omitempty"`       // Status response code (text optional)
	Location     *string         `json:"location,omitempty"`     // The location (if the operation returns a location)
	Etag         *string         `json:"etag,omitempty"`         // The Etag for the resource (if relevant)
	LastModified *string         `json:"lastModified,omitempty"` // Server's date time modified
	Outcome      json.RawMessage `json:"outcome,omitempty"`      // OperationOutcome with hints and warnings (for batch/transaction)
}

type BundleEntryLink struct {
}
