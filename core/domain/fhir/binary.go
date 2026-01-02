package fhir

type Binary struct {
	ContentType     *Element
	Data            *Element
	ImplicitRules   *Element
	Language        *Element
	Meta            *Meta
	SecurityContext *Reference
}
