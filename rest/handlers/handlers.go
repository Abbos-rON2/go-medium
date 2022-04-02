package handlers

import (
	"net/http"

	"github.com/abbos-ron2/go-medium/config"
	"github.com/abbos-ron2/go-medium/models"
	"github.com/abbos-ron2/go-medium/storage"
	"github.com/gin-gonic/gin"
)

type handler struct {
	cfg     config.Config
	storage storage.StorageI
}

func NewHandler(cfg config.Config, storage storage.StorageI) *handler {
	return &handler{
		cfg:     cfg,
		storage: storage,
	}
}

func (h *handler) handleSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, models.Response{
		Code:    http.StatusOK,
		Message: "success",
		Data:    data,
	})
}
func (h *handler) handleError(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, models.Response{
		Code:    http.StatusInternalServerError,
		Message: err.Error(),
	})
}
