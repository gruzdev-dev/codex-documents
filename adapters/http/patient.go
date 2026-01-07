package http

import (
	"encoding/json"
	"net/http"

	"codex-documents/pkg/ptr"
	models "github.com/gruzdev-dev/fhir/r5"
	"github.com/gorilla/mux"
)

func (h *Handler) CreatePatient(w http.ResponseWriter, r *http.Request) {
	var patient models.Patient
	if err := json.NewDecoder(r.Body).Decode(&patient); err != nil {
		h.respondWithError(w, err)
		return
	}

	if err := patient.Validate(); err != nil {
		h.respondWithError(w, err)
		return
	}

	if err := h.patientService.Create(r.Context(), &patient); err != nil {
		h.respondWithError(w, err)
		return
	}

	h.respondWithResource(w, http.StatusCreated, &patient)
}

func (h *Handler) GetPatient(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	patient, err := h.patientService.Get(r.Context(), id)
	if err != nil {
		h.respondWithError(w, err)
		return
	}

	h.respondWithResource(w, http.StatusOK, patient)
}

func (h *Handler) UpdatePatient(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var patient models.Patient
	if err := json.NewDecoder(r.Body).Decode(&patient); err != nil {
		h.respondWithError(w, err)
		return
	}

	patient.Id = ptr.To(id)

	if err := patient.Validate(); err != nil {
		h.respondWithError(w, err)
		return
	}

	if err := h.patientService.Update(r.Context(), &patient); err != nil {
		h.respondWithError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}