package tests

import (
	"bytes"
	"cmd/main.go/internal/handler"
	"cmd/main.go/internal/repository"
	"cmd/main.go/internal/service"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type testSetup struct {
	handler *handler.ReceiptHandler
	router  *mux.Router
	testID  string
}

func setup() *testSetup {
	mockRepo := repository.NewRepository()
	mockService := service.NewReceiptService(mockRepo)
	handler := handler.NewReceiptHandler(mockService)

	router := mux.NewRouter()
	router.HandleFunc("/receipts/{id}/points", handler.GetPointsById).Methods("GET")
	router.HandleFunc("/receipts/process", handler.ProcessReceipt).Methods("POST")

	testID, _ := mockRepo.AddEntry(100)

	return &testSetup{
		handler: handler,
		router:  router,
		testID:  testID,
	}
}

func TestReceiptHandler_ProcessReceipt(t *testing.T) {
	ts := setup()

	tests := []struct {
		name           string
		payload        []byte
		expectedStatus int
		expectedBody   string
		checkJSONField string
	}{
		{
			name: "Valid Receipt",
			payload: []byte(`{
				"retailer":"Walmart",
				"purchaseDate":"2024-06-26",
				"purchaseTime":"10:25",
				"items":[{"shortDescription":"xyz","price":"10.01"}],
				"total":"10.01"
			}`),
			expectedStatus: http.StatusOK,
			checkJSONField: "id",
		},
		{
			name:           "Invalid JSON",
			payload:        []byte(`{"invalid_json":`),
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "The receipt is invalid\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("POST", "/receipts/process", bytes.NewBuffer(tt.payload))
			require.NoError(t, err)

			rr := httptest.NewRecorder()
			ts.router.ServeHTTP(rr, req)

			assert.Equal(t, tt.expectedStatus, rr.Code)

			if tt.expectedBody != "" {
				assert.Equal(t, tt.expectedBody, rr.Body.String())
			}

			if tt.checkJSONField != "" {
				var responseBody map[string]interface{}
				err = json.Unmarshal(rr.Body.Bytes(), &responseBody)
				require.NoError(t, err, "Response body should be valid JSON")
				_, exists := responseBody[tt.checkJSONField]
				assert.True(t, exists, "Response should contain a '%s' field", tt.checkJSONField)
			}
		})
	}
}

func TestReceiptHandler_GetPointsByID(t *testing.T) {
	ts := setup()

	tests := []struct {
		name           string
		id             string
		expectedStatus int
		expectedBody   string
		checkJSONField string
	}{
		{
			name:           "Valid ID",
			id:             ts.testID,
			expectedStatus: http.StatusOK,
			checkJSONField: "points",
		},
		{
			name:           "Invalid ID",
			id:             "invalidID",
			expectedStatus: http.StatusNotFound,
			expectedBody:   "Receipt not found\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "/receipts/"+tt.id+"/points", nil)
			require.NoError(t, err)

			rr := httptest.NewRecorder()
			ts.router.ServeHTTP(rr, req)

			assert.Equal(t, tt.expectedStatus, rr.Code)

			if tt.expectedBody != "" {
				assert.Equal(t, tt.expectedBody, rr.Body.String())
			}

			if tt.checkJSONField != "" {
				var responseBody map[string]interface{}
				err = json.Unmarshal(rr.Body.Bytes(), &responseBody)
				require.NoError(t, err, "Response body should be valid JSON")
				_, exists := responseBody[tt.checkJSONField]
				assert.True(t, exists, "Response should contain a '%s' field", tt.checkJSONField)
			}
		})
	}
}
