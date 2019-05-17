package test

import (
	"net/http/httptest"
	"testing"
)

func AssertResponseHeader(t *testing.T, response *httptest.ResponseRecorder, header string, want string) {
	t.Helper()
	h := response.Header().Get(header)
	if h != want {
		t.Errorf("response header '%s' is incorrect, got '%s', wanted '%s'", header, h, want)
	}
}

func AssertResponseBody(t *testing.T, response *httptest.ResponseRecorder, want string) {
	t.Helper()
	b := response.Body.String()
	if b != want {
		t.Errorf("response body is incorrect, got '%s', wanted '%s'", b, want)
	}
}

func AssertResponseContentType(t *testing.T, response *httptest.ResponseRecorder, want string) {
	t.Helper()
	AssertResponseHeader(t, response, "content-type", want)
}

func AssertResponseAuthorization(t *testing.T, response *httptest.ResponseRecorder, want string) {
	t.Helper()
	AssertResponseHeader(t, response, "authorization", want)
}

func AssertResponseStatus(t *testing.T, response *httptest.ResponseRecorder, want int) {
	t.Helper()
	s := response.Code
	if s != want {
		t.Errorf("response code is incorrect, got %d, wanted %d", s, want)
	}
}
