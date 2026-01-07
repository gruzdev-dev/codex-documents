package http

import (
	"encoding/json"
	"errors"
	"net/http"

	"codex-documents/core/domain"
	"codex-documents/pkg/ptr"
	models "github.com/gruzdev-dev/fhir/r5"
)

func (h *Handler) respondWithError(w http.ResponseWriter, err error) {
	status, severity, issueCode := h.mapErrorToFhir(err)

	outcome := models.OperationOutcome{
		Issue: []models.OperationOutcomeIssue{
			{
				Severity:    string(severity),
				Code:        string(issueCode),
				Diagnostics: ptr.To(err.Error()),
			},
		},
	}

	if valErr := outcome.Validate(); valErr != nil {
		h.respondWithInternalError(w, valErr)
		return
	}

	w.Header().Set("Content-Type", "application/fhir+json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(outcome)
}

func (h *Handler) mapErrorToFhir(err error) (int, models.IssueSeverity, models.IssueType) {
	switch {
	case errors.Is(err, domain.ErrPatientNotFound):
		return http.StatusNotFound, models.IssueSeverityError, models.IssueTypeNotFound

	case errors.Is(err, domain.ErrAccessDenied):
		return http.StatusForbidden, models.IssueSeverityError, models.IssueTypeForbidden

	case errors.Is(err, domain.ErrInvalidInput):
		return http.StatusUnprocessableEntity, models.IssueSeverityError, models.IssueTypeInvalid

	case errors.Is(err, domain.ErrPatientIDRequired):
		return http.StatusBadRequest, models.IssueSeverityError, models.IssueTypeRequired

	default:
		return http.StatusInternalServerError, models.IssueSeverityFatal, models.IssueTypeException
	}
}

func (h *Handler) respondWithInternalError(w http.ResponseWriter, valErr error) {
	w.Header().Set("Content-Type", "application/fhir+json")
	w.WriteHeader(http.StatusInternalServerError)

	json.NewEncoder(w).Encode(models.OperationOutcome{
		Issue: []models.OperationOutcomeIssue{
			{
				Severity:    string(models.IssueSeverityFatal),
				Code:        string(models.IssueTypeException),
				Diagnostics: ptr.To("Internal error during error response generation: " + valErr.Error()),
			},
		},
	})
}
