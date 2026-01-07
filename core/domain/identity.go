package domain

import "slices"

type Identity struct {
	UserID    string
	PatientID string
	Scopes    []string
}

func (i *Identity) HasScope(scope string) bool {
	return slices.Contains(i.Scopes, scope)
}

func (i *Identity) IsPatient(id string) bool {
	return i.PatientID == id
}