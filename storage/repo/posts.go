package repo

import (
	"context"

	"github.com/abbos-ron2/go-medium/models"
)

type PostI interface {
	CreatePost(ctx context.Context, post models.Post) error
	GetPost(ctx context.Context, id string, post *models.Post) error
	GetAllPosts(ctx context.Context, posts *[]models.Post) error
	GetPostsByUser(ctx context.Context, userID string, posts *[]models.Post) error
	GetPostComments(ctx context.Context, postID string, comments *[]models.Comment) error
	GetPostLikes(ctx context.Context, postID string, likes *[]models.Like) error
	GetPostLikesCount(ctx context.Context, postID string, likesCount *int) error
}
