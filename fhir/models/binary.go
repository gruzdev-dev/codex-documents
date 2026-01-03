package models

// A resource that represents the data of a single raw artifact as digital content accessible in its native format.  A Binary resource can contain any content, whether text, image, pdf, zip archive, etc.
type Binary struct {
	Id              *string    `json:"id,omitempty"`              // Logical id of this artifact
	Meta            *Meta      `json:"meta,omitempty"`            // Metadata about the resource
	ImplicitRules   *string    `json:"implicitRules,omitempty"`   // A set of rules under which this content was created
	Language        *string    `json:"language,omitempty"`        // Language of the resource content
	ContentType     string     `json:"contentType,omitempty"`     // MimeType of the binary content
	SecurityContext *Reference `json:"securityContext,omitempty"` // Identifies another resource to use as proxy when enforcing access control
	Data            *string    `json:"data,omitempty"`            // The actual content
}
