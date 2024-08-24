package controllers

import (
	"github.com/EstebanHorn/RecipesAPI/services"
	"github.com/EstebanHorn/RecipesAPI/models"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"fmt"
)

func Home(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Welcome to the Recipes API"))
}

func CreateRecipe(w http.ResponseWriter, r *http.Request) {
	var recipe models.Recipe
	if err := json.NewDecoder(r.Body).Decode(&recipe); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Mostrar la receta recibida en los logs para verificación

	if err := services.CreateRecipe(recipe); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(recipe)
}


func GetRecipe(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	// Mostrar el ID en los logs para verificación

	recipe, err := services.GetRecipe(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(recipe)
}
func GetRecipes(w http.ResponseWriter, r *http.Request) {
	recipes, err := services.GetRecipes()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(recipes)
}

func UpdateRecipe(w http.ResponseWriter, r *http.Request) {
	var recipe models.Recipe
	if err := json.NewDecoder(r.Body).Decode(&recipe); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id := mux.Vars(r)["id"]

	if err := services.UpdateRecipe(id, recipe); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(recipe)
}

func DeleteRecipe(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]


	if err := services.DeleteRecipe(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w).Encode(fmt.Sprintf("Recipe with ID %s was deleted", id))
	
}