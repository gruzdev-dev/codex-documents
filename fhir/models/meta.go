package fhir

// Meta Type: The metadata about a resource. This is content in the resource that is maintained by the infrastructure. Changes to the content might not always be associated with version changes to the resource.
type Meta struct {
	Id          *string  `json:"id,omitempty"`          // Unique id for inter-element referencing
	VersionId   *string  `json:"versionId,omitempty"`   // Version specific identifier
	LastUpdated *string  `json:"lastUpdated,omitempty"` // When the resource version last changed
	Source      *string  `json:"source,omitempty"`      // Identifies where the resource comes from
	Profile     []string `json:"profile,omitempty"`     // Profiles this resource claims to conform to
	Security    []Coding `json:"security,omitempty"`    // Security Labels applied to this resource
	Tag         []Coding `json:"tag,omitempty"`         // Tags applied to this resource
}
