package main

import (
	"log"
	"projectapi/database"
	"projectapi/database/migration"
	"projectapi/server"
)

func main() {
	err := database.StartDB()
	logError(err)
	database, err := database.GetDatabase()
	logError(err)
	err = migration.RunMigrations(database)
	logError(err)
	server := server.NewServer()
	server.Run()
}

func logError(err error) {
	if err != nil {
		log.Fatalln(err.Error())
	}
}
