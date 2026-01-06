package models

// The definition of a value set is used to create a simple collection of codes suitable for use for data entry or validation.   If the operation is not called at the instance level, one of the in parameters url, context or valueSet must be provided.  An expanded value set will be returned, or an OperationOutcome with an error message.
type Expand struct {
}

func (r *Expand) Validate() error {
	return nil
}
