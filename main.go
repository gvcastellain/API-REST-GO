package main

import (
	"api_rest/database"
	"api_rest/server"
)

func main() {
	database.StartDB()

	server := server.NewServer()

	server.Run()
}
