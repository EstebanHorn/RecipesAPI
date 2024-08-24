package main

import (
	"fmt"
	//"github.com/EstebanHorn/RecipesAPI/database"
	"github.com/EstebanHorn/RecipesAPI/routes"
	"net/http"
	"github.com/gorilla/mux"
	"log"
	//"context"
)

func main() {
	fmt.Println("Recipes api")


	r := mux.NewRouter()
	routes.ConfigRoutes(r)

	http.Handle("/", r)
	log.Println("Server running on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server")
	}

}