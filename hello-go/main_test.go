package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/go-chi/chi/v5"
)

var r *chi.Mux

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func setup() {
	r = chi.NewRouter()
	r.Get("/hello", helloHandler)
}

func teardown() {
	// Perform any necessary cleanup after the tests.
}

func TestHelloHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/hello", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	recorder := httptest.NewRecorder()
	r.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, recorder.Code)
	}

	expectedResponse := "Hello World!"
	if recorder.Body.String() != expectedResponse {
		t.Errorf("Expected response body to be '%s', but got '%s'", expectedResponse, recorder.Body.String())
	}
}
