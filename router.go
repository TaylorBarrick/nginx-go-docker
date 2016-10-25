package main

import (
	"encoding/json"
	"net/http"
)

type Router struct {
	mux *http.ServeMux
}

func NewRouter() *Router {
	var r Router
	r.mux = http.NewServeMux()
	var s Stuff
	var t Thing
	r.mux.HandleFunc("/stuff", Route(s))
	r.mux.HandleFunc("/thing", Route(t))
	return &r
}

func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router.mux.ServeHTTP(w, r)
}

func Route(i interface{}) func(w http.ResponseWriter, r *http.Request) {
	var api RestAPI
	switch i.(type) {
	case Stuff:
		api = NewStuffAPI(nil)
	case Thing:
		api = NewThingAPI(nil)
	default:
		api = NewThingAPI(nil)
	}
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			//TODO: Determine if params are present
			v := api.GetAll()
			json.NewEncoder(w).Encode(v)
		default:
			v := api.GetAll()
			json.NewEncoder(w).Encode(v)
		}
	}
}
