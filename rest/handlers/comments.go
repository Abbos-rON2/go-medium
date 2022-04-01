package handlers

import (
	"net/http"

	"github.com/abbos-ron2/go-medium/models"
	"github.com/gin-gonic/gin"
)

func (h *handler) CreateComment(c *gin.Context) {
	var comment models.Comment

	if err := c.ShouldBind(&comment); err != nil {
		return
	}

	if err := h.storage.CreateComment(c, comment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "comment created",
	})
}
