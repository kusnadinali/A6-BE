package service

import (
	"context"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

// NewHTTPServer...
func NewHTTPServer(ctx context.Context, endpoints Endpoints) http.Handler {
	r := mux.NewRouter()

	r.Use(JSONHeader)

	//tar tambahin parameter path pas udah nambahin function di database --> /postingan/{id}/{end}
	r.Methods(http.MethodGet).Path("/postingan").Handler(httptransport.NewServer(
		endpoints.GetFeeds,
		decodeGetFeeds,
		encodeResponse,
	))

	return r
}

// JSONHeader ...
func JSONHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
