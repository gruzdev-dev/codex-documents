package models

import (
	"encoding/json"
	"fmt"
)

// A record of a healthcare consumer’s  choices  or choices made on their behalf by a third party, which permits or denies identified recipient(s) or recipient role(s) to perform one or more actions within a given policy context, for specific purposes and periods of time.
type Consent struct {
	Id               *string               `json:"id,omitempty" bson:"id,omitempty"`                              // Logical id of this artifact
	Meta             *Meta                 `json:"meta,omitempty" bson:"meta,omitempty"`                          // Metadata about the resource
	ImplicitRules    *string               `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`       // A set of rules under which this content was created
	Language         *string               `json:"language,omitempty" bson:"language,omitempty"`                  // Language of the resource content
	Text             *Narrative            `json:"text,omitempty" bson:"text,omitempty"`                          // Text summary of the resource, for human interpretation
	Contained        []json.RawMessage     `json:"contained,omitempty" bson:"contained,omitempty"`                // Contained, inline Resources
	Identifier       []Identifier          `json:"identifier,omitempty" bson:"identifier,omitempty"`              // Identifier for this record (external references)
	Status           string                `json:"status" bson:"status"`                                          // draft | active | inactive | not-done | entered-in-error | unknown
	Category         []CodeableConcept     `json:"category,omitempty" bson:"category,omitempty"`                  // Classification of the consent statement - for indexing/retrieval
	Subject          *Reference            `json:"subject,omitempty" bson:"subject,omitempty"`                    // Who the consent applies to
	Date             *string               `json:"date,omitempty" bson:"date,omitempty"`                          // Fully executed date of the consent
	Period           *Period               `json:"period,omitempty" bson:"period,omitempty"`                      // Effective period for this Consent
	Grantor          []Reference           `json:"grantor,omitempty" bson:"grantor,omitempty"`                    // Who is granting rights according to the policy and rules
	Grantee          []Reference           `json:"grantee,omitempty" bson:"grantee,omitempty"`                    // Who is agreeing to the policy and rules
	Manager          []Reference           `json:"manager,omitempty" bson:"manager,omitempty"`                    // Consent workflow management
	Controller       []Reference           `json:"controller,omitempty" bson:"controller,omitempty"`              // Consent Enforcer
	SourceAttachment []Attachment          `json:"sourceAttachment,omitempty" bson:"source_attachment,omitempty"` // Source from which this consent is taken
	SourceReference  []Reference           `json:"sourceReference,omitempty" bson:"source_reference,omitempty"`   // Source from which this consent is taken
	RegulatoryBasis  []CodeableConcept     `json:"regulatoryBasis,omitempty" bson:"regulatory_basis,omitempty"`   // Regulations establishing base Consent
	PolicyBasis      *ConsentPolicyBasis   `json:"policyBasis,omitempty" bson:"policy_basis,omitempty"`           // Computable version of the backing policy
	PolicyText       []Reference           `json:"policyText,omitempty" bson:"policy_text,omitempty"`             // Human Readable Policy
	Verification     []ConsentVerification `json:"verification,omitempty" bson:"verification,omitempty"`          // Consent Verified by patient or family
	Decision         *string               `json:"decision,omitempty" bson:"decision,omitempty"`                  // deny | permit
	Provision        []ConsentProvision    `json:"provision,omitempty" bson:"provision,omitempty"`                // Constraints to the base Consent.policyRule/Consent.policy
}

func (r *Consent) Validate() error {
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
	for i, item := range r.Category {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Category[%d]: %w", i, err)
		}
	}
	if r.Subject != nil {
		if err := r.Subject.Validate(); err != nil {
			return fmt.Errorf("Subject: %w", err)
		}
	}
	if r.Period != nil {
		if err := r.Period.Validate(); err != nil {
			return fmt.Errorf("Period: %w", err)
		}
	}
	for i, item := range r.Grantor {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Grantor[%d]: %w", i, err)
		}
	}
	for i, item := range r.Grantee {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Grantee[%d]: %w", i, err)
		}
	}
	for i, item := range r.Manager {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Manager[%d]: %w", i, err)
		}
	}
	for i, item := range r.Controller {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Controller[%d]: %w", i, err)
		}
	}
	for i, item := range r.SourceAttachment {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("SourceAttachment[%d]: %w", i, err)
		}
	}
	for i, item := range r.SourceReference {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("SourceReference[%d]: %w", i, err)
		}
	}
	for i, item := range r.RegulatoryBasis {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("RegulatoryBasis[%d]: %w", i, err)
		}
	}
	if r.PolicyBasis != nil {
		if err := r.PolicyBasis.Validate(); err != nil {
			return fmt.Errorf("PolicyBasis: %w", err)
		}
	}
	for i, item := range r.PolicyText {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("PolicyText[%d]: %w", i, err)
		}
	}
	for i, item := range r.Verification {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Verification[%d]: %w", i, err)
		}
	}
	for i, item := range r.Provision {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Provision[%d]: %w", i, err)
		}
	}
	return nil
}

type ConsentProvisionData struct {
	Id        *string    `json:"id,omitempty" bson:"id,omitempty"` // Unique id for inter-element referencing
	Meaning   string     `json:"meaning" bson:"meaning"`           // instance | related | dependents | authoredby
	Reference *Reference `json:"reference" bson:"reference"`       // The actual data reference
}

func (r *ConsentProvisionData) Validate() error {
	var emptyString string
	if r.Meaning == emptyString {
		return fmt.Errorf("field 'Meaning' is required")
	}
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

type ConsentPolicyBasis struct {
	Id        *string    `json:"id,omitempty" bson:"id,omitempty"`               // Unique id for inter-element referencing
	Reference *Reference `json:"reference,omitempty" bson:"reference,omitempty"` // Reference backing policy resource
	Uri       *string    `json:"uri,omitempty" bson:"uri,omitempty"`             // URI to a computable backing policy
}

func (r *ConsentPolicyBasis) Validate() error {
	if r.Reference != nil {
		if err := r.Reference.Validate(); err != nil {
			return fmt.Errorf("Reference: %w", err)
		}
	}
	return nil
}

type ConsentVerification struct {
	Id           *string          `json:"id,omitempty" bson:"id,omitempty"`                      // Unique id for inter-element referencing
	Verified     bool             `json:"verified" bson:"verified"`                              // Has been verified
	Type         *CodeableConcept `json:"type,omitempty" bson:"type,omitempty"`                  // Business case of verification
	VerifiedBy   *Reference       `json:"verifiedBy,omitempty" bson:"verified_by,omitempty"`     // Person conducting verification
	VerifiedWith *Reference       `json:"verifiedWith,omitempty" bson:"verified_with,omitempty"` // Person who verified
	Date         []string         `json:"date,omitempty" bson:"date,omitempty"`                  // When consent verified
}

func (r *ConsentVerification) Validate() error {
	if r.Type != nil {
		if err := r.Type.Validate(); err != nil {
			return fmt.Errorf("Type: %w", err)
		}
	}
	if r.VerifiedBy != nil {
		if err := r.VerifiedBy.Validate(); err != nil {
			return fmt.Errorf("VerifiedBy: %w", err)
		}
	}
	if r.VerifiedWith != nil {
		if err := r.VerifiedWith.Validate(); err != nil {
			return fmt.Errorf("VerifiedWith: %w", err)
		}
	}
	return nil
}

type ConsentProvision struct {
	Id            *string                 `json:"id,omitempty" bson:"id,omitempty"`                        // Unique id for inter-element referencing
	Period        *Period                 `json:"period,omitempty" bson:"period,omitempty"`                // Timeframe for this provision
	Actor         []ConsentProvisionActor `json:"actor,omitempty" bson:"actor,omitempty"`                  // Who|what controlled by this provision (or group, by role)
	Action        []CodeableConcept       `json:"action,omitempty" bson:"action,omitempty"`                // Actions controlled by this provision
	SecurityLabel []Coding                `json:"securityLabel,omitempty" bson:"security_label,omitempty"` // Security Labels that define affected resources
	Purpose       []Coding                `json:"purpose,omitempty" bson:"purpose,omitempty"`              // Context of activities covered by this provision
	DocumentType  []Coding                `json:"documentType,omitempty" bson:"document_type,omitempty"`   // e.g. Resource Type, Profile, CDA, etc
	ResourceType  []Coding                `json:"resourceType,omitempty" bson:"resource_type,omitempty"`   // e.g. Resource Type, Profile, etc
	Code          []CodeableConcept       `json:"code,omitempty" bson:"code,omitempty"`                    // e.g. LOINC or SNOMED CT code, etc. in the content
	DataPeriod    *Period                 `json:"dataPeriod,omitempty" bson:"data_period,omitempty"`       // Timeframe for data controlled by this provision
	Data          []ConsentProvisionData  `json:"data,omitempty" bson:"data,omitempty"`                    // Data controlled by this provision
	Expression    *Expression             `json:"expression,omitempty" bson:"expression,omitempty"`        // A computable expression of the consent
	Provision     []ConsentProvision      `json:"provision,omitempty" bson:"provision,omitempty"`          // Nested Exception Provisions
}

func (r *ConsentProvision) Validate() error {
	if r.Period != nil {
		if err := r.Period.Validate(); err != nil {
			return fmt.Errorf("Period: %w", err)
		}
	}
	for i, item := range r.Actor {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Actor[%d]: %w", i, err)
		}
	}
	for i, item := range r.Action {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Action[%d]: %w", i, err)
		}
	}
	for i, item := range r.SecurityLabel {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("SecurityLabel[%d]: %w", i, err)
		}
	}
	for i, item := range r.Purpose {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Purpose[%d]: %w", i, err)
		}
	}
	for i, item := range r.DocumentType {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("DocumentType[%d]: %w", i, err)
		}
	}
	for i, item := range r.ResourceType {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("ResourceType[%d]: %w", i, err)
		}
	}
	for i, item := range r.Code {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Code[%d]: %w", i, err)
		}
	}
	if r.DataPeriod != nil {
		if err := r.DataPeriod.Validate(); err != nil {
			return fmt.Errorf("DataPeriod: %w", err)
		}
	}
	for i, item := range r.Data {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Data[%d]: %w", i, err)
		}
	}
	if r.Expression != nil {
		if err := r.Expression.Validate(); err != nil {
			return fmt.Errorf("Expression: %w", err)
		}
	}
	for i, item := range r.Provision {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Provision[%d]: %w", i, err)
		}
	}
	return nil
}

type ConsentProvisionActor struct {
	Id        *string          `json:"id,omitempty" bson:"id,omitempty"`               // Unique id for inter-element referencing
	Role      *CodeableConcept `json:"role,omitempty" bson:"role,omitempty"`           // How the actor is involved
	Reference *Reference       `json:"reference,omitempty" bson:"reference,omitempty"` // Resource for the actor (or group, by role)
}

func (r *ConsentProvisionActor) Validate() error {
	if r.Role != nil {
		if err := r.Role.Validate(); err != nil {
			return fmt.Errorf("Role: %w", err)
		}
	}
	if r.Reference != nil {
		if err := r.Reference.Validate(); err != nil {
			return fmt.Errorf("Reference: %w", err)
		}
	}
	return nil
}
