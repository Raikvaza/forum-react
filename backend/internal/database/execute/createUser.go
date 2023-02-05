package execute

import (
	"database/sql"
	"fmt"
	"forum-backend/internal/models"
)

func CreateUserSql(User models.NewUser, db *sql.DB) (string, bool) {
	stmt, err := db.Prepare("INSERT INTO User(username, password,email,token,expiresAt) values(?,?,?,?,?)")
	if err != nil {
		return "SQL INJECTION", false
	}
	// fmt.Println(User.Username)
	// fmt.Println(User.Email)
	// fmt.Println(User.Password)
	if _, err := stmt.Exec(User.Username, User.Password, User.Email, "null", "null"); err != nil {
		fmt.Println(err.Error())
		return "Error with creation of new user", false
	}
	return "", true
}
