package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// PostgreSQL connection string
const (
	host     = "localhost"
	port     = 5432
	user     = "kiran"
	password = "kiran0612"
	dbname   = "testdb"
)

// Struct to model the PostgreSQL table
type User struct {
	ID      int    `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Age     int    `json:"age,omitempty"`
	Gender  string `json:"gender,omitempty"`
	Email   string `json:"email,omitempty"`
	Mobile  string `json:"mobile,omitempty"`
	Address string `json:"address,omitempty"`
}

var db *sql.DB

func main() {
	// Initialize PostgreSQL connection
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to PostgreSQL!")

	// Initialize router
	r := mux.NewRouter()

	// Define endpoints
	r.HandleFunc("/users", createUser).Methods("POST")
	r.HandleFunc("/users/bulk", createUsersBulk).Methods("POST")
	r.HandleFunc("/users/{id}", getUser).Methods("GET")
	r.HandleFunc("/users/{id}", updateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", deleteUser).Methods("DELETE")

	// Start server
	fmt.Println("Starting server on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", r))
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	sqlStatement := `
        INSERT INTO users (name, age, gender, email, mobile, address)
        VALUES ($1, $2, $3, $4, $5, $6)
        RETURNING id`

	err = db.QueryRow(sqlStatement, user.Name, user.Age, user.Gender, user.Email, user.Mobile, user.Address).Scan(&user.ID)
	if err != nil {
		log.Printf("Failed to create user: %v", err)
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func createUsersBulk(w http.ResponseWriter, r *http.Request) {
	var users []User
	err := json.NewDecoder(r.Body).Decode(&users)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	// Start a PostgreSQL transaction
	tx, err := db.Begin()
	if err != nil {
		http.Error(w, "Failed to start transaction", http.StatusInternalServerError)
		return
	}
	defer tx.Rollback()

	// Prepare statement for bulk insert
	stmt, err := tx.Prepare("INSERT INTO users(name, age, gender, email, mobile, address) VALUES($1, $2, $3, $4, $5, $6)")
	if err != nil {
		http.Error(w, "Failed to prepare statement", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	// Execute statement for each user
	for _, user := range users {
		_, err := stmt.Exec(user.Name, user.Age, user.Gender, user.Email, user.Mobile, user.Address)
		if err != nil {
			http.Error(w, "Failed to execute statement", http.StatusInternalServerError)
			return
		}
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		http.Error(w, "Failed to commit transaction", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(users)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idStr := params["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var user User
	sqlStatement := `SELECT id, name, age, gender, email, mobile, address FROM users WHERE id=$1`
	row := db.QueryRow(sqlStatement, id)
	err = row.Scan(&user.ID, &user.Name, &user.Age, &user.Gender, &user.Email, &user.Mobile, &user.Address)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idStr := params["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var user User
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	sqlStatement := `
		UPDATE users
		SET name=$2, age=$3, gender=$4, email=$5, mobile=$6, address=$7
		WHERE id=$1`
	_, err = db.Exec(sqlStatement, id, user.Name, user.Age, user.Gender, user.Email, user.Mobile, user.Address)
	if err != nil {
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}

	user.ID = id
	json.NewEncoder(w).Encode(user)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idStr := params["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	sqlStatement := `DELETE FROM users WHERE id=$1`
	_, err = db.Exec(sqlStatement, id)
	if err != nil {
		http.Error(w, "Failed to delete user", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode("User deleted")
}
