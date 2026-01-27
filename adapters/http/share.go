package http

import (
	"encoding/json"
	"net/http"

	"github.com/gruzdev-dev/codex-documents/core/domain"
)

type CreateShareRequest struct {
	ResourceIDs []string `json:"resource_ids"`
	TTLSeconds  int64    `json:"ttl_seconds"`
}

func (h *Handler) CreateShare(w http.ResponseWriter, r *http.Request) {
	var req CreateShareRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.respondWithError(w, err)
		return
	}

	shareReq := domain.ShareRequest{
		ResourceIDs: req.ResourceIDs,
		TTLSeconds:  req.TTLSeconds,
	}

	resp, err := h.shareService.Share(r.Context(), shareReq)
	if err != nil {
		h.respondWithError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
	}
}

func (h *Handler) GetSharedResources(w http.ResponseWriter, r *http.Request) {
	resp, err := h.shareService.GetSharedResources(r.Context())
	if err != nil {
		h.respondWithError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
	}
}
