package execute

import (
	"database/sql"
	"fmt"
	"forum-backend/internal/models"
	"time"
)

func CheckPasswordSql(User models.CheckUser, db *sql.DB) (models.User, bool) {
	var fullUser models.User
	query := `SELECT * FROM user WHERE username=$1 and password=$2`
	row := db.QueryRow(query, User.Username, User.Password)
	if err := row.Scan(&fullUser.UserId, &fullUser.Username, &fullUser.Password, &fullUser.Email, &fullUser.Token, &fullUser.ExpiresAt); err != nil {
		fmt.Print(err.Error())
		return fullUser, false
	}
	fullUser.Token = "Token"
	fullUser.ExpiresAt = time.Now().AddDate(0, 0, 14).String()
	return fullUser, true
}
