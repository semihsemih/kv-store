package api

import (
	"github.com/semihsemih/kv-store/internal/storage"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSet(t *testing.T) {
	// Initialize storage
	var s storage.Storage = storage.Storage{}

	// Request for set a new key-value
	req, err := http.NewRequest("GET", "/set?key=foo&value=bar", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Response recorder to handle response from request
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Set(s))

	handler.ServeHTTP(rr, req)

	// Check response status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check expected value
	expected := `{"foo":"bar"}`
	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestGet(t *testing.T) {
	// Initialize storage
	var s storage.Storage = storage.Storage{
		"foo": "bar",
	}

	// Request for get value from storage given by key
	req, err := http.NewRequest("GET", "/get?key=foo", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Response recorder to handle response from request
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Get(s))

	handler.ServeHTTP(rr, req)

	// Check response status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check expected value
	expected := `{"foo":"bar"}`
	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
