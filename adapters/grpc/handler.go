package grpc

import (
	"context"
	"fmt"

	"codex-documents/api/proto"
	"codex-documents/core/ports"
	"github.com/gruzdev-dev/fhir/r5"
)

type AuthHandler struct {
	proto.UnimplementedAuthIntegrationServer
	patientService ports.PatientService
}

func NewAuthHandler(ps ports.PatientService) *AuthHandler {
	return &AuthHandler{
		patientService: ps,
	}
}

func (h *AuthHandler) CreatePatient(ctx context.Context, req *proto.CreatePatientRequest) (*proto.CreatePatientResponse, error) {
	if req.Email == "" {
		return nil, fmt.Errorf("email is required")
	}

	emailSystem := "email"
	patient := &models.Patient{
		Telecom: []models.ContactPoint{
			{
				System: &emailSystem,
				Value:  &req.Email,
			},
		},
	}

	created, err := h.patientService.Create(ctx, patient)
	if err != nil {
		return nil, fmt.Errorf("failed to create patient profile: %v", err)
	}

	return &proto.CreatePatientResponse{
		PatientId: *created.Id,
		Scopes: []string{
			"patient/*.read",
			"patient/*.write",
		},
	}, nil
}
