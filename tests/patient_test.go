package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	nethttp "net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang-jwt/jwt/v5"
	models "github.com/gruzdev-dev/fhir/r5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createTestJWTToken(secret string) (string, error) {
	claims := jwt.MapClaims{
		"sub":        "test-user",
		"patient_id": "",
		"scope":      "patient/*.read patient/*.write",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func TestPatientIntegration(t *testing.T) {
	env := SetupTestEnv(t)
	defer env.Cleanup()

	ts := httptest.NewServer(env.Handler)
	defer ts.Close()

	token, err := createTestJWTToken("secret-key")
	require.NoError(t, err)

	client := &nethttp.Client{}

	var patientID string

	t.Run("Create Patient", func(t *testing.T) {
		createJSON := `{
			"resourceType": "Patient",
			"active": true,
			"name": [
				{
					"family": "Ivanov",
					"given": ["Ivan"]
				}
			],
			"gender": "male",
			"birthDate": "1990-01-01",
			"telecom": [
				{
					"system": "email",
					"value": "ivan.ivanov@example.com"
				}
			]
		}`

		req, err := nethttp.NewRequest("POST", ts.URL+"/api/v1/Patient", bytes.NewBufferString(createJSON))
		require.NoError(t, err)
		req.Header.Set("Content-Type", "application/fhir+json")
		req.Header.Set("Authorization", "Bearer "+token)

		resp, err := client.Do(req)
		require.NoError(t, err)
		defer resp.Body.Close()

		require.Equal(t, nethttp.StatusCreated, resp.StatusCode, "Expected status 201 Created")

		body, err := io.ReadAll(resp.Body)
		require.NoError(t, err)

		var createdPatient models.Patient
		err = json.Unmarshal(body, &createdPatient)
		require.NoError(t, err)

		require.NotNil(t, createdPatient.Id, "Patient ID should be present")
		require.NotEmpty(t, *createdPatient.Id, "Patient ID should not be empty")
		patientID = *createdPatient.Id
	})

	t.Run("Update Patient", func(t *testing.T) {
		require.NotEmpty(t, patientID, "Patient ID should be set from Create step")

		updateJSON := `{
			"resourceType": "Patient",
			"active": false,
			"name": [
				{
					"family": "Ivanov",
					"given": ["Ivan"]
				},
				{
					"family": "Petrov",
					"given": ["Ivan"]
				}
			],
			"gender": "male",
			"birthDate": "1990-01-01",
			"telecom": [
				{
					"system": "email",
					"value": "ivan.ivanov@example.com"
				}
			]
		}`

		req, err := nethttp.NewRequest("PUT", ts.URL+"/api/v1/Patient/"+patientID, bytes.NewBufferString(updateJSON))
		require.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+token)

		resp, err := client.Do(req)
		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, nethttp.StatusOK, resp.StatusCode, "Expected status 200 OK")
	})

	t.Run("Verify FHIR JSON", func(t *testing.T) {
		require.NotEmpty(t, patientID, "Patient ID should be set from Create step")

		req, err := nethttp.NewRequest("GET", ts.URL+"/api/v1/Patient/"+patientID, nil)
		require.NoError(t, err)
		req.Header.Set("Authorization", "Bearer "+token)

		resp, err := client.Do(req)
		require.NoError(t, err)
		defer resp.Body.Close()

		require.Equal(t, nethttp.StatusOK, resp.StatusCode, "Expected status 200 OK")

		body, err := io.ReadAll(resp.Body)
		require.NoError(t, err)

		expectedJSON := fmt.Sprintf(`{
			"resourceType": "Patient",
			"id": "%s",
			"active": false,
			"name": [
				{
					"family": "Ivanov",
					"given": ["Ivan"]
				},
				{
					"family": "Petrov",
					"given": ["Ivan"]
				}
			],
			"gender": "male",
			"birthDate": "1990-01-01",
			"telecom": [
				{
					"system": "email",
					"value": "ivan.ivanov@example.com"
				}
			]
		}`, patientID)

		assert.JSONEq(t, expectedJSON, string(body), "Response JSON should match expected structure")
	})
}
