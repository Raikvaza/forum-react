package app

import (
	"forum-backend/internal/database"
	"forum-backend/internal/handler"
	"log"
)

func Run() error {
	configDB := database.NewConfDB()

	db, err := database.InitDB(configDB)
	log.Println("db have been crated")
	if err != nil {
		return err
	}
	if err := database.CreateTables(db); err != nil {
		return err
	}
	log.Println("Table have been crated")
	apiServer := handler.NewApiServer(db)

	return apiServer.Start()
}
