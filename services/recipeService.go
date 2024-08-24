package services

import (
	"go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
	"context"
	"github.com/EstebanHorn/RecipesAPI/database"
	"github.com/EstebanHorn/RecipesAPI/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//database.DB is the connection to the database
var collection *mongo.Collection

func init() () {
	client, err := database.Connect()
	if err != nil {
		panic(err) // Manejo adecuado del error en producción
	}

	collection = client.Database("recipes_db").Collection("recetas")
}


func CreateRecipe(recipe models.Recipe) error {
	_, err := collection.InsertOne(context.TODO(), recipe)
	if err != nil {
		return err
	}
	return nil
}

func GetRecipe(id string) (models.Recipe, error) {
	var recipe models.Recipe
	
	// Convertir el ID a ObjectID si es necesario
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return recipe, err
	}

	err = collection.FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&recipe)
	if err != nil {
		return models.Recipe{}, err
	}

	return recipe, nil
}

func GetRecipes() ([]models.Recipe, error) {
	var recipes []models.Recipe
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var recipe models.Recipe
		if err := cursor.Decode(&recipe); err != nil {
			return nil, err
		}
		recipes = append(recipes, recipe)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return recipes, nil
}
