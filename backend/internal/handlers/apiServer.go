package handlers

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"runtime"

	"forum-backend/internal/database/execute"
	"forum-backend/internal/loger"
)

type apiServer struct {
	DB     *sql.DB
	Router *http.ServeMux
	log    *loger.ErrStr
}

func NewApiServer(db *sql.DB, log *loger.ErrStr) *apiServer {
	return &apiServer{
		DB:     db,
		Router: http.NewServeMux(),
		log:    log,
	}
}

func CorsHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Cookie, Content-Length, Accept-Encoding, X-CSRF-Token, charset, Credentials, Accept, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if r.Method == http.MethodOptions {
			w.WriteHeader(200)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (s *apiServer) validateTokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Cookies())
		var tokenClient *http.Cookie
		var err error
		if len(r.Cookies()) == 0 {
			next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "isAuthenticated", false)))
			return
		}
		for _, cookie := range r.Cookies() {
			if cookie.Name == "token" {
				tokenClient, err = r.Cookie("token")
				if err != nil {
					_, fileName, lineNum, _ := runtime.Caller(0)
					errStr := fmt.Sprintf("%s, %s(%s)", err.Error(), fileName, lineNum)
					s.log.Output(errStr)
					w.WriteHeader(http.StatusBadRequest)
					return
				}
			}
		}
		// if !validTokens[token] {
		// 	next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "isAuthenticated", false)))
		// 	return
		// }
		booll := execute.CheckByTokenLogin(s.DB, tokenClient.Value)
		log.Println(booll)
		if !booll {
			log.Println("Not Authenticated")
			next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "isAuthenticated", false)))
			return
			// log.Println("No didn't work")
			// w.WriteHeader(http.StatusBadRequest)
			// w.Write([]byte("Bad Request"))
			// return
		} else {
			log.Println("Authenticated")
			next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "isAuthenticated", true)))
			return
		}
	})
}

func (s *apiServer) Start() error {
	log.Println("Starting the server at port localhost:8080")
	s.Router.Handle("/api/home", CorsHeaders(http.HandlerFunc(s.HomeHandler)))
	s.Router.Handle("/api/signin", CorsHeaders(http.HandlerFunc(s.SignInHandler)))
	// s.Router.Handle("/api/checkUserByToken", CorsHeaders(http.HandlerFunc(s.CheckUSerByToken)))
	s.Router.Handle("/api/signup", CorsHeaders(s.validateTokenMiddleware(http.HandlerFunc(s.SignupHandler))))
	s.Router.Handle("/api/createPost", CorsHeaders(http.HandlerFunc(s.CreatePost)))
	// s.Router.HandleFunc("/api/createComment", s.CommentHandler)
	// s.Router.HandleFunc("/api/newLike", s.LikeHandler)
	s.Router.Handle("/api/checkUser", CorsHeaders(http.HandlerFunc(s.CheckUser)))
	//	s.Router.Handle("/api/getAllpost", CorsHeaders(http.HandlerFunc(s.GetAllPost)))
	return http.ListenAndServe(":8080", s.Router)
}
