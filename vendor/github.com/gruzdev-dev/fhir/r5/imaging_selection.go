package models

import (
	"encoding/json"
	"fmt"
)

// A selection of DICOM SOP instances within a single Study and Series. This might include additional specifics such as a set of frames or an image region, allowing linkage to an Observation Resource.
type ImagingSelection struct {
	ResourceType        string                          `json:"resourceType" bson:"resource_type"`                                     // Type of resource
	Id                  *string                         `json:"id,omitempty" bson:"id,omitempty"`                                      // Logical id of this artifact
	Meta                *Meta                           `json:"meta,omitempty" bson:"meta,omitempty"`                                  // Metadata about the resource
	ImplicitRules       *string                         `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`               // A set of rules under which this content was created
	Language            *string                         `json:"language,omitempty" bson:"language,omitempty"`                          // Language of the resource content
	Text                *Narrative                      `json:"text,omitempty" bson:"text,omitempty"`                                  // Text summary of the resource, for human interpretation
	Contained           []json.RawMessage               `json:"contained,omitempty" bson:"contained,omitempty"`                        // Contained, inline Resources
	Identifier          []Identifier                    `json:"identifier,omitempty" bson:"identifier,omitempty"`                      // Business identifier for imaging selection
	Status              string                          `json:"status" bson:"status"`                                                  // available | entered-in-error | inactive | unknown
	Category            []CodeableConcept               `json:"category,omitempty" bson:"category,omitempty"`                          // Classifies the imaging selection
	Modality            *CodeableConcept                `json:"modality,omitempty" bson:"modality,omitempty"`                          // The distinct modality
	Code                *CodeableConcept                `json:"code" bson:"code"`                                                      // Imaging Selection purpose text or code
	Subject             *Reference                      `json:"subject,omitempty" bson:"subject,omitempty"`                            // Who or what is the subject of the imaging selection
	Issued              *string                         `json:"issued,omitempty" bson:"issued,omitempty"`                              // When the imaging selection was created
	Performer           []ImagingSelectionPerformer     `json:"performer,omitempty" bson:"performer,omitempty"`                        // Who performed imaging selection and what they did
	BasedOn             []Reference                     `json:"basedOn,omitempty" bson:"based_on,omitempty"`                           // Fulfills plan or order
	DerivedFrom         *Reference                      `json:"derivedFrom,omitempty" bson:"derived_from,omitempty"`                   // The imaging study from which the imaging selection is derived
	StudyUid            *string                         `json:"studyUid,omitempty" bson:"study_uid,omitempty"`                         // DICOM Study Instance UID
	SeriesUid           *string                         `json:"seriesUid,omitempty" bson:"series_uid,omitempty"`                       // DICOM Series Instance UID
	SeriesNumber        *int                            `json:"seriesNumber,omitempty" bson:"series_number,omitempty"`                 // Numeric identifier of the selected series
	FrameOfReferenceUid *string                         `json:"frameOfReferenceUid,omitempty" bson:"frame_of_reference_uid,omitempty"` // The Frame of Reference UID for the selected images
	BodySite            []CodeableReference             `json:"bodySite,omitempty" bson:"body_site,omitempty"`                         // Selected anatomic structure
	Focus               []Reference                     `json:"focus,omitempty" bson:"focus,omitempty"`                                // Related resources that are the focus for the imaging selection
	Endpoint            []Reference                     `json:"endpoint,omitempty" bson:"endpoint,omitempty"`                          // The network services providing access for the subset of the study
	Instance            []ImagingSelectionInstance      `json:"instance,omitempty" bson:"instance,omitempty"`                          // The selected instances
	ImageRegion3D       []ImagingSelectionImageRegion3D `json:"imageRegion3D,omitempty" bson:"image_region3_d,omitempty"`              // A 3D region in a DICOM frame of reference
}

func (r *ImagingSelection) Validate() error {
	if r.ResourceType != "ImagingSelection" {
		return fmt.Errorf("invalid resourceType: expected 'ImagingSelection', got '%s'", r.ResourceType)
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
	for i, item := range r.Category {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Category[%d]: %w", i, err)
		}
	}
	if r.Modality != nil {
		if err := r.Modality.Validate(); err != nil {
			return fmt.Errorf("Modality: %w", err)
		}
	}
	if r.Code == nil {
		return fmt.Errorf("field 'Code' is required")
	}
	if r.Code != nil {
		if err := r.Code.Validate(); err != nil {
			return fmt.Errorf("Code: %w", err)
		}
	}
	if r.Subject != nil {
		if err := r.Subject.Validate(); err != nil {
			return fmt.Errorf("Subject: %w", err)
		}
	}
	for i, item := range r.Performer {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Performer[%d]: %w", i, err)
		}
	}
	for i, item := range r.BasedOn {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("BasedOn[%d]: %w", i, err)
		}
	}
	if r.DerivedFrom != nil {
		if err := r.DerivedFrom.Validate(); err != nil {
			return fmt.Errorf("DerivedFrom: %w", err)
		}
	}
	for i, item := range r.BodySite {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("BodySite[%d]: %w", i, err)
		}
	}
	for i, item := range r.Focus {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Focus[%d]: %w", i, err)
		}
	}
	for i, item := range r.Endpoint {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Endpoint[%d]: %w", i, err)
		}
	}
	for i, item := range r.Instance {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Instance[%d]: %w", i, err)
		}
	}
	for i, item := range r.ImageRegion3D {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("ImageRegion3D[%d]: %w", i, err)
		}
	}
	return nil
}

type ImagingSelectionPerformer struct {
	Id       *string          `json:"id,omitempty" bson:"id,omitempty"`             // Unique id for inter-element referencing
	Function *CodeableConcept `json:"function,omitempty" bson:"function,omitempty"` // Type of performance
	Actor    *Reference       `json:"actor" bson:"actor"`                           // Who performed the imaging selection
}

func (r *ImagingSelectionPerformer) Validate() error {
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

type ImagingSelectionInstance struct {
	Id                              *string                                 `json:"id,omitempty" bson:"id,omitempty"`                                                              // Unique id for inter-element referencing
	Uid                             string                                  `json:"uid" bson:"uid"`                                                                                // DICOM SOP Instance UID
	Number                          *int                                    `json:"number,omitempty" bson:"number,omitempty"`                                                      // The number of this instance in the series
	SopClass                        *string                                 `json:"sopClass,omitempty" bson:"sop_class,omitempty"`                                                 // DICOM class type
	FrameNumber                     []int                                   `json:"frameNumber,omitempty" bson:"frame_number,omitempty"`                                           // Selected frames
	ReferencedContentItemIdentifier []int                                   `json:"referencedContentItemIdentifier,omitempty" bson:"referenced_content_item_identifier,omitempty"` // Selected content items
	SegmentNumber                   []int                                   `json:"segmentNumber,omitempty" bson:"segment_number,omitempty"`                                       // Selected segments
	RegionOfInterest                []int                                   `json:"regionOfInterest,omitempty" bson:"region_of_interest,omitempty"`                                // Selected regions of interest
	WaveFormChannel                 []int                                   `json:"waveFormChannel,omitempty" bson:"wave_form_channel,omitempty"`                                  // Selected waveform channel
	ImageRegion2D                   []ImagingSelectionInstanceImageRegion2D `json:"imageRegion2D,omitempty" bson:"image_region2_d,omitempty"`                                      // A 2D region in an image
}

func (r *ImagingSelectionInstance) Validate() error {
	var emptyString string
	if r.Uid == emptyString {
		return fmt.Errorf("field 'Uid' is required")
	}
	for i, item := range r.ImageRegion2D {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("ImageRegion2D[%d]: %w", i, err)
		}
	}
	return nil
}

type ImagingSelectionInstanceImageRegion2D struct {
	Id         *string   `json:"id,omitempty" bson:"id,omitempty"` // Unique id for inter-element referencing
	RegionType string    `json:"regionType" bson:"region_type"`    // point | polyline | multipoint | circle | ellipse
	Coordinate []float64 `json:"coordinate" bson:"coordinate"`     // The coordinates that define the image region
}

func (r *ImagingSelectionInstanceImageRegion2D) Validate() error {
	var emptyString string
	if r.RegionType == emptyString {
		return fmt.Errorf("field 'RegionType' is required")
	}
	if len(r.Coordinate) < 1 {
		return fmt.Errorf("field 'Coordinate' must have at least 1 elements")
	}
	return nil
}

type ImagingSelectionImageRegion3D struct {
	Id         *string   `json:"id,omitempty" bson:"id,omitempty"` // Unique id for inter-element referencing
	RegionType string    `json:"regionType" bson:"region_type"`    // point | multipoint | polyline | polygon | ellipse | ellipsoid
	Coordinate []float64 `json:"coordinate" bson:"coordinate"`     // Specifies the coordinates that define the image region
}

func (r *ImagingSelectionImageRegion3D) Validate() error {
	var emptyString string
	if r.RegionType == emptyString {
		return fmt.Errorf("field 'RegionType' is required")
	}
	if len(r.Coordinate) < 1 {
		return fmt.Errorf("field 'Coordinate' must have at least 1 elements")
	}
	return nil
}
