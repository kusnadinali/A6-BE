package feeds

import (
	"context"
	"time"
)

type Feeds struct {
	post_id       string    `json:"post_id,omitempty"`
	username      string    `json:"username"`
	profile_image string    `json:"profile_image"`
	post_image    string    `json:"post_image"`
	post_caption  string    `json:"post_caption"`
	post_date     time.Time `json:"post_date"`
	total_like    int       `json:"total_like"`
	total_comment int       `json:"total_comment"`
}

type Repository interface {
	GetPost(ctx context.Context, post_id string) (string, error)
}
