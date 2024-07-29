package account

import (
	"context"
	"net/http"
	"github.com/gorilla/mux"
	httptransport "github.com/go-kit/kit/transport/http"
)

func NewHTTPServer(ctx context.Context, endpoints Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Use(commonMiddleware)

	r.Methods("POST").Path("user").Handler(httptransport.NewServer (
		endpoints.CreateUser,
		decodeUserReq,
		encodeResponse,
	))

	r.Methods("GET").Path("/user/{id}").handler(httpteransport.NewServer (
		endpoints.GetUser,
		decodeEmailReq,
		encodeResponse,
	))
}

func CommonMiddleWear(next http.handler) http.Handler {
	return http.HandlerFunc(func(w http.ResonseWriter, r *http.Request) {
		w.Header().Add("content-type", "appliction/json")
		ext.serveHTTP(w, r)
	})
}