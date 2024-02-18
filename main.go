package main

import (
	"teste/database"
	"teste/routes"
)

func main() {
	database.StartDB()

	routes.RunRoutes()
}
