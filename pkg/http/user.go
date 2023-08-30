package http

import (
	"imguessr/pkg/domain"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	UserSvc domain.UserSvc
}

func (h *UserHandler) GetAll(c *gin.Context) {
	userList, err := h.UserSvc.GetAll()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusOK, userList)
}

func (h *UserHandler) Get(c *gin.Context) {
	id := c.Param("id")

	userId, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Not valid UUID")
		return
	}

	user, err := h.UserSvc.Get(userId.String())
	if err != nil {
		c.JSON(http.StatusNotFound, "Can't find user with this id")
		return
	}

	c.IndentedJSON(http.StatusOK, user)
}

func (h *UserHandler) Create(c *gin.Context) {
	// Get the email and the password from the body
	var newUser domain.User

	if err := c.BindJSON(&newUser); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Hash the password and create the User
	hash, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), 10)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	newUser.Password = string(hash)

	err = h.UserSvc.Create(&newUser)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, newUser)
}

func (h *UserHandler) Update(c *gin.Context) {
	// Get the user at the given id
	id := c.Param("id")

	userId, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Not valid UUID")
		return
	}

	user, err := h.UserSvc.Get(userId.String())
	if err != nil {
		c.JSON(http.StatusNotFound, "Can't find user with this id")
		return
	}

	// Get the request body containing the updated user and modify the current user
	err = c.BindJSON(&user.UpdateUser)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Save the modified user
	err = h.UserSvc.Update(user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) Delete(c *gin.Context) {
	// Get the user at the given id
	id := c.Param("id")

	err := h.UserSvc.Delete(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
	}

	c.Status(http.StatusNoContent)
}
