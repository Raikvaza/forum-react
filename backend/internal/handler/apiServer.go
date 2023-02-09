package handler

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"strings"
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
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Cookie, Content-Length, Accept-Encoding, X-CSRF-Token, charset, Credentials, Accept")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if r.Method == http.MethodOptions {
			w.WriteHeader(200)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func validateTokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "isAuthenticated", false)))
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == "" {
			next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "isAuthenticated", false)))
			return
		}

		// Validate the token here. You can use a database or a map to store the valid tokens and
		// check if the token in the header is present in the stored tokens.
		// For example:
		validTokens := map[string]bool{
			"9f9c2fcd-c51f-4de0-b22c-4f4a99d4ad79": true,
			"7f9c2fcd-c51f-4de0-b22c-4f4a99d4ad79": true,
		}
		if !validTokens[token] {
			next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "isAuthenticated", false)))
			return
		}

		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "isAuthenticated", true)))
	})
}

func (s *apiServer) Start() error {
	log.Println("Starting the server at port localhost:8080")

	s.Router.Handle("/api/checkPassword", CorsHeaders(http.HandlerFunc(s.CheckPassword)))
	s.Router.Handle("/api/checkUserByToken", CorsHeaders(http.HandlerFunc(s.CheckUSerByToken)))
	s.Router.Handle("/api/createUser", CorsHeaders(http.HandlerFunc(s.CreateUser)))
	s.Router.Handle("/api/createPost", CorsHeaders(http.HandlerFunc(s.CreatePost)))
	s.Router.HandleFunc("/api/createComment", s.NewComment)
	s.Router.HandleFunc("/api/newLike", s.Like)

	s.Router.Handle("/api/getAllpost", CorsHeaders(http.HandlerFunc(s.GetAllpost)))
	return http.ListenAndServe(":8080", s.Router)
}
