package routes

import (
	"log"
	"net/http"
	"teste/controllers"
)

func RunRoutes() {
	http.HandleFunc("/usuarios/", controllers.UserHandler)
	log.Println("Rodando")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
