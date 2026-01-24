package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gruzdev-dev/codex-documents/core/domain"
	"github.com/gruzdev-dev/codex-documents/pkg/ptr"
	models "github.com/gruzdev-dev/fhir/r5"
)

func (h *Handler) CreateObservation(w http.ResponseWriter, r *http.Request) {
	var obs models.Observation
	if err := json.NewDecoder(r.Body).Decode(&obs); err != nil {
		h.respondWithError(w, err)
		return
	}

	if err := obs.Validate(); err != nil {
		h.respondWithError(w, fmt.Errorf("%w: %v", domain.ErrInvalidInput, err))
		return
	}

	result, err := h.observationService.Create(r.Context(), &obs)
	if err != nil {
		h.respondWithError(w, err)
		return
	}

	h.respondWithResource(w, http.StatusCreated, result)
}

func (h *Handler) GetObservation(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	obs, err := h.observationService.Get(r.Context(), id)
	if err != nil {
		h.respondWithError(w, err)
		return
	}

	h.respondWithResource(w, http.StatusOK, obs)
}

func (h *Handler) UpdateObservation(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var obs models.Observation
	if err := json.NewDecoder(r.Body).Decode(&obs); err != nil {
		h.respondWithError(w, err)
		return
	}

	if err := obs.Validate(); err != nil {
		h.respondWithError(w, fmt.Errorf("%w: %v", domain.ErrInvalidInput, err))
		return
	}

	if obs.Id == nil || *obs.Id != id {
		h.respondWithError(w, fmt.Errorf("%w: id in body must match URL", domain.ErrInvalidInput))
		return
	}

	result, err := h.observationService.Update(r.Context(), &obs)
	if err != nil {
		h.respondWithError(w, err)
		return
	}

	h.respondWithResource(w, http.StatusOK, result)
}

func (h *Handler) DeleteObservation(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	if err := h.observationService.Delete(r.Context(), id); err != nil {
		h.respondWithError(w, err)
		return
	}

	h.respondWithResource(w, http.StatusOK, nil)
}

func (h *Handler) ListObservations(w http.ResponseWriter, r *http.Request) {
	patientID := r.URL.Query().Get("patient")
	if patientID == "" {
		h.respondWithError(w, fmt.Errorf("%w: patient query parameter is required", domain.ErrInvalidInput))
		return
	}

	limit, offset := h.parsePagination(r)

	res, err := h.observationService.List(r.Context(), patientID, limit, offset)
	if err != nil {
		h.respondWithError(w, err)
		return
	}

	bundle := h.wrapObservationsInBundle(res.Items, res.Total)
	h.respondWithResource(w, http.StatusOK, bundle)
}

func (h *Handler) wrapObservationsInBundle(observations []models.Observation, total int64) *models.Bundle {
	bundleID := fmt.Sprintf("bundle-%d", total)

	bundle := &models.Bundle{
		ResourceType: "Bundle",
		Id:           &bundleID,
		Type:         "searchset",
		Total:        ptr.To(int(total)),
		Entry:        make([]models.BundleEntry, 0, len(observations)),
	}

	for i := range observations {
		resourceRaw, err := json.Marshal(observations[i])
		if err != nil {
			continue
		}

		bundle.Entry = append(bundle.Entry, models.BundleEntry{
			Resource: resourceRaw,
		})
	}

	return bundle
}
