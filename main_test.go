package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/joho/godotenv"
)

func TestAPICEPInvalido(t *testing.T) {

	req, err := http.NewRequest("GET", "/temperatura?cep=80530", nil)

	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	// Faz a chamada para o handler
	handler := http.HandlerFunc(climaHandler)

	handler.ServeHTTP(rr, req)

	// Check status code
	if status := rr.Code; status != http.StatusUnprocessableEntity {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusUnprocessableEntity)
	}

	expected := `{"message": "invalid zipcode"}`

	if responseBody := strings.TrimSpace(rr.Body.String()); responseBody != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			responseBody, expected)
	}
}

func TestAPIOK(t *testing.T) {

	_ = godotenv.Load()
	weatherAPIKey = os.Getenv("WEATHER_API_KEY")

	req, err := http.NewRequest("GET", "/temperatura?cep=80530000", nil)

	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	// Faz a chamada para o handler
	handler := http.HandlerFunc(climaHandler)

	handler.ServeHTTP(rr, req)

	// Check status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

func TestAPINotFound(t *testing.T) {

	req, err := http.NewRequest("GET", "/temperatura?cep=80530999", nil)

	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	// Faz a chamada para o handler
	handler := http.HandlerFunc(climaHandler)

	handler.ServeHTTP(rr, req)

	// Check status code
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}

	expected := `{"message": "cannot find zipcode"}`

	if responseBody := strings.TrimSpace(rr.Body.String()); responseBody != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			responseBody, expected)
	}
}
