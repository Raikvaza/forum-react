package execute

import (
	"database/sql"
	"log"
)

func CheckByTokenLogin(db *sql.DB, clientToken string) bool {
	var id int
	query := `SELECT userId FROM user_sessions WHERE token=$1`
	err := db.QueryRow(query, clientToken).Scan(&id)
	if err == sql.ErrNoRows {
		log.Println(err.Error())
		return false
	}
	if err != nil {
		return false
	}

	return true
}
