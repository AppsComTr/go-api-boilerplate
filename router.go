package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kayalardanmehmet/wordtagapi/middleware"
)

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc

		if route.Middlewares != nil {

			for _, mid := range route.Middlewares {

				handler = mid(handler)

			}

		}

		handler = middleware.Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}

	return router
}
