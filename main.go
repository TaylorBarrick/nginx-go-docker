package main

import (
	"fmt"
	"net/http"
)

func main() {
	router := NewRouter()
	fmt.Println("Listening on :8081...")
	http.ListenAndServe(":8081", router)
}
