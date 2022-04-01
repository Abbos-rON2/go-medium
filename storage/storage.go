package storage

import (
	"github.com/abbos-ron2/go-medium/config"
	"github.com/abbos-ron2/go-medium/storage/repo"
	"github.com/jackc/pgx/v4"
)

type storage struct {
	cfg config.Config
	db  *pgx.Conn
}

type StorageI interface {
	repo.PostI
	repo.CommentI
	repo.UserI
	repo.LikeI
}

func New(cfg config.Config, db *pgx.Conn) *storage {
	return &storage{
		cfg: cfg,
		db:  db,
	}
}
