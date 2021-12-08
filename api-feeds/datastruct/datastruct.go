package datastruct

import (
	"context"
)

type Feeds struct {
	Username     string `json:"username"`
	Image_file   string `json:"image_file"`
	Feed_id      int    `json:"feed_id"`
	User_id      int    `json:"user_id"`
	Image_feed   string `json:"image_feed"`
	Caption_feed string `json:"caption_feed"`
	Post_date    string `json:"post_date"`
}

type FeedsRepository interface {
	GetFeeds(ctx context.Context) ([]Feeds, error)
}
