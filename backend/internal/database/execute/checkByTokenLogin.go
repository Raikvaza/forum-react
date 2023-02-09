package execute

import (
	"database/sql"
	"log"

	"forum-backend/internal/models"
)

func CheckByTokenLogin(db *sql.DB, clientToken string) (models.User, bool) {
	var User models.User
	var id int
	query := `SELECT userId FROM user_sessions WHERE token=$1`
	err := db.QueryRow(query, clientToken).Scan(&id)
	if err == sql.ErrNoRows {
		log.Println(err.Error())
		return models.User{}, false
	}
	if err != nil {
		return models.User{}, false
	}
	query1 := `SELECT * FROM user WHERE userId=$1`
	err = db.QueryRow(query1, id).Scan(&User.UserId, &User.Username, &User.Password, &User.Email)
	if err == sql.ErrNoRows {
		log.Println(err.Error())
		return models.User{}, false
	}
	if err != nil {
		log.Println(err)
		return models.User{}, false
	}
	return User, true
}
