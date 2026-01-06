package models

// Test the subsumption relationship between code/Coding A and code/Coding B given the semantics of subsumption in the underlying code system (see [hierarchyMeaning](codesystem-definitions.html#CodeSystem.hierarchyMeaning)).  When invoking this operation, a client SHALL provide both A and B codes, either as code or Coding parameters. The system parameter is required unless the operation is invoked on an instance of a code system resource. Other parameters are optional
type Subsumes struct {
}

func (r *Subsumes) Validate() error {
	return nil
}
