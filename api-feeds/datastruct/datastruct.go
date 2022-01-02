package datastruct

import (
	"context"
)

type Feeds struct {
	User_id        int    `json:"user_id"`
	Username       string `json:"username"`
	Image_file     string `json:"image_file"`
	Feed_id        int    `json:"feed_id"`
	Image_feed     string `json:"image_feed"`
	Caption_feed   string `json:"caption_feed"`
	Post_date      string `json:"post_date"`
	Total_Likes    int64  `json:"total_likes"`
	Total_Comments int64  `json:"total_comments"`
}

type FeedsRepository interface {
	GetFeeds(ctx context.Context, id string, endRow string) ([]Feeds, error)
}
