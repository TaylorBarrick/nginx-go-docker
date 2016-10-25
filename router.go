package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

//Router is a mux for handling a couple of API items (Stufgf and Thing)
type Router struct {
	mux *http.ServeMux
}

//NewRouter provides a pointer to a new router with configured mux routes for Thing and Stuff
func NewRouter() *Router {
	var r Router
	var s Stuff
	var t Thing
	r.mux = http.NewServeMux()
	r.mux.HandleFunc("/stuff", route(s))
	r.mux.HandleFunc("/thing", route(t))
	return &r
}

func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router.mux.ServeHTTP(w, r)
}

func route(i interface{}) func(w http.ResponseWriter, r *http.Request) {
	var api RestAPI
	switch i.(type) {
	case Stuff:
		api = NewStuffAPI(nil)
	case Thing:
		api = NewThingAPI(nil)
	}
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			tokens := strings.Split(r.RequestURI, "/")
			fmt.Println(tokens)
			//TODO: Determine if params are present
			v := api.GetAll()
			json.NewEncoder(w).Encode(v)
		default:
			v := api.GetAll()
			json.NewEncoder(w).Encode(v)
		}
	}
}
