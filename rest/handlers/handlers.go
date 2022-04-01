package handlers

import (
	"net/http"
	"time"

	"github.com/abbos-ron2/go-medium/config"
	"github.com/abbos-ron2/go-medium/models"
	"github.com/abbos-ron2/go-medium/storage"
	"github.com/dgrijalva/jwt-go"
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

func (h *handler) Login(c *gin.Context) {
	var user models.User

	if err := c.ShouldBind(&user); err != nil {
		return
	}
	// create jwt token

	dbUser := models.User{}
	if err := h.storage.GetUserByEmail(c, user.Email, &dbUser); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	if user.Password != dbUser.Password {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid password",
		})
		return
	}

	// create jwt token
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &models.Token{
		UserID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(h.cfg.JwtSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Header("Authorization", tokenString)

	c.JSON(200, user)
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

func (h *handler) AuthMiddleware(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "no token",
		})
		c.Abort()
		return
	}
	claims := &models.Token{}
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		c.Abort()
		return
	}
	if !tkn.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid token",
		})
		c.Abort()
		return
	}
	c.Next()
}
