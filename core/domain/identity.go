package domain

import (
	"fmt"
	"slices"
)

type Identity struct {
	UserID    string
	PatientID string
	Scopes    []string
}

func (i *Identity) HasScope(scope string) bool {
	if i.Scopes == nil {
		return false
	}
	return slices.Contains(i.Scopes, scope)
}

func (i *Identity) IsPatient(id string) bool {
	return i.PatientID == id
}

func (i *Identity) HasResourceScope(service, resource, id, action string) bool {
	if i.Scopes == nil {
		return false
	}
	scope := fmt.Sprintf("%s:%s:%s:%s", service, resource, id, action)
	return slices.Contains(i.Scopes, scope)
}

func (i *Identity) IsTmpToken() bool {
	return i.UserID == "" && i.PatientID == ""
}
