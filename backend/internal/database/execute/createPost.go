package execute

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"forum-backend/internal/models"
)

func CreatePostSql(post models.NewPost, db *sql.DB) (string, bool) {
	fmt.Println(post)
	stmt, err := db.Prepare("INSERT INTO posts(author, title,content,creationDate) values(?,?,?,?)")
	if err != nil {
		log.Println(err.Error())
		return "SQL INJECTION", false
	}
	fmt.Println(post)
	if _, err := stmt.Exec(post.Author, post.Title, post.Content, time.Now().Format("01-02-2006 15:04:05")); err != nil {
		// fmt.Println(asd)
		log.Println(err.Error())
		return "Error with creation of new user", false
	}
	return "", true
}
