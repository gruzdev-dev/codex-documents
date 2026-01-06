package models

import (
	"encoding/json"
	"fmt"
)

// The Care Team includes all the people, organizations, and care teams who participate or plan to participate in the coordination and delivery of care.
type CareTeam struct {
	Id                   *string               `json:"id,omitempty" bson:"id,omitempty"`                                      // Logical id of this artifact
	Meta                 *Meta                 `json:"meta,omitempty" bson:"meta,omitempty"`                                  // Metadata about the resource
	ImplicitRules        *string               `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`               // A set of rules under which this content was created
	Language             *string               `json:"language,omitempty" bson:"language,omitempty"`                          // Language of the resource content
	Text                 *Narrative            `json:"text,omitempty" bson:"text,omitempty"`                                  // Text summary of the resource, for human interpretation
	Contained            []json.RawMessage     `json:"contained,omitempty" bson:"contained,omitempty"`                        // Contained, inline Resources
	Identifier           []Identifier          `json:"identifier,omitempty" bson:"identifier,omitempty"`                      // External Ids for this team
	Status               *string               `json:"status,omitempty" bson:"status,omitempty"`                              // proposed | active | suspended | inactive | entered-in-error
	Category             []CodeableConcept     `json:"category,omitempty" bson:"category,omitempty"`                          // Type of team
	Name                 *string               `json:"name,omitempty" bson:"name,omitempty"`                                  // Name of the team, such as crisis assessment team
	Subject              *Reference            `json:"subject,omitempty" bson:"subject,omitempty"`                            // Who care team is for
	Period               *Period               `json:"period,omitempty" bson:"period,omitempty"`                              // Time period team covers
	Participant          []CareTeamParticipant `json:"participant,omitempty" bson:"participant,omitempty"`                    // Members of the team
	Reason               []CodeableReference   `json:"reason,omitempty" bson:"reason,omitempty"`                              // Why the care team exists
	ManagingOrganization []Reference           `json:"managingOrganization,omitempty" bson:"managing_organization,omitempty"` // Organization responsible for the care team
	Telecom              []ContactPoint        `json:"telecom,omitempty" bson:"telecom,omitempty"`                            // A contact detail for the care team (that applies to all members)
	Note                 []Annotation          `json:"note,omitempty" bson:"note,omitempty"`                                  // Comments made about the CareTeam
}

func (r *CareTeam) Validate() error {
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
	for i, item := range r.Participant {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Participant[%d]: %w", i, err)
		}
	}
	for i, item := range r.Reason {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Reason[%d]: %w", i, err)
		}
	}
	for i, item := range r.ManagingOrganization {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("ManagingOrganization[%d]: %w", i, err)
		}
	}
	for i, item := range r.Telecom {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Telecom[%d]: %w", i, err)
		}
	}
	for i, item := range r.Note {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Note[%d]: %w", i, err)
		}
	}
	return nil
}

type CareTeamParticipant struct {
	Id              *string          `json:"id,omitempty" bson:"id,omitempty"`                            // Unique id for inter-element referencing
	Role            *CodeableConcept `json:"role,omitempty" bson:"role,omitempty"`                        // Type of involvement
	Member          *Reference       `json:"member,omitempty" bson:"member,omitempty"`                    // Who is involved
	OnBehalfOf      *Reference       `json:"onBehalfOf,omitempty" bson:"on_behalf_of,omitempty"`          // Entity that the participant is acting as a proxy of, or an agent of, or in the interest of, or as a representative of
	EffectivePeriod *Period          `json:"effectivePeriod,omitempty" bson:"effective_period,omitempty"` // When the member is generally available within this care team
	EffectiveTiming *Timing          `json:"effectiveTiming,omitempty" bson:"effective_timing,omitempty"` // When the member is generally available within this care team
	SupportingInfo  []Reference      `json:"supportingInfo,omitempty" bson:"supporting_info,omitempty"`   // Basis for the member's participation
}

func (r *CareTeamParticipant) Validate() error {
	if r.Role != nil {
		if err := r.Role.Validate(); err != nil {
			return fmt.Errorf("Role: %w", err)
		}
	}
	if r.Member != nil {
		if err := r.Member.Validate(); err != nil {
			return fmt.Errorf("Member: %w", err)
		}
	}
	if r.OnBehalfOf != nil {
		if err := r.OnBehalfOf.Validate(); err != nil {
			return fmt.Errorf("OnBehalfOf: %w", err)
		}
	}
	if r.EffectivePeriod != nil {
		if err := r.EffectivePeriod.Validate(); err != nil {
			return fmt.Errorf("EffectivePeriod: %w", err)
		}
	}
	if r.EffectiveTiming != nil {
		if err := r.EffectiveTiming.Validate(); err != nil {
			return fmt.Errorf("EffectiveTiming: %w", err)
		}
	}
	for i, item := range r.SupportingInfo {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("SupportingInfo[%d]: %w", i, err)
		}
	}
	return nil
}
