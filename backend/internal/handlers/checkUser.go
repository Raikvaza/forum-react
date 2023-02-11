package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"forum-backend/internal/database/execute"
)

func (s *apiServer) CheckUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		// w.Write([]byte("Allowed method is GET"))
		return
	}

	tokenClient, err := r.Cookie("token")
	if err != nil {
		// _, fileName, lineNum, _ := runtime.Caller(0)
		// errStr := fmt.Sprintf("%s, %s(%s)", err.Error(), fileName, lineNum)
		// s.log.Output(errStr)
		w.WriteHeader(http.StatusUnauthorized)
		// w.Write([]byte("Unauthorized"))
		return
	}

	User, booll, err := execute.GetByToken(s.DB, tokenClient.Value)
	if !booll {
		w.WriteHeader(http.StatusUnauthorized)
		// w.Write([]byte("Unauthorized"))
		return
	}
	if err != nil {
		// _, fileName, lineNum, _ := runtime.Caller(0)
		// errStr := fmt.Sprintf("%s, %s(%s)", err.Error(), fileName, lineNum)
		// s.log.Output(errStr)
		log.Println(err)
		return
	}
	err = json.NewEncoder(w).Encode(User)
	if err != nil {
		// _, fileName, lineNum, _ := runtime.Caller(0)
		// errStr := fmt.Sprintf("%s, %s(%s)", err.Error(), fileName, lineNum)
		// s.log.Output(errStr)
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		// w.Write([]byte("Bad Request"))
		return
	}
	w.WriteHeader(http.StatusOK)
	return
}
