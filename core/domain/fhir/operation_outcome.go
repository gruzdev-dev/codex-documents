package fhir

type OperationOutcome struct {
	ImplicitRules     *Element
	Language          *Element
	Extension         []Extension
	Meta              *Meta
	ModifierExtension []Extension
}
