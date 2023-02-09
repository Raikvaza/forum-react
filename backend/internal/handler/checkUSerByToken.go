package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"forum-backend/internal/database/execute"
)

func (s *apiServer) CheckUSerByToken(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	tokenClient, err := r.Cookie("token")
	//	log.Println(tokenClient.Value)
	if err != nil {
		log.Println(err.Error)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	user, booll := execute.CheckByTokenLogin(s.DB, tokenClient.Value)
	if !booll {
		log.Println("execute didn't work")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request"))
		return
	}
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		log.Println("Error Parsing JSON after confirmation of userCheck")
		w.Write([]byte("StatusBadRequest"))
		w.WriteHeader(http.StatusBadRequest)
	}
}
