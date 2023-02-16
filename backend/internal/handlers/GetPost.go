package handlers

import (
	"fmt"
	"net/http"
)

func GetPostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	id := r.URL.Query().Get("id")
	fmt.Printf("The ID is: %s\n", id)
}
