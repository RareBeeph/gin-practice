package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type testContents struct {
	verb     string
	endpoint string
	bodyIn   string
	codeOut  int
	bodyOut  string
}

func failCheck(t *testing.T, checkRecorder *httptest.ResponseRecorder, checkInstance testContents) {
	if checkRecorder.Code != checkInstance.codeOut {
		t.Fail()
		log.Printf("Unexpected response code: expected %d, received %d", checkInstance.codeOut, checkRecorder.Code)
	}
	if checkRecorder.Body.String() != checkInstance.bodyOut {
		t.Fail()
		log.Printf("Unexpected response body: expected %s, received %s", checkInstance.bodyOut, checkRecorder.Body.String())
	}
}

func TestEndpoints(t *testing.T) {
	router := setupRouter()

	var testList = []testContents{
		{"GET", "/getendpoint", "", http.StatusOK, "Placeholder response (Get)"},
		{"POST", "/postendpoint", "Placeholder input", http.StatusOK, "Placeholder input"},
		{"DELETE", "/deleteendpoint", "", http.StatusOK, "Placeholder response (Delete)"},
	}

	for _, testInstance := range testList {
		t.Run("Test"+testInstance.verb, func(inner *testing.T) {
			testRecorder := httptest.NewRecorder()

			testRequest := httptest.NewRequest(testInstance.verb, testInstance.endpoint, strings.NewReader(testInstance.bodyIn))
			router.ServeHTTP(testRecorder, testRequest)

			failCheck(inner, testRecorder, testInstance)
		})
	}
}
