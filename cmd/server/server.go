package main

import (
	"net/http"

	v1 "github.com/okeeffed/go-1-22-mux/api/v1"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /v1/hello", v1.HelloHandler)
	mux.HandleFunc("GET /v1/goodbye", v1.GoodbyeHandler)

	http.ListenAndServe(":8080", mux)
}