package test

import (
	"io"
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

		// Get the body
		body, err := io.ReadAll(w.Body)
		if err != nil {
			t.Fatal(err)
		}
		bodyStr := string(body)

		// Check the status code is what we expect.
		if status := w.Code; status != r.wantedCode {
			t.Errorf("handler returned wrong status code: got %v want %v. Body: %v",
				status, http.StatusOK, bodyStr)
		}
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

		// Get the body
		body, err := io.ReadAll(w.Body)
		if err != nil {
			t.Fatal(err)
		}
		bodyStr := string(body)

		// Check the status code is what we expect.
		if status := w.Code; status != r.wantedCode {
			t.Errorf("handler returned wrong status code: got %v want %v. Body: %v",
				status, http.StatusOK, bodyStr)
		}

		// Check that the body contains a token
		if !strings.Contains(bodyStr, r.bodyIncl) {
			t.Errorf("body doesn't contain %v. Body: %v", r.bodyIncl, bodyStr)
		}
	}
}
