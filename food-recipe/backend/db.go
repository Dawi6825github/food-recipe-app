import (
    "database/sql"
    _ "github.com/lib/pq" // PostgreSQL driver
    "errors"
)

// Database struct to hold the connection
type Database struct {
    Connection *sql.DB
}

// Method to get hashed password from the database
func (db *Database) GetHashedPassword(username string) (string, error) {
    if db.Connection == nil {
        return "", errors.New("database connection is not established")
    }
    var hashedPassword string
    query := "SELECT password FROM users WHERE username = $1"
    err := db.Connection.QueryRow(query, username).Scan(&hashedPassword)
    if err != nil {
        if err == sql.ErrNoRows {
            return "", errors.New("user not found")
        }
        return "", err
    }
    return hashedPassword, nil
}

// Method to save user to the database
func (db *Database) SaveUser(username, hashedPassword string) error {
    if db.Connection == nil {
        return errors.New("database connection is not established")
    }
    if username == "" || hashedPassword == "" {
        return errors.New("username and hashedPassword are required")
    }
    query := "INSERT INTO users (username, password) VALUES ($1, $2)"
    _, err := db.Connection.Exec(query, username, hashedPassword)
    return err
}

// Method to create a new recipe in the database
func (db *Database) CreateRecipe(title, description, category string, preparationTime int) error {
    if db.Connection == nil {
        return errors.New("database connection is not established")
    }
    query := "INSERT INTO recipes (title, description, category, preparation_time) VALUES ($1, $2, $3, $4)"
    _, err := db.Connection.Exec(query, title, description, category, preparationTime)
    return err
}

// Method to get all recipes from the database
func (db *Database) GetAllRecipes() ([]Recipe, error) {
    if db.Connection == nil {
        return nil, errors.New("database connection is not established")
    }
    rows, err := db.Connection.Query("SELECT id, title, description, category, preparation_time FROM recipes")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var recipes []Recipe
    for rows.Next() {
        var recipe Recipe
        if err := rows.Scan(&recipe.ID, &recipe.Title, &recipe.Description, &recipe.Category, &recipe.PreparationTime); err != nil {
            return nil, err
        }
        recipes = append(recipes, recipe)
    }
    return recipes, nil
}

// Recipe struct to represent a recipe
type Recipe struct {
    ID              int
    Title           string
    Description     string
    Category        string
    PreparationTime int
}
