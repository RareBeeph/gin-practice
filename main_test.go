package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func failIf(t *testing.T, b bool) {
	if b {
		t.Fail()
	}
}

// TestParametric is not a very good name for this, I would use something more descriptive
func TestParametric(t *testing.T) {
	router := setupRouter()

	var testList = []struct {
		verb     string
		endpoint string
		bodyIn   string
		codeOut  int
		bodyOut  string
	}{
		{"GET", "/GetEndpoint", "", http.StatusOK, "Placeholder response (Get)"},
		{"POST", "/PostEndpoint", "Placeholder input", http.StatusOK, "Placeholder input"},
		{"DELETE", "/DeleteEndpoint", "", http.StatusOK, "Placeholder response (Delete)"},
	}

	for _, testInstance := range testList {
		// You're passing an empty string for the name of every test, don't do that
		// Should one fail you'll have an unclear debug log as to which one and why
		t.Run("", func(inner *testing.T) {
			// Excellent use of the standard library here
			w := httptest.NewRecorder()

			req := httptest.NewRequest(testInstance.verb, testInstance.endpoint, strings.NewReader(testInstance.bodyIn))
			router.ServeHTTP(w, req)

			// Instead of splitting these up into two method calls, I would create a
			// function that takes a test instance and performs all relevant validations on it
			// so you only have to update one spot
			failIf(inner, (w.Code != testInstance.codeOut))
			failIf(inner, (w.Body.String() != testInstance.bodyOut))
		})
	}
}
