package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

var routes = Routes{
	Route{"/stuff/", NewStuffAPI()},
	Route{"/thing/", NewThingAPI()},
}

func main() {
	flag.Parse()
	router := NewRouter(&routes)
	port := os.Getenv("PORT")
	fmt.Printf("Listening on PORT %s...\n", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), router)
}
