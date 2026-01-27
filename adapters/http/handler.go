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
	cfg                *configs.Config
	patientService     ports.PatientService
	documentService    ports.DocumentService
	observationService ports.ObservationService
	shareService       ports.ShareService
}

func NewHandler(cfg *configs.Config, ps ports.PatientService, ds ports.DocumentService, os ports.ObservationService, ss ports.ShareService) *Handler {
	return &Handler{
		cfg:                cfg,
		patientService:     ps,
		documentService:    ds,
		observationService: os,
		shareService:       ss,
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
	d.HandleFunc("/{id}", h.DeleteDocument).Methods("DELETE")

	o := api.PathPrefix("/Observation").Subrouter()
	o.HandleFunc("", h.CreateObservation).Methods("POST")
	o.HandleFunc("", h.ListObservations).Methods("GET")
	o.HandleFunc("/{id}", h.GetObservation).Methods("GET")
	o.HandleFunc("/{id}", h.UpdateObservation).Methods("PUT")
	o.HandleFunc("/{id}", h.DeleteObservation).Methods("DELETE")

	api.HandleFunc("/share", h.CreateShare).Methods("POST")
	api.HandleFunc("/shared", h.GetSharedResources).Methods("GET")
}

func (h *Handler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprintln(w, "OK")
}

func (h *Handler) respondWithResource(w http.ResponseWriter, status int, resource any) {
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
