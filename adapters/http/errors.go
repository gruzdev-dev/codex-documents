package http

import (
	"errors"
	"net/http"

	"github.com/gruzdev-dev/codex-documents/core/domain"
	"github.com/gruzdev-dev/codex-documents/pkg/ptr"
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
		h.respondWithResource(w, http.StatusInternalServerError, models.OperationOutcome{
			Issue: []models.OperationOutcomeIssue{
				{
					Severity:    string(models.IssueSeverityFatal),
					Code:        string(models.IssueTypeException),
					Diagnostics: ptr.To("Internal error during error response generation: " + valErr.Error()),
				},
			},
		})
		return
	}

	h.respondWithResource(w, status, outcome)
}

func (h *Handler) mapErrorToFhir(err error) (int, models.IssueSeverity, models.IssueType) {
	switch {
	case errors.Is(err, domain.ErrPatientNotFound):
		return http.StatusNotFound, models.IssueSeverityError, models.IssueTypeNotFound

	case errors.Is(err, domain.ErrDocumentNotFound):
		return http.StatusNotFound, models.IssueSeverityError, models.IssueTypeNotFound

	case errors.Is(err, domain.ErrDocumentIDRequired):
		return http.StatusUnprocessableEntity, models.IssueSeverityError, models.IssueTypeRequired

	case errors.Is(err, domain.ErrObservationNotFound):
		return http.StatusNotFound, models.IssueSeverityError, models.IssueTypeNotFound

	case errors.Is(err, domain.ErrObservationIDRequired):
		return http.StatusBadRequest, models.IssueSeverityError, models.IssueTypeRequired

	case errors.Is(err, domain.ErrInvalidDerivedFromRef):
		return http.StatusUnprocessableEntity, models.IssueSeverityError, models.IssueTypeInvalid

	case errors.Is(err, domain.ErrDerivedFromDocNotFound):
		return http.StatusUnprocessableEntity, models.IssueSeverityError, models.IssueTypeNotFound

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
