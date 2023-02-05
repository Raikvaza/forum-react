package handler

import (
	"encoding/json"
	"fmt"
	"forum-backend/internal/database/execute"
	"forum-backend/internal/models"
	"io"
	"net/http"
)

func (s *apiServer) CreatePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil || len(body) == 0 {
		w.WriteHeader(400)
		return
	}
	var post models.NewPost
	err = json.Unmarshal(body, &post)
	if err != nil {
		fmt.Print(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if res, booll := execute.CreatePostSql(post, s.DB); !booll {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(res))
		return
	}
	w.WriteHeader(http.StatusCreated)
	return
}
