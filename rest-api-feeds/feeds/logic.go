package feeds

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

type service struct {
	repository Repository
	logger     log.Logger
}

func NewService(rep Repository, logger log.Logger) Service {
	return &service{
		repository: rep,
		logger:     logger,
	}
}

func (s service) GetPost(ctx context.Context, post_id string) (string, error) {
	logger := log.With(s.logger, "method", "GetPost")

	username, err := s.repository.GetPost(ctx, post_id)

	if err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}

	logger.Log("Get Post", post_id)
	return username, nil
}
