package main

import (
    "encoding/json"
    "net/http"
    "github.com/dgrijalva/jwt-go"
    "golang.org/x/crypto/bcrypt"
    "time"
)

// User structure
type User struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

// Generate JWT token
func generateToken(username string) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "username": username,
        "exp":      time.Now().Add(time.Hour * 72).Unix(),
    })

    tokenString, err := token.SignedString([]byte("your_secret_key"))
    if err != nil {
        return "", err
    }
    return tokenString, nil
}

// Handler for user login
func loginHandler(w http.ResponseWriter, r *http.Request) {
    var user User
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, "Invalid request", http.StatusBadRequest)
        return
    }

    // Logic to retrieve the stored hashed password from the database
    storedHashedPassword, err := db.GetHashedPassword(user.Username)
    if err != nil {
        http.Error(w, "User not found", http.StatusUnauthorized)
        return
    }


    // Compare the hashed password
    err = bcrypt.CompareHashAndPassword([]byte(storedHashedPassword), []byte(user.Password))
    if err != nil {
        http.Error(w, "Invalid credentials", http.StatusUnauthorized)
        return
    }

    token, err := generateToken(user.Username)
    if err != nil {
        http.Error(w, "Could not generate token", http.StatusInternalServerError)
        return
    }
    w.Write([]byte(token))
}

// Handler for user registration
func registerHandler(w http.ResponseWriter, r *http.Request) {
    var user User
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, "Invalid request", http.StatusBadRequest)
        return
    }

    // Logic to save user to the database (this should check for existing users)
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        http.Error(w, "Could not hash password", http.StatusInternalServerError)
        return
    }

    // Save user to the database
    err = db.SaveUser(user.Username, string(hashedPassword))
    if err != nil {
        http.Error(w, "Could not save user", http.StatusInternalServerError)
        returnpackage main

import (
    "database/sql"
    "encoding/json"
    "net/http"
    "github.com/dgrijalva/jwt-go"
    "golang.org/x/crypto/bcrypt"
    "time"
)

// User structure
type User struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

// Database interface
type Database interface {
    GetHashedPassword(username string) (string, error)
    SaveUser(username string, hashedPassword string) error
}

// SQLite database implementation
type SQLiteDB struct {
    db *sql.DB
}

func (db *SQLiteDB) GetHashedPassword(username string) (string, error) {
    var hashedPassword string
    err := db.db.QueryRow("SELECT password FROM users WHERE username = ?", username).Scan(&hashedPassword)
    if err != nil {
        return "", err
    }
    return hashedPassword, nil
}

func (db *SQLiteDB) SaveUser(username string, hashedPassword string) error {
    _, err := db.db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", username, hashedPassword)
    return err
}

// Generate JWT token
func generateToken(username string) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "username": username,
        "exp":      time.Now().Add(time.Hour * 72).Unix(),
    })

    tokenString, err := token.SignedString([]byte("your_secret_key"))
    if err != nil {
        return "", err
    }
    return tokenString, nil
}

// Handler for user login
func loginHandler(w http.ResponseWriter, r *http.Request) {
    var user User
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, "Invalid request", http.StatusBadRequest)
        return
    }

    // Logic to retrieve the stored hashed password from the database
    storedHashedPassword, err := db.GetHashedPassword(user.Username)
    if err != nil {
        if err == sql.ErrNoRows {
            http.Error(w, "User not found", http.StatusUnauthorized)
        } else {
            http.Error(w, "Database error", http.StatusInternalServerError)
        }
        return
    }

    // Compare the hashed password
    err = bcrypt.CompareHashAndPassword([]byte(storedHashedPassword), []byte(user.Password))
    if err != nil {
        http.Error(w, "Invalid credentials", http.StatusUnauthorized)
        return
    }

    token, err := generateToken(user.Username)
    if err != nil {
        http.Error(w, "Could not generate token", http.StatusInternalServerError)
        return
    }
    w.Write([]byte(token))
}

// Handler for user registration
func registerHandler(w http.ResponseWriter, r *http.Request) {
    var user User
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, "Invalid request", http.StatusBadRequest)
        return
    }

    // Logic to save user to the database (this should check for existing users)
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        http.Error(w, "Could not hash password", http.StatusInternalServerError)
        return
    }

    // Save user to the database
    err = db.SaveUser(user.Username, string(hashedPassword))
    if err != nil {
        if err == sql.ErrNoRows {
            http.Error(w, "User already exists", http.StatusConflict)
        } else {
            http.Error(w, "Could not save user", http.StatusInternalServerError)
        }
        return
    }

    w.Write([]byte("User registered successfully"))
}

func main() {
    // Initialize the database
    db, err := sql.Open("sqlite3", "./users.db")
    if err != nil {
        panic(err)
    }
    defer db.Close()

    // Create the users table if it doesn't exist
    _, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS users (
            username TEXT PRIMARY KEY,
            password TEXT NOT NULL
        );
    `)
    if err != nil {
        panic(err)
    }

    // Initialize the database interface
    dbInterface := &SQLiteDB{db: db}

    // Register the handlers
    http.HandleFunc("/login", loginHandler)
    http.HandleFunc("/register", registerHandler)

    // Start the server
    http.ListenAndServe(":8080", nil)
}
    }

    w.Write([]byte("User registered successfully"))
}
