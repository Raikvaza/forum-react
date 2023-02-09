package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"forum-backend/internal/database/execute"
)

func (s *apiServer) GetAllpost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Write([]byte("Allowed method is GET"))
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	tokenClient, err := r.Cookie("token")
	log.Println(tokenClient)
	if err != nil {
		log.Println(err.Error)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if booll := execute.CheckByToken(s.DB, tokenClient.Value); !booll {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request"))
		return
	}
	allPost, err := execute.GetAllpostSql(s.DB)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request"))
		return
	}
	err = json.NewEncoder(w).Encode(allPost)
	log.Println(allPost)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request"))
		return
	}
	w.WriteHeader(http.StatusOK)
	return
}
