package feeds

import (
	"github.com/go-kit/kit/endpoint"

	"context"
)

type Endpoints struct {
	GetPost endpoint.Endpoint
}

func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		GetPost: makeGetPostEndpoint(s),
	}
}

func makeGetPostEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetPostRequest)
		username, err := s.GetPost(ctx, req.post_id)

		return GetPostResponse{
			username: username,
		}, err

	}

}
