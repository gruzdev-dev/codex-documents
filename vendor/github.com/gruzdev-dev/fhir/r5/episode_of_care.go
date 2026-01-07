package models

import (
	"encoding/json"
	"fmt"
)

// An association between a patient and an organization / healthcare provider(s) during which time encounters may occur. The managing organization assumes a level of responsibility for the patient during this time.
type EpisodeOfCare struct {
	Id                   *string                      `json:"id,omitempty" bson:"id,omitempty"`                                      // Logical id of this artifact
	Meta                 *Meta                        `json:"meta,omitempty" bson:"meta,omitempty"`                                  // Metadata about the resource
	ImplicitRules        *string                      `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`               // A set of rules under which this content was created
	Language             *string                      `json:"language,omitempty" bson:"language,omitempty"`                          // Language of the resource content
	Text                 *Narrative                   `json:"text,omitempty" bson:"text,omitempty"`                                  // Text summary of the resource, for human interpretation
	Contained            []json.RawMessage            `json:"contained,omitempty" bson:"contained,omitempty"`                        // Contained, inline Resources
	Identifier           []Identifier                 `json:"identifier,omitempty" bson:"identifier,omitempty"`                      // Business Identifier(s) relevant for this EpisodeOfCare
	Status               string                       `json:"status" bson:"status"`                                                  // planned | waitlist | active | onhold | finished | cancelled | entered-in-error
	StatusHistory        []EpisodeOfCareStatusHistory `json:"statusHistory,omitempty" bson:"status_history,omitempty"`               // Past list of status codes (the current status may be included to cover the start date of the status)
	Type                 []CodeableConcept            `json:"type,omitempty" bson:"type,omitempty"`                                  // Type/class  - e.g. specialist referral, disease management
	Reason               []EpisodeOfCareReason        `json:"reason,omitempty" bson:"reason,omitempty"`                              // The list of medical reasons that are expected to be addressed during the episode of care
	Diagnosis            []EpisodeOfCareDiagnosis     `json:"diagnosis,omitempty" bson:"diagnosis,omitempty"`                        // The list of medical conditions that were addressed during the episode of care
	Subject              *Reference                   `json:"subject" bson:"subject"`                                                // The patient/group who is the focus of this episode of care
	ManagingOrganization *Reference                   `json:"managingOrganization,omitempty" bson:"managing_organization,omitempty"` // Organization that assumes responsibility for care coordination
	Period               *Period                      `json:"period,omitempty" bson:"period,omitempty"`                              // Interval during responsibility is assumed
	ReferralRequest      []Reference                  `json:"referralRequest,omitempty" bson:"referral_request,omitempty"`           // Originating Referral Request(s)
	CareManager          *Reference                   `json:"careManager,omitempty" bson:"care_manager,omitempty"`                   // Care manager/care coordinator for the patient
	CareTeam             []Reference                  `json:"careTeam,omitempty" bson:"care_team,omitempty"`                         // Other practitioners facilitating this episode of care
	Account              []Reference                  `json:"account,omitempty" bson:"account,omitempty"`                            // The set of accounts that may be used for billing for this EpisodeOfCare
}

func (r *EpisodeOfCare) Validate() error {
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
	for i, item := range r.StatusHistory {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("StatusHistory[%d]: %w", i, err)
		}
	}
	for i, item := range r.Type {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Type[%d]: %w", i, err)
		}
	}
	for i, item := range r.Reason {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Reason[%d]: %w", i, err)
		}
	}
	for i, item := range r.Diagnosis {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Diagnosis[%d]: %w", i, err)
		}
	}
	if r.Subject == nil {
		return fmt.Errorf("field 'Subject' is required")
	}
	if r.Subject != nil {
		if err := r.Subject.Validate(); err != nil {
			return fmt.Errorf("Subject: %w", err)
		}
	}
	if r.ManagingOrganization != nil {
		if err := r.ManagingOrganization.Validate(); err != nil {
			return fmt.Errorf("ManagingOrganization: %w", err)
		}
	}
	if r.Period != nil {
		if err := r.Period.Validate(); err != nil {
			return fmt.Errorf("Period: %w", err)
		}
	}
	for i, item := range r.ReferralRequest {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("ReferralRequest[%d]: %w", i, err)
		}
	}
	if r.CareManager != nil {
		if err := r.CareManager.Validate(); err != nil {
			return fmt.Errorf("CareManager: %w", err)
		}
	}
	for i, item := range r.CareTeam {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("CareTeam[%d]: %w", i, err)
		}
	}
	for i, item := range r.Account {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Account[%d]: %w", i, err)
		}
	}
	return nil
}

type EpisodeOfCareReason struct {
	Id    *string             `json:"id,omitempty" bson:"id,omitempty"`       // Unique id for inter-element referencing
	Use   []CodeableConcept   `json:"use,omitempty" bson:"use,omitempty"`     // What the reason value should be used for/as
	Value []CodeableReference `json:"value,omitempty" bson:"value,omitempty"` // Medical reason to be addressed
}

func (r *EpisodeOfCareReason) Validate() error {
	for i, item := range r.Use {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Use[%d]: %w", i, err)
		}
	}
	for i, item := range r.Value {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Value[%d]: %w", i, err)
		}
	}
	return nil
}

type EpisodeOfCareDiagnosis struct {
	Id        *string             `json:"id,omitempty" bson:"id,omitempty"`               // Unique id for inter-element referencing
	Condition []CodeableReference `json:"condition,omitempty" bson:"condition,omitempty"` // The medical condition that was addressed during the episode of care
	Use       []CodeableConcept   `json:"use,omitempty" bson:"use,omitempty"`             // Role that this diagnosis has within the episode of care (e.g. admission, billing, discharge …)
}

func (r *EpisodeOfCareDiagnosis) Validate() error {
	for i, item := range r.Condition {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Condition[%d]: %w", i, err)
		}
	}
	for i, item := range r.Use {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Use[%d]: %w", i, err)
		}
	}
	return nil
}

type EpisodeOfCareStatusHistory struct {
	Id     *string `json:"id,omitempty" bson:"id,omitempty"` // Unique id for inter-element referencing
	Status string  `json:"status" bson:"status"`             // planned | waitlist | active | onhold | finished | cancelled | entered-in-error
	Period *Period `json:"period" bson:"period"`             // Duration the EpisodeOfCare was in the specified status
}

func (r *EpisodeOfCareStatusHistory) Validate() error {
	var emptyString string
	if r.Status == emptyString {
		return fmt.Errorf("field 'Status' is required")
	}
	if r.Period == nil {
		return fmt.Errorf("field 'Period' is required")
	}
	if r.Period != nil {
		if err := r.Period.Validate(); err != nil {
			return fmt.Errorf("Period: %w", err)
		}
	}
	return nil
}
