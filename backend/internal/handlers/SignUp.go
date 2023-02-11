package handlers

import (
	"encoding/json"
	"fmt"
	"forum-backend/internal/database/execute"
	"forum-backend/internal/models"
	"io"
	"net/http"
	"runtime"
)

func (s *apiServer) SignupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method Not Allowed"))
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		_, fileName, lineNum, _ := runtime.Caller(0)
		errStr := fmt.Sprintf("%s, %s(%s)", err.Error(), fileName, lineNum)
		s.log.Output(errStr)
	}
	if len(body) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request"))
		return
	}
	var usr models.NewUser
	err = json.Unmarshal(body, &usr)
	if err != nil {
		_, fileName, lineNum, _ := runtime.Caller(0)
		errStr := fmt.Sprintf("%s, %s(%s)", err.Error(), fileName, lineNum)
		s.log.Output(errStr)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request"))
		return
	}
	fmt.Println(usr)
	if res, booll := execute.CreateUserSql(usr, s.DB); !booll {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(res))
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Succesfully created"))
	return
}
