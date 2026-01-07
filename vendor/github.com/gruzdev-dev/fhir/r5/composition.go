package models

import (
	"encoding/json"
	"fmt"
)

// A set of healthcare-related information that is assembled together into a single logical package that provides a single coherent statement of meaning, establishes its own context and has traceability to the author who is making the statement. A Composition defines the structure and narrative content necessary for a document. However, a Composition alone does not constitute a document. Rather, the Composition must be the first entry in a Bundle where Bundle.type=document, and any other resources referenced from Composition must be included as subsequent entries in the Bundle (for example Patient, Practitioner, Encounter, etc.).
type Composition struct {
	ResourceType  string                   `json:"resourceType" bson:"resource_type"`                       // Type of resource
	Id            *string                  `json:"id,omitempty" bson:"id,omitempty"`                        // Logical id of this artifact
	Meta          *Meta                    `json:"meta,omitempty" bson:"meta,omitempty"`                    // Metadata about the resource
	ImplicitRules *string                  `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"` // A set of rules under which this content was created
	Language      *string                  `json:"language,omitempty" bson:"language,omitempty"`            // Language of the resource content
	Text          *Narrative               `json:"text,omitempty" bson:"text,omitempty"`                    // Text summary of the resource, for human interpretation
	Contained     []json.RawMessage        `json:"contained,omitempty" bson:"contained,omitempty"`          // Contained, inline Resources
	Url           *string                  `json:"url,omitempty" bson:"url,omitempty"`                      // Canonical identifier for this Composition, represented as a URI (globally unique)
	Identifier    []Identifier             `json:"identifier,omitempty" bson:"identifier,omitempty"`        // Version-independent identifier for the Composition
	Version       *string                  `json:"version,omitempty" bson:"version,omitempty"`              // An explicitly assigned identifier of a variation of the content in the Composition
	Consent       []Reference              `json:"consent,omitempty" bson:"consent,omitempty"`              // Represents informed consents and medico-legal transactions
	BasedOn       []Reference              `json:"basedOn,omitempty" bson:"based_on,omitempty"`             // Fulfills plan, proposal or order
	Status        string                   `json:"status" bson:"status"`                                    // registered | partial | preliminary | final | amended | corrected | appended | cancelled | entered-in-error | deprecated | unknown
	Type          *CodeableConcept         `json:"type" bson:"type"`                                        // Kind of composition (LOINC if possible)
	Category      []CodeableConcept        `json:"category,omitempty" bson:"category,omitempty"`            // Categorization of Composition
	Subject       []Reference              `json:"subject,omitempty" bson:"subject,omitempty"`              // Who and/or what the composition is about
	Encounter     *Reference               `json:"encounter,omitempty" bson:"encounter,omitempty"`          // Context of the Composition
	Date          string                   `json:"date" bson:"date"`                                        // Composition editing time
	UseContext    []UsageContext           `json:"useContext,omitempty" bson:"use_context,omitempty"`       // The context that the content is intended to support
	Author        []Reference              `json:"author,omitempty" bson:"author,omitempty"`                // Who and/or what authored the composition
	Participant   []CompositionParticipant `json:"participant,omitempty" bson:"participant,omitempty"`      // Identifies supporting entities, including parents, relatives, caregivers, insurance policyholders, guarantors, and others related in some way to the patient
	Name          *string                  `json:"name,omitempty" bson:"name,omitempty"`                    // Name for this Composition (computer friendly)
	Title         *string                  `json:"title,omitempty" bson:"title,omitempty"`                  // Human Readable name/title
	Note          []Annotation             `json:"note,omitempty" bson:"note,omitempty"`                    // For any additional notes
	Attester      []CompositionAttester    `json:"attester,omitempty" bson:"attester,omitempty"`            // Attests to accuracy of composition
	Custodian     *Reference               `json:"custodian,omitempty" bson:"custodian,omitempty"`          // Organization which maintains the composition
	RelatesTo     []CompositionRelatesTo   `json:"relatesTo,omitempty" bson:"relates_to,omitempty"`         // Relationships to other compositions/documents
	Event         []CompositionEvent       `json:"event,omitempty" bson:"event,omitempty"`                  // The clinical service(s) being documented
	Section       []CompositionSection     `json:"section,omitempty" bson:"section,omitempty"`              // Composition is broken into sections
}

func (r *Composition) Validate() error {
	if r.ResourceType != "Composition" {
		return fmt.Errorf("invalid resourceType: expected 'Composition', got '%s'", r.ResourceType)
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
	for i, item := range r.Consent {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Consent[%d]: %w", i, err)
		}
	}
	for i, item := range r.BasedOn {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("BasedOn[%d]: %w", i, err)
		}
	}
	var emptyString string
	if r.Status == emptyString {
		return fmt.Errorf("field 'Status' is required")
	}
	if r.Type == nil {
		return fmt.Errorf("field 'Type' is required")
	}
	if r.Type != nil {
		if err := r.Type.Validate(); err != nil {
			return fmt.Errorf("Type: %w", err)
		}
	}
	for i, item := range r.Category {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Category[%d]: %w", i, err)
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
	if r.Date == emptyString {
		return fmt.Errorf("field 'Date' is required")
	}
	for i, item := range r.UseContext {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("UseContext[%d]: %w", i, err)
		}
	}
	for i, item := range r.Author {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Author[%d]: %w", i, err)
		}
	}
	for i, item := range r.Participant {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Participant[%d]: %w", i, err)
		}
	}
	for i, item := range r.Note {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Note[%d]: %w", i, err)
		}
	}
	for i, item := range r.Attester {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Attester[%d]: %w", i, err)
		}
	}
	if r.Custodian != nil {
		if err := r.Custodian.Validate(); err != nil {
			return fmt.Errorf("Custodian: %w", err)
		}
	}
	for i, item := range r.RelatesTo {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("RelatesTo[%d]: %w", i, err)
		}
	}
	for i, item := range r.Event {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Event[%d]: %w", i, err)
		}
	}
	for i, item := range r.Section {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Section[%d]: %w", i, err)
		}
	}
	return nil
}

type CompositionRelatesTo struct {
	Id               *string          `json:"id,omitempty" bson:"id,omitempty"`          // Unique id for inter-element referencing
	Type             *CodeableConcept `json:"type" bson:"type"`                          // documentation | justification | citation | predecessor | successor | derived-from | depends-on | composed-of | part-of | amends | amended-with | appends | appended-with | cites | cited-by | comments-on | comment-in | contains | contained-in | corrects | correction-in | replaces | replaced-with | retracts | retracted-by | signs | similar-to | supports | supported-with | transforms | transformed-into | transformed-with | documents | specification-of | created-with | cite-as | reprint | reprint-of | summarizes
	TargetUri        *string          `json:"targetUri" bson:"target_uri"`               // The artifact that is related to this Composition
	TargetAttachment *Attachment      `json:"targetAttachment" bson:"target_attachment"` // The artifact that is related to this Composition
	TargetCanonical  *string          `json:"targetCanonical" bson:"target_canonical"`   // The artifact that is related to this Composition
	TargetReference  *Reference       `json:"targetReference" bson:"target_reference"`   // The artifact that is related to this Composition
	TargetMarkdown   *string          `json:"targetMarkdown" bson:"target_markdown"`     // The artifact that is related to this Composition
}

func (r *CompositionRelatesTo) Validate() error {
	if r.Type == nil {
		return fmt.Errorf("field 'Type' is required")
	}
	if r.Type != nil {
		if err := r.Type.Validate(); err != nil {
			return fmt.Errorf("Type: %w", err)
		}
	}
	if r.TargetUri == nil {
		return fmt.Errorf("field 'TargetUri' is required")
	}
	if r.TargetAttachment == nil {
		return fmt.Errorf("field 'TargetAttachment' is required")
	}
	if r.TargetAttachment != nil {
		if err := r.TargetAttachment.Validate(); err != nil {
			return fmt.Errorf("TargetAttachment: %w", err)
		}
	}
	if r.TargetCanonical == nil {
		return fmt.Errorf("field 'TargetCanonical' is required")
	}
	if r.TargetReference == nil {
		return fmt.Errorf("field 'TargetReference' is required")
	}
	if r.TargetReference != nil {
		if err := r.TargetReference.Validate(); err != nil {
			return fmt.Errorf("TargetReference: %w", err)
		}
	}
	if r.TargetMarkdown == nil {
		return fmt.Errorf("field 'TargetMarkdown' is required")
	}
	return nil
}

type CompositionEvent struct {
	Id     *string             `json:"id,omitempty" bson:"id,omitempty"`         // Unique id for inter-element referencing
	Period *Period             `json:"period,omitempty" bson:"period,omitempty"` // The period covered by the documentation
	Detail []CodeableReference `json:"detail,omitempty" bson:"detail,omitempty"` // The event(s) being documented, as code(s), reference(s), or both
}

func (r *CompositionEvent) Validate() error {
	if r.Period != nil {
		if err := r.Period.Validate(); err != nil {
			return fmt.Errorf("Period: %w", err)
		}
	}
	for i, item := range r.Detail {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Detail[%d]: %w", i, err)
		}
	}
	return nil
}

type CompositionSection struct {
	Id          *string              `json:"id,omitempty" bson:"id,omitempty"`                    // Unique id for inter-element referencing
	Title       *string              `json:"title,omitempty" bson:"title,omitempty"`              // Label for section (e.g. for ToC)
	Code        *CodeableConcept     `json:"code,omitempty" bson:"code,omitempty"`                // Classification of section (recommended)
	Author      []Reference          `json:"author,omitempty" bson:"author,omitempty"`            // Who and/or what authored the section, when the section is authored by someone other than the composition.author
	Focus       *Reference           `json:"focus,omitempty" bson:"focus,omitempty"`              // Who/what the section is about, when it is not about the subject of composition
	Text        *Narrative           `json:"text,omitempty" bson:"text,omitempty"`                // Text summary of the section, for human interpretation
	Note        []Annotation         `json:"note,omitempty" bson:"note,omitempty"`                // Information about the section contents that is not represented by any of the section entries
	OrderedBy   *CodeableConcept     `json:"orderedBy,omitempty" bson:"ordered_by,omitempty"`     // Order of section entries
	Entry       []Reference          `json:"entry,omitempty" bson:"entry,omitempty"`              // A reference to data that supports this section
	EmptyReason *CodeableConcept     `json:"emptyReason,omitempty" bson:"empty_reason,omitempty"` // Why the section is empty
	Section     []CompositionSection `json:"section,omitempty" bson:"section,omitempty"`          // Nested Section
}

func (r *CompositionSection) Validate() error {
	if r.Code != nil {
		if err := r.Code.Validate(); err != nil {
			return fmt.Errorf("Code: %w", err)
		}
	}
	for i, item := range r.Author {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Author[%d]: %w", i, err)
		}
	}
	if r.Focus != nil {
		if err := r.Focus.Validate(); err != nil {
			return fmt.Errorf("Focus: %w", err)
		}
	}
	if r.Text != nil {
		if err := r.Text.Validate(); err != nil {
			return fmt.Errorf("Text: %w", err)
		}
	}
	for i, item := range r.Note {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Note[%d]: %w", i, err)
		}
	}
	if r.OrderedBy != nil {
		if err := r.OrderedBy.Validate(); err != nil {
			return fmt.Errorf("OrderedBy: %w", err)
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
	for i, item := range r.Section {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Section[%d]: %w", i, err)
		}
	}
	return nil
}

type CompositionParticipant struct {
	Id       *string           `json:"id,omitempty" bson:"id,omitempty"` // Unique id for inter-element referencing
	Type     []CodeableConcept `json:"type" bson:"type"`                 // AUT | AUTHEN | CST | LA | RCT | SBJ
	Function []CodeableConcept `json:"function,omitempty" bson:"function,omitempty"`
	Time     *Period           `json:"time,omitempty" bson:"time,omitempty"` // Time period of participation
	Party    *Reference        `json:"party" bson:"party"`                   // Who the participant is
}

func (r *CompositionParticipant) Validate() error {
	if len(r.Type) < 1 {
		return fmt.Errorf("field 'Type' must have at least 1 elements")
	}
	for i, item := range r.Type {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Type[%d]: %w", i, err)
		}
	}
	for i, item := range r.Function {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Function[%d]: %w", i, err)
		}
	}
	if r.Time != nil {
		if err := r.Time.Validate(); err != nil {
			return fmt.Errorf("Time: %w", err)
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
	return nil
}

type CompositionAttester struct {
	Id    *string          `json:"id,omitempty" bson:"id,omitempty"`       // Unique id for inter-element referencing
	Mode  *CodeableConcept `json:"mode" bson:"mode"`                       // personal | professional | legal | official
	Time  *string          `json:"time,omitempty" bson:"time,omitempty"`   // When the composition was attested
	Party *Reference       `json:"party,omitempty" bson:"party,omitempty"` // Who attested the composition
}

func (r *CompositionAttester) Validate() error {
	if r.Mode == nil {
		return fmt.Errorf("field 'Mode' is required")
	}
	if r.Mode != nil {
		if err := r.Mode.Validate(); err != nil {
			return fmt.Errorf("Mode: %w", err)
		}
	}
	if r.Party != nil {
		if err := r.Party.Validate(); err != nil {
			return fmt.Errorf("Party: %w", err)
		}
	}
	return nil
}
