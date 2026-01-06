package models

// markdown type: A string that may contain Github Flavored Markdown syntax for optional processing by a mark down presentation engine
type markdown struct {
}

func (r *markdown) Validate() error {
	return nil
}

type Markdown struct {
	Id    *string `json:"id,omitempty" bson:"id,omitempty"`       // xml:id (or equivalent in JSON)
	Value *string `json:"value,omitempty" bson:"value,omitempty"` // Primitive value for markdown
}

func (r *Markdown) Validate() error {
	return nil
}
