package services

import (
	"context"
	"fmt"
	"strings"

	"github.com/gruzdev-dev/codex-documents/core/domain"
	"github.com/gruzdev-dev/codex-documents/core/ports"
	"github.com/gruzdev-dev/codex-documents/pkg/identity"
	models "github.com/gruzdev-dev/fhir/r5"
)

type ShareService struct {
	obsRepo         ports.ObservationRepository
	docRepo         ports.DocumentRepository
	tmpAccessClient ports.TmpAccessClient
}

func NewShareService(
	obsRepo ports.ObservationRepository,
	docRepo ports.DocumentRepository,
	tmpAccessClient ports.TmpAccessClient,
) *ShareService {
	return &ShareService{
		obsRepo:         obsRepo,
		docRepo:         docRepo,
		tmpAccessClient: tmpAccessClient,
	}
}

func (s *ShareService) Share(ctx context.Context, req domain.ShareRequest) (*domain.ShareResponse, error) {
	user, ok := identity.FromCtx(ctx)
	if !ok {
		return nil, domain.ErrAccessDenied
	}

	if !user.HasScope("patient/*.read") || user.PatientID == "" {
		return nil, domain.ErrAccessDenied
	}

	if len(req.ResourceIDs) == 0 {
		return nil, domain.ErrNoResourcesToShare
	}

	obsIDs, docIDs := s.classifyResourceIDs(req.ResourceIDs)

	allObs, err := s.obsRepo.GetByIDs(ctx, obsIDs)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", domain.ErrInternal, err)
	}

	allDocs, err := s.docRepo.GetByIDs(ctx, docIDs)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", domain.ErrInternal, err)
	}

	foundIDs := make(map[string]bool)
	for _, obs := range allObs {
		if obs.Id != nil {
			foundIDs[*obs.Id] = true
		}
	}
	for _, doc := range allDocs {
		if doc.Id != nil {
			foundIDs[*doc.Id] = true
		}
	}

	for _, resourceID := range req.ResourceIDs {
		id := resourceID
		if strings.HasPrefix(id, "Observation/") {
			id = strings.TrimPrefix(id, "Observation/")
		} else if strings.HasPrefix(id, "DocumentReference/") {
			id = strings.TrimPrefix(id, "DocumentReference/")
		}
		if !foundIDs[id] {
			return nil, domain.ErrResourceNotOwned
		}
	}

	expectedPatientRef := fmt.Sprintf("Patient/%s", user.PatientID)

	for _, obs := range allObs {
		if obs.Subject == nil || obs.Subject.Reference == nil {
			return nil, domain.ErrResourceNotOwned
		}
		if *obs.Subject.Reference != expectedPatientRef {
			return nil, domain.ErrResourceNotOwned
		}
	}

	for _, doc := range allDocs {
		if doc.Subject == nil || doc.Subject.Reference == nil {
			return nil, domain.ErrResourceNotOwned
		}
		if *doc.Subject.Reference != expectedPatientRef {
			return nil, domain.ErrResourceNotOwned
		}
	}

	docIDsFromObs := s.extractDocumentReferencesFromObservations(allObs)
	if len(docIDsFromObs) > 0 {
		additionalDocs, err := s.docRepo.GetByIDs(ctx, docIDsFromObs)
		if err != nil {
			return nil, fmt.Errorf("%w: %v", domain.ErrInternal, err)
		}
		allDocs = append(allDocs, additionalDocs...)
	}

	fileIDs := s.extractFileIDsFromDocuments(allDocs)

	scopes := s.buildScopes(allObs, allDocs, fileIDs)

	scopesStr := strings.Join(scopes, ",")
	resp, err := s.tmpAccessClient.GenerateTmpToken(ctx, domain.GenerateTmpTokenRequest{
		Payload: map[string]string{
			"scopes": scopesStr,
		},
		TtlSeconds: req.TTLSeconds,
	})
	if err != nil {
		return nil, fmt.Errorf("%w: %v", domain.ErrInternal, err)
	}

	return &domain.ShareResponse{
		Token:       resp.TmpToken,
		ResourceURL: "/api/v1/shared",
	}, nil
}

func (s *ShareService) GetSharedResources(ctx context.Context) (*domain.SharedResourcesResponse, error) {
	user, ok := identity.FromCtx(ctx)
	if !ok {
		return nil, domain.ErrAccessDenied
	}

	if !user.IsTmpToken() {
		return nil, domain.ErrAccessDenied
	}

	var observations []string
	var documentReferences []string

	for _, scope := range user.Scopes {
		parts := strings.Split(scope, ":")
		if len(parts) != 4 {
			continue
		}

		service, resource, id, action := parts[0], parts[1], parts[2], parts[3]
		if service != "docs" || action != "read" {
			continue
		}

		switch resource {
		case "observation":
			observations = append(observations, fmt.Sprintf("/api/v1/Observation/%s", id))
		case "document_reference":
			documentReferences = append(documentReferences, fmt.Sprintf("/api/v1/DocumentReference/%s", id))
		}
	}

	return &domain.SharedResourcesResponse{
		Observations:       observations,
		DocumentReferences: documentReferences,
	}, nil
}

func (s *ShareService) classifyResourceIDs(resourceIDs []string) (obsIDs []string, docIDs []string) {
	for _, id := range resourceIDs {
		if strings.HasPrefix(id, "Observation/") {
			obsIDs = append(obsIDs, strings.TrimPrefix(id, "Observation/"))
		} else if strings.HasPrefix(id, "DocumentReference/") {
			docIDs = append(docIDs, strings.TrimPrefix(id, "DocumentReference/"))
		} else {
			obsIDs = append(obsIDs, id)
			docIDs = append(docIDs, id)
		}
	}
	return obsIDs, docIDs
}

func (s *ShareService) extractDocumentReferencesFromObservations(observations []models.Observation) []string {
	docIDMap := make(map[string]bool)
	for _, obs := range observations {
		if obs.DerivedFrom == nil {
			continue
		}
		for _, ref := range obs.DerivedFrom {
			if ref.Reference == nil {
				continue
			}
			refStr := *ref.Reference
			if strings.HasPrefix(refStr, "DocumentReference/") {
				docID := strings.TrimPrefix(refStr, "DocumentReference/")
				if docID != "" {
					docIDMap[docID] = true
				}
			}
		}
	}

	docIDs := make([]string, 0, len(docIDMap))
	for docID := range docIDMap {
		docIDs = append(docIDs, docID)
	}
	return docIDs
}

func (s *ShareService) extractFileIDsFromDocuments(documents []models.DocumentReference) []string {
	fileIDMap := make(map[string]bool)
	for _, doc := range documents {
		if doc.Content == nil {
			continue
		}
		for _, content := range doc.Content {
			if content.Attachment == nil || content.Attachment.Id == nil {
				continue
			}
			fileID := *content.Attachment.Id
			if fileID != "" {
				fileIDMap[fileID] = true
			}
		}
	}

	fileIDs := make([]string, 0, len(fileIDMap))
	for fileID := range fileIDMap {
		fileIDs = append(fileIDs, fileID)
	}
	return fileIDs
}

func (s *ShareService) buildScopes(observations []models.Observation, documents []models.DocumentReference, fileIDs []string) []string {
	var scopes []string

	for _, obs := range observations {
		if obs.Id != nil && *obs.Id != "" {
			scopes = append(scopes, fmt.Sprintf("docs:observation:%s:read", *obs.Id))
		}
	}

	for _, doc := range documents {
		if doc.Id != nil && *doc.Id != "" {
			scopes = append(scopes, fmt.Sprintf("docs:document_reference:%s:read", *doc.Id))
		}
	}

	for _, fileID := range fileIDs {
		if fileID != "" {
			scopes = append(scopes, fmt.Sprintf("files:file:%s:read", fileID))
		}
	}

	return scopes
}
