package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"forum-backend/internal/database/execute"
	"forum-backend/internal/models"
)

func (s *apiServer) CheckPassword(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint reached")
	if r.Method != http.MethodPost {
		fmt.Println("Wrong Method", r.Method)
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil || len(body) == 0 {
		log.Println("Empty body")
		w.Write([]byte("StatusBadRequest"))
		// w.WriteHeader(http.StatusBadRequest)
		return
	}

	// var usr map[string]interface{}
	// err = json.Unmarshal(body, &usr)
	// fmt.Println(usr)
	// data := map[string]interface{}{
	// 	"1": "one",
	// 	"2": "two",
	// 	"3": "three",
	// }
	// err = json.NewEncoder(w).Encode(&data)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	var usr models.CheckUser
	err = json.Unmarshal(body, &usr)
	if err != nil {
		fmt.Print(err.Error())
		w.Write([]byte("StatusBadRequest"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if userModel, booll := execute.CheckPasswordSql(usr, s.DB); booll {
		jData, err := json.Marshal(userModel)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jData)
		return
	}
	w.WriteHeader(http.StatusBadRequest)
	fmt.Println("User mpt found")
	return
}
