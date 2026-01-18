package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gruzdev-dev/codex-documents/core/domain"
	"github.com/gruzdev-dev/codex-documents/pkg/ptr"
	models "github.com/gruzdev-dev/fhir/r5"
)

func (h *Handler) CreateDocument(w http.ResponseWriter, r *http.Request) {
	var doc models.DocumentReference
	if err := json.NewDecoder(r.Body).Decode(&doc); err != nil {
		h.respondWithError(w, err)
		return
	}

	if err := doc.Validate(); err != nil {
		h.respondWithError(w, fmt.Errorf("%w: %v", domain.ErrInvalidInput, err))
		return
	}

	createdDoc, err := h.documentService.CreateDocument(r.Context(), &doc)
	if err != nil {
		h.respondWithError(w, err)
		return
	}

	h.respondWithResource(w, http.StatusCreated, createdDoc)
}

func (h *Handler) GetDocument(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	doc, err := h.documentService.GetDocument(r.Context(), id)
	if err != nil {
		h.respondWithError(w, err)
		return
	}

	h.respondWithResource(w, http.StatusOK, doc)
}

func (h *Handler) UpdateDocument(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var doc models.DocumentReference
	if err := json.NewDecoder(r.Body).Decode(&doc); err != nil {
		h.respondWithError(w, err)
		return
	}

	doc.Id = ptr.To(id)

	if err := doc.Validate(); err != nil {
		h.respondWithError(w, fmt.Errorf("%w: %v", domain.ErrInvalidInput, err))
		return
	}

	updatedDoc, err := h.documentService.UpdateDocument(r.Context(), &doc)
	if err != nil {
		h.respondWithError(w, err)
		return
	}

	h.respondWithResource(w, http.StatusOK, updatedDoc)
}

func (h *Handler) DeleteDocument(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	if err := h.documentService.DeleteDocument(r.Context(), id); err != nil {
		h.respondWithError(w, err)
		return
	}

	h.respondWithResource(w, http.StatusOK, nil)
}

func (h *Handler) ListDocuments(w http.ResponseWriter, r *http.Request) {
	patientID := r.URL.Query().Get("patient")
	if patientID == "" {
		h.respondWithError(w, fmt.Errorf("%w: patient query parameter is required", domain.ErrInvalidInput))
		return
	}

	limit, offset := h.parsePagination(r)

	res, err := h.documentService.ListDocuments(r.Context(), patientID, limit, offset)
	if err != nil {
		h.respondWithError(w, err)
		return
	}

	bundle := h.wrapInBundle(res.Items, res.Total)
	h.respondWithResource(w, http.StatusOK, bundle)
}

func (h *Handler) parsePagination(r *http.Request) (limit int, offset int) {
	limit, _ = strconv.Atoi(r.URL.Query().Get("_count"))
	if limit <= 0 || limit > 100 {
		limit = 20
	}

	offset, _ = strconv.Atoi(r.URL.Query().Get("_offset"))
	if offset < 0 {
		offset = 0
	}
	return limit, offset
}

func (h *Handler) wrapInBundle(docs []models.DocumentReference, total int64) *models.Bundle {
	bundleID := fmt.Sprintf("bundle-%s", strconv.FormatInt(total, 10))

	bundle := &models.Bundle{
		ResourceType: "Bundle",
		Id:           &bundleID,
		Type:         "searchset",
		Total:        ptr.To(int(total)),
		Entry:        make([]models.BundleEntry, 0, len(docs)),
	}

	for i := range docs {
		resourceRaw, err := json.Marshal(docs[i])
		if err != nil {
			continue
		}

		bundle.Entry = append(bundle.Entry, models.BundleEntry{
			Resource: resourceRaw,
		})
	}

	return bundle
}
