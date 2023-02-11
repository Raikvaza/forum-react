package app

import (
	"forum-backend/internal/database"
	"forum-backend/internal/handlers"
	"forum-backend/internal/loger"
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
	logError := loger.NewLogger()
	apiServer := handlers.NewApiServer(db, logError)

	return apiServer.Start()
}
