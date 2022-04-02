package handlers

import (
	"github.com/abbos-ron2/go-medium/models"
	"github.com/gin-gonic/gin"
)

// @Summary Create like
// @Description Create like
// @Tags likes
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param like body models.CreateLikeRequest true "like"
// @Success 200 {object} models.Response{data=string}
// @Failure 400 {object} models.Response
// @Failure 404 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /likes [post]
func (h *handler) CreateLike(c *gin.Context) {
	var like models.CreateLikeRequest

	if err := c.ShouldBind(&like); err != nil {
		return
	}

	if err := h.storage.CreateLike(c, like); err != nil {
		h.handleError(c, err)
		return
	}

	h.handleSuccess(c, "like created")
}

// @Summary Delete like
// @Description Delete like
// @Tags likes
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param post_id path string true "post_id"
// @Param user_id path string true "user_id"
// @Success 200 {object} models.Response{data=string}
// @Failure 400 {object} models.Response
// @Failure 404 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /likes/{post_id}/{user_id} [delete]
func (h *handler) DeleteLike(c *gin.Context) {
	postId := c.Param("post_id")
	userId := c.Param("user_id")

	if err := h.storage.DeleteLike(c, postId, userId); err != nil {
		h.handleError(c, err)
		return
	}

	h.handleSuccess(c, "like deleted")
}
