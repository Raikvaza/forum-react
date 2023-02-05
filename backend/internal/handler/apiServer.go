package handler

import (
	"database/sql"
	"fmt"
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

func CorsHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, charset")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		fmt.Println("Cors applied")
		next.ServeHTTP(w, r)
	})
}

func (s *apiServer) Start() error {
	// fmt.Println("Router Start")
	s.Router.Handle("/api/checkPassword", CorsHeaders(http.HandlerFunc(s.CheckPassword)))
	s.Router.HandleFunc("/api/createUser", s.CreateUser)
	s.Router.HandleFunc("/api/createPost", s.CreatePost)
	s.Router.HandleFunc("/api/createComment", s.NewComment)
	s.Router.HandleFunc("/api/newLike", s.Like)
	return http.ListenAndServe(":8080", s.Router)
}
