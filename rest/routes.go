package rest

import (
	"github.com/abbos-ron2/go-medium/config"
	"github.com/gin-gonic/gin"
)

type handler struct {
	cfg config.Config
}

func newHandler(cfg config.Config) *handler {
	return &handler{
		cfg: cfg,
	}
}

func (h *handler) Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
