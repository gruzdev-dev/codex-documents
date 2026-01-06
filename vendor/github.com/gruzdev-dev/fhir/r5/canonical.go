package models

// canonical type: A URI that is a reference to a canonical URL on a FHIR resource
type canonical struct {
}

func (r *canonical) Validate() error {
	return nil
}

type Canonical struct {
	Id    *string `json:"id,omitempty" bson:"id,omitempty"`       // xml:id (or equivalent in JSON)
	Value *string `json:"value,omitempty" bson:"value,omitempty"` // Primitive value for canonical
}

func (r *Canonical) Validate() error {
	return nil
}
