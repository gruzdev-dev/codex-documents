package models

import (
	"encoding/json"
	"fmt"
)

// An event (i.e. any change to current patient status) that may be related to unintended effects on a patient or research participant. The unintended effects may require additional monitoring, treatment, hospitalization, or may result in death. The AdverseEvent resource also extends to potential or avoided events that could have had such effects. There are two major domains where the AdverseEvent resource is expected to be used. One is in clinical care reported adverse events and the other is in reporting adverse events in clinical  research trial management.  Adverse events can be reported by healthcare providers, patients, caregivers or by medical products manufacturers.  Given the differences between these two concepts, we recommend consulting the domain specific implementation guides when implementing the AdverseEvent Resource. The implementation guides include specific extensions, value sets and constraints.
type AdverseEvent struct {
	ResourceType            string                      `json:"resourceType" bson:"resource_type"`                                             // Type of resource
	Id                      *string                     `json:"id,omitempty" bson:"id,omitempty"`                                              // Logical id of this artifact
	Meta                    *Meta                       `json:"meta,omitempty" bson:"meta,omitempty"`                                          // Metadata about the resource
	ImplicitRules           *string                     `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`                       // A set of rules under which this content was created
	Language                *string                     `json:"language,omitempty" bson:"language,omitempty"`                                  // Language of the resource content
	Text                    *Narrative                  `json:"text,omitempty" bson:"text,omitempty"`                                          // Text summary of the resource, for human interpretation
	Contained               []json.RawMessage           `json:"contained,omitempty" bson:"contained,omitempty"`                                // Contained, inline Resources
	Identifier              []Identifier                `json:"identifier,omitempty" bson:"identifier,omitempty"`                              // Business identifier for the event
	Status                  string                      `json:"status" bson:"status"`                                                          // in-progress | completed | entered-in-error | unknown
	Actuality               string                      `json:"actuality" bson:"actuality"`                                                    // actual | potential
	Category                []CodeableConcept           `json:"category,omitempty" bson:"category,omitempty"`                                  // wrong-patient | procedure-mishap | medication-mishap | device | unsafe-physical-environment | hospital-aquired-infection | wrong-body-site
	Code                    *CodeableConcept            `json:"code,omitempty" bson:"code,omitempty"`                                          // Event or incident that occurred or was averted
	Subject                 *Reference                  `json:"subject" bson:"subject"`                                                        // Subject impacted by event
	Encounter               *Reference                  `json:"encounter,omitempty" bson:"encounter,omitempty"`                                // The Encounter associated with the start of the AdverseEvent
	EffectDateTime          *string                     `json:"effectDateTime,omitempty" bson:"effect_date_time,omitempty"`                    // When the effect of the AdverseEvent occurred
	EffectPeriod            *Period                     `json:"effectPeriod,omitempty" bson:"effect_period,omitempty"`                         // When the effect of the AdverseEvent occurred
	Detected                *string                     `json:"detected,omitempty" bson:"detected,omitempty"`                                  // When the event was detected
	RecordedDate            *string                     `json:"recordedDate,omitempty" bson:"recorded_date,omitempty"`                         // When the event was recorded
	ResultingEffect         []CodeableReference         `json:"resultingEffect,omitempty" bson:"resulting_effect,omitempty"`                   // Effect on the subject due to this event
	Location                *Reference                  `json:"location,omitempty" bson:"location,omitempty"`                                  // Location where adverse event occurred
	Seriousness             *CodeableConcept            `json:"seriousness,omitempty" bson:"seriousness,omitempty"`                            // Seriousness or gravity of the event
	Outcome                 []CodeableConcept           `json:"outcome,omitempty" bson:"outcome,omitempty"`                                    // Type of outcome from the adverse event
	Recorder                *Reference                  `json:"recorder,omitempty" bson:"recorder,omitempty"`                                  // Who recorded the adverse event
	Participant             []AdverseEventParticipant   `json:"participant,omitempty" bson:"participant,omitempty"`                            // Who was involved in the adverse event or the potential adverse event and what they did
	Study                   []Reference                 `json:"study,omitempty" bson:"study,omitempty"`                                        // Research study that the subject is enrolled in
	ExpectedInResearchStudy *bool                       `json:"expectedInResearchStudy,omitempty" bson:"expected_in_research_study,omitempty"` // Considered likely or probable or anticipated in the research study
	SuspectEntity           []AdverseEventSuspectEntity `json:"suspectEntity,omitempty" bson:"suspect_entity,omitempty"`                       // The suspected agent causing the adverse event
	ContributingFactor      []CodeableReference         `json:"contributingFactor,omitempty" bson:"contributing_factor,omitempty"`             // Contributing factors suspected to have increased the probability or severity of the adverse event
	PreventiveAction        []CodeableReference         `json:"preventiveAction,omitempty" bson:"preventive_action,omitempty"`                 // Preventive actions that contributed to avoiding the adverse event
	MitigatingAction        []CodeableReference         `json:"mitigatingAction,omitempty" bson:"mitigating_action,omitempty"`                 // Ameliorating actions taken after the adverse event occurred in order to reduce the extent of harm
	SupportingInfo          []CodeableReference         `json:"supportingInfo,omitempty" bson:"supporting_info,omitempty"`                     // Subject medical history or document relevant to this adverse event
	Note                    []Annotation                `json:"note,omitempty" bson:"note,omitempty"`                                          // Comment on adverse event
}

func (r *AdverseEvent) Validate() error {
	if r.ResourceType != "AdverseEvent" {
		return fmt.Errorf("invalid resourceType: expected 'AdverseEvent', got '%s'", r.ResourceType)
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
	if r.Actuality == emptyString {
		return fmt.Errorf("field 'Actuality' is required")
	}
	for i, item := range r.Category {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Category[%d]: %w", i, err)
		}
	}
	if r.Code != nil {
		if err := r.Code.Validate(); err != nil {
			return fmt.Errorf("Code: %w", err)
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
	if r.Encounter != nil {
		if err := r.Encounter.Validate(); err != nil {
			return fmt.Errorf("Encounter: %w", err)
		}
	}
	if r.EffectPeriod != nil {
		if err := r.EffectPeriod.Validate(); err != nil {
			return fmt.Errorf("EffectPeriod: %w", err)
		}
	}
	for i, item := range r.ResultingEffect {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("ResultingEffect[%d]: %w", i, err)
		}
	}
	if r.Location != nil {
		if err := r.Location.Validate(); err != nil {
			return fmt.Errorf("Location: %w", err)
		}
	}
	if r.Seriousness != nil {
		if err := r.Seriousness.Validate(); err != nil {
			return fmt.Errorf("Seriousness: %w", err)
		}
	}
	for i, item := range r.Outcome {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Outcome[%d]: %w", i, err)
		}
	}
	if r.Recorder != nil {
		if err := r.Recorder.Validate(); err != nil {
			return fmt.Errorf("Recorder: %w", err)
		}
	}
	for i, item := range r.Participant {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Participant[%d]: %w", i, err)
		}
	}
	for i, item := range r.Study {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Study[%d]: %w", i, err)
		}
	}
	for i, item := range r.SuspectEntity {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("SuspectEntity[%d]: %w", i, err)
		}
	}
	for i, item := range r.ContributingFactor {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("ContributingFactor[%d]: %w", i, err)
		}
	}
	for i, item := range r.PreventiveAction {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("PreventiveAction[%d]: %w", i, err)
		}
	}
	for i, item := range r.MitigatingAction {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("MitigatingAction[%d]: %w", i, err)
		}
	}
	for i, item := range r.SupportingInfo {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("SupportingInfo[%d]: %w", i, err)
		}
	}
	for i, item := range r.Note {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Note[%d]: %w", i, err)
		}
	}
	return nil
}

type AdverseEventParticipant struct {
	Id       *string          `json:"id,omitempty" bson:"id,omitempty"`             // Unique id for inter-element referencing
	Function *CodeableConcept `json:"function,omitempty" bson:"function,omitempty"` // Type of involvement
	Actor    *Reference       `json:"actor" bson:"actor"`                           // Who was involved in the adverse event or the potential adverse event
}

func (r *AdverseEventParticipant) Validate() error {
	if r.Function != nil {
		if err := r.Function.Validate(); err != nil {
			return fmt.Errorf("Function: %w", err)
		}
	}
	if r.Actor == nil {
		return fmt.Errorf("field 'Actor' is required")
	}
	if r.Actor != nil {
		if err := r.Actor.Validate(); err != nil {
			return fmt.Errorf("Actor: %w", err)
		}
	}
	return nil
}

type AdverseEventSuspectEntity struct {
	Id                 *string                             `json:"id,omitempty" bson:"id,omitempty"`                                   // Unique id for inter-element referencing
	Instance           *CodeableReference                  `json:"instance" bson:"instance"`                                           // Refers to the specific entity that caused the adverse event
	Causality          *AdverseEventSuspectEntityCausality `json:"causality,omitempty" bson:"causality,omitempty"`                     // Information on the possible cause of the event
	OccurrenceDateTime *string                             `json:"occurrenceDateTime,omitempty" bson:"occurrence_date_time,omitempty"` // When the suspect entity occurred
	OccurrencePeriod   *Period                             `json:"occurrencePeriod,omitempty" bson:"occurrence_period,omitempty"`      // When the suspect entity occurred
}

func (r *AdverseEventSuspectEntity) Validate() error {
	if r.Instance == nil {
		return fmt.Errorf("field 'Instance' is required")
	}
	if r.Instance != nil {
		if err := r.Instance.Validate(); err != nil {
			return fmt.Errorf("Instance: %w", err)
		}
	}
	if r.Causality != nil {
		if err := r.Causality.Validate(); err != nil {
			return fmt.Errorf("Causality: %w", err)
		}
	}
	if r.OccurrencePeriod != nil {
		if err := r.OccurrencePeriod.Validate(); err != nil {
			return fmt.Errorf("OccurrencePeriod: %w", err)
		}
	}
	return nil
}

type AdverseEventSuspectEntityCausality struct {
	Id                *string          `json:"id,omitempty" bson:"id,omitempty"`                                // Unique id for inter-element referencing
	AssessmentMethod  *CodeableConcept `json:"assessmentMethod,omitempty" bson:"assessment_method,omitempty"`   // Method of evaluating the relatedness of the suspected entity to the event
	EntityRelatedness *CodeableConcept `json:"entityRelatedness,omitempty" bson:"entity_relatedness,omitempty"` // Result of the assessment regarding the relatedness of the suspected entity to the event
	Author            *Reference       `json:"author,omitempty" bson:"author,omitempty"`                        // Author of the information on the possible cause of the event
}

func (r *AdverseEventSuspectEntityCausality) Validate() error {
	if r.AssessmentMethod != nil {
		if err := r.AssessmentMethod.Validate(); err != nil {
			return fmt.Errorf("AssessmentMethod: %w", err)
		}
	}
	if r.EntityRelatedness != nil {
		if err := r.EntityRelatedness.Validate(); err != nil {
			return fmt.Errorf("EntityRelatedness: %w", err)
		}
	}
	if r.Author != nil {
		if err := r.Author.Validate(); err != nil {
			return fmt.Errorf("Author: %w", err)
		}
	}
	return nil
}
