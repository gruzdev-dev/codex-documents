package models

// Validate that a coded value is in the set of codes allowed by a value set.  If the operation is not called at the instance level, one of the in parameters url, context or valueSet must be provided.  One (and only one) of the in parameters code, coding, or codeableConcept must be provided. If a code is provided, either a system or inferSystem **SHOULD** be provided. The operation returns a result (true / false), an error message, and the recommended display for the code. When validating a code or a coding, then the code, system and version output parameters **SHOULD** be populated when possible. When a validating a CodeableConcept, then the codeableConcept output parameter **SHOULD** be populated when possible.
type ValidateCode struct {
}

func (r *ValidateCode) Validate() error {
	return nil
}
