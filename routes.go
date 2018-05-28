package main

import (
	"net/http"

	"github.com/kayalardanmehmet/wordtagapi/handler"
	"github.com/kayalardanmehmet/wordtagapi/middleware"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
	Middlewares []middleware.MiddlewareFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		handler.Index,
		[]middleware.MiddlewareFunc{
			middleware.AuthMiddleware,
		},
	},
	Route{
		"TodoIndex",
		"GET",
		"/todos",
		handler.TodoIndex,
		nil,
	},
	Route{
		"TodoCreate",
		"POST",
		"/todos",
		handler.TodoCreate,
		nil,
	},
	Route{
		"TodoShow",
		"GET",
		"/todos/{todoId}",
		handler.TodoShow,
		nil,
	},
}
