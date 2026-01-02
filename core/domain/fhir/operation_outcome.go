package fhir

type OperationOutcome struct {
	ImplicitRules     *Element
	Language          *Element
	Contained         []OperationOutcomeContainedElem
	Extension         []Extension
	ID                *ID
	ImplicitRules_2   *Uri
	Issue             []OperationOutcomeIssue
	Language_2        *Code
	Meta              *Meta
	ModifierExtension []Extension
	ResourceType      interface{}
	Text              *Narrative
}
