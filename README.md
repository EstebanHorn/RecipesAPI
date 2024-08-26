
---

# Recipes API

This is a RESTful API for managing recipes, allowing users to create, retrieve, update, and delete recipes. The API is built with Go, MongoDB, and the Gorilla Mux router.

## Table of Contents

- [Features](#features)
- [Installation](#installation)
- [Configuration](#configuration)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)

## Features

- Create new recipes with ingredients, instructions, and preparation time.
- Retrieve existing recipes by ID.
- List all available recipes.
- Update or delete recipes by ID.
- MongoDB as the database for storing recipes.

## Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/EstebanHorn/RecipesAPI.git
    cd RecipesAPI
    ```

2. Install dependencies:

    ```bash
    go mod tidy
    ```

3. Create a `.env` file in the root directory and add your MongoDB URI:

    ```bash
    MONGOURI=mongodb+srv://yourusername:yourpassword@cluster.mongodb.net/?retryWrites=true&w=majority
    ```

4. Build and run the application:

    ```bash
    go run main.go
    ```

## Configuration

Make sure you have a `.env` file in the root directory of the project with the following variable:

```plaintext
MONGOURI=mongodb+srv://<username>:<password>@<cluster-url>/<dbname>?retryWrites=true&w=majority
```

Replace `<username>`, `<password>`, `<cluster-url>`, and `<dbname>` with your MongoDB credentials and database information.

## Usage

Once the API is running, you can interact with it using tools like [Postman](https://www.postman.com/) or `curl`.

### Example

To create a new recipe, you can send a POST request to the `/recipes` endpoint:

```bash
curl -X POST http://localhost:8080/recipes -H "Content-Type: application/json" -d '{
  "name": "Chocolate Cake",
  "ingredients": ["Flour", "Sugar", "Cocoa powder", "Eggs", "Butter", "Milk"],
  "instructions": "Preheat the oven to 350°F. Mix the ingredients and bake for 30 minutes.",
  "preparationTime": 45
}'
```

## API Endpoints

| Method | Endpoint       | Description                |
|--------|----------------|----------------------------|
| POST   | `/recipes`     | Create a new recipe        |
| GET    | `/recipes`     | Retrieve all recipes       |
| GET    | `/recipes/{id}`| Retrieve a recipe by ID    |
| PUT    | `/recipes/{id}`| Update a recipe by ID      |
| DELETE | `/recipes/{id}`| Delete a recipe by ID      |

---
