package main

import "projectapi/server"

func main() {
	server := server.NewServer()
	server.Run()
}
