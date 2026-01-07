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

const (
	CREATE_JSON = `{
		"resourceType": "Patient",
		"active": true,
		"name": [
			{
				"use": "official",
				"family": "Ivanov",
				"given": ["Ivan", "Sergeevich"]
			}
		],
		"telecom": [
			{
				"system": "phone",
				"value": "+79001234567",
				"use": "mobile"
			},
			{
				"system": "email",
				"value": "ivan@example.com"
			}
		],
		"gender": "male",
		"birthDate": "1985-10-25",
		"deceasedBoolean": false,
		"address": [
			{
				"use": "home",
				"line": ["ul. Lenina, 1"],
				"city": "Moscow",
				"postalCode": "101000"
			}
		],
		"maritalStatus": {
			"coding": [
				{
					"system": "http://terminology.hl7.org/CodeSystem/v3-MaritalStatus",
					"code": "U",
					"display": "unmarried"
				}
			]
		},
		"communication": [
			{
				"language": {
					"coding": [
						{
							"system": "urn:ietf:bcp:47",
							"code": "ru",
							"display": "Russian"
						}
					]
				},
				"preferred": true
			}
		]
	}`
	UPDATE_JSON = `{
		"resourceType": "Patient",
		"id": "%s",
		"active": false,
		"name": [
			{
				"use": "official",
				"family": "Ivanov",
				"given": ["Ivan", "Sergeevich"]
			}
		],
		"telecom": [
			{
				"system": "phone",
				"value": "+79001234567",
				"use": "mobile"
			},
			{
				"system": "email",
				"value": "ivan@example.com"
			}
		],
		"gender": "female",
		"birthDate": "1985-10-25",
		"deceasedBoolean": false,
		"address": [
			{
				"use": "home",
				"line": ["ul. Lenina, 1"],
				"city": "Moscow",
				"postalCode": "101000"
			}
		],
		"maritalStatus": {
			"coding": [
				{
					"system": "http://terminology.hl7.org/CodeSystem/v3-MaritalStatus",
					"code": "M",
					"display": "Married"
				}
			]
		},
		"communication": [
			{
				"language": {
					"coding": [
						{
							"system": "urn:ietf:bcp:47",
							"code": "ru",
							"display": "Russian"
						}
					]
				},
				"preferred": true
			}
		]
	}`
)

func createTestJWTToken(secret string, patientID string) (string, error) {
	claims := jwt.MapClaims{
		"sub":        "test-user",
		"patient_id": patientID,
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

	client := &nethttp.Client{}

	var patientID string
	var token string

	t.Run("Create Patient", func(t *testing.T) {
		req, err := nethttp.NewRequest("POST", ts.URL+"/api/v1/Patient", bytes.NewBufferString(CREATE_JSON))
		require.NoError(t, err)
		req.Header.Set("Content-Type", "application/fhir+json")

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

		token, err = createTestJWTToken("secret-key", patientID)
		require.NoError(t, err)
	})

	t.Run("Update Patient", func(t *testing.T) {
		require.NotEmpty(t, patientID, "Patient ID should be set from Create step")

		updateJSON := fmt.Sprintf(UPDATE_JSON, patientID)

		req, err := nethttp.NewRequest("PUT", ts.URL+"/api/v1/Patient/"+patientID, bytes.NewBufferString(updateJSON))
		require.NoError(t, err)
		req.Header.Set("Content-Type", "application/fhir+json")
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

		expectedJSON := fmt.Sprintf(UPDATE_JSON, patientID)

		assert.JSONEq(t, expectedJSON, string(body), "Response JSON should match expected structure")
	})
}
