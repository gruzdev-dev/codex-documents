package models

// oid type: An OID represented as a URI
type FHIRString struct {
	Id    *string `json:"id,omitempty" bson:"id,omitempty"`       // xml:id (or equivalent in JSON)
	Value *string `json:"value,omitempty" bson:"value,omitempty"` // Primitive value for oid
}

func (r *FHIRString) Validate() error {
	return nil
}
