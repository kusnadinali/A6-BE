package repository

import (
	"context"
	"database/sql"
	"tes/datastruct"

	"github.com/go-kit/kit/log"
)

type repo struct {
	db     *sql.DB
	logger log.Logger
}

// NewRepo handle all db operation
func NewRepo(db *sql.DB, logger log.Logger) datastruct.FeedsRepository {
	return &repo{
		db:     db,
		logger: log.With(logger, "repo", "postgres"),
	}
}

const (
	queryGetFeeds = "select z.username,z.image_file,f.user_id,f.feed_id,f.image_feed,f.caption_feed, to_char(post_date, 'DD-MM-YYYY HH24:MI:SS') as date from friends x join members y on y.user_id=x.user_id join members z on z.user_id = x.following_id join feeds f on z.user_id=f.user_id where y.user_id = 1 order by post_date desc"
)

func (repo *repo) GetFeeds(ctx context.Context) ([]datastruct.Feeds, error) {
	var feeds []datastruct.Feeds
	var err error

	rows, err := repo.db.QueryContext(ctx, queryGetFeeds)
	if err != nil {
		//http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return nil, err
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
			&feed.Post_date)

		if err != nil {
			//log.Println(err)
			//http.Error(w, http.StatusText(500), 500)
			//log.Fatal(err)
			return nil, err
		}

		feeds = append(feeds, feed)
	}

	if err = rows.Err(); err != nil {
		// http.Error(w, http.StatusText(500), 500)
		// log.Fatal(err)
		return nil, err
	}

	return feeds, nil
}
