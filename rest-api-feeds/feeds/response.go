package feeds

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type (
	GetPostResponse struct {
		username string `json:username`
	}

	GetPostRequest struct {
		post_id string `json:post_id`
	}
)

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func decodeUsernameReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req GetPostRequest
	vars := mux.Vars(r)

	req = GetPostRequest{
		post_id: vars["post_id"],
	}
	return req, nil
}
