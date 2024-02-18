package api

import (
	"net/http"

	model "github.com/okeeffed/go-1-22-mux/internal/model"
)

// HelloHandler is a simple HTTP handler that writes a response.
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	// Create a new user.
	user := model.User{
		Name: "World",
	}

	// Write a response to the client.
	w.Write([]byte("Hello, " + user.Name + "!"))
}