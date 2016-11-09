package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

//Router is a mux for handling a couple of API items (Stuff and Thing)
type Router struct {
	mux *http.ServeMux
}

//Route represents a http pattern and supporting RestAPI
type Route struct {
	Path string
	API  RestAPI
}

//Routes is a collection of Route
type Routes []Route

//NewRouter provides a pointer to a new router with configured mux routes for Thing and Stuff
func NewRouter(routes *Routes) *Router {
	var r Router
	r.mux = http.NewServeMux()
	for _, e := range *routes {
		r.mux.HandleFunc(e.Path, route(e))
	}
	return &r
}

func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router.mux.ServeHTTP(w, r)
}

func route(route Route) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := -1
		params := getURIParams(r)
		if len(params) > 1 {
			id, _ = strconv.Atoi(params[1])
		}
		switch r.Method {
		case "GET":
			var v interface{}
			if id >= 0 {
				v = route.API.Get(id)
			} else {
				v = route.API.GetAll()
			}
			json.NewEncoder(w).Encode(v)
		case "DELETE":
			route.API.Delete(id)
		default:
			v := route.API.GetAll()
			json.NewEncoder(w).Encode(v)
		}
	}
}

func getURIParams(r *http.Request) []string {
	path := strings.TrimSuffix(strings.TrimPrefix(r.RequestURI, "/"), "/")
	return strings.Split(path, "/")
}
