package http

import "github.com/gin-gonic/gin"

func GetRoutes(r *gin.Engine, h *Handler) {
	v1 := r.Group("/v1")
	addUserRoutes(v1, h)
}

func addUserRoutes(rg *gin.RouterGroup, h *Handler) {
	user := rg.Group("/user")

	user.GET("/", h.GetUser)
}