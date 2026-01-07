package models

import (
	"encoding/json"
	"fmt"
)

// Representation of the content produced in a DICOM imaging study. A study comprises a set of series, each of which includes a set of images or other data objects (called Service-Object Pair Instances or SOP Instances) acquired or produced in a common context.
type ImagingStudy struct {
	Id                *string              `json:"id,omitempty" bson:"id,omitempty"`                                 // Logical id of this artifact
	Meta              *Meta                `json:"meta,omitempty" bson:"meta,omitempty"`                             // Metadata about the resource
	ImplicitRules     *string              `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`          // A set of rules under which this content was created
	Language          *string              `json:"language,omitempty" bson:"language,omitempty"`                     // Language of the resource content
	Text              *Narrative           `json:"text,omitempty" bson:"text,omitempty"`                             // Text summary of the resource, for human interpretation
	Contained         []json.RawMessage    `json:"contained,omitempty" bson:"contained,omitempty"`                   // Contained, inline Resources
	Identifier        []Identifier         `json:"identifier,omitempty" bson:"identifier,omitempty"`                 // Business identifier for imaging study
	Status            string               `json:"status" bson:"status"`                                             // registered | available | cancelled | entered-in-error | unknown | inactive
	Modality          []CodeableConcept    `json:"modality,omitempty" bson:"modality,omitempty"`                     // The distinct values for series' modalities
	Subject           *Reference           `json:"subject" bson:"subject"`                                           // Who or what is the subject of the study
	Encounter         *Reference           `json:"encounter,omitempty" bson:"encounter,omitempty"`                   // Encounter with which this imaging study is associated
	Started           *string              `json:"started,omitempty" bson:"started,omitempty"`                       // When the study was started
	BasedOn           []Reference          `json:"basedOn,omitempty" bson:"based_on,omitempty"`                      // Fulfills plan or order
	Procedure         []Reference          `json:"procedure,omitempty" bson:"procedure,omitempty"`                   // Imaging performed procedure(s)
	Referrer          *Reference           `json:"referrer,omitempty" bson:"referrer,omitempty"`                     // Referring physician
	Endpoint          []Reference          `json:"endpoint,omitempty" bson:"endpoint,omitempty"`                     // Study access endpoint
	Location          *Reference           `json:"location,omitempty" bson:"location,omitempty"`                     // Where imaging study occurred
	Reason            []CodeableReference  `json:"reason,omitempty" bson:"reason,omitempty"`                         // Why was imaging study performed?
	Note              []Annotation         `json:"note,omitempty" bson:"note,omitempty"`                             // Comments made about the imaging study
	Description       *string              `json:"description,omitempty" bson:"description,omitempty"`               // Study Description or Classification
	NumberOfSeries    *int                 `json:"numberOfSeries,omitempty" bson:"number_of_series,omitempty"`       // Number of Study Related Series
	NumberOfInstances *int                 `json:"numberOfInstances,omitempty" bson:"number_of_instances,omitempty"` // Number of Study Related Instances
	Series            []ImagingStudySeries `json:"series,omitempty" bson:"series,omitempty"`                         // The set of Series belonging to the study
}

func (r *ImagingStudy) Validate() error {
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
	for i, item := range r.Modality {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Modality[%d]: %w", i, err)
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
	for i, item := range r.BasedOn {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("BasedOn[%d]: %w", i, err)
		}
	}
	for i, item := range r.Procedure {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Procedure[%d]: %w", i, err)
		}
	}
	if r.Referrer != nil {
		if err := r.Referrer.Validate(); err != nil {
			return fmt.Errorf("Referrer: %w", err)
		}
	}
	for i, item := range r.Endpoint {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Endpoint[%d]: %w", i, err)
		}
	}
	if r.Location != nil {
		if err := r.Location.Validate(); err != nil {
			return fmt.Errorf("Location: %w", err)
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
	for i, item := range r.Series {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Series[%d]: %w", i, err)
		}
	}
	return nil
}

type ImagingStudySeries struct {
	Id                *string                       `json:"id,omitempty" bson:"id,omitempty"`                                 // Unique id for inter-element referencing
	Uid               string                        `json:"uid" bson:"uid"`                                                   // DICOM Series Instance UID
	Number            *int                          `json:"number,omitempty" bson:"number,omitempty"`                         // Numeric identifier of this series
	Modality          *CodeableConcept              `json:"modality" bson:"modality"`                                         // The modality used for this series
	Description       *string                       `json:"description,omitempty" bson:"description,omitempty"`               // Series Description or Classification
	NumberOfInstances *int                          `json:"numberOfInstances,omitempty" bson:"number_of_instances,omitempty"` // Number of Series Related Instances
	Endpoint          []Reference                   `json:"endpoint,omitempty" bson:"endpoint,omitempty"`                     // Series access endpoint
	BodySite          *CodeableReference            `json:"bodySite,omitempty" bson:"body_site,omitempty"`                    // Body part examined
	Specimen          []Reference                   `json:"specimen,omitempty" bson:"specimen,omitempty"`                     // Specimen imaged
	Started           *string                       `json:"started,omitempty" bson:"started,omitempty"`                       // When the series started
	Performer         []ImagingStudySeriesPerformer `json:"performer,omitempty" bson:"performer,omitempty"`                   // Who performed the series
	Instance          []ImagingStudySeriesInstance  `json:"instance,omitempty" bson:"instance,omitempty"`                     // A single SOP instance from the series
}

func (r *ImagingStudySeries) Validate() error {
	var emptyString string
	if r.Uid == emptyString {
		return fmt.Errorf("field 'Uid' is required")
	}
	if r.Modality == nil {
		return fmt.Errorf("field 'Modality' is required")
	}
	if r.Modality != nil {
		if err := r.Modality.Validate(); err != nil {
			return fmt.Errorf("Modality: %w", err)
		}
	}
	for i, item := range r.Endpoint {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Endpoint[%d]: %w", i, err)
		}
	}
	if r.BodySite != nil {
		if err := r.BodySite.Validate(); err != nil {
			return fmt.Errorf("BodySite: %w", err)
		}
	}
	for i, item := range r.Specimen {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Specimen[%d]: %w", i, err)
		}
	}
	for i, item := range r.Performer {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Performer[%d]: %w", i, err)
		}
	}
	for i, item := range r.Instance {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Instance[%d]: %w", i, err)
		}
	}
	return nil
}

type ImagingStudySeriesPerformer struct {
	Id       *string          `json:"id,omitempty" bson:"id,omitempty"`             // Unique id for inter-element referencing
	Function *CodeableConcept `json:"function,omitempty" bson:"function,omitempty"` // Type of performance
	Actor    *Reference       `json:"actor" bson:"actor"`                           // Who performed imaging study
}

func (r *ImagingStudySeriesPerformer) Validate() error {
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

type ImagingStudySeriesInstance struct {
	Id       *string `json:"id,omitempty" bson:"id,omitempty"`         // Unique id for inter-element referencing
	Uid      string  `json:"uid" bson:"uid"`                           // DICOM SOP Instance UID
	SopClass string  `json:"sopClass" bson:"sop_class"`                // DICOM class type
	Number   *int    `json:"number,omitempty" bson:"number,omitempty"` // The number of this instance in the series
	Title    *string `json:"title,omitempty" bson:"title,omitempty"`   // Name or title of the instance
}

func (r *ImagingStudySeriesInstance) Validate() error {
	var emptyString string
	if r.Uid == emptyString {
		return fmt.Errorf("field 'Uid' is required")
	}
	if r.SopClass == emptyString {
		return fmt.Errorf("field 'SopClass' is required")
	}
	return nil
}
