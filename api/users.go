package api

import (
	"encoding/json"
	"go-scim/pkg/handlers"
	"go-scim/pkg/models"
	"net/http"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("unable to read body"))
		return
	}

}

func UpsertUsers(w http.ResponseWriter, r *http.Request) {
	var b []byte
	_, err := r.Body.Read(b)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unable to read body"))
		return
	}
	var users []models.User
	err = json.Unmarshal(b, &users)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unable to marshal payload into go struct"))
		return
	}
	err = handlers.UpsertUsers(users)
	if err != nil {
		if err.
	}
}
