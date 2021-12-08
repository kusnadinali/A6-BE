package service

import (
	"context"
	"tes/datastruct"

	"github.com/go-kit/kit/log"
)

//In Go kit, we model a service as an interface.
// An Interface is an abstract type. Must implements.
// Interface describes all the methods of a method set and provides the signatures for each method.

type (
	Service interface {
		GetFeeds(ctx context.Context) ([]datastruct.Feeds, error)
	}

	service struct {
		repository datastruct.FeedsRepository
		logger     log.Logger
	}
)

func NewService(repo datastruct.FeedsRepository, logger log.Logger) Service {
	return &service{
		repository: repo,
		logger:     log.With(logger, "repo", "service"),
	}
}

func (s *service) GetFeeds(ctx context.Context) ([]datastruct.Feeds, error) {
	return s.repository.GetFeeds(ctx)
}
