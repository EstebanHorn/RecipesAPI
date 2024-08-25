package database

import (
	"context"
	"fmt"
	"os"
	"log"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connect establishes a connection to the MongoDB database and returns the client.
func Connect() (*mongo.Client, error) {
	
	// Cargar el archivo .env
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("Error loading .env file: %v", err) // Usa Printf en lugar de Fatalf para no detener la ejecución
	}

	// Obtener la URI de MongoDB desde la variable de entorno
	mongoURI := os.Getenv("MONGOURI")
	if mongoURI == "" {
		log.Fatal("MONGOURI environment variable is not set")
	}

	fmt.Println("Mongo URI: ", mongoURI)

	// Configurar las opciones del cliente de MongoDB
	opts := options.Client().
		ApplyURI(mongoURI).
		SetServerAPIOptions(options.ServerAPI(options.ServerAPIVersion1))

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %v", err)
	}

	// Verificar la conexión a MongoDB
	if err := client.Ping(context.TODO(), nil); err != nil {
		return nil, fmt.Errorf("failed to ping MongoDB: %v", err)
	}

	fmt.Println("Successfully connected to MongoDB!")
	return client, nil
}
