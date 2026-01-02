package fhir

import (
	domain "codex-documents/core/domain/fhir"
	dto "codex-documents/fhir"
)

// goverter:converter
// goverter:output:file ./generated.go
// goverter:output:package codex-documents/adapters/fhir
// goverter:extend MapAny
type Converter interface {
	ToDomainOperationOutcome(source *dto.OperationOutcome) *domain.OperationOutcome
	FromDomainOperationOutcome(source *domain.OperationOutcome) *dto.OperationOutcome

	ToDomainEncounter(source *dto.Encounter) *domain.Encounter
	FromDomainEncounter(source *domain.Encounter) *dto.Encounter

	ToDomainPractitioner(source *dto.Practitioner) *domain.Practitioner
	FromDomainPractitioner(source *domain.Practitioner) *dto.Practitioner

	ToDomainPatient(source *dto.Patient) *domain.Patient
	FromDomainPatient(source *domain.Patient) *dto.Patient

	ToDomainDocumentReference(source *dto.DocumentReference) *domain.DocumentReference
	FromDomainDocumentReference(source *domain.DocumentReference) *dto.DocumentReference

	ToDomainBinary(source *dto.Binary) *domain.Binary
	FromDomainBinary(source *domain.Binary) *dto.Binary

	ToDomainObservation(source *dto.Observation) *domain.Observation
	FromDomainObservation(source *domain.Observation) *dto.Observation

	ToDomainBundle(source *dto.Bundle) *domain.Bundle
	FromDomainBundle(source *domain.Bundle) *dto.Bundle
}

func MapAny(v interface{}) interface{} { return v }
