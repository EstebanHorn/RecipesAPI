package routes

import (
	"github.com/EstebanHorn/RecipesAPI/controllers"
	"github.com/gorilla/mux"
)


func ConfigRoutes(router *mux.Router) {
	router.HandleFunc("/", controllers.Home).Methods("GET")
	router.HandleFunc("/recipes", controllers.GetRecipes).Methods("GET")
	router.HandleFunc("/recipes", controllers.CreateRecipe).Methods("POST")
	router.HandleFunc("/recipes/{id}", controllers.GetRecipe).Methods("GET")
	router.HandleFunc("/recipes/{id}", controllers.UpdateRecipe).Methods("PUT")
	router.HandleFunc("/recipes/{id}", controllers.DeleteRecipe).Methods("DELETE")
}