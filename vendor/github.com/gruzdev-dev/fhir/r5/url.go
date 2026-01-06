package models

// url type: A URI that is a literal reference
type url struct {
}

func (r *url) Validate() error {
	return nil
}

type Url struct {
	Id    *string `json:"id,omitempty" bson:"id,omitempty"`       // xml:id (or equivalent in JSON)
	Value *string `json:"value,omitempty" bson:"value,omitempty"` // Primitive value for url
}

func (r *Url) Validate() error {
	return nil
}
