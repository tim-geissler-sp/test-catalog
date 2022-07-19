// Copyright (c) 2020, SailPoint Technologies, Inc. All rights reserved.
package infra

import (
	"github.com/gorilla/mux"
	"github.com/sailpoint/${{ values.name }}/internal/${{ values.name }}/cmd"
	"github.com/sailpoint/atlas-go/atlas/web"
	"net/http"
)

// buildRoutes configures all of the HTTP endpoints for the service.
func (s *Service) buildRoutes() *mux.Router {
	r := web.NewRouter(web.DefaultAuthenticationConfig(s.TokenValidator))

	r.Handle("/hello-world", s.returnHelloWorld()).Methods("GET")

	return r
}

// write http handler to return hello world (code 200)
func (s *Service) returnHelloWorld() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		cmd, err := cmd.NewHelloWorld()

		str, err := cmd.Handle(ctx) //output string

		if err != nil {
			web.WriteJSONWithError(ctx, w, err)
			return
		}

		web.WriteJSON(ctx, w, str)
	}
}

// requireRight is a middleware function that ensures that the current request
// has the specified right before calling the next handler in the chain.
// Requests that are missing the specified right will be terminated with
// a 403 Forbidden response.
func (s *Service) requireRight(right string, next http.Handler) http.Handler {
	m := web.RequireRights(s.AccessSummarizer, right)
	return m.Middleware(next)
}
