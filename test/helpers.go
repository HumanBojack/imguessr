package test

import (
	ihttp "imguessr/pkg/http"
	"imguessr/pkg/service"
	"net/http"
	"net/http/httptest"

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
