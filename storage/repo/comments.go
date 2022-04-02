package repo

import (
	"context"

	"github.com/abbos-ron2/go-medium/models"
	"github.com/jackc/pgx/v4"
)

type commentRepo struct {
	db *pgx.Conn
}
type CommentI interface {
	CreateComment(ctx context.Context, comment models.CreateCommentRequest) error
}

func NewCommentRepo(db *pgx.Conn) CommentI {
	return &commentRepo{db: db}
}

func (r *commentRepo) CreateComment(ctx context.Context, comment models.CreateCommentRequest) error {
	_, err := r.db.Exec(ctx, `
		INSERT INTO comments (
			post_id,
			user_id,
			content,
			reply_id,
		) VALUES (
			$1,
			$2,
			$3,
			$4, 
			$5
		)
	`, comment.PostID, comment.UserID, comment.Content, comment.ReplyID)
	return err
}
