package feeds

import (
	"context"
	"database/sql"
	"errors"

	"github.com/go-kit/kit/log"
)

var RepoErr = errors.New("Unable to handle Repo Request")

type repo struct {
	db     *sql.DB
	logger log.Logger
}

func NewRepo(db *sql.DB, logger log.Logger) Repository {
	return &repo{
		db:     db,
		logger: log.With(logger, "repo", "sql"),
	}
}

func (repo *repo) GetPost(ctx context.Context, post_id string) (string, error) {
	var username string
	err := repo.db.QueryRow("SELECT username FROM postingan WHERE post_id=$2", post_id)

	if err != nil {
		return "", RepoErr
	}

	return username, nil
}
