package storage

import (
	"context"

	"github.com/abbos-ron2/go-medium/models"
)

// Posts
func (s *storage) CreatePost(ctx context.Context, post models.Post) error {
	_, err := s.db.Exec(
		ctx,
		`INSERT INTO posts (title, content, author_id) VALUES ($1, $2, $3)`,
		post.Title,
		post.Content,
		post.AuthorID,
	)
	return err
}
func (s *storage) GetPost(ctx context.Context, id string, post *models.Post) error {
	return s.db.QueryRow(
		ctx,
		`SELECT id, title, content, author_id, created_at FROM posts WHERE id = $1`,
		id,
	).Scan(
		&post.ID,
		&post.Title,
		&post.Content,
		&post.AuthorID,
		&post.CreatedAt,
	)
}
func (s *storage) GetAllPosts(ctx context.Context, posts *[]models.Post) error {
	rows, err := s.db.Query(
		ctx,
		`SELECT id, title, content, author_id, created_at FROM posts`,
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
		); err != nil {
			return err
		}
		*posts = append(*posts, post)
	}
	return nil
}
func (s *storage) GetPostsByUser(ctx context.Context, userID string, posts *[]models.Post) error {
	rows, err := s.db.Query(
		ctx,
		`SELECT id, title, content, author_id, created_at FROM posts WHERE author_id = $1`,
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
		); err != nil {
			return err
		}
		*posts = append(*posts, post)
	}
	return nil
}

// Comments
func (s *storage) GetPostComments(ctx context.Context, postID string, comments *[]models.Comment) error {
	rows, err := s.db.Query(
		ctx,
		`SELECT id, user_id, content, reply_id, created_at FROM comments WHERE post_id = $1`,
		postID,
	)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var comment models.Comment
		if err := rows.Scan(
			&comment.ID,
			&comment.UserID,
			&comment.Content,
			&comment.ReplyID,
			&comment.CreatedAt,
		); err != nil {
			return err
		}
		*comments = append(*comments, comment)
	}
	return nil
}

//Likes
func (s *storage) GetPostLikes(ctx context.Context, postID string, likes *[]models.Like) error {
	rows, err := s.db.Query(ctx, `
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
func (s *storage) GetPostLikesCount(ctx context.Context, postID string, count *int) error {

	if err := s.db.QueryRow(ctx, `
		SELECT COUNT(*) FROM likes WHERE post_id = $1
	`, postID).Scan(&count); err != nil {
		return err
	}
	return nil
}
