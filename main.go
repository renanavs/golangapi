package main

import (
	"projectapi/server"
)

func main() {
	// database.StartDB()
	server := server.NewServer()
	server.Run()
}
