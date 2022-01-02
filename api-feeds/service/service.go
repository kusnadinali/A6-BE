package service

import (
	"context"

	"be-feeds/datastruct"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

type (
	Service interface {
		GetFeeds(ctx context.Context, id string, endRow string) ([]datastruct.Feeds, error)
	}

	service struct {
		repository datastruct.FeedsRepository
		logger     log.Logger
	}
)

func NewService(repo datastruct.FeedsRepository, logger log.Logger) Service {
	return &service{
		repository: repo,
		logger:     logger,
	}
}

func (s service) GetFeeds(ctx context.Context, id string, endRow string) ([]datastruct.Feeds, error) {
	logger := log.With(s.logger, "method", "GetUser")

	dataFeed, err := s.repository.GetFeeds(ctx, id, endRow)

	if err != nil {
		level.Error(logger).Log("err", err)
		return nil, err
	}

	logger.Log("Get Feeds", "userID: "+id+", MaxData: "+endRow)

	return dataFeed, nil
}
