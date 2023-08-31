package http

import "github.com/gin-gonic/gin"

func GetRoutes(r *gin.Engine, h *Handler) {
	v1 := r.Group("/v1")
	addUserRoutes(v1, h)
	addAuthRoutes(v1, h)
}

func addUserRoutes(rg *gin.RouterGroup, h *Handler) {
	user := rg.Group("/user")
	user.Use(authMiddleware())

	user.GET("/", authorizeAdmin(), h.UserHandler.GetAll)
	user.GET("/:id", h.UserHandler.Get)
	user.POST("/", h.UserHandler.Create)
	user.PUT("/:id", h.UserHandler.Update)
	user.DELETE("/:id", h.UserHandler.Delete)
}

func addAuthRoutes(rg *gin.RouterGroup, h *Handler) {
	auth := rg.Group("/auth")

	auth.POST("/login", h.AuthHandler.Login)
	auth.POST("/register", h.UserHandler.Create)
}
