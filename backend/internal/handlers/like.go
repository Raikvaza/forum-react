package handlers

import (
	"encoding/json"
	"fmt"
	"forum-backend/internal/models"
	"io"
	"net/http"
	"runtime"
)

func (s *apiServer) Like(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		_, fileName, lineNum, _ := runtime.Caller(0)
		errStr := fmt.Sprintf("%s, %s(%s)", err.Error(), fileName, lineNum)
		s.log.Output(errStr)
		w.WriteHeader(400)
		return
	}

	if len(body) == 0 {
		w.WriteHeader(400)
		return
	}
	var newLike models.Like
	err = json.Unmarshal(body, &newLike)
	if err != nil {
		_, fileName, lineNum, _ := runtime.Caller(0)
		errStr := fmt.Sprintf("%s, %s(%s)", err.Error(), fileName, lineNum)
		s.log.Output(errStr)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// checking for existing comment
	if newLike.CommentID > 8 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	// TODO INSTERT THIST TO DB usr
	///////////////////////////
	w.WriteHeader(http.StatusOK)
	return
}
