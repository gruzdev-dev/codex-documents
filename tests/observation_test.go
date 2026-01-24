//go:build integration

package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	nethttp "net/http"
	"testing"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gruzdev-dev/codex-documents/core/domain"
	"github.com/gruzdev-dev/codex-documents/proto"
	models "github.com/gruzdev-dev/fhir/r5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"google.golang.org/grpc/metadata"
)

func TestObservationIntegration(t *testing.T) {
	env := SetupTestEnv(t)
	defer env.Cleanup()

	env.MockFileProvider.EXPECT().GetPresignedUrls(gomock.Any(), gomock.Any()).Return(&domain.PresignedUrlsResponse{
		FileId:      "test-file-id",
		UploadUrl:   "http://test/upload",
		DownloadUrl: "http://test/download",
	}, nil).AnyTimes()

	client := &nethttp.Client{}

	var patientID string
	var token string
	var doc1ID, doc2ID, doc3ID string
	var obs1ID, obs2ID, obs3ID string

	t.Run("Step 1: Create Patient via gRPC", func(t *testing.T) {
		md := metadata.Pairs("x-internal-token", "test-secret")
		ctx := metadata.NewOutgoingContext(context.Background(), md)

		resp, err := env.GRPCClient.CreatePatient(ctx, &proto.CreatePatientRequest{
			Email: "test@example.com",
		})
		require.NoError(t, err)

		patientID = resp.PatientId
		require.NotEmpty(t, patientID)

		claims := jwt.MapClaims{
			"sub":        "test-user",
			"patient_id": patientID,
			"scope":      "patient/*.read patient/*.write",
		}
		jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		token, err = jwtToken.SignedString([]byte("secret-key"))
		require.NoError(t, err)
	})

	t.Run("Step 2: Create Document 1", func(t *testing.T) {
		doc := models.DocumentReference{
			ResourceType: "DocumentReference",
			Status:       "current",
			Content: []models.DocumentReferenceContent{
				{
					Attachment: &models.Attachment{
						ContentType: strPtr("application/pdf"),
						Size:        int64Ptr(1024),
						Title:       strPtr("Document 1"),
					},
				},
			},
		}

		docJSON, err := json.Marshal(doc)
		require.NoError(t, err)

		req, err := nethttp.NewRequest("POST", env.ServerURL+"/api/v1/DocumentReference", bytes.NewBuffer(docJSON))
		require.NoError(t, err)
		req.Header.Set("Content-Type", "application/fhir+json")
		req.Header.Set("Authorization", "Bearer "+token)

		resp, err := client.Do(req)
		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, nethttp.StatusCreated, resp.StatusCode)

		body, err := io.ReadAll(resp.Body)
		require.NoError(t, err)

		var createdDoc models.DocumentReference
		err = json.Unmarshal(body, &createdDoc)
		require.NoError(t, err)
		require.NotNil(t, createdDoc.Id)
		doc1ID = *createdDoc.Id
	})

	t.Run("Step 3: Create Document 2", func(t *testing.T) {
		doc := models.DocumentReference{
			ResourceType: "DocumentReference",
			Status:       "current",
			Content: []models.DocumentReferenceContent{
				{
					Attachment: &models.Attachment{
						ContentType: strPtr("application/pdf"),
						Size:        int64Ptr(2048),
						Title:       strPtr("Document 2"),
					},
				},
			},
		}

		docJSON, err := json.Marshal(doc)
		require.NoError(t, err)

		req, err := nethttp.NewRequest("POST", env.ServerURL+"/api/v1/DocumentReference", bytes.NewBuffer(docJSON))
		require.NoError(t, err)
		req.Header.Set("Content-Type", "application/fhir+json")
		req.Header.Set("Authorization", "Bearer "+token)

		resp, err := client.Do(req)
		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, nethttp.StatusCreated, resp.StatusCode)

		body, err := io.ReadAll(resp.Body)
		require.NoError(t, err)

		var createdDoc models.DocumentReference
		err = json.Unmarshal(body, &createdDoc)
		require.NoError(t, err)
		require.NotNil(t, createdDoc.Id)
		doc2ID = *createdDoc.Id
	})

	t.Run("Step 4: Create Document 3", func(t *testing.T) {
		doc := models.DocumentReference{
			ResourceType: "DocumentReference",
			Status:       "current",
			Content: []models.DocumentReferenceContent{
				{
					Attachment: &models.Attachment{
						ContentType: strPtr("application/pdf"),
						Size:        int64Ptr(3072),
						Title:       strPtr("Document 3"),
					},
				},
			},
		}

		docJSON, err := json.Marshal(doc)
		require.NoError(t, err)

		req, err := nethttp.NewRequest("POST", env.ServerURL+"/api/v1/DocumentReference", bytes.NewBuffer(docJSON))
		require.NoError(t, err)
		req.Header.Set("Content-Type", "application/fhir+json")
		req.Header.Set("Authorization", "Bearer "+token)

		resp, err := client.Do(req)
		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, nethttp.StatusCreated, resp.StatusCode)

		body, err := io.ReadAll(resp.Body)
		require.NoError(t, err)

		var createdDoc models.DocumentReference
		err = json.Unmarshal(body, &createdDoc)
		require.NoError(t, err)
		require.NotNil(t, createdDoc.Id)
		doc3ID = *createdDoc.Id
	})

	t.Run("Step 5: Create Observation 1 - simple, no DerivedFrom", func(t *testing.T) {
		obs := models.Observation{
			ResourceType: "Observation",
			Status:       "final",
			Code: &models.CodeableConcept{
				Coding: []models.Coding{
					{
						System:  strPtr("http://loinc.org"),
						Code:    strPtr("85354-9"),
						Display: strPtr("Blood pressure"),
					},
				},
			},
			EffectiveDateTime: strPtr("2024-01-15T10:30:00Z"),
			ValueQuantity: &models.Quantity{
				Value:  float64Ptr(120.0),
				Unit:   strPtr("mmHg"),
				System: strPtr("http://unitsofmeasure.org"),
				Code:   strPtr("mm[Hg]"),
			},
		}

		obsJSON, err := json.Marshal(obs)
		require.NoError(t, err)

		req, err := nethttp.NewRequest("POST", env.ServerURL+"/api/v1/Observation", bytes.NewBuffer(obsJSON))
		require.NoError(t, err)
		req.Header.Set("Content-Type", "application/fhir+json")
		req.Header.Set("Authorization", "Bearer "+token)

		resp, err := client.Do(req)
		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, nethttp.StatusCreated, resp.StatusCode)

		body, err := io.ReadAll(resp.Body)
		require.NoError(t, err)

		var createdObs models.Observation
		err = json.Unmarshal(body, &createdObs)
		require.NoError(t, err)
		require.NotNil(t, createdObs.Id)
		obs1ID = *createdObs.Id

		assert.NotNil(t, createdObs.Subject)
		assert.Contains(t, *createdObs.Subject.Reference, patientID)
		assert.Nil(t, createdObs.DerivedFrom)
	})

	t.Run("Step 6: Create Observation 2 - with DerivedFrom to Document 1", func(t *testing.T) {
		doc1Ref := fmt.Sprintf("DocumentReference/%s", doc1ID)
		obs := models.Observation{
			ResourceType: "Observation",
			Status:       "final",
			Code: &models.CodeableConcept{
				Coding: []models.Coding{
					{
						System:  strPtr("http://loinc.org"),
						Code:    strPtr("718-7"),
						Display: strPtr("Hemoglobin"),
					},
				},
			},
			EffectiveDateTime: strPtr("2024-01-15T11:00:00Z"),
			ValueQuantity: &models.Quantity{
				Value:  float64Ptr(14.5),
				Unit:   strPtr("g/dL"),
				System: strPtr("http://unitsofmeasure.org"),
				Code:   strPtr("g/dL"),
			},
			DerivedFrom: []models.Reference{
				{
					Reference: &doc1Ref,
				},
			},
		}

		obsJSON, err := json.Marshal(obs)
		require.NoError(t, err)

		req, err := nethttp.NewRequest("POST", env.ServerURL+"/api/v1/Observation", bytes.NewBuffer(obsJSON))
		require.NoError(t, err)
		req.Header.Set("Content-Type", "application/fhir+json")
		req.Header.Set("Authorization", "Bearer "+token)

		resp, err := client.Do(req)
		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, nethttp.StatusCreated, resp.StatusCode)

		body, err := io.ReadAll(resp.Body)
		require.NoError(t, err)

		var createdObs models.Observation
		err = json.Unmarshal(body, &createdObs)
		require.NoError(t, err)
		require.NotNil(t, createdObs.Id)
		obs2ID = *createdObs.Id

		assert.NotNil(t, createdObs.DerivedFrom)
		assert.Len(t, createdObs.DerivedFrom, 1)
		assert.Equal(t, doc1Ref, *createdObs.DerivedFrom[0].Reference)
	})

	t.Run("Step 7: Create Observation 3 - with DerivedFrom to Documents 1 and 2", func(t *testing.T) {
		doc1Ref := fmt.Sprintf("DocumentReference/%s", doc1ID)
		doc2Ref := fmt.Sprintf("DocumentReference/%s", doc2ID)
		obs := models.Observation{
			ResourceType: "Observation",
			Status:       "preliminary",
			Code: &models.CodeableConcept{
				Coding: []models.Coding{
					{
						System:  strPtr("http://loinc.org"),
						Code:    strPtr("2093-3"),
						Display: strPtr("Cholesterol"),
					},
				},
			},
			EffectiveDateTime: strPtr("2024-01-15T12:00:00Z"),
			ValueQuantity: &models.Quantity{
				Value:  float64Ptr(200.0),
				Unit:   strPtr("mg/dL"),
				System: strPtr("http://unitsofmeasure.org"),
				Code:   strPtr("mg/dL"),
			},
			DerivedFrom: []models.Reference{
				{
					Reference: &doc1Ref,
				},
				{
					Reference: &doc2Ref,
				},
			},
		}

		obsJSON, err := json.Marshal(obs)
		require.NoError(t, err)

		req, err := nethttp.NewRequest("POST", env.ServerURL+"/api/v1/Observation", bytes.NewBuffer(obsJSON))
		require.NoError(t, err)
		req.Header.Set("Content-Type", "application/fhir+json")
		req.Header.Set("Authorization", "Bearer "+token)

		resp, err := client.Do(req)
		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, nethttp.StatusCreated, resp.StatusCode)

		body, err := io.ReadAll(resp.Body)
		require.NoError(t, err)

		var createdObs models.Observation
		err = json.Unmarshal(body, &createdObs)
		require.NoError(t, err)
		require.NotNil(t, createdObs.Id)
		obs3ID = *createdObs.Id

		assert.NotNil(t, createdObs.DerivedFrom)
		assert.Len(t, createdObs.DerivedFrom, 2)
	})

	t.Run("Step 8: Update Observation 2 - change DerivedFrom to Document 3", func(t *testing.T) {
		doc3Ref := fmt.Sprintf("DocumentReference/%s", doc3ID)
		obs := models.Observation{
			ResourceType: "Observation",
			Id:           &obs2ID,
			Status:       "final",
			Code: &models.CodeableConcept{
				Coding: []models.Coding{
					{
						System:  strPtr("http://loinc.org"),
						Code:    strPtr("718-7"),
						Display: strPtr("Hemoglobin"),
					},
				},
			},
			EffectiveDateTime: strPtr("2024-01-15T11:00:00Z"),
			ValueQuantity: &models.Quantity{
				Value:  float64Ptr(14.5),
				Unit:   strPtr("g/dL"),
				System: strPtr("http://unitsofmeasure.org"),
				Code:   strPtr("g/dL"),
			},
			DerivedFrom: []models.Reference{
				{
					Reference: &doc3Ref,
				},
			},
		}

		obsJSON, err := json.Marshal(obs)
		require.NoError(t, err)

		req, err := nethttp.NewRequest("PUT", env.ServerURL+"/api/v1/Observation/"+obs2ID, bytes.NewBuffer(obsJSON))
		require.NoError(t, err)
		req.Header.Set("Content-Type", "application/fhir+json")
		req.Header.Set("Authorization", "Bearer "+token)

		resp, err := client.Do(req)
		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, nethttp.StatusOK, resp.StatusCode)

		body, err := io.ReadAll(resp.Body)
		require.NoError(t, err)

		var updatedObs models.Observation
		err = json.Unmarshal(body, &updatedObs)
		require.NoError(t, err)

		assert.NotNil(t, updatedObs.DerivedFrom)
		assert.Len(t, updatedObs.DerivedFrom, 1)
		assert.Equal(t, doc3Ref, *updatedObs.DerivedFrom[0].Reference)
	})

	t.Run("Step 9: Get Observation 2 by ID", func(t *testing.T) {
		req, err := nethttp.NewRequest("GET", env.ServerURL+"/api/v1/Observation/"+obs2ID, nil)
		require.NoError(t, err)
		req.Header.Set("Authorization", "Bearer "+token)

		resp, err := client.Do(req)
		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, nethttp.StatusOK, resp.StatusCode)

		body, err := io.ReadAll(resp.Body)
		require.NoError(t, err)

		var obs models.Observation
		err = json.Unmarshal(body, &obs)
		require.NoError(t, err)

		assert.Equal(t, obs2ID, *obs.Id)
		doc3Ref := fmt.Sprintf("DocumentReference/%s", doc3ID)
		assert.NotNil(t, obs.DerivedFrom)
		assert.Len(t, obs.DerivedFrom, 1)
		assert.Equal(t, doc3Ref, *obs.DerivedFrom[0].Reference)
	})

	t.Run("Step 10: Delete Observation 1", func(t *testing.T) {
		req, err := nethttp.NewRequest("DELETE", env.ServerURL+"/api/v1/Observation/"+obs1ID, nil)
		require.NoError(t, err)
		req.Header.Set("Authorization", "Bearer "+token)

		resp, err := client.Do(req)
		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, nethttp.StatusOK, resp.StatusCode)
	})

	t.Run("Step 11: List Observations - should have 2", func(t *testing.T) {
		req, err := nethttp.NewRequest("GET", env.ServerURL+"/api/v1/Observation?patient="+patientID, nil)
		require.NoError(t, err)
		req.Header.Set("Authorization", "Bearer "+token)

		resp, err := client.Do(req)
		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, nethttp.StatusOK, resp.StatusCode)

		body, err := io.ReadAll(resp.Body)
		require.NoError(t, err)

		var bundle models.Bundle
		err = json.Unmarshal(body, &bundle)
		require.NoError(t, err)

		assert.Equal(t, 2, *bundle.Total, "Should have 2 observations after deleting one")
		assert.Len(t, bundle.Entry, 2)

		var foundObs2, foundObs3 bool
		for _, entry := range bundle.Entry {
			var obs models.Observation
			err := json.Unmarshal(entry.Resource, &obs)
			require.NoError(t, err)

			if *obs.Id == obs2ID {
				foundObs2 = true
				assert.NotNil(t, obs.Subject)
				assert.Contains(t, *obs.Subject.Reference, patientID)
				doc3Ref := fmt.Sprintf("DocumentReference/%s", doc3ID)
				assert.NotNil(t, obs.DerivedFrom)
				assert.Equal(t, doc3Ref, *obs.DerivedFrom[0].Reference)
			}
			if *obs.Id == obs3ID {
				foundObs3 = true
				assert.NotNil(t, obs.Subject)
				assert.Contains(t, *obs.Subject.Reference, patientID)
				assert.NotNil(t, obs.DerivedFrom)
				assert.Len(t, obs.DerivedFrom, 2)
			}
		}

		assert.True(t, foundObs2, "Observation 2 should be in the list")
		assert.True(t, foundObs3, "Observation 3 should be in the list")
	})
}

func float64Ptr(f float64) *float64 {
	return &f
}
