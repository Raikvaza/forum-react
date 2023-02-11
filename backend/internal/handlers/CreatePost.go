package handlers

// import (
// 	"encoding/json"
// 	"forum-backend/internal/database/execute"
// 	"log"
// 	"net/http"
// )

// func (s *apiServer) CreatePost(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodPost {
// 		w.WriteHeader(http.StatusMethodNotAllowed)
// 		w.Write([]byte("Allowed method is GET"))
// 		return
// 	}
// 	tokenClient, err := r.Cookie("token")
// 	// log.Println(tokenClient)
// 	if err != nil {
// 		log.Println(err)
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}
// 	if booll := execute.CheckByToken(s.DB, tokenClient.Value); !booll {
// 		w.WriteHeader(http.StatusBadRequest)
// 		w.Write([]byte("Bad Request"))
// 		return
// 	}
// 	allPost, err := execute.GetAllpostSql(s.DB)
// 	if err != nil {
// 		log.Println(err)
// 		w.WriteHeader(http.StatusBadRequest)
// 		w.Write([]byte("Bad Request"))
// 		return
// 	}
// 	err = json.NewEncoder(w).Encode(allPost)
// 	log.Println(allPost)
// 	if err != nil {
// 		log.Println(err)
// 		w.WriteHeader(http.StatusBadRequest)
// 		w.Write([]byte("Bad Request"))
// 		return
// 	}
// }

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"forum-backend/internal/database/execute"
	"forum-backend/internal/models"
)

func (s *apiServer) CreatePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		fmt.Println("Wrong Method", r.Method)
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil || len(body) == 0 {
		w.WriteHeader(400)
		return
	}
	tokenClient, err := r.Cookie("token")
	log.Println(tokenClient)
	if err != nil {
		log.Println(err.Error)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if booll := execute.CheckByToken(s.DB, tokenClient.Value); !booll {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request"))
		return
	}
	var post models.NewPost
	err = json.Unmarshal(body, &post)
	fmt.Println(post)
	if err != nil {
		log.Print(err.Error())
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
