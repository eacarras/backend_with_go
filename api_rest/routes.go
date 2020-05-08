package main

import (
	"net/http"
	"github.com/gorilla/mux"	
)

type Route struct {
	name string
	method string
	path string
	handle http.HandlerFunc
}

type Routes [] Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true);

	for _, route := range routes {
		router.
			Name(route.name).
			Methods(route.method).
			Path(route.path).
			Handler(route.handle)
	}
	return router;
}

var routes = Routes{
	Route{
		"Health",
		"GET",
		"/health",
		Health,
	},
	Route{
		"Get Persons",
		"GET",
		"/persons",
		PersonsList,
	},
	Route{
		"Get Person Details",
		"GET",
		"/person/{id}",
		PersonDetails,
	},
}
