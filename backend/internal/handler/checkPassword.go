package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"forum-backend/internal/database/execute"
	"forum-backend/internal/models"
)

func (s *apiServer) CheckPassword(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint reached")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, charset")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil || len(body) == 0 {
		w.Write([]byte("StatusBadRequest"))
		// w.WriteHeader(http.StatusBadRequest)
		return
	}

	var usr models.CheckUser
	err = json.Unmarshal(body, &usr)
	if err != nil {
		fmt.Print(err.Error())
		w.Write([]byte("StatusBadRequest"))
		// w.WriteHeader(http.StatusBadRequest)
		return
	}

	if userModel, booll := execute.CheckPasswordSql(usr, s.DB); booll {
		jData, err := json.Marshal(userModel)
		if err != nil {
			//	w.WriteHeader(http.StatusInternalServerError)
			return
		}
		// w.WriteHeader(http.StatusOK)
		// w.Header().Set("Content-Type", "application/json")
		w.Write(jData)
		fmt.Println(jData)
		return
	}

	// w.WriteHeader(http.StatusBadRequest)
	return
}
