package fhir

type Binary struct {
	ContentType     *Element
	Data            *Element
	ImplicitRules   *Element
	Language        *Element
	ContentType_2   *Code
	Data_2          *Base64Binary
	ID              *ID
	ImplicitRules_2 *Uri
	Language_2      *Code
	Meta            *Meta
	ResourceType    interface{}
	SecurityContext *Reference
}
