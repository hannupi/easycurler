package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func testServer(statusCode int, resBody string) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(statusCode)
		io.WriteString(w, resBody)
	}))
	return server
}

func TestFetch_OK(t *testing.T) {
	const expectedResBody = "hello world"
	const expectedStatusCode = http.StatusOK
	server := testServer(expectedStatusCode, expectedResBody)
	defer server.Close()

	res := fetchURL(server.URL)
	msg := res()

	switch msg := msg.(type) {
	case httpResMsg:
		if string(msg) != expectedResBody {
			t.Errorf("Expected '%s' but got '%s'", expectedResBody, msg)
		}
	default:
		t.Errorf("Expected httpResMsg but got %T", msg)
	}
}

func TestFetch_Error(t *testing.T) {
	const expectedErrorMessage = "HTTP error: 404 Not Found"
	const expectedStatusCode = http.StatusNotFound

	server := testServer(expectedStatusCode, expectedErrorMessage)
	defer server.Close()

	res := fetchURL(server.URL)
	msg := res()

	switch msg := msg.(type) {
	case errMsg:
		if msg.Error() != expectedErrorMessage {
			t.Errorf("Expected error message '%s' but got '%v'", expectedErrorMessage, msg)
		}
	default:
		t.Errorf("Expected errMsg but got %T", msg)
	}
}
