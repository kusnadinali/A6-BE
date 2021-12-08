package service

import (
	"context"
	"encoding/json"
	"net/http"
	"tes/datastruct"

	"github.com/go-kit/kit/endpoint"
)

type (
	// Endpoints define all endpoint
	Endpoints struct {
		GetFeeds endpoint.Endpoint
	}

	Response struct {
		Status  bool               `json:"status"`
		Message string             `json:"msg"`
		Data    []datastruct.Feeds `json: "data"`
	}

	GetFeedsReq struct {
		ID int `json:"id"`
	}
)

func MakeFeedsEndpoints(svc Service) Endpoints {
	return Endpoints{
		GetFeeds: makeGetFeeds(svc),
	}
}

func makeGetFeeds(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		res, err := svc.GetFeeds(ctx)
		if err != nil {
			return Response{Status: false, Message: err.Error()}, nil
		}
		return Response{Status: true, Message: "succes", Data: res}, nil
	}
}

// r teh apa isinya aih
func decodeGetFeeds(_ context.Context, r *http.Request) (request interface{}, err error) {

	return r, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	res := response.(Response)

	// ga pake package util, gatau buat apaan hshs

	//sc := util.StatusCode(res.Message)
	// if sc == 0 {
	// 	sc = 500
	// }
	//w.WriteHeader(sc)

	return json.NewEncoder(w).Encode(&res)
}
