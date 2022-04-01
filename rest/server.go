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

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth
func New(cfg config.Config, s storage.StorageI) (srv *http.Server) {
	h := handlers.NewHandler(cfg, s)

	r := gin.Default()
	// r.Use(AuthMiddleware)
	r.GET("login", h.Login)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	users := r.Group("/users")
	{
		users.POST("/", h.CreateUser)
		users.GET("/:id", h.GetUser)
		users.GET("/email/:email", h.GetUserByEmail)
		users.GET("/", h.GetAllUsers)
		users.GET("/:id/posts", h.GetPostsByUser)

	}
	posts := r.Group("/posts")
	{
		posts.POST("/", h.CreatePost)
		posts.GET("/:id", h.GetPost)
		posts.GET("/", h.GetAllPosts)
		posts.GET("/:id/comments", h.GetPostComments)
		posts.GET("/:id/likes", h.GetPostLikes)
		posts.GET("/:id/likes_count", h.GetPostLikesCount)
	}
	likes := r.Group("/likes")
	{
		likes.POST("/", h.CreateLike)
		likes.DELETE("/:post_id/:user_id", h.DeleteLike)
	}
	comments := r.Group("/comments")
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

// func (h *handler) AuthMiddleware(c *gin.Context) {
// 	token := c.Request.Header.Get("Authorization")
// 	if token == "" {
// 		c.JSON(http.StatusUnauthorized, gin.H{
// 			"error": "no token",
// 		})
// 		c.Abort()
// 		return
// 	}
// 	claims := &models.Token{}
// 	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
// 		return []byte("secret"), nil
// 	})
// 	if err != nil {
// 		c.JSON(http.StatusUnauthorized, gin.H{
// 			"error": err.Error(),
// 		})
// 		c.Abort()
// 		return
// 	}
// 	if !tkn.Valid {
// 		c.JSON(http.StatusUnauthorized, gin.H{
// 			"error": "invalid token",
// 		})
// 		c.Abort()
// 		return
// 	}
// 	c.Next()
// }
