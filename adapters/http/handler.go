package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gruzdev-dev/codex-documents/configs"
	"github.com/gruzdev-dev/codex-documents/core/ports"

	"github.com/gorilla/mux"
)

type Handler struct {
	cfg             *configs.Config
	patientService  ports.PatientService
	documentService ports.DocumentService
}

func NewHandler(cfg *configs.Config, ps ports.PatientService, ds ports.DocumentService) *Handler {
	return &Handler{
		cfg:             cfg,
		patientService:  ps,
		documentService: ds,
	}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	authMid := NewAuthMiddleware(h.cfg.Auth.JWTSecret)

	router.HandleFunc("/health", h.HealthCheck).Methods("GET")

	api := router.PathPrefix("/api/v1").Subrouter()
	api.Use(authMid.Handler)

	p := api.PathPrefix("/Patient").Subrouter()
	p.HandleFunc("/{id}", h.GetPatient).Methods("GET")
	p.HandleFunc("/{id}", h.UpdatePatient).Methods("PUT")

	d := api.PathPrefix("/DocumentReference").Subrouter()
	d.HandleFunc("", h.CreateDocument).Methods("POST")
	d.HandleFunc("", h.ListDocuments).Methods("GET")
	d.HandleFunc("/{id}", h.GetDocument).Methods("GET")
	d.HandleFunc("/{id}", h.UpdateDocument).Methods("PUT")
	d.HandleFunc("/{id}", h.DeleteDocument).Methods("DELETE")
}

func (h *Handler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprintln(w, "OK")
}

func (h *Handler) respondWithResource(w http.ResponseWriter, status int, resource interface{}) {
	w.Header().Set("Content-Type", "application/fhir+json")
	w.WriteHeader(status)
	if resource != nil {
		err := json.NewEncoder(w).Encode(resource)
		if err != nil {
			w.WriteHeader(http.StatusNoContent)
		}
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}
