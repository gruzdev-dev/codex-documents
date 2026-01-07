package models

import (
	"encoding/json"
	"fmt"
)

// A scientific study intended to increase health-related knowledge. For example, clinical trials are research studies that involve people. These studies may be related to new ways to screen, prevent, diagnose, and treat disease. They may also study certain outcomes and certain groups of people by looking at data collected in the past or future.
type ResearchStudy struct {
	Id                 *string                        `json:"id,omitempty" bson:"id,omitempty"`                                   // Logical id of this artifact
	Meta               *Meta                          `json:"meta,omitempty" bson:"meta,omitempty"`                               // Metadata about the resource
	ImplicitRules      *string                        `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`            // A set of rules under which this content was created
	Language           *string                        `json:"language,omitempty" bson:"language,omitempty"`                       // Language of the resource content
	Text               *Narrative                     `json:"text,omitempty" bson:"text,omitempty"`                               // Text summary of the resource, for human interpretation
	Contained          []json.RawMessage              `json:"contained,omitempty" bson:"contained,omitempty"`                     // Contained, inline Resources
	Url                *string                        `json:"url,omitempty" bson:"url,omitempty"`                                 // Canonical identifier for this study resource
	Identifier         []Identifier                   `json:"identifier,omitempty" bson:"identifier,omitempty"`                   // Business Identifier for study
	Version            *string                        `json:"version,omitempty" bson:"version,omitempty"`                         // The business version for the study record
	Name               *string                        `json:"name,omitempty" bson:"name,omitempty"`                               // Name for this study (computer friendly)
	Title              *string                        `json:"title,omitempty" bson:"title,omitempty"`                             // Human readable name of the study
	Label              []ResearchStudyLabel           `json:"label,omitempty" bson:"label,omitempty"`                             // Additional names for the study
	Protocol           []Reference                    `json:"protocol,omitempty" bson:"protocol,omitempty"`                       // Steps followed in executing study
	PartOf             []Reference                    `json:"partOf,omitempty" bson:"part_of,omitempty"`                          // Part of larger study
	CiteAs             *string                        `json:"citeAs,omitempty" bson:"cite_as,omitempty"`                          // How to cite this ResearchStudy
	RelatesTo          []ResearchStudyRelatesTo       `json:"relatesTo,omitempty" bson:"relates_to,omitempty"`                    // Relationships to other Resources
	Date               *string                        `json:"date,omitempty" bson:"date,omitempty"`                               // Date the resource last changed
	Status             string                         `json:"status" bson:"status"`                                               // draft | active | retired | unknown
	PrimaryPurposeType *CodeableConcept               `json:"primaryPurposeType,omitempty" bson:"primary_purpose_type,omitempty"` // treatment | prevention | diagnostic | supportive-care | screening | health-services-research | basic-science | device-feasibility
	Phase              *CodeableConcept               `json:"phase,omitempty" bson:"phase,omitempty"`                             // Classifier used for clinical trials
	StudyDesign        []CodeableConcept              `json:"studyDesign,omitempty" bson:"study_design,omitempty"`                // Classifications of the study design characteristics
	Focus              []CodeableReference            `json:"focus,omitempty" bson:"focus,omitempty"`                             // Drugs, devices, etc. under study
	Condition          []CodeableConcept              `json:"condition,omitempty" bson:"condition,omitempty"`                     // Condition being studied
	Keyword            []CodeableConcept              `json:"keyword,omitempty" bson:"keyword,omitempty"`                         // Used to search for the study
	Region             []CodeableConcept              `json:"region,omitempty" bson:"region,omitempty"`                           // Geographic area for the study
	DescriptionSummary *string                        `json:"descriptionSummary,omitempty" bson:"description_summary,omitempty"`  // Brief text explaining the study
	Description        *string                        `json:"description,omitempty" bson:"description,omitempty"`                 // Detailed narrative of the study
	Period             *Period                        `json:"period,omitempty" bson:"period,omitempty"`                           // When the study began and ended
	Site               []Reference                    `json:"site,omitempty" bson:"site,omitempty"`                               // Facility where study activities are conducted
	Note               []Annotation                   `json:"note,omitempty" bson:"note,omitempty"`                               // Comments made about the study
	Classifier         []CodeableConcept              `json:"classifier,omitempty" bson:"classifier,omitempty"`                   // Classification for the study
	AssociatedParty    []ResearchStudyAssociatedParty `json:"associatedParty,omitempty" bson:"associated_party,omitempty"`        // Sponsors, collaborators, and other parties
	ProgressStatus     []ResearchStudyProgressStatus  `json:"progressStatus,omitempty" bson:"progress_status,omitempty"`          // Status of study with time for that status
	WhyStopped         *CodeableConcept               `json:"whyStopped,omitempty" bson:"why_stopped,omitempty"`                  // accrual-goal-met | closed-due-to-toxicity | closed-due-to-lack-of-study-progress | temporarily-closed-per-study-design
	Recruitment        *ResearchStudyRecruitment      `json:"recruitment,omitempty" bson:"recruitment,omitempty"`                 // Target or actual group of participants enrolled in study
	ComparisonGroup    []ResearchStudyComparisonGroup `json:"comparisonGroup,omitempty" bson:"comparison_group,omitempty"`        // Defined path through the study for a subject
	Objective          []ResearchStudyObjective       `json:"objective,omitempty" bson:"objective,omitempty"`                     // A goal for the study
	Result             []Reference                    `json:"result,omitempty" bson:"result,omitempty"`                           // Link to results generated during the study
}

func (r *ResearchStudy) Validate() error {
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
	for i, item := range r.Label {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Label[%d]: %w", i, err)
		}
	}
	for i, item := range r.Protocol {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Protocol[%d]: %w", i, err)
		}
	}
	for i, item := range r.PartOf {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("PartOf[%d]: %w", i, err)
		}
	}
	for i, item := range r.RelatesTo {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("RelatesTo[%d]: %w", i, err)
		}
	}
	var emptyString string
	if r.Status == emptyString {
		return fmt.Errorf("field 'Status' is required")
	}
	if r.PrimaryPurposeType != nil {
		if err := r.PrimaryPurposeType.Validate(); err != nil {
			return fmt.Errorf("PrimaryPurposeType: %w", err)
		}
	}
	if r.Phase != nil {
		if err := r.Phase.Validate(); err != nil {
			return fmt.Errorf("Phase: %w", err)
		}
	}
	for i, item := range r.StudyDesign {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("StudyDesign[%d]: %w", i, err)
		}
	}
	for i, item := range r.Focus {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Focus[%d]: %w", i, err)
		}
	}
	for i, item := range r.Condition {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Condition[%d]: %w", i, err)
		}
	}
	for i, item := range r.Keyword {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Keyword[%d]: %w", i, err)
		}
	}
	for i, item := range r.Region {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Region[%d]: %w", i, err)
		}
	}
	if r.Period != nil {
		if err := r.Period.Validate(); err != nil {
			return fmt.Errorf("Period: %w", err)
		}
	}
	for i, item := range r.Site {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Site[%d]: %w", i, err)
		}
	}
	for i, item := range r.Note {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Note[%d]: %w", i, err)
		}
	}
	for i, item := range r.Classifier {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Classifier[%d]: %w", i, err)
		}
	}
	for i, item := range r.AssociatedParty {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("AssociatedParty[%d]: %w", i, err)
		}
	}
	for i, item := range r.ProgressStatus {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("ProgressStatus[%d]: %w", i, err)
		}
	}
	if r.WhyStopped != nil {
		if err := r.WhyStopped.Validate(); err != nil {
			return fmt.Errorf("WhyStopped: %w", err)
		}
	}
	if r.Recruitment != nil {
		if err := r.Recruitment.Validate(); err != nil {
			return fmt.Errorf("Recruitment: %w", err)
		}
	}
	for i, item := range r.ComparisonGroup {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("ComparisonGroup[%d]: %w", i, err)
		}
	}
	for i, item := range r.Objective {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Objective[%d]: %w", i, err)
		}
	}
	for i, item := range r.Result {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Result[%d]: %w", i, err)
		}
	}
	return nil
}

type ResearchStudyAssociatedParty struct {
	Id         *string           `json:"id,omitempty" bson:"id,omitempty"`                 // Unique id for inter-element referencing
	Name       *string           `json:"name,omitempty" bson:"name,omitempty"`             // Name of associated party
	Role       *CodeableConcept  `json:"role" bson:"role"`                                 // sponsor | lead-sponsor | sponsor-investigator | primary-investigator | collaborator | funding-source | general-contact | recruitment-contact | sub-investigator | study-chair | irb | data-monitoring
	Period     []Period          `json:"period,omitempty" bson:"period,omitempty"`         // When active in the role
	Classifier []CodeableConcept `json:"classifier,omitempty" bson:"classifier,omitempty"` // nih | fda | government | nonprofit | academic | industry
	Party      *Reference        `json:"party,omitempty" bson:"party,omitempty"`           // Individual or organization associated with study (use practitionerRole to specify their organisation)
}

func (r *ResearchStudyAssociatedParty) Validate() error {
	if r.Role == nil {
		return fmt.Errorf("field 'Role' is required")
	}
	if r.Role != nil {
		if err := r.Role.Validate(); err != nil {
			return fmt.Errorf("Role: %w", err)
		}
	}
	for i, item := range r.Period {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Period[%d]: %w", i, err)
		}
	}
	for i, item := range r.Classifier {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Classifier[%d]: %w", i, err)
		}
	}
	if r.Party != nil {
		if err := r.Party.Validate(); err != nil {
			return fmt.Errorf("Party: %w", err)
		}
	}
	return nil
}

type ResearchStudyComparisonGroup struct {
	Id            *string    `json:"id,omitempty" bson:"id,omitempty"`                        // Unique id for inter-element referencing
	TargetNumber  *int       `json:"targetNumber,omitempty" bson:"target_number,omitempty"`   // Estimated total number of participants to be enrolled in the comparison group
	ActualNumber  *int       `json:"actualNumber,omitempty" bson:"actual_number,omitempty"`   // Actual total number of participants enrolled in the comparison group
	Eligibility   *Reference `json:"eligibility,omitempty" bson:"eligibility,omitempty"`      // Inclusion and exclusion criteria for the comparison group
	ObservedGroup *Reference `json:"observedGroup,omitempty" bson:"observed_group,omitempty"` // Group of participants who were enrolled in the comparison group
	Description   *string    `json:"description,omitempty" bson:"description,omitempty"`      // Description of the comparison Group
}

func (r *ResearchStudyComparisonGroup) Validate() error {
	if r.Eligibility != nil {
		if err := r.Eligibility.Validate(); err != nil {
			return fmt.Errorf("Eligibility: %w", err)
		}
	}
	if r.ObservedGroup != nil {
		if err := r.ObservedGroup.Validate(); err != nil {
			return fmt.Errorf("ObservedGroup: %w", err)
		}
	}
	return nil
}

type ResearchStudyObjective struct {
	Id             *string                                `json:"id,omitempty" bson:"id,omitempty"`                          // Unique id for inter-element referencing
	Name           *string                                `json:"name,omitempty" bson:"name,omitempty"`                      // Label for the objective
	Type           *CodeableConcept                       `json:"type,omitempty" bson:"type,omitempty"`                      // primary | secondary | exploratory
	Description    *string                                `json:"description,omitempty" bson:"description,omitempty"`        // Description of the objective
	OutcomeMeasure []ResearchStudyObjectiveOutcomeMeasure `json:"outcomeMeasure,omitempty" bson:"outcome_measure,omitempty"` // A variable measured during the study
}

func (r *ResearchStudyObjective) Validate() error {
	if r.Type != nil {
		if err := r.Type.Validate(); err != nil {
			return fmt.Errorf("Type: %w", err)
		}
	}
	for i, item := range r.OutcomeMeasure {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("OutcomeMeasure[%d]: %w", i, err)
		}
	}
	return nil
}

type ResearchStudyObjectiveOutcomeMeasure struct {
	Id                   *string                                             `json:"id,omitempty" bson:"id,omitempty"`                                       // Unique id for inter-element referencing
	Name                 *string                                             `json:"name,omitempty" bson:"name,omitempty"`                                   // Label for the outcome measure
	Type                 *CodeableConcept                                    `json:"type,omitempty" bson:"type,omitempty"`                                   // primary | secondary | exploratory
	Description          *string                                             `json:"description,omitempty" bson:"description,omitempty"`                     // Description of the outcome measure
	Endpoint             *Reference                                          `json:"endpoint" bson:"endpoint"`                                               // Definition of the outcome measure
	Population           *Reference                                          `json:"population,omitempty" bson:"population,omitempty"`                       // Population for this estimand
	Intervention         *Reference                                          `json:"intervention,omitempty" bson:"intervention,omitempty"`                   // Comparison group of interest
	Comparator           *Reference                                          `json:"comparator,omitempty" bson:"comparator,omitempty"`                       // Comparison group for comparison
	SummaryMeasure       *CodeableConcept                                    `json:"summaryMeasure,omitempty" bson:"summary_measure,omitempty"`              // Statistical measure for treatment effect estimate
	EndpointAnalysisPlan *Reference                                          `json:"endpointAnalysisPlan,omitempty" bson:"endpoint_analysis_plan,omitempty"` // Statistical analysis plan for a single endpoint
	EventHandling        []ResearchStudyObjectiveOutcomeMeasureEventHandling `json:"eventHandling,omitempty" bson:"event_handling,omitempty"`                // Handling of intercurrent event
}

func (r *ResearchStudyObjectiveOutcomeMeasure) Validate() error {
	if r.Type != nil {
		if err := r.Type.Validate(); err != nil {
			return fmt.Errorf("Type: %w", err)
		}
	}
	if r.Endpoint == nil {
		return fmt.Errorf("field 'Endpoint' is required")
	}
	if r.Endpoint != nil {
		if err := r.Endpoint.Validate(); err != nil {
			return fmt.Errorf("Endpoint: %w", err)
		}
	}
	if r.Population != nil {
		if err := r.Population.Validate(); err != nil {
			return fmt.Errorf("Population: %w", err)
		}
	}
	if r.Intervention != nil {
		if err := r.Intervention.Validate(); err != nil {
			return fmt.Errorf("Intervention: %w", err)
		}
	}
	if r.Comparator != nil {
		if err := r.Comparator.Validate(); err != nil {
			return fmt.Errorf("Comparator: %w", err)
		}
	}
	if r.SummaryMeasure != nil {
		if err := r.SummaryMeasure.Validate(); err != nil {
			return fmt.Errorf("SummaryMeasure: %w", err)
		}
	}
	if r.EndpointAnalysisPlan != nil {
		if err := r.EndpointAnalysisPlan.Validate(); err != nil {
			return fmt.Errorf("EndpointAnalysisPlan: %w", err)
		}
	}
	for i, item := range r.EventHandling {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("EventHandling[%d]: %w", i, err)
		}
	}
	return nil
}

type ResearchStudyProgressStatus struct {
	Id     *string          `json:"id,omitempty" bson:"id,omitempty"`         // Unique id for inter-element referencing
	State  *CodeableConcept `json:"state" bson:"state"`                       // Label for status or state (e.g. recruitment status)
	Actual bool             `json:"actual,omitempty" bson:"actual,omitempty"` // Actual if true else anticipated
	Period *Period          `json:"period,omitempty" bson:"period,omitempty"` // Date range
}

func (r *ResearchStudyProgressStatus) Validate() error {
	if r.State == nil {
		return fmt.Errorf("field 'State' is required")
	}
	if r.State != nil {
		if err := r.State.Validate(); err != nil {
			return fmt.Errorf("State: %w", err)
		}
	}
	if r.Period != nil {
		if err := r.Period.Validate(); err != nil {
			return fmt.Errorf("Period: %w", err)
		}
	}
	return nil
}

type ResearchStudyRecruitment struct {
	Id           *string    `json:"id,omitempty" bson:"id,omitempty"`                      // Unique id for inter-element referencing
	TargetNumber *int       `json:"targetNumber,omitempty" bson:"target_number,omitempty"` // Estimated total number of participants to be enrolled
	ActualNumber *int       `json:"actualNumber,omitempty" bson:"actual_number,omitempty"` // Actual total number of participants enrolled in study
	Eligibility  *Reference `json:"eligibility,omitempty" bson:"eligibility,omitempty"`    // Inclusion and exclusion criteria
	ActualGroup  *Reference `json:"actualGroup,omitempty" bson:"actual_group,omitempty"`   // Group of participants who were enrolled in study
	Description  *string    `json:"description,omitempty" bson:"description,omitempty"`    // Description of the recruitment
}

func (r *ResearchStudyRecruitment) Validate() error {
	if r.Eligibility != nil {
		if err := r.Eligibility.Validate(); err != nil {
			return fmt.Errorf("Eligibility: %w", err)
		}
	}
	if r.ActualGroup != nil {
		if err := r.ActualGroup.Validate(); err != nil {
			return fmt.Errorf("ActualGroup: %w", err)
		}
	}
	return nil
}

type ResearchStudyObjectiveOutcomeMeasureEventHandling struct {
	Id          *string          `json:"id,omitempty" bson:"id,omitempty"`                   // Unique id for inter-element referencing
	Event       *CodeableConcept `json:"event,omitempty" bson:"event,omitempty"`             // The event
	Group       *CodeableConcept `json:"group,omitempty" bson:"group,omitempty"`             // The group that is affected by this event handling
	Handling    *CodeableConcept `json:"handling,omitempty" bson:"handling,omitempty"`       // How the data is handled
	Description *string          `json:"description,omitempty" bson:"description,omitempty"` // Text summary of event handling
}

func (r *ResearchStudyObjectiveOutcomeMeasureEventHandling) Validate() error {
	if r.Event != nil {
		if err := r.Event.Validate(); err != nil {
			return fmt.Errorf("Event: %w", err)
		}
	}
	if r.Group != nil {
		if err := r.Group.Validate(); err != nil {
			return fmt.Errorf("Group: %w", err)
		}
	}
	if r.Handling != nil {
		if err := r.Handling.Validate(); err != nil {
			return fmt.Errorf("Handling: %w", err)
		}
	}
	return nil
}

type ResearchStudyLabel struct {
	Id       *string          `json:"id,omitempty" bson:"id,omitempty"`             // Unique id for inter-element referencing
	Type     *CodeableConcept `json:"type,omitempty" bson:"type,omitempty"`         // primary | official | scientific | plain-language | subtitle | short-title | acronym | earlier-title | language | auto-translated | human-use | machine-use | duplicate-uid
	Value    *string          `json:"value,omitempty" bson:"value,omitempty"`       // The name
	Language *string          `json:"language,omitempty" bson:"language,omitempty"` // Used to express the specific language
}

func (r *ResearchStudyLabel) Validate() error {
	if r.Type != nil {
		if err := r.Type.Validate(); err != nil {
			return fmt.Errorf("Type: %w", err)
		}
	}
	return nil
}

type ResearchStudyRelatesTo struct {
	Id               *string          `json:"id,omitempty" bson:"id,omitempty"`          // Unique id for inter-element referencing
	Type             *CodeableConcept `json:"type" bson:"type"`                          // documentation | justification | citation | predecessor | successor | derived-from | depends-on | composed-of | part-of | amends | amended-with | appends | appended-with | cites | cited-by | comments-on | comment-in | contains | contained-in | corrects | correction-in | replaces | replaced-with | retracts | retracted-by | signs | similar-to | supports | supported-with | transforms | transformed-into | transformed-with | documents | specification-of | created-with | cite-as | reprint | reprint-of | summarizes
	TargetUri        *string          `json:"targetUri" bson:"target_uri"`               // The artifact that is related to this ResearchStudy
	TargetAttachment *Attachment      `json:"targetAttachment" bson:"target_attachment"` // The artifact that is related to this ResearchStudy
	TargetCanonical  *string          `json:"targetCanonical" bson:"target_canonical"`   // The artifact that is related to this ResearchStudy
	TargetReference  *Reference       `json:"targetReference" bson:"target_reference"`   // The artifact that is related to this ResearchStudy
	TargetMarkdown   *string          `json:"targetMarkdown" bson:"target_markdown"`     // The artifact that is related to this ResearchStudy
}

func (r *ResearchStudyRelatesTo) Validate() error {
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
