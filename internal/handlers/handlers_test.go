package handlers

// import (
// 	"io"
// 	"net/http"
// 	"net/http/httptest"
// 	"strings"
// 	"testing"
// )

// func TestCalculusHandler(t *testing.T) {
// 	tests := []struct {
// 		name             string
// 		request          string
// 		expectedStatus   int
// 		expectedResponse string
// 	}{
// 		{
// 			name:             "Valid request",
// 			request:          `{"expression": "1+2"}`,
// 			expectedStatus:   http.StatusOK,
// 			expectedResponse: "Result: 3.000000",
// 		},
// 		{
// 			name:             "Invalid request",
// 			request:          `{"expression": "1++"}`,
// 			expectedStatus:   http.StatusUnprocessableEntity,
// 			expectedResponse: "ERROR: invalid expression",
// 		},
// 		{
// 			name:             "Empty request",
// 			request:          "",
// 			expectedStatus:   http.StatusInternalServerError,
// 			expectedResponse: "ERROR: internal server error",
// 		},
// 	}

// 	for _, test := range tests {
// 		t.Run(test.name, func(t *testing.T) {
// 			r := httptest.NewRequest("POST", "/api/v1/calculate", strings.NewReader(test.request))
// 			w := httptest.NewRecorder()

// 			CalculusHandler(w, r)
// 			resp := w.Result()
// 			defer resp.Body.Close()

// 			body, err := io.ReadAll(resp.Body)
// 			if err != nil {
// 				t.Fatalf("Error reading response body: %v", err)
// 			}

// 			if resp.StatusCode != test.expectedStatus {
// 				t.Errorf("expected status code %d, got %d", test.expectedStatus, resp.StatusCode)
// 			}

// 			if string(body) != test.expectedResponse {
// 				t.Errorf("expected response body %s, got %s", test.expectedResponse, string(body))
// 			}
// 		})
// 	}
// }
