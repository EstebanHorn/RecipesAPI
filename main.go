package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/EstebanHorn/RecipesAPI/routes"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Recipes api")


	// Configurar las rutas del servidor
	r := mux.NewRouter()
	routes.ConfigRoutes(r)

	// Iniciar el servidor
	log.Println("Server running on port 8080...")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
