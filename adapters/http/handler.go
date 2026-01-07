package http

import (
	"codex-documents/configs"
	"codex-documents/core/ports"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Handler struct {
	cfg            *configs.Config
	patientService ports.PatientService
}

func NewHandler(cfg *configs.Config, ps ports.PatientService) *Handler {
	return &Handler{
		cfg:            cfg,
		patientService: ps,
	}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	authMid := NewAuthMiddleware(h.cfg.Auth.JWTSecret)

	router.HandleFunc("/health", h.HealthCheck).Methods("GET")

	api := router.PathPrefix("/api/v1").Subrouter()
	api.Use(authMid.Handler)

	p := api.PathPrefix("/Patient").Subrouter()
	p.HandleFunc("", h.CreatePatient).Methods("POST")
	p.HandleFunc("/{id}", h.GetPatient).Methods("GET")
	p.HandleFunc("/{id}", h.UpdatePatient).Methods("PUT")
}

func (h *Handler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "OK")
}

func (h *Handler) respondWithResource(w http.ResponseWriter, status int, resource interface{}) {
	w.Header().Set("Content-Type", "application/fhir+json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(resource)
}
