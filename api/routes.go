package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Root",
		"GET",
		"/",
		Ping,
	},
	Route{
		"Ping",
		"GET",
		"/ping",
		Ping,
	},
	Route{
		Name:        "Dummy response",
		Method:      "GET",
		Pattern:     "/dummy",
		HandlerFunc: Dummy,
	},
}
