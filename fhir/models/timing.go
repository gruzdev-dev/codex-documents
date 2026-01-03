package models

// Timing Type: Specifies an event that may occur multiple times. Timing schedules are used to record when things are planned, expected or requested to occur. The most common usage is in dosage instructions for medications. They are also used when planning care of various kinds, and may be used for reporting the schedule to which past regular activities were carried out.
type Timing struct {
	Id     *string          `json:"id,omitempty"`     // Unique id for inter-element referencing
	Event  []string         `json:"event,omitempty"`  // When the event occurs
	Repeat *TimingRepeat    `json:"repeat,omitempty"` // When the event is to occur
	Code   *CodeableConcept `json:"code,omitempty"`   // C | BID | TID | QID | AM | PM | QD | QOD | +
}

type TimingRepeat struct {
	Id             *string   `json:"id,omitempty"`             // Unique id for inter-element referencing
	BoundsDuration *Duration `json:"boundsDuration,omitempty"` // Length/Range of lengths, or (Start and/or end) limits
	BoundsRange    *Range    `json:"boundsRange,omitempty"`    // Length/Range of lengths, or (Start and/or end) limits
	BoundsPeriod   *Period   `json:"boundsPeriod,omitempty"`   // Length/Range of lengths, or (Start and/or end) limits
	Count          *int      `json:"count,omitempty"`          // Number of times to repeat
	CountMax       *int      `json:"countMax,omitempty"`       // Maximum number of times to repeat
	Duration       *float64  `json:"duration,omitempty"`       // How long when it happens
	DurationMax    *float64  `json:"durationMax,omitempty"`    // How long when it happens (Max)
	DurationUnit   *string   `json:"durationUnit,omitempty"`   // s | min | h | d | wk | mo | a - unit of time (UCUM)
	Frequency      *int      `json:"frequency,omitempty"`      // Indicates the number of repetitions that should occur within a period. I.e. Event occurs frequency times per period
	FrequencyMax   *int      `json:"frequencyMax,omitempty"`   // Event occurs up to frequencyMax times per period
	Period         *float64  `json:"period,omitempty"`         // The duration to which the frequency applies. I.e. Event occurs frequency times per period
	PeriodMax      *float64  `json:"periodMax,omitempty"`      // Upper limit of period (3-4 hours)
	PeriodUnit     *string   `json:"periodUnit,omitempty"`     // s | min | h | d | wk | mo | a - unit of time (UCUM)
	StartOffset    *Quantity `json:"startOffset,omitempty"`    // Events within the repeat period do not start until startOffset has elapsed
	EndOffset      *Quantity `json:"endOffset,omitempty"`      // Events within the repeat period step once endOffset before the end of the period
	DayOfWeek      []string  `json:"dayOfWeek,omitempty"`      // mon | tue | wed | thu | fri | sat | sun
	TimeOfDay      []string  `json:"timeOfDay,omitempty"`      // Time of day for action
	When           []string  `json:"when,omitempty"`           // Code for time period of occurrence
	Offset         *int      `json:"offset,omitempty"`         // Minutes from event (before or after)
}
