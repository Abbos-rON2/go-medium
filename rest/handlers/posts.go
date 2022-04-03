package handlers

import (
	"github.com/abbos-ron2/go-medium/models"
	"github.com/gin-gonic/gin"
)

// Posts

// @Summary Create a new post
// @Description Create a new post
// @Tags posts
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param post body models.CreatePostRequest true "Post"
// @Success 200 {object} models.Response{data=models.CreatePostRequest}
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /posts [post]
func (h *handler) CreatePost(c *gin.Context) {
	var post models.CreatePostRequest

	if err := c.ShouldBind(&post); err != nil {
		return
	}
	post.AuthorID = c.GetInt("user_id")

	if err := h.storage.Post().CreatePost(c, post); err != nil {
		h.handleError(c, err)
		return
	}

	h.handleSuccess(c, post)
}

// @Summary Get a post
// @Description Get a post
// @Tags posts
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param id path string true "Post ID"
// @Success 200 {object} models.Response{data=models.Post}
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /posts/{id} [get]
func (h *handler) GetPost(c *gin.Context) {
	var post models.Post

	err := h.storage.Post().GetPost(c, c.Param("id"), &post)
	if err != nil {
		h.handleError(c, err)
		return
	}
	h.handleSuccess(c, post)
}

// @Summary Get all posts
// @Description Get all posts
// @Tags posts
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {object} models.Response{data=[]models.Post}
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /posts [get]
func (h *handler) GetAllPosts(c *gin.Context) {
	var posts []models.Post

	err := h.storage.Post().GetAllPosts(c, &posts)
	if err != nil {
		h.handleError(c, err)
		return
	}
	h.handleSuccess(c, posts)
}

// @Summary Get post by user ID
// @Description Get post by user ID
// @Tags posts
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param id path string true "User ID"
// @Success 200 {object} models.Response{data=[]models.Post}
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /users/{id}/posts [get]
func (h *handler) GetPostsByUser(c *gin.Context) {
	var posts []models.Post

	err := h.storage.Post().GetPostsByUser(c, c.Param("id"), &posts)
	if err != nil {
		h.handleError(c, err)
		return
	}
	h.handleSuccess(c, posts)
}

// @Summary Get post comments
// @Description Get post comments
// @Tags posts
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param id path string true "Post ID"
// @Success 200 {object} models.Response{data=[]models.Comment}
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /posts/{id}/comments [get]
func (h *handler) GetPostComments(c *gin.Context) {
	var comments []models.Comment

	err := h.storage.Post().GetPostComments(c, c.Param("id"), &comments)
	if err != nil {
		h.handleError(c, err)
		return
	}
	h.handleSuccess(c, comments)
}

// @Summary Get post likes
// @Description Get post likes
// @Tags posts
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param id path string true "Post ID"
// @Success 200 {object} models.Response{data=[]models.Like}
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /posts/{id}/likes [get]
func (h *handler) GetPostLikes(c *gin.Context) {
	var likes []models.Like

	err := h.storage.Post().GetPostLikes(c, c.Param("id"), &likes)
	if err != nil {
		h.handleError(c, err)
		return
	}

	h.handleSuccess(c, likes)
}

// @Summary Get post likes count
// @Description Get post likes count
// @Tags posts
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param id path string true "Post ID"
// @Success 200 {object} models.Response{data=models.Count}
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /posts/{id}/likes_count [get]
func (h *handler) GetPostLikesCount(c *gin.Context) {
	var count models.Count

	err := h.storage.Post().GetPostLikesCount(c, c.Param("id"), &count)
	if err != nil {
		h.handleError(c, err)
		return
	}
	h.handleSuccess(c, count)
}
