package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"forum-backend/internal/database/execute"
)

func GetPosts(DB *sql.DB, w http.ResponseWriter) {
	allPost, err := execute.GetAllpostSql(DB)
	if err != nil {
		log.Println("GetAllpostSql in GetPosts")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request"))
		return
	}
	err = json.NewEncoder(w).Encode(allPost)
	log.Println(allPost)
	if err != nil {
		log.Println("Encoder in GetPosts")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request"))
		return
	}
}
