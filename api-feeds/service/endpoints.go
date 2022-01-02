package service

import (
	"context"
	"encoding/json"
	"net/http"

	"be-feeds/datastruct"

	"github.com/go-kit/kit/endpoint"
	"github.com/gorilla/mux"
)

type (
	Endpoints struct {
		GetFeeds endpoint.Endpoint
	}

	GetFeedsRequest struct {
		Id     string `json:"id"`
		EndRow string `json:"endRow"`
	}

	GetFeedsResponse struct {
		Status    bool               `json:"status"`
		Message   string             `json:"message"`
		DataFeeds []datastruct.Feeds `json:"data"`
	}
)

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func decodeFeedsRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req GetFeedsRequest
	vars := mux.Vars(r)

	req = GetFeedsRequest{
		Id:     vars["id"],
		EndRow: vars["endRow"],
	}

	return req, nil
}

func MakeFeedsEndpoints(svc Service) Endpoints {
	return Endpoints{
		GetFeeds: makeGetFeeds(svc),
	}
}

func makeGetFeeds(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetFeedsRequest)
		res, err := svc.GetFeeds(ctx, req.Id, req.EndRow)

		if err != nil {
			return GetFeedsResponse{Status: false, Message: err.Error(), DataFeeds: nil}, err
		}
		return GetFeedsResponse{Status: true, Message: "success", DataFeeds: res}, nil
	}
}
