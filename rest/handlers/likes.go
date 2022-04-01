package handlers

import (
	"net/http"

	"github.com/abbos-ron2/go-medium/models"
	"github.com/gin-gonic/gin"
)

func (h *handler) CreateLike(c *gin.Context) {
	var like models.Like

	if err := c.ShouldBind(&like); err != nil {
		return
	}

	if err := h.storage.CreateLike(c, like); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "like created",
	})
}
func (h *handler) DeleteLike(c *gin.Context) {
	postId := c.Param("post_id")
	userId := c.Param("user_id")

	if err := h.storage.DeleteLike(c, postId, userId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "like deleted",
	})
}
