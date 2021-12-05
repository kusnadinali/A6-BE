package feeds

import "context"

type Service interface {
	GetPost(ctx context.Context, post_id string) (string, error)
}
