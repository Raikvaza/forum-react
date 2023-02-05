package handler

import (
	"database/sql"
	"net/http"
)

type apiServer struct {
	DB     *sql.DB
	Router *http.ServeMux
}

func NewApiServer(db *sql.DB) *apiServer {
	return &apiServer{
		DB:     db,
		Router: http.NewServeMux(),
	}
}

func (s *apiServer) Start() error {
	// fmt.Println("Router Start")
	s.Router.HandleFunc("/api/checkPassword", s.CheckPassword)
	s.Router.HandleFunc("/api/createUser", s.CreateUser)
	s.Router.HandleFunc("/api/createPost", s.CreatePost)
	s.Router.HandleFunc("/api/createComment", s.NewComment)
	s.Router.HandleFunc("/api/newLike", s.Like)
	return http.ListenAndServe(":8080", s.Router)
}
