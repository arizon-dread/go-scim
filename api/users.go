package api

import (
	"net/http"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("unable to read body"))
	}

}
