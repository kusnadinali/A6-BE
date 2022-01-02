package service

import (
	"context"
	"database/sql"
	"errors"
	"net/http"

	"be-feeds/datastruct"

	"github.com/go-kit/kit/log"
)

const (
	queryGetFeeds = "SELECT * FROM GetFeedsData($1, $2)"
)

var ErrFoo = errors.New(http.StatusText(500))

type repo struct {
	db     *sql.DB
	logger log.Logger
}

func NewRepo(db *sql.DB, logger log.Logger) datastruct.FeedsRepository {
	return &repo{
		db:     db,
		logger: log.With(logger, "repo", "postgres"),
	}
}

func (repo *repo) GetFeeds(ctx context.Context, id string, endRow string) ([]datastruct.Feeds, error) {
	var feeds []datastruct.Feeds

	rows, err := repo.db.QueryContext(ctx, queryGetFeeds, id, endRow)
	if err != nil {
		return nil, ErrFoo
	}

	defer rows.Close()

	for rows.Next() {
		var feed datastruct.Feeds
		err := rows.Scan(
			&feed.Username,
			&feed.Image_file,
			&feed.User_id,
			&feed.Feed_id,
			&feed.Image_feed,
			&feed.Caption_feed,
			&feed.Post_date,
			&feed.Total_Likes,
			&feed.Total_Comments)

		if err != nil {
			return nil, ErrFoo
		}

		feeds = append(feeds, feed)
	}

	if err = rows.Err(); err != nil {
		return nil, ErrFoo
	}

	return feeds, nil
}
