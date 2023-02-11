package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"runtime"
	"time"

	"forum-backend/internal/database/execute"
	"forum-backend/internal/models"

	"github.com/google/uuid"
)

func (s *apiServer) SignInHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		body, err := io.ReadAll(r.Body)
		if err != nil || len(body) == 0 {
			log.Println("Empty body")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		var usr models.CheckUser
		err = json.Unmarshal(body, &usr)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if userModel, booll, _ := execute.CheckPasswordSql(usr, s.DB); booll {
			sessionNotExists, sessionErr := sessionNotExists(s.DB, userModel.UserId)
			if sessionErr != nil {
				log.Println(sessionErr)
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			log.Println("session: ", sessionNotExists)
			if sessionNotExists {
				tokenStat, err := s.DB.Prepare(`INSERT INTO user_sessions (token, expiresAt, userId) VALUES (?, ?, ?);`)
				// (SELECT userId FROM user WHERE username = ?)
				if err != nil {
					_, fileName, lineNum, _ := runtime.Caller(0)
					errStr := fmt.Sprintf("%s, %s(%s)", err.Error(), fileName, lineNum)
					s.log.Output(errStr)
					return
				}
				expiresAt := time.Now().Add(time.Hour * 24).Format("2006-01-02 15:04:05")
				token := generateToken()

				_, err = tokenStat.Exec(token, expiresAt, userModel.UserId)
				if err != nil {
					log.Println(err)
					return
				}
				cookie := &http.Cookie{
					Name:     "token",
					Value:    token,
					Expires:  time.Now().Add(time.Hour),
					HttpOnly: false,
					Path:     "/",
				}
				http.SetCookie(w, cookie)
			}
			// w.Header().Set("Content-Type", "application/json")
			err := json.NewEncoder(w).Encode(userModel)
			if err != nil {
				log.Println(err)
				w.WriteHeader(http.StatusBadRequest)
			}
			// w.Write(jData)
			w.WriteHeader(http.StatusOK)
			return
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)

		return
	}
}

// 	w.WriteHeader(http.StatusBadRequest)
// 	fmt.Println("User not found")
// 	return
// }

func generateToken() string {
	token := uuid.New().String()
	fmt.Println("Generated token:", token)
	return token
}

func sessionNotExists(db *sql.DB, userID int) (bool, error) {
	// check if the token already exists in the sessions table
	selectRecord := "SELECT token FROM user_sessions WHERE userId = ?"
	var token string
	err := db.QueryRow(selectRecord, userID).Scan(&token)
	if err == sql.ErrNoRows {
		// Handle case where no token exists for provided userId
		log.Println("Not in sessions")
		return true, nil
	} else if err != nil {
		// Handle other errors
		return false, err
	}
	return false, nil
}
