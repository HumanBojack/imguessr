package test

import (
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
		checkResponse(t, w, r.wantedCode, r.bodyIncl)
	}
}

// Test the GetAllUsers function as an admin and non-admin
func TestGetAllUsers(t *testing.T) {
	// Get the auth token for the user
	userToken := getAndParseToken(t, `{"name":"testLogin","password":"test1"}`)
	t.Logf("User token : %v", userToken)
	// Get the auth token for the admin
	adminToken := getAndParseToken(t, `{"name":"testAdmin","password":"test1"}`)
	t.Logf("Admin token : %v", adminToken)

	// Set the different requests to test
	requests := []struct {
		wantedCode int
		bodyIncl   string
		token      string
	}{
		// Test without auth token
		{http.StatusUnauthorized, "Authorization header is missing", ""},
		// Test with invalid auth token
		{http.StatusUnauthorized, "error", "invalid token"},
		// Test as user
		{http.StatusForbidden, `Unauthorized`, userToken},
		// Test as admin
		{http.StatusOK, `"testLogin"`, adminToken},
	}

	for _, r := range requests {
		// Create a request to pass to our handler.
		req, err := http.NewRequest("GET", "/v1/user/", nil)

		if r.token != "" {
			req.Header.Set("Authorization", "Bearer "+r.token)
			t.Logf("Request with token : %v", r.token)
		}

		if err != nil {
			t.Fatal(err)
		}

		w := mockRequest(req)
		checkResponse(t, w, r.wantedCode, r.bodyIncl)
	}
}

// Test the CreateUser function as an admin and non-admin
// Test if a user created by a non-admin is not an admin
func TestCreateUser(t *testing.T) {
	// Get the auth token for the user
	userToken := getAndParseToken(t, `{"name":"testLogin","password":"test1"}`)
	t.Logf("User token : %v", userToken)
	// Get the auth token for the admin
	adminToken := getAndParseToken(t, `{"name":"testAdmin","password":"test1"}`)
	t.Logf("Admin token : %v", adminToken)

	// Set the different requests to test
	requests := []struct {
		body       string
		wantedCode int
		bodyIncl   string
		token      string
	}{
		// Test without auth token
		{`{"name":"test", "password":"test"}`, http.StatusUnauthorized, "Authorization header is missing", ""},
		// Test with invalid auth token
		{`{"name":"test", "password":"test"}`, http.StatusUnauthorized, "error", "invalid token"},
		// Test as user, creating a non-admin user
		{`{"name":"test-na-na", "password":"test", "isAdmin":false}`, http.StatusOK, `"isAdmin":false`, userToken},
		// Test as user, creating an admin user
		{`{"name":"test-na-a", "password":"test", "isAdmin":true}`, http.StatusOK, `"isAdmin":false`, userToken},
		// Test as admin creating a non-admin user
		{`{"name":"test-a-na", "password":"test", "isAdmin":false}`, http.StatusOK, `"isAdmin":false`, adminToken},
		// Test as admin creating an admin user
		{`{"name":"test-a-a", "password":"test", "isAdmin":true}`, http.StatusOK, `"isAdmin":true`, adminToken},
	}

	for _, r := range requests {
		// Create a request to pass to our handler.
		req, err := http.NewRequest("POST", "/v1/user/", strings.NewReader(r.body))

		if r.token != "" {
			req.Header.Set("Authorization", "Bearer "+r.token)
			t.Logf("Request with token : %v", r.token)
		}

		if err != nil {
			t.Fatal(err)
		}

		w := mockRequest(req)
		checkResponse(t, w, r.wantedCode, r.bodyIncl)
	}
}

// Test the UpdateUser function as an admin and non-admin
// A non-admin can't update it's own admin status
// A non-admin can't update other users
func TestUpdateUser(t *testing.T) {
	// Get the auth token for the user
	userToken := getAndParseToken(t, `{"name":"testLogin","password":"test1"}`)
	t.Logf("User token : %v", userToken)
	// Get the auth token for the admin
	adminToken := getAndParseToken(t, `{"name":"testAdmin","password":"test1"}`)
	t.Logf("Admin token : %v", adminToken)

	// Set the different requests to test
	requests := []struct {
		id         string
		body       string
		wantedCode int
		bodyIncl   string
		token      string
	}{
		// Test without auth token
		{"39c71853-6206-4eef-9f5b-7a1a90830ccc", `{"isAdmin":true}`, http.StatusUnauthorized, "Authorization header is missing", ""},
		// Test with invalid auth token
		{"39c71853-6206-4eef-9f5b-7a1a90830ccc", `{"isAdmin":true}`, http.StatusUnauthorized, "error", "invalid token"},
		// Test as user, promoting itself to admin
		{"39c71853-6206-4eef-9f5b-7a1a90830ccc", `{"isAdmin":true}`, http.StatusOK, `"isAdmin":false`, userToken},
		// Test as user, changing password
		{"39c71853-6206-4eef-9f5b-7a1a90830ccc", `{"password":"$2y$10$/bWLJytUWkDIVYJsZTvg4O6.fTtBjDw9BOx50YSER2Iqo0dEGc2Be"}`, http.StatusOK, `"password":"$2y$10$/bWLJytUWkDIVYJsZTvg4O6.fTtBjDw9BOx50YSER2Iqo0dEGc2Be"`, userToken},
		// Test as user, changing another user
		{"ec4e2897-4ca4-4694-94d7-96db81ec223f", `{"password":"newPwd"}`, http.StatusUnauthorized, `Unauthorized`, userToken},
		// Test as admin, changing password
		{"ec4e2897-4ca4-4694-94d7-96db81ec223f", `{"password":"$2y$10$/bWLJytUWkDIVYJsZTvg4O6.fTtBjDw9BOx50YSER2Iqo0dEGc2Be"}`, http.StatusOK, `"password":"$2y$10$/bWLJytUWkDIVYJsZTvg4O6.fTtBjDw9BOx50YSER2Iqo0dEGc2Be"`, adminToken},
		// Test as admin, promoting another user to admin
		{"39c71853-6206-4eef-9f5b-7a1a90830ccc", `{"isAdmin":true}`, http.StatusOK, `"isAdmin":true`, adminToken},
		// Test as admin, downgrading another admin to user
		{"39c71853-6206-4eef-9f5b-7a1a90830ccc", `{"isAdmin":false}`, http.StatusOK, `"isAdmin":false`, adminToken},
	}

	for _, r := range requests {
		// Create a request to pass to our handler.
		req, err := http.NewRequest("PUT", "/v1/user/"+r.id, strings.NewReader(r.body))

		if r.token != "" {
			req.Header.Set("Authorization", "Bearer "+r.token)
			t.Logf("Request with token : %v", r.token)
		}

		if err != nil {
			t.Fatal(err)
		}

		w := mockRequest(req)
		checkResponse(t, w, r.wantedCode, r.bodyIncl)
	}
}

// Test the DeleteUser function as an admin and non-admin
// A non-admin can't delete other users
func TestDeleteUser(t *testing.T) {
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
		// Test as user, deleting another user
		{"ec4e2897-4ca4-4694-94d7-96db81ec223f", http.StatusUnauthorized, `Unauthorized`, userToken},
		// Test as user, deleting itself
		{"39c71853-6206-4eef-9f5b-7a1a90830ccc", http.StatusNoContent, "", userToken},
		// Test as admin, deleting non existing user
		{"39c71853-6206-4eef-9f5b-7a1a90830ccc", http.StatusNotFound, `can't find user`, adminToken},
		// Test as admin, deleting another user
		{"eee42492-e135-4c9a-89dc-923b9239b816", http.StatusNoContent, "", adminToken},
		// Test as admin, deleting itself
		{"ec4e2897-4ca4-4694-94d7-96db81ec223f", http.StatusNoContent, "", adminToken},
	}

	for _, r := range requests {
		// Create a request to pass to our handler.
		req, err := http.NewRequest("DELETE", "/v1/user/"+r.id, nil)

		if r.token != "" {
			req.Header.Set("Authorization", "Bearer "+r.token)
			t.Logf("Request with token : %v", r.token)
		}

		if err != nil {
			t.Fatal(err)
		}

		w := mockRequest(req)
		checkResponse(t, w, r.wantedCode, r.bodyIncl)
	}
}
