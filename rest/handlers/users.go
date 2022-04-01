package handlers

import (
	"github.com/abbos-ron2/go-medium/models"
	"github.com/gin-gonic/gin"
)

// @Summary Create a new user
// @Description Create a new user
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body models.User true "User"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /users [post]
func (h *handler) CreateUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBind(&user); err != nil {
		return
	}

	if err := h.storage.CreateUser(c, user); err != nil {
		h.handleError(c, err)
		return
	}

	h.handleSuccess(c, "user created")
}

// @Summary Get a user
// @Description Get a user
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /users/{id} [get]
func (h *handler) GetUser(c *gin.Context) {
	var user models.User

	err := h.storage.GetUser(c, c.Param("id"), &user)
	if err != nil {
		h.handleError(c, err)
		return
	}
	h.handleSuccess(c, user)
}

// @Summary Get user by email
// @Description Get user by email
// @Tags users
// @Accept  json
// @Produce  json
// @Param email path string true "User email"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /users/email/{email} [get]
func (h *handler) GetUserByEmail(c *gin.Context) {
	var user models.User

	err := h.storage.GetUserByEmail(c, c.Param("email"), &user)
	if err != nil {
		h.handleError(c, err)
		return
	}
	h.handleSuccess(c, user)
}

// @Summary Get all users
// @Description Get all users
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /users [get]
func (h *handler) GetAllUsers(c *gin.Context) {
	var users []models.User

	err := h.storage.GetAllUsers(c, &users)
	if err != nil {
		h.handleError(c, err)
		return
	}

	h.handleSuccess(c, users)
}
