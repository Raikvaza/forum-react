package handler

import (
	"encoding/json"
	"forum-backend/internal/models"
	"io"
	"net/http"
	"strconv"
)

func (s *apiServer) NewComment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	postID, err := strconv.Atoi(r.URL.Query().Get("post_id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// check post is exist ?
	if postID > 8 {

		w.WriteHeader(http.StatusNotFound)
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil || len(body) == 0 {
		w.WriteHeader(400)
		return
	}
	var comment models.NewComment
	err = json.Unmarshal(body, &comment)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// Insert into database comment
	w.WriteHeader(http.StatusCreated)
}
