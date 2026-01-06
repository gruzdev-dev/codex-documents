package models

import (
	"encoding/json"
	"fmt"
)

// Legally enforceable, formally recorded unilateral or bilateral directive i.e., a policy or agreement.
type Contract struct {
	Id                       *string                    `json:"id,omitempty" bson:"id,omitempty"`                                               // Logical id of this artifact
	Meta                     *Meta                      `json:"meta,omitempty" bson:"meta,omitempty"`                                           // Metadata about the resource
	ImplicitRules            *string                    `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`                        // A set of rules under which this content was created
	Language                 *string                    `json:"language,omitempty" bson:"language,omitempty"`                                   // Language of the resource content
	Text                     *Narrative                 `json:"text,omitempty" bson:"text,omitempty"`                                           // Text summary of the resource, for human interpretation
	Contained                []json.RawMessage          `json:"contained,omitempty" bson:"contained,omitempty"`                                 // Contained, inline Resources
	Identifier               []Identifier               `json:"identifier,omitempty" bson:"identifier,omitempty"`                               // Contract number
	Url                      *string                    `json:"url,omitempty" bson:"url,omitempty"`                                             // Basal definition
	Version                  *string                    `json:"version,omitempty" bson:"version,omitempty"`                                     // Business edition
	Status                   *string                    `json:"status,omitempty" bson:"status,omitempty"`                                       // amended | appended | cancelled | disputed | entered-in-error | executable +
	LegalState               *CodeableConcept           `json:"legalState,omitempty" bson:"legal_state,omitempty"`                              // Negotiation status
	InstantiatesCanonical    *Reference                 `json:"instantiatesCanonical,omitempty" bson:"instantiates_canonical,omitempty"`        // Source Contract Definition
	InstantiatesUri          *string                    `json:"instantiatesUri,omitempty" bson:"instantiates_uri,omitempty"`                    // External Contract Definition
	ContentDerivative        *CodeableConcept           `json:"contentDerivative,omitempty" bson:"content_derivative,omitempty"`                // Content derived from the basal information
	Issued                   *string                    `json:"issued,omitempty" bson:"issued,omitempty"`                                       // When this Contract was issued
	Applies                  *Period                    `json:"applies,omitempty" bson:"applies,omitempty"`                                     // Effective time
	ExpirationType           *CodeableConcept           `json:"expirationType,omitempty" bson:"expiration_type,omitempty"`                      // Contract cessation cause
	Subject                  []Reference                `json:"subject,omitempty" bson:"subject,omitempty"`                                     // Contract Target Entity
	Authority                []Reference                `json:"authority,omitempty" bson:"authority,omitempty"`                                 // Authority under which this Contract has standing
	Domain                   []Reference                `json:"domain,omitempty" bson:"domain,omitempty"`                                       // A sphere of control governed by an authoritative jurisdiction, organization, or person
	Site                     []Reference                `json:"site,omitempty" bson:"site,omitempty"`                                           // Specific Location
	Name                     *string                    `json:"name,omitempty" bson:"name,omitempty"`                                           // Computer friendly designation
	Title                    *string                    `json:"title,omitempty" bson:"title,omitempty"`                                         // Human Friendly name
	Subtitle                 *string                    `json:"subtitle,omitempty" bson:"subtitle,omitempty"`                                   // Subordinate Friendly name
	Alias                    []string                   `json:"alias,omitempty" bson:"alias,omitempty"`                                         // Acronym or short name
	Author                   *Reference                 `json:"author,omitempty" bson:"author,omitempty"`                                       // Source of Contract
	Scope                    *CodeableConcept           `json:"scope,omitempty" bson:"scope,omitempty"`                                         // Range of Legal Concerns
	TopicCodeableConcept     *CodeableConcept           `json:"topicCodeableConcept,omitempty" bson:"topic_codeable_concept,omitempty"`         // Focus of contract interest
	TopicReference           *Reference                 `json:"topicReference,omitempty" bson:"topic_reference,omitempty"`                      // Focus of contract interest
	Type                     *CodeableConcept           `json:"type,omitempty" bson:"type,omitempty"`                                           // Legal instrument category
	SubType                  []CodeableConcept          `json:"subType,omitempty" bson:"sub_type,omitempty"`                                    // Subtype within the context of type
	ContentDefinition        *ContractContentDefinition `json:"contentDefinition,omitempty" bson:"content_definition,omitempty"`                // Contract precursor content
	Term                     []ContractTerm             `json:"term,omitempty" bson:"term,omitempty"`                                           // Contract Term List
	SupportingInfo           []Reference                `json:"supportingInfo,omitempty" bson:"supporting_info,omitempty"`                      // Extra Information
	RelevantHistory          []Reference                `json:"relevantHistory,omitempty" bson:"relevant_history,omitempty"`                    // Key event in Contract History
	Signer                   []ContractSigner           `json:"signer,omitempty" bson:"signer,omitempty"`                                       // Contract Signatory
	Friendly                 []ContractFriendly         `json:"friendly,omitempty" bson:"friendly,omitempty"`                                   // Contract Friendly Language
	Legal                    []ContractLegal            `json:"legal,omitempty" bson:"legal,omitempty"`                                         // Contract Legal Language
	Rule                     []ContractRule             `json:"rule,omitempty" bson:"rule,omitempty"`                                           // Computable Contract Language
	LegallyBindingAttachment *Attachment                `json:"legallyBindingAttachment,omitempty" bson:"legally_binding_attachment,omitempty"` // Binding Contract
	LegallyBindingReference  *Reference                 `json:"legallyBindingReference,omitempty" bson:"legally_binding_reference,omitempty"`   // Binding Contract
}

func (r *Contract) Validate() error {
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
	if r.LegalState != nil {
		if err := r.LegalState.Validate(); err != nil {
			return fmt.Errorf("LegalState: %w", err)
		}
	}
	if r.InstantiatesCanonical != nil {
		if err := r.InstantiatesCanonical.Validate(); err != nil {
			return fmt.Errorf("InstantiatesCanonical: %w", err)
		}
	}
	if r.ContentDerivative != nil {
		if err := r.ContentDerivative.Validate(); err != nil {
			return fmt.Errorf("ContentDerivative: %w", err)
		}
	}
	if r.Applies != nil {
		if err := r.Applies.Validate(); err != nil {
			return fmt.Errorf("Applies: %w", err)
		}
	}
	if r.ExpirationType != nil {
		if err := r.ExpirationType.Validate(); err != nil {
			return fmt.Errorf("ExpirationType: %w", err)
		}
	}
	for i, item := range r.Subject {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Subject[%d]: %w", i, err)
		}
	}
	for i, item := range r.Authority {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Authority[%d]: %w", i, err)
		}
	}
	for i, item := range r.Domain {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Domain[%d]: %w", i, err)
		}
	}
	for i, item := range r.Site {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Site[%d]: %w", i, err)
		}
	}
	if r.Author != nil {
		if err := r.Author.Validate(); err != nil {
			return fmt.Errorf("Author: %w", err)
		}
	}
	if r.Scope != nil {
		if err := r.Scope.Validate(); err != nil {
			return fmt.Errorf("Scope: %w", err)
		}
	}
	if r.TopicCodeableConcept != nil {
		if err := r.TopicCodeableConcept.Validate(); err != nil {
			return fmt.Errorf("TopicCodeableConcept: %w", err)
		}
	}
	if r.TopicReference != nil {
		if err := r.TopicReference.Validate(); err != nil {
			return fmt.Errorf("TopicReference: %w", err)
		}
	}
	if r.Type != nil {
		if err := r.Type.Validate(); err != nil {
			return fmt.Errorf("Type: %w", err)
		}
	}
	for i, item := range r.SubType {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("SubType[%d]: %w", i, err)
		}
	}
	if r.ContentDefinition != nil {
		if err := r.ContentDefinition.Validate(); err != nil {
			return fmt.Errorf("ContentDefinition: %w", err)
		}
	}
	for i, item := range r.Term {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Term[%d]: %w", i, err)
		}
	}
	for i, item := range r.SupportingInfo {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("SupportingInfo[%d]: %w", i, err)
		}
	}
	for i, item := range r.RelevantHistory {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("RelevantHistory[%d]: %w", i, err)
		}
	}
	for i, item := range r.Signer {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Signer[%d]: %w", i, err)
		}
	}
	for i, item := range r.Friendly {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Friendly[%d]: %w", i, err)
		}
	}
	for i, item := range r.Legal {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Legal[%d]: %w", i, err)
		}
	}
	for i, item := range r.Rule {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Rule[%d]: %w", i, err)
		}
	}
	if r.LegallyBindingAttachment != nil {
		if err := r.LegallyBindingAttachment.Validate(); err != nil {
			return fmt.Errorf("LegallyBindingAttachment: %w", err)
		}
	}
	if r.LegallyBindingReference != nil {
		if err := r.LegallyBindingReference.Validate(); err != nil {
			return fmt.Errorf("LegallyBindingReference: %w", err)
		}
	}
	return nil
}

type ContractTerm struct {
	Id                   *string                     `json:"id,omitempty" bson:"id,omitempty"`                                       // Unique id for inter-element referencing
	Identifier           *Identifier                 `json:"identifier,omitempty" bson:"identifier,omitempty"`                       // Contract Term Number
	Issued               *string                     `json:"issued,omitempty" bson:"issued,omitempty"`                               // Contract Term Issue Date Time
	Applies              *Period                     `json:"applies,omitempty" bson:"applies,omitempty"`                             // Contract Term Effective Time
	TopicCodeableConcept *CodeableConcept            `json:"topicCodeableConcept,omitempty" bson:"topic_codeable_concept,omitempty"` // Term Concern
	TopicReference       *Reference                  `json:"topicReference,omitempty" bson:"topic_reference,omitempty"`              // Term Concern
	Type                 *CodeableConcept            `json:"type,omitempty" bson:"type,omitempty"`                                   // Contract Term Type or Form
	SubType              *CodeableConcept            `json:"subType,omitempty" bson:"sub_type,omitempty"`                            // Contract Term Type specific classification
	Text                 *string                     `json:"text,omitempty" bson:"text,omitempty"`                                   // Term Statement
	SecurityLabel        []ContractTermSecurityLabel `json:"securityLabel,omitempty" bson:"security_label,omitempty"`                // Protection for the Term
	Offer                *ContractTermOffer          `json:"offer" bson:"offer"`                                                     // Context of the Contract term
	Asset                []ContractTermAsset         `json:"asset,omitempty" bson:"asset,omitempty"`                                 // Contract Term Asset List
	Action               []ContractTermAction        `json:"action,omitempty" bson:"action,omitempty"`                               // Entity being ascribed responsibility
	Group                []ContractTerm              `json:"group,omitempty" bson:"group,omitempty"`                                 // Nested Contract Term Group
}

func (r *ContractTerm) Validate() error {
	if r.Identifier != nil {
		if err := r.Identifier.Validate(); err != nil {
			return fmt.Errorf("Identifier: %w", err)
		}
	}
	if r.Applies != nil {
		if err := r.Applies.Validate(); err != nil {
			return fmt.Errorf("Applies: %w", err)
		}
	}
	if r.TopicCodeableConcept != nil {
		if err := r.TopicCodeableConcept.Validate(); err != nil {
			return fmt.Errorf("TopicCodeableConcept: %w", err)
		}
	}
	if r.TopicReference != nil {
		if err := r.TopicReference.Validate(); err != nil {
			return fmt.Errorf("TopicReference: %w", err)
		}
	}
	if r.Type != nil {
		if err := r.Type.Validate(); err != nil {
			return fmt.Errorf("Type: %w", err)
		}
	}
	if r.SubType != nil {
		if err := r.SubType.Validate(); err != nil {
			return fmt.Errorf("SubType: %w", err)
		}
	}
	for i, item := range r.SecurityLabel {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("SecurityLabel[%d]: %w", i, err)
		}
	}
	if r.Offer == nil {
		return fmt.Errorf("field 'Offer' is required")
	}
	if r.Offer != nil {
		if err := r.Offer.Validate(); err != nil {
			return fmt.Errorf("Offer: %w", err)
		}
	}
	for i, item := range r.Asset {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Asset[%d]: %w", i, err)
		}
	}
	for i, item := range r.Action {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Action[%d]: %w", i, err)
		}
	}
	for i, item := range r.Group {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Group[%d]: %w", i, err)
		}
	}
	return nil
}

type ContractTermSecurityLabel struct {
	Id             *string  `json:"id,omitempty" bson:"id,omitempty"`             // Unique id for inter-element referencing
	Number         []int    `json:"number,omitempty" bson:"number,omitempty"`     // Link to Security Labels
	Classification *Coding  `json:"classification" bson:"classification"`         // Confidentiality Protection
	Category       []Coding `json:"category,omitempty" bson:"category,omitempty"` // Applicable Policy
	Control        []Coding `json:"control,omitempty" bson:"control,omitempty"`   // Handling Instructions
}

func (r *ContractTermSecurityLabel) Validate() error {
	if r.Classification == nil {
		return fmt.Errorf("field 'Classification' is required")
	}
	if r.Classification != nil {
		if err := r.Classification.Validate(); err != nil {
			return fmt.Errorf("Classification: %w", err)
		}
	}
	for i, item := range r.Category {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Category[%d]: %w", i, err)
		}
	}
	for i, item := range r.Control {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Control[%d]: %w", i, err)
		}
	}
	return nil
}

type ContractTermAssetContext struct {
	Id        *string           `json:"id,omitempty" bson:"id,omitempty"`               // Unique id for inter-element referencing
	Reference *Reference        `json:"reference,omitempty" bson:"reference,omitempty"` // Creator,custodian or owner
	Code      []CodeableConcept `json:"code,omitempty" bson:"code,omitempty"`           // Codeable asset context
	Text      *string           `json:"text,omitempty" bson:"text,omitempty"`           // Context description
}

func (r *ContractTermAssetContext) Validate() error {
	if r.Reference != nil {
		if err := r.Reference.Validate(); err != nil {
			return fmt.Errorf("Reference: %w", err)
		}
	}
	for i, item := range r.Code {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Code[%d]: %w", i, err)
		}
	}
	return nil
}

type ContractFriendly struct {
	Id                *string     `json:"id,omitempty" bson:"id,omitempty"`            // Unique id for inter-element referencing
	ContentAttachment *Attachment `json:"contentAttachment" bson:"content_attachment"` // Easily comprehended representation of this Contract
	ContentReference  *Reference  `json:"contentReference" bson:"content_reference"`   // Easily comprehended representation of this Contract
}

func (r *ContractFriendly) Validate() error {
	if r.ContentAttachment == nil {
		return fmt.Errorf("field 'ContentAttachment' is required")
	}
	if r.ContentAttachment != nil {
		if err := r.ContentAttachment.Validate(); err != nil {
			return fmt.Errorf("ContentAttachment: %w", err)
		}
	}
	if r.ContentReference == nil {
		return fmt.Errorf("field 'ContentReference' is required")
	}
	if r.ContentReference != nil {
		if err := r.ContentReference.Validate(); err != nil {
			return fmt.Errorf("ContentReference: %w", err)
		}
	}
	return nil
}

type ContractRule struct {
	Id                *string     `json:"id,omitempty" bson:"id,omitempty"`            // Unique id for inter-element referencing
	ContentAttachment *Attachment `json:"contentAttachment" bson:"content_attachment"` // Computable Contract Rules
	ContentReference  *Reference  `json:"contentReference" bson:"content_reference"`   // Computable Contract Rules
}

func (r *ContractRule) Validate() error {
	if r.ContentAttachment == nil {
		return fmt.Errorf("field 'ContentAttachment' is required")
	}
	if r.ContentAttachment != nil {
		if err := r.ContentAttachment.Validate(); err != nil {
			return fmt.Errorf("ContentAttachment: %w", err)
		}
	}
	if r.ContentReference == nil {
		return fmt.Errorf("field 'ContentReference' is required")
	}
	if r.ContentReference != nil {
		if err := r.ContentReference.Validate(); err != nil {
			return fmt.Errorf("ContentReference: %w", err)
		}
	}
	return nil
}

type ContractTermOfferParty struct {
	Id        *string          `json:"id,omitempty" bson:"id,omitempty"` // Unique id for inter-element referencing
	Reference []Reference      `json:"reference" bson:"reference"`       // Referenced entity
	Role      *CodeableConcept `json:"role" bson:"role"`                 // Participant engagement type
}

func (r *ContractTermOfferParty) Validate() error {
	if len(r.Reference) < 1 {
		return fmt.Errorf("field 'Reference' must have at least 1 elements")
	}
	for i, item := range r.Reference {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Reference[%d]: %w", i, err)
		}
	}
	if r.Role == nil {
		return fmt.Errorf("field 'Role' is required")
	}
	if r.Role != nil {
		if err := r.Role.Validate(); err != nil {
			return fmt.Errorf("Role: %w", err)
		}
	}
	return nil
}

type ContractSigner struct {
	Id        *string     `json:"id,omitempty" bson:"id,omitempty"` // Unique id for inter-element referencing
	Type      *Coding     `json:"type" bson:"type"`                 // Contract Signatory Role
	Party     *Reference  `json:"party" bson:"party"`               // Contract Signatory Party
	Signature []Signature `json:"signature" bson:"signature"`       // Contract Documentation Signature
}

func (r *ContractSigner) Validate() error {
	if r.Type == nil {
		return fmt.Errorf("field 'Type' is required")
	}
	if r.Type != nil {
		if err := r.Type.Validate(); err != nil {
			return fmt.Errorf("Type: %w", err)
		}
	}
	if r.Party == nil {
		return fmt.Errorf("field 'Party' is required")
	}
	if r.Party != nil {
		if err := r.Party.Validate(); err != nil {
			return fmt.Errorf("Party: %w", err)
		}
	}
	if len(r.Signature) < 1 {
		return fmt.Errorf("field 'Signature' must have at least 1 elements")
	}
	for i, item := range r.Signature {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Signature[%d]: %w", i, err)
		}
	}
	return nil
}

type ContractLegal struct {
	Id                *string     `json:"id,omitempty" bson:"id,omitempty"`            // Unique id for inter-element referencing
	ContentAttachment *Attachment `json:"contentAttachment" bson:"content_attachment"` // Contract Legal Text
	ContentReference  *Reference  `json:"contentReference" bson:"content_reference"`   // Contract Legal Text
}

func (r *ContractLegal) Validate() error {
	if r.ContentAttachment == nil {
		return fmt.Errorf("field 'ContentAttachment' is required")
	}
	if r.ContentAttachment != nil {
		if err := r.ContentAttachment.Validate(); err != nil {
			return fmt.Errorf("ContentAttachment: %w", err)
		}
	}
	if r.ContentReference == nil {
		return fmt.Errorf("field 'ContentReference' is required")
	}
	if r.ContentReference != nil {
		if err := r.ContentReference.Validate(); err != nil {
			return fmt.Errorf("ContentReference: %w", err)
		}
	}
	return nil
}

type ContractContentDefinition struct {
	Id                *string          `json:"id,omitempty" bson:"id,omitempty"`                            // Unique id for inter-element referencing
	Type              *CodeableConcept `json:"type" bson:"type"`                                            // Content structure and use
	SubType           *CodeableConcept `json:"subType,omitempty" bson:"sub_type,omitempty"`                 // Detailed Content Type Definition
	Publisher         *Reference       `json:"publisher,omitempty" bson:"publisher,omitempty"`              // Publisher Entity
	PublicationDate   *string          `json:"publicationDate,omitempty" bson:"publication_date,omitempty"` // When published
	PublicationStatus string           `json:"publicationStatus" bson:"publication_status"`                 // amended | appended | cancelled | disputed | entered-in-error | executable +
	Copyright         *string          `json:"copyright,omitempty" bson:"copyright,omitempty"`              // Publication Ownership
}

func (r *ContractContentDefinition) Validate() error {
	if r.Type == nil {
		return fmt.Errorf("field 'Type' is required")
	}
	if r.Type != nil {
		if err := r.Type.Validate(); err != nil {
			return fmt.Errorf("Type: %w", err)
		}
	}
	if r.SubType != nil {
		if err := r.SubType.Validate(); err != nil {
			return fmt.Errorf("SubType: %w", err)
		}
	}
	if r.Publisher != nil {
		if err := r.Publisher.Validate(); err != nil {
			return fmt.Errorf("Publisher: %w", err)
		}
	}
	var emptyString string
	if r.PublicationStatus == emptyString {
		return fmt.Errorf("field 'PublicationStatus' is required")
	}
	return nil
}

type ContractTermOffer struct {
	Id                  *string                   `json:"id,omitempty" bson:"id,omitempty"`                                     // Unique id for inter-element referencing
	Identifier          []Identifier              `json:"identifier,omitempty" bson:"identifier,omitempty"`                     // Offer business ID
	Party               []ContractTermOfferParty  `json:"party,omitempty" bson:"party,omitempty"`                               // Offer Recipient
	Topic               *Reference                `json:"topic,omitempty" bson:"topic,omitempty"`                               // Negotiable offer asset
	Type                *CodeableConcept          `json:"type,omitempty" bson:"type,omitempty"`                                 // Contract Offer Type or Form
	Decision            *CodeableConcept          `json:"decision,omitempty" bson:"decision,omitempty"`                         // Accepting party choice
	DecisionMode        []CodeableConcept         `json:"decisionMode,omitempty" bson:"decision_mode,omitempty"`                // How decision is conveyed
	Answer              []ContractTermOfferAnswer `json:"answer,omitempty" bson:"answer,omitempty"`                             // Response to offer text
	Text                *string                   `json:"text,omitempty" bson:"text,omitempty"`                                 // Human readable offer text
	LinkId              []string                  `json:"linkId,omitempty" bson:"link_id,omitempty"`                            // Pointer to text
	SecurityLabelNumber []int                     `json:"securityLabelNumber,omitempty" bson:"security_label_number,omitempty"` // Offer restriction numbers
}

func (r *ContractTermOffer) Validate() error {
	for i, item := range r.Identifier {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Identifier[%d]: %w", i, err)
		}
	}
	for i, item := range r.Party {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Party[%d]: %w", i, err)
		}
	}
	if r.Topic != nil {
		if err := r.Topic.Validate(); err != nil {
			return fmt.Errorf("Topic: %w", err)
		}
	}
	if r.Type != nil {
		if err := r.Type.Validate(); err != nil {
			return fmt.Errorf("Type: %w", err)
		}
	}
	if r.Decision != nil {
		if err := r.Decision.Validate(); err != nil {
			return fmt.Errorf("Decision: %w", err)
		}
	}
	for i, item := range r.DecisionMode {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("DecisionMode[%d]: %w", i, err)
		}
	}
	for i, item := range r.Answer {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Answer[%d]: %w", i, err)
		}
	}
	return nil
}

type ContractTermOfferAnswer struct {
	Id              *string     `json:"id,omitempty" bson:"id,omitempty"`        // Unique id for inter-element referencing
	ValueBoolean    *bool       `json:"valueBoolean" bson:"value_boolean"`       // The actual answer response
	ValueDecimal    *float64    `json:"valueDecimal" bson:"value_decimal"`       // The actual answer response
	ValueInteger    *int        `json:"valueInteger" bson:"value_integer"`       // The actual answer response
	ValueDate       *string     `json:"valueDate" bson:"value_date"`             // The actual answer response
	ValueDateTime   *string     `json:"valueDateTime" bson:"value_date_time"`    // The actual answer response
	ValueTime       *string     `json:"valueTime" bson:"value_time"`             // The actual answer response
	ValueString     *string     `json:"valueString" bson:"value_string"`         // The actual answer response
	ValueUri        *string     `json:"valueUri" bson:"value_uri"`               // The actual answer response
	ValueAttachment *Attachment `json:"valueAttachment" bson:"value_attachment"` // The actual answer response
	ValueCoding     *Coding     `json:"valueCoding" bson:"value_coding"`         // The actual answer response
	ValueQuantity   *Quantity   `json:"valueQuantity" bson:"value_quantity"`     // The actual answer response
	ValueReference  *Reference  `json:"valueReference" bson:"value_reference"`   // The actual answer response
}

func (r *ContractTermOfferAnswer) Validate() error {
	if r.ValueBoolean == nil {
		return fmt.Errorf("field 'ValueBoolean' is required")
	}
	if r.ValueDecimal == nil {
		return fmt.Errorf("field 'ValueDecimal' is required")
	}
	if r.ValueInteger == nil {
		return fmt.Errorf("field 'ValueInteger' is required")
	}
	if r.ValueDate == nil {
		return fmt.Errorf("field 'ValueDate' is required")
	}
	if r.ValueDateTime == nil {
		return fmt.Errorf("field 'ValueDateTime' is required")
	}
	if r.ValueTime == nil {
		return fmt.Errorf("field 'ValueTime' is required")
	}
	if r.ValueString == nil {
		return fmt.Errorf("field 'ValueString' is required")
	}
	if r.ValueUri == nil {
		return fmt.Errorf("field 'ValueUri' is required")
	}
	if r.ValueAttachment == nil {
		return fmt.Errorf("field 'ValueAttachment' is required")
	}
	if r.ValueAttachment != nil {
		if err := r.ValueAttachment.Validate(); err != nil {
			return fmt.Errorf("ValueAttachment: %w", err)
		}
	}
	if r.ValueCoding == nil {
		return fmt.Errorf("field 'ValueCoding' is required")
	}
	if r.ValueCoding != nil {
		if err := r.ValueCoding.Validate(); err != nil {
			return fmt.Errorf("ValueCoding: %w", err)
		}
	}
	if r.ValueQuantity == nil {
		return fmt.Errorf("field 'ValueQuantity' is required")
	}
	if r.ValueQuantity != nil {
		if err := r.ValueQuantity.Validate(); err != nil {
			return fmt.Errorf("ValueQuantity: %w", err)
		}
	}
	if r.ValueReference == nil {
		return fmt.Errorf("field 'ValueReference' is required")
	}
	if r.ValueReference != nil {
		if err := r.ValueReference.Validate(); err != nil {
			return fmt.Errorf("ValueReference: %w", err)
		}
	}
	return nil
}

type ContractTermAction struct {
	Id                  *string                     `json:"id,omitempty" bson:"id,omitempty"`                                     // Unique id for inter-element referencing
	DoNotPerform        bool                        `json:"doNotPerform,omitempty" bson:"do_not_perform,omitempty"`               // True if the term prohibits the  action
	Type                *CodeableConcept            `json:"type" bson:"type"`                                                     // Type or form of the action
	Subject             []ContractTermActionSubject `json:"subject,omitempty" bson:"subject,omitempty"`                           // Entity of the action
	Intent              *CodeableConcept            `json:"intent" bson:"intent"`                                                 // Purpose for the Contract Term Action
	LinkId              []string                    `json:"linkId,omitempty" bson:"link_id,omitempty"`                            // Pointer to specific item
	Status              *CodeableConcept            `json:"status" bson:"status"`                                                 // State of the action
	Context             *Reference                  `json:"context,omitempty" bson:"context,omitempty"`                           // Episode associated with action
	ContextLinkId       []string                    `json:"contextLinkId,omitempty" bson:"context_link_id,omitempty"`             // Pointer to specific item
	OccurrenceDateTime  *string                     `json:"occurrenceDateTime,omitempty" bson:"occurrence_date_time,omitempty"`   // When action happens
	OccurrencePeriod    *Period                     `json:"occurrencePeriod,omitempty" bson:"occurrence_period,omitempty"`        // When action happens
	OccurrenceTiming    *Timing                     `json:"occurrenceTiming,omitempty" bson:"occurrence_timing,omitempty"`        // When action happens
	Requester           []Reference                 `json:"requester,omitempty" bson:"requester,omitempty"`                       // Who asked for action
	RequesterLinkId     []string                    `json:"requesterLinkId,omitempty" bson:"requester_link_id,omitempty"`         // Pointer to specific item
	PerformerType       []CodeableConcept           `json:"performerType,omitempty" bson:"performer_type,omitempty"`              // Kind of service performer
	PerformerRole       *CodeableConcept            `json:"performerRole,omitempty" bson:"performer_role,omitempty"`              // Competency of the performer
	Performer           *Reference                  `json:"performer,omitempty" bson:"performer,omitempty"`                       // Actor that wil execute (or not) the action
	PerformerLinkId     []string                    `json:"performerLinkId,omitempty" bson:"performer_link_id,omitempty"`         // Pointer to specific item
	Reason              []CodeableReference         `json:"reason,omitempty" bson:"reason,omitempty"`                             // Why is action (not) needed?
	ReasonLinkId        []string                    `json:"reasonLinkId,omitempty" bson:"reason_link_id,omitempty"`               // Pointer to specific item
	Note                []Annotation                `json:"note,omitempty" bson:"note,omitempty"`                                 // Comments about the action
	SecurityLabelNumber []int                       `json:"securityLabelNumber,omitempty" bson:"security_label_number,omitempty"` // Action restriction numbers
}

func (r *ContractTermAction) Validate() error {
	if r.Type == nil {
		return fmt.Errorf("field 'Type' is required")
	}
	if r.Type != nil {
		if err := r.Type.Validate(); err != nil {
			return fmt.Errorf("Type: %w", err)
		}
	}
	for i, item := range r.Subject {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Subject[%d]: %w", i, err)
		}
	}
	if r.Intent == nil {
		return fmt.Errorf("field 'Intent' is required")
	}
	if r.Intent != nil {
		if err := r.Intent.Validate(); err != nil {
			return fmt.Errorf("Intent: %w", err)
		}
	}
	if r.Status == nil {
		return fmt.Errorf("field 'Status' is required")
	}
	if r.Status != nil {
		if err := r.Status.Validate(); err != nil {
			return fmt.Errorf("Status: %w", err)
		}
	}
	if r.Context != nil {
		if err := r.Context.Validate(); err != nil {
			return fmt.Errorf("Context: %w", err)
		}
	}
	if r.OccurrencePeriod != nil {
		if err := r.OccurrencePeriod.Validate(); err != nil {
			return fmt.Errorf("OccurrencePeriod: %w", err)
		}
	}
	if r.OccurrenceTiming != nil {
		if err := r.OccurrenceTiming.Validate(); err != nil {
			return fmt.Errorf("OccurrenceTiming: %w", err)
		}
	}
	for i, item := range r.Requester {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Requester[%d]: %w", i, err)
		}
	}
	for i, item := range r.PerformerType {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("PerformerType[%d]: %w", i, err)
		}
	}
	if r.PerformerRole != nil {
		if err := r.PerformerRole.Validate(); err != nil {
			return fmt.Errorf("PerformerRole: %w", err)
		}
	}
	if r.Performer != nil {
		if err := r.Performer.Validate(); err != nil {
			return fmt.Errorf("Performer: %w", err)
		}
	}
	for i, item := range r.Reason {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Reason[%d]: %w", i, err)
		}
	}
	for i, item := range r.Note {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Note[%d]: %w", i, err)
		}
	}
	return nil
}

type ContractTermActionSubject struct {
	Id        *string          `json:"id,omitempty" bson:"id,omitempty"`     // Unique id for inter-element referencing
	Reference []Reference      `json:"reference" bson:"reference"`           // Entity of the action
	Role      *CodeableConcept `json:"role,omitempty" bson:"role,omitempty"` // Role type of the agent
}

func (r *ContractTermActionSubject) Validate() error {
	if len(r.Reference) < 1 {
		return fmt.Errorf("field 'Reference' must have at least 1 elements")
	}
	for i, item := range r.Reference {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Reference[%d]: %w", i, err)
		}
	}
	if r.Role != nil {
		if err := r.Role.Validate(); err != nil {
			return fmt.Errorf("Role: %w", err)
		}
	}
	return nil
}

type ContractTermAsset struct {
	Id                  *string                       `json:"id,omitempty" bson:"id,omitempty"`                                     // Unique id for inter-element referencing
	Scope               *CodeableConcept              `json:"scope,omitempty" bson:"scope,omitempty"`                               // Range of asset
	Type                []CodeableConcept             `json:"type,omitempty" bson:"type,omitempty"`                                 // Asset category
	TypeReference       []Reference                   `json:"typeReference,omitempty" bson:"type_reference,omitempty"`              // Associated entities
	Subtype             []CodeableConcept             `json:"subtype,omitempty" bson:"subtype,omitempty"`                           // Asset sub-category
	Relationship        *Coding                       `json:"relationship,omitempty" bson:"relationship,omitempty"`                 // Kinship of the asset
	Context             []ContractTermAssetContext    `json:"context,omitempty" bson:"context,omitempty"`                           // Circumstance of the asset
	Condition           *string                       `json:"condition,omitempty" bson:"condition,omitempty"`                       // Quality desctiption of asset
	PeriodType          []CodeableConcept             `json:"periodType,omitempty" bson:"period_type,omitempty"`                    // Asset availability types
	Period              []Period                      `json:"period,omitempty" bson:"period,omitempty"`                             // Time period of the asset
	UsePeriod           []Period                      `json:"usePeriod,omitempty" bson:"use_period,omitempty"`                      // Time period
	Text                *string                       `json:"text,omitempty" bson:"text,omitempty"`                                 // Asset clause or question text
	LinkId              []string                      `json:"linkId,omitempty" bson:"link_id,omitempty"`                            // Pointer to asset text
	Answer              []ContractTermOfferAnswer     `json:"answer,omitempty" bson:"answer,omitempty"`                             // Response to assets
	SecurityLabelNumber []int                         `json:"securityLabelNumber,omitempty" bson:"security_label_number,omitempty"` // Asset restriction numbers
	ValuedItem          []ContractTermAssetValuedItem `json:"valuedItem,omitempty" bson:"valued_item,omitempty"`                    // Contract Valued Item List
}

func (r *ContractTermAsset) Validate() error {
	if r.Scope != nil {
		if err := r.Scope.Validate(); err != nil {
			return fmt.Errorf("Scope: %w", err)
		}
	}
	for i, item := range r.Type {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Type[%d]: %w", i, err)
		}
	}
	for i, item := range r.TypeReference {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("TypeReference[%d]: %w", i, err)
		}
	}
	for i, item := range r.Subtype {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Subtype[%d]: %w", i, err)
		}
	}
	if r.Relationship != nil {
		if err := r.Relationship.Validate(); err != nil {
			return fmt.Errorf("Relationship: %w", err)
		}
	}
	for i, item := range r.Context {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Context[%d]: %w", i, err)
		}
	}
	for i, item := range r.PeriodType {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("PeriodType[%d]: %w", i, err)
		}
	}
	for i, item := range r.Period {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Period[%d]: %w", i, err)
		}
	}
	for i, item := range r.UsePeriod {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("UsePeriod[%d]: %w", i, err)
		}
	}
	for i, item := range r.Answer {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Answer[%d]: %w", i, err)
		}
	}
	for i, item := range r.ValuedItem {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("ValuedItem[%d]: %w", i, err)
		}
	}
	return nil
}

type ContractTermAssetValuedItem struct {
	Id                    *string          `json:"id,omitempty" bson:"id,omitempty"`                                         // Unique id for inter-element referencing
	EntityCodeableConcept *CodeableConcept `json:"entityCodeableConcept,omitempty" bson:"entity_codeable_concept,omitempty"` // Contract Valued Item Type
	EntityReference       *Reference       `json:"entityReference,omitempty" bson:"entity_reference,omitempty"`              // Contract Valued Item Type
	Identifier            *Identifier      `json:"identifier,omitempty" bson:"identifier,omitempty"`                         // Contract Valued Item Number
	EffectiveTime         *string          `json:"effectiveTime,omitempty" bson:"effective_time,omitempty"`                  // Contract Valued Item Effective Tiem
	Quantity              *Quantity        `json:"quantity,omitempty" bson:"quantity,omitempty"`                             // Count of Contract Valued Items
	UnitPrice             *Money           `json:"unitPrice,omitempty" bson:"unit_price,omitempty"`                          // Contract Valued Item fee, charge, or cost
	Factor                *float64         `json:"factor,omitempty" bson:"factor,omitempty"`                                 // Contract Valued Item Price Scaling Factor
	Points                *float64         `json:"points,omitempty" bson:"points,omitempty"`                                 // Contract Valued Item Difficulty Scaling Factor
	Net                   *Money           `json:"net,omitempty" bson:"net,omitempty"`                                       // Total Contract Valued Item Value
	Payment               *string          `json:"payment,omitempty" bson:"payment,omitempty"`                               // Terms of valuation
	PaymentDate           *string          `json:"paymentDate,omitempty" bson:"payment_date,omitempty"`                      // When payment is due
	Responsible           *Reference       `json:"responsible,omitempty" bson:"responsible,omitempty"`                       // Who will make payment
	Recipient             *Reference       `json:"recipient,omitempty" bson:"recipient,omitempty"`                           // Who will receive payment
	LinkId                []string         `json:"linkId,omitempty" bson:"link_id,omitempty"`                                // Pointer to specific item
	SecurityLabelNumber   []int            `json:"securityLabelNumber,omitempty" bson:"security_label_number,omitempty"`     // Security Labels that define affected terms
}

func (r *ContractTermAssetValuedItem) Validate() error {
	if r.EntityCodeableConcept != nil {
		if err := r.EntityCodeableConcept.Validate(); err != nil {
			return fmt.Errorf("EntityCodeableConcept: %w", err)
		}
	}
	if r.EntityReference != nil {
		if err := r.EntityReference.Validate(); err != nil {
			return fmt.Errorf("EntityReference: %w", err)
		}
	}
	if r.Identifier != nil {
		if err := r.Identifier.Validate(); err != nil {
			return fmt.Errorf("Identifier: %w", err)
		}
	}
	if r.Quantity != nil {
		if err := r.Quantity.Validate(); err != nil {
			return fmt.Errorf("Quantity: %w", err)
		}
	}
	if r.UnitPrice != nil {
		if err := r.UnitPrice.Validate(); err != nil {
			return fmt.Errorf("UnitPrice: %w", err)
		}
	}
	if r.Net != nil {
		if err := r.Net.Validate(); err != nil {
			return fmt.Errorf("Net: %w", err)
		}
	}
	if r.Responsible != nil {
		if err := r.Responsible.Validate(); err != nil {
			return fmt.Errorf("Responsible: %w", err)
		}
	}
	if r.Recipient != nil {
		if err := r.Recipient.Validate(); err != nil {
			return fmt.Errorf("Recipient: %w", err)
		}
	}
	return nil
}
