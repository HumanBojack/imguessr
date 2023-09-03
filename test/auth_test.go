package test

import (
	"net/http"
	"strings"
	"testing"
)

func TestRegister(t *testing.T) {
	// Set the different requests to test
	requests := []struct {
		body       string
		wantedCode int
	}{
		{`{"name":"test", "password":"test"}`, http.StatusOK},
		{`{"name":"test"}`, http.StatusBadRequest},
	}

	for _, r := range requests {
		// Create a request to pass to our handler. The body is the JSON user
		req, err := http.NewRequest("POST", "/v1/auth/register",
			strings.NewReader(r.body),
		)
		if err != nil {
			t.Fatal(err)
		}

		w := mockRequest(req)
		checkResponse(t, w, r.wantedCode, "")
	}
}

func TestLogin(t *testing.T) {
	// Set the different requests to test
	requests := []struct {
		body       string
		wantedCode int
		bodyIncl   string
	}{
		{`{"name":"testLogin", "password":"test1"}`, http.StatusOK, "token"},
		{`{"name":"testLogin", "password":"wrongPwd"}`, http.StatusBadRequest, "Invalid password"},
		{`{"name":"NonExistingUser", "password":"test"}`, http.StatusBadRequest, "can't find user with name :"},
	}

	for _, r := range requests {
		// Create a request to pass to our handler. The body is the JSON user
		req, err := http.NewRequest("POST", "/v1/auth/login",
			strings.NewReader(r.body),
		)

		if err != nil {
			t.Fatal(err)
		}

		w := mockRequest(req)
		checkResponse(t, w, r.wantedCode, r.bodyIncl)
	}
}
