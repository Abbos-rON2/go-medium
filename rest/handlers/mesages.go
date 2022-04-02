package handlers

import (
	"github.com/abbos-ron2/go-medium/models"
	"github.com/gin-gonic/gin"
)

// @Summary Create a message
// @Description Create a message
// @Tags messages
// @Accept  json
// @Produce  json
// @Param message body models.CreateMessageRequest true "message"
// @Success 200 {object} models.Response{data=string}
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /messages [post]
func (h *handler) CreateMessage(c *gin.Context) {
	var message models.CreateMessageRequest
	if err := c.BindJSON(&message); err != nil {
		h.handleError(c, err)
		return
	}

	if err := h.storage.Message().CreateMessage(c.Request.Context(), message); err != nil {
		h.handleError(c, err)
		return
	}

	h.handleSuccess(c, "message created")
}

// @Summary Get a message
// @Description Get a message
// @Tags messages
// @Accept  json
// @Produce  json
// @Param id path string true "message id"
// @Success 200 {object} models.Response{data=models.Message}
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /messages/{id} [get]
func (h *handler) GetMessage(c *gin.Context) {
	var message models.Message
	if err := h.storage.Message().GetMessage(c.Request.Context(), c.Param("id"), &message); err != nil {
		h.handleError(c, err)
		return
	}

	h.handleSuccess(c, message)
}

// @Summary Get all messages
// @Description Get all messages
// @Tags messages
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Response{data=[]models.Message}
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /messages [get]
func (h *handler) GetAllMessages(c *gin.Context) {
	var messages []models.Message
	if err := h.storage.Message().GetAllMessages(c.Request.Context(), &messages); err != nil {
		h.handleError(c, err)
		return
	}

	h.handleSuccess(c, messages)
}

// @Summary Get all messages from a chat
// @Description Get all messages from a chat
// @Tags messages
// @Accept  json
// @Produce  json
// @Param chat_id path string true "chat id"
// @Success 200 {object} models.Response{data=[]models.Message}
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /chats/{chat_id}/messages [get]
func (h *handler) GetChatMessages(c *gin.Context) {
	var messages []models.Message
	if err := h.storage.Message().GetChatMessages(c.Request.Context(), c.Param("chat_id"), &messages); err != nil {
		h.handleError(c, err)
		return
	}

	h.handleSuccess(c, messages)
}
