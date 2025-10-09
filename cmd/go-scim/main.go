package main

import (
	"go-scim/api"
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	//User endpoints
	mux.HandleFunc("GET /users/{id}", api.GetUsers)
	//mux.HandleFunc("PATCH /users/{id}, api.PatchUsers

	//Group endpoints
}
