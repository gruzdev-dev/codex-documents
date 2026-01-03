package models

// SampledData Type: A series of measurements taken by a device, with upper and lower limits. There may be more than one dimension in the data.
type SampledData struct {
	Id           *string  `json:"id,omitempty"`           // Unique id for inter-element referencing
	Origin       Quantity `json:"origin,omitempty"`       // Zero value and units
	Interval     *float64 `json:"interval,omitempty"`     // Number of intervalUnits between samples
	IntervalUnit string   `json:"intervalUnit,omitempty"` // The measurement unit of the interval between samples
	Factor       *float64 `json:"factor,omitempty"`       // Multiply data by this before adding to origin
	LowerLimit   *float64 `json:"lowerLimit,omitempty"`   // Lower limit of detection
	UpperLimit   *float64 `json:"upperLimit,omitempty"`   // Upper limit of detection
	Dimensions   int      `json:"dimensions,omitempty"`   // Number of sample points at each time point
	CodeMap      *string  `json:"codeMap,omitempty"`      // Defines the codes used in the data
	Offsets      *string  `json:"offsets,omitempty"`      // Offsets, typically in time, at which data values were taken
	Data         *string  `json:"data,omitempty"`         // Decimal values with spaces, or "E" | "U" | "L", or another code
}
