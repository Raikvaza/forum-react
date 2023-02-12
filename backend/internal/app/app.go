package app

import (
	"forum-backend/internal/Log"
	"forum-backend/internal/database"
	"forum-backend/internal/handlers"
)

func Run() error {
	configDB := database.NewConfDB()
	logger, err := Log.CreateLogger() // Setting the logger
	if err != nil {
		return err
	}
	defer Log.CloseLogger(logger)
	if err != nil {
		return err
	}
	db, err := database.InitDB(configDB)

	Log.LogInfo("Successfully Initiated the Data Base")

	if err != nil {
		return err
	}
	if err := database.CreateTables(db); err != nil {
		return err
	}

	Log.LogInfo("Tables have been created")

	apiServer := handlers.NewApiServer(db)

	Log.LogInfo("NewApiServer has been created")

	return apiServer.Start()
}
