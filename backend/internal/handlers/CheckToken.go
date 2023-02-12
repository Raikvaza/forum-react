package handlers

import (
	"encoding/json"
	"forum-backend/internal/Log"
	"net/http"

	"forum-backend/internal/database/execute"
)

func (s *apiServer) CheckToken(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	tokenClient, err := r.Cookie("token")
	if err != nil {

		Log.LogError(err.Error())
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	User, booll, err := execute.GetByToken(s.DB, tokenClient.Value)
	if !booll {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if err != nil {

		Log.LogError(err.Error())
		return
	}
	err = json.NewEncoder(w).Encode(User)
	if err != nil {

		Log.LogError(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}
