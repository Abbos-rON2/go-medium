package rest

import (
	"net/http"

	"github.com/abbos-ron2/go-medium/config"
	"github.com/gin-gonic/gin"
)

func New(cfg config.Config) (srv *http.Server) {
	r := gin.Default()

	group := r.Group("/")

	h := newHandler(cfg)

	group.GET("/", h.Ping)

	srv = &http.Server{
		Addr:    cfg.HTTPPort,
		Handler: r,
	}

	return
}
