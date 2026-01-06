package models

// This operation takes a resource in one form, and returns to in another form. Both the 'resource' and 'return' parameters are a single resource. The primary use of this operation is to convert between formats (e.g. (XML -> JSON or vice versa)
type Convert struct {
}

func (r *Convert) Validate() error {
	return nil
}
