package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetTemperatura(t *testing.T) {
	tests := []struct {
		name           string
		localidade     string
		mockResponse   string
		mockStatusCode int
		expectedTempC  float64
		expectedTempF  float64
		expectedTempK  float64
		expectError    bool
	}{
		{
			name:           "Successful response",
			localidade:     "Sao Paulo",
			mockResponse:   `{"current": {"temp_c": 25.0}}`,
			mockStatusCode: http.StatusOK,
			expectedTempC:  25.0,
			expectedTempF:  77.0,
			expectedTempK:  298.15,
			expectError:    false,
		},
		{
			name:           "Non-200 status code",
			localidade:     "Sao Paulo",
			mockResponse:   `{"current": {"temp_c": 25.0}}`,
			mockStatusCode: http.StatusInternalServerError,
			expectedTempC:  0,
			expectedTempF:  0,
			expectedTempK:  0,
			expectError:    true,
		},
		{
			name:           "JSON unmarshalling error",
			localidade:     "Sao Paulo",
			mockResponse:   `invalid json`,
			mockStatusCode: http.StatusOK,
			expectedTempC:  0,
			expectedTempF:  0,
			expectedTempK:  0,
			expectError:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tt.mockStatusCode)
				w.Write([]byte(tt.mockResponse))
			}))
			defer server.Close()

			weatherURL = server.URL + "?key=%s&q=%s&aqi=no"

			tempC, tempF, tempK, err := getTemperatura(tt.localidade)
			if (err != nil) != tt.expectError {
				t.Errorf("expected error: %v, got: %v", tt.expectError, err)
			}
			if tempC != tt.expectedTempC {
				t.Errorf("expected tempC: %v, got: %v", tt.expectedTempC, tempC)
			}
			if tempF != tt.expectedTempF {
				t.Errorf("expected tempF: %v, got: %v", tt.expectedTempF, tempF)
			}
			if tempK != tt.expectedTempK {
				t.Errorf("expected tempK: %v, got: %v", tt.expectedTempK, tempK)
			}
		})
	}
}
