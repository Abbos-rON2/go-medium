package repo

import (
	"context"

	"github.com/abbos-ron2/go-medium/models"
	"github.com/jackc/pgx/v4"
)

type postRepo struct {
	db *pgx.Conn
}
type PostI interface {
	CreatePost(ctx context.Context, post models.CreatePostRequest) error
	GetPost(ctx context.Context, id string, post *models.Post) error
	GetAllPosts(ctx context.Context, posts *[]models.Post) error
	GetPostsByUser(ctx context.Context, userID string, posts *[]models.Post) error
	GetPostComments(ctx context.Context, postID string, comments *[]models.Comment) error
	GetPostLikes(ctx context.Context, postID string, likes *[]models.Like) error
	GetPostLikesCount(ctx context.Context, postID string, likesCount *models.Count) error

	DeletePost(ctx context.Context, id string) error
}

func NewPostRepo(db *pgx.Conn) PostI {
	return &postRepo{
		db: db,
	}
}

// Posts
func (r *postRepo) CreatePost(ctx context.Context, post models.CreatePostRequest) error {
	_, err := r.db.Exec(
		ctx,
		`INSERT INTO posts (title, content, author_id) VALUES ($1, $2, $3)`,
		post.Title,
		post.Content,
		post.AuthorID,
	)
	return err
}
func (r *postRepo) GetPost(ctx context.Context, id string, post *models.Post) error {
	return r.db.QueryRow(
		ctx,
		`SELECT id, title, content, author_id, created_at, updated_at FROM posts WHERE id = $1`,
		id,
	).Scan(
		&post.ID,
		&post.Title,
		&post.Content,
		&post.AuthorID,
		&post.CreatedAt,
		&post.UpdatedAt,
	)
}
func (r *postRepo) GetAllPosts(ctx context.Context, posts *[]models.Post) error {
	rows, err := r.db.Query(
		ctx,
		`SELECT id, title, content, author_id, created_at, updated_at FROM posts`,
	)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var post models.Post
		if err := rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.CreatedAt,
			&post.UpdatedAt,
		); err != nil {
			return err
		}
		*posts = append(*posts, post)
	}
	return nil
}
func (r *postRepo) GetPostsByUser(ctx context.Context, userID string, posts *[]models.Post) error {
	rows, err := r.db.Query(
		ctx,
		`SELECT id, title, content, author_id, created_at, updated_at FROM posts WHERE author_id = $1`,
		userID,
	)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var post models.Post
		if err := rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.CreatedAt,
			&post.UpdatedAt,
		); err != nil {
			return err
		}
		*posts = append(*posts, post)
	}
	return nil
}

// Comments
func (r *postRepo) GetPostComments(ctx context.Context, postID string, comments *[]models.Comment) error {
	rows, err := r.db.Query(
		ctx,
		`SELECT id, author_id, content, reply_id, created_at FROM comments WHERE post_id = $1`,
		postID,
	)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var comment models.Comment
		var replyID *int
		if err := rows.Scan(
			&comment.ID,
			&comment.AuthorID,
			&comment.Content,
			&replyID,
			&comment.CreatedAt,
		); err != nil {
			return err
		}
		if replyID != nil {
			comment.ReplyID = *replyID
		}
		*comments = append(*comments, comment)
	}
	return nil
}

//Likes
func (r *postRepo) GetPostLikes(ctx context.Context, postID string, likes *[]models.Like) error {
	rows, err := r.db.Query(ctx, `
		SELECT id, post_id, user_id, created_at FROM likes WHERE post_id = $1
	`, postID)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var like models.Like
		if err := rows.Scan(
			&like.ID,
			&like.PostID,
			&like.UserID,
			&like.CreatedAt,
		); err != nil {
			return err
		}
		*likes = append(*likes, like)
	}
	return nil
}
func (r *postRepo) GetPostLikesCount(ctx context.Context, postID string, count *models.Count) error {

	if err := r.db.QueryRow(ctx, `
		SELECT COUNT(*) FROM likes WHERE post_id = $1
	`, postID).Scan(&count.Count); err != nil {
		return err
	}
	return nil
}

func (r *postRepo) DeletePost(ctx context.Context, id string) error {
	_, err := r.db.Exec(
		ctx,
		`DELETE FROM posts WHERE id = $1`,
		id,
	)
	return err
}
