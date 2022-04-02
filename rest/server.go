package rest

import (
	"fmt"
	"net/http"

	_ "github.com/abbos-ron2/go-medium/api/docs"
	"github.com/abbos-ron2/go-medium/config"
	"github.com/abbos-ron2/go-medium/rest/handlers"
	"github.com/abbos-ron2/go-medium/storage"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title           Swagger API
// @version         1.0
// @description     Social media pet project.
// @contact.name   API Support
// @contact.url    t.me/rON2_webdev
// @contact.email  abbosamritdidnov@gmail.com
// @host      localhost:8080
// @BasePath  /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func New(cfg config.Config, s storage.StorageI) (srv *http.Server) {
	h := handlers.NewHandler(cfg, s)

	r := gin.Default()
	r.POST("login", h.Login)
	r.POST("/register", h.CreateUser)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	users := r.Group("/users").Use(h.AuthMiddleware)
	{
		users.GET("/:id", h.GetUser)
		users.GET("/email/:email", h.GetUserByEmail)
		users.GET("/", h.GetAllUsers)
		users.GET("/:id/posts", h.GetPostsByUser)
	}
	posts := r.Group("/posts").Use(h.AuthMiddleware)
	{
		posts.POST("/", h.CreatePost)
		posts.GET("/:id", h.GetPost)
		posts.GET("/", h.GetAllPosts)
		posts.GET("/:id/comments", h.GetPostComments)
		posts.GET("/:id/likes", h.GetPostLikes)
		posts.GET("/:id/likes_count", h.GetPostLikesCount)
	}

	likes := r.Group("/likes").Use(h.AuthMiddleware)
	{
		likes.POST("/", h.CreateLike)
		likes.DELETE("/:post_id/:user_id", h.DeleteLike)
	}

	comments := r.Group("/comments").Use(h.AuthMiddleware)
	{
		comments.POST("/", h.CreateComment)
	}

	srv = &http.Server{
		Addr:    cfg.HTTPPort,
		Handler: r,
	}
	fmt.Println("Running on port " + cfg.HTTPPort)
	return
}
