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
		t.Run("", func(inner *testing.T) {
			w := httptest.NewRecorder()

			req := httptest.NewRequest(testInstance.verb, testInstance.endpoint, strings.NewReader(testInstance.bodyIn))
			router.ServeHTTP(w, req)

			failIf(inner, (w.Code != testInstance.codeOut))
			failIf(inner, (w.Body.String() != testInstance.bodyOut))
		})
	}
}
