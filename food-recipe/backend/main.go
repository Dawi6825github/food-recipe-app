package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	_ "github.com/lib/pq" // PostgreSQL driver
)

var db *sql.DB

// JWT Middleware to protect routes
func jwtMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "Missing authorization token", http.StatusUnauthorized)
			return
		}

		// Validate the token
		_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte("your_secret_key"), nil
		})
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	var err error
	// Initialize database connection
	db, err = sql.Open("postgres", "user=yourusername dbname=yourdbname sslmode=disable")
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}
	defer db.Close()

	r := setupRoutes() // Use the setupRoutes function to define routes

	// Apply JWT middleware to protected routes
	protectedRoutes := r.PathPrefix("/api").Subrouter()
	protectedRoutes.Use(jwtMiddleware)

	log.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
