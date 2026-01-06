package models

// uri Type: String of characters used to identify a name or a resource
type uri struct {
}

func (r *uri) Validate() error {
	return nil
}

type Uri struct {
	Id    *string `json:"id,omitempty" bson:"id,omitempty"`       // xml:id (or equivalent in JSON)
	Value *string `json:"value,omitempty" bson:"value,omitempty"` // Primitive value for uri
}

func (r *Uri) Validate() error {
	return nil
}
