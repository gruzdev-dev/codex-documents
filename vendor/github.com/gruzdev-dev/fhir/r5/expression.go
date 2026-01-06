package models

// Expression Type: A expression that is evaluated in a specified context and returns a value. The context of use of the expression must specify the context in which the expression is evaluated, and how the result of the expression is used.
type Expression struct {
	Id          *string `json:"id,omitempty" bson:"id,omitempty"`                   // Unique id for inter-element referencing
	Description *string `json:"description,omitempty" bson:"description,omitempty"` // Natural language description of the condition
	Name        *string `json:"name,omitempty" bson:"name,omitempty"`               // Short name assigned to expression for reuse
	Language    *string `json:"language,omitempty" bson:"language,omitempty"`       // text/cql | text/fhirpath | application/x-fhir-query | etc.
	Expression  *string `json:"expression,omitempty" bson:"expression,omitempty"`   // Expression in specified language
	Reference   *string `json:"reference,omitempty" bson:"reference,omitempty"`     // Where the expression is found
}

func (r *Expression) Validate() error {
	return nil
}
