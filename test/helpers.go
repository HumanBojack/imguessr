package test

import (
	ihttp "imguessr/pkg/http"
	"imguessr/pkg/service"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

// Mock a request to the server
func mockRequest(req *http.Request) *httptest.ResponseRecorder {
	// Create a router
	r := gin.Default()

	// Mock the database interface
	db := mockDB{}

	// Create Services
	uSvc := service.NewUserSvc(db)
	aSvc := service.NewAuthSvc()

	// Create a handler and add the routes
	h := ihttp.NewHandler(uSvc, aSvc)
	ihttp.GetRoutes(r, h)

	// Create a response recorder so you can inspect the response
	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)

	return w
}

// Get the generated token from the response body
func getAndParseToken(t *testing.T, requestBdy string) string {
	req, err := http.NewRequest("POST", "/v1/auth/login",
		strings.NewReader(requestBdy),
	)

	if err != nil {
		t.Fatal(err)
	}

	w := mockRequest(req)
	// Get the "token" from the body
	body, err := io.ReadAll(w.Body)
	if err != nil {
		t.Fatal(err)
	}
	bodyStr := string(body)
	token := strings.Split(bodyStr, ":")[1]
	token = strings.ReplaceAll(token, "\"", "")
	token = strings.ReplaceAll(token, "}", "")

	return token
}
