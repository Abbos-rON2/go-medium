package handlers

import (
	"strconv"

	"github.com/abbos-ron2/go-medium/models"
	"github.com/gin-gonic/gin"
)

// @Summary Subscribe to user
// @Description Subscribe to user
// @Tags subscriptions
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param target_id path string true "Target ID"
// @Success 200 {object} models.Response{data=string}
// @Failure 400 {object} models.Response
// @Failure 404 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /subscriptions/{target_id} [post]
func (h *handler) Subscribe(c *gin.Context) {
	userID := strconv.Itoa(c.GetInt("user_id"))
	targetID := c.Param("target_id")

	if err := h.storage.Subscription().Subscribe(c.Request.Context(), userID, targetID); err != nil {
		h.handleError(c, err)
		return
	}

	h.handleSuccess(c, "subscribed")
}

// @Summary Unsubscribe from user
// @Description Unsubscribe from user
// @Tags subscriptions
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param target_id path string true "Target ID"
// @Success 200 {object} models.Response{data=string}
// @Failure 400 {object} models.Response
// @Failure 404 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /subscriptions/{target_id} [delete]
func (h *handler) Unsubscribe(c *gin.Context) {
	userID := strconv.Itoa(c.GetInt("user_id"))
	targetID := c.Param("target_id")

	if err := h.storage.Subscription().Unsubscribe(c.Request.Context(), userID, targetID); err != nil {
		h.handleError(c, err)

		return
	}

	h.handleSuccess(c, "unsubscribed")
}

// @Summary Get user subscribers
// @Description Get user subscribers
// @Tags subscriptions
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param user_id path string true "User ID"
// @Success 200 {object} models.Response{data=[]models.Subscription}
// @Failure 400 {object} models.Response
// @Failure 404 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /subscriptions/{user_id}/subscribers [get]
func (h *handler) GetSubscribers(c *gin.Context) {
	userID := c.Param("user_id")
	subscribers := []models.Subscription{}

	if err := h.storage.Subscription().GetSubscribers(c.Request.Context(), userID, &subscribers); err != nil {
		h.handleError(c, err)
		return
	}

	h.handleSuccess(c, subscribers)
}

// @Summary Get user subscribers count
// @Description Get user subscribers count
// @Tags subscriptions
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param user_id path string true "User ID"
// @Success 200 {object} models.Response{data=models.Count}
// @Failure 400 {object} models.Response
// @Failure 404 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /subscriptions/{user_id}/subscribers_count [get]
func (h *handler) GetSubscribersCount(c *gin.Context) {
	userID := c.Param("user_id")
	count := 0

	if err := h.storage.Subscription().GetSubscribersCount(c.Request.Context(), userID, &count); err != nil {
		h.handleError(c, err)
		return
	}

	h.handleSuccess(c, count)
}

// @Summary Get user subscriptions
// @Description Get user subscriptions
// @Tags subscriptions
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param user_id path string true "User ID"
// @Success 200 {object} models.Response{data=[]models.Subscription}
// @Failure 400 {object} models.Response
// @Failure 404 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /subscriptions/{user_id}/subscriptions [get]
func (h *handler) GetSubscriptions(c *gin.Context) {
	userID := c.Param("user_id")
	subscriptions := []models.Subscription{}

	if err := h.storage.Subscription().GetSubscriptions(c.Request.Context(), userID, &subscriptions); err != nil {
		h.handleError(c, err)
		return
	}

	h.handleSuccess(c, subscriptions)
}

// @Summary Get user subscriptions count
// @Description Get user subscriptions count
// @Tags subscriptions
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param user_id path string true "User ID"
// @Success 200 {object} models.Response{data=models.Count}
// @Failure 400 {object} models.Response
// @Failure 404 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /subscriptions/{user_id}/subscriptions_count [get]
func (h *handler) GetSubscriptionsCount(c *gin.Context) {
	userID := c.Param("user_id")
	count := 0

	if err := h.storage.Subscription().GetSubscriptionsCount(c.Request.Context(), userID, &count); err != nil {
		h.handleError(c, err)
		return
	}

	h.handleSuccess(c, count)
}
