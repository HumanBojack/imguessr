package test

import (
	"io"
	"net/http"
	"strings"
	"testing"
)

// Test the GetUser function as an admin and non-admin
func TestGetUser(t *testing.T) {
	// Get the auth token for the user
	userToken := getAndParseToken(t, `{"name":"testLogin","password":"test1"}`)
	t.Logf("User token : %v", userToken)
	// Get the auth token for the admin
	adminToken := getAndParseToken(t, `{"name":"testAdmin","password":"test1"}`)
	t.Logf("Admin token : %v", adminToken)

	// Set the different requests to test
	requests := []struct {
		id         string
		wantedCode int
		bodyIncl   string
		token      string
	}{
		// Test without auth token
		{"39c71853-6206-4eef-9f5b-7a1a90830ccc", http.StatusUnauthorized, "Authorization header is missing", ""},
		// Test with invalid auth token
		{"39c71853-6206-4eef-9f5b-7a1a90830ccc", http.StatusUnauthorized, "error", "invalid token"},
		// Test accessing same user with valid auth token as user
		{"39c71853-6206-4eef-9f5b-7a1a90830ccc", http.StatusOK, `"name": "testLogin"`, userToken},
		// Test accessing other user with valid auth token as user
		{"ec4e2897-4ca4-4694-94d7-96db81ec223f", http.StatusOK, `"name": "testAdmin"`, userToken},
		// Test accessing same user with valid auth token as admin
		{"ec4e2897-4ca4-4694-94d7-96db81ec223f", http.StatusOK, `"name": "testAdmin"`, adminToken},
		// Test accessing other user with valid auth token as admin
		{"39c71853-6206-4eef-9f5b-7a1a90830ccc", http.StatusOK, `"name": "testLogin"`, adminToken},
	}

	for _, r := range requests {
		// Create a request to pass to our handler.
		req, err := http.NewRequest("GET", "/v1/user/"+r.id, nil)

		if r.token != "" {
			req.Header.Set("Authorization", "Bearer "+r.token)
			t.Logf("Request with token : %v", r.token)
		}

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
				status, r.wantedCode, bodyStr)
		}

		// Check that the body contains a token
		if !strings.Contains(bodyStr, r.bodyIncl) {
			t.Errorf("body doesn't contain '%v'. Body: %v", r.bodyIncl, bodyStr)
		}
	}
}
