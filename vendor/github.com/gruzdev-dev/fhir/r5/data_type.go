package models

// DataType Type: The base class for all re-useable types defined as part of the FHIR Specification.
type DataType struct {
	Id *string `json:"id,omitempty" bson:"id,omitempty"` // Unique id for inter-element referencing
}

func (r *DataType) Validate() error {
	return nil
}
