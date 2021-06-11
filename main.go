package main

import (
	"log"
	"projectapi/database"
	"projectapi/migration"
	"projectapi/server"
)

func main() {
	err := database.StartDB()
	if err != nil {
		log.Fatalln(err.Error())
	}
	database, err := database.GetDatabase()
	if err != nil {
		log.Fatalln(err.Error())
	}
	err = migration.RunMigrations(database)
	if err != nil {
		log.Fatalln(err.Error())
	}
	server := server.NewServer()
	server.Run()
}
