package routes

import (
	"net/http"

	handle "github.com/dvbnrg/ToyotaCodingChallenge/handlers"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		handle.Index,
	},
	Route{
		"Create",
		"GET",
		"/currency/{symbol}",
		handle.TodoCreate,
	},
	Route{
		"ReadAll",
		"GET",
		"/currency/all",
		handle.TodoGetAll,
	},
}
