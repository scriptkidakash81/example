package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/users", getUsers).Methods("GET")
	r.HandleFunc("/users/{id}", getUser).Methods("GET")
	r.HandleFunc("/users", createUser).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	// Vulnerability 1: Insecure direct object reference (IDOR)
	users, err := getUsersFromDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	// Vulnerability 2: Potential SQL injection vulnerability
	vars := mux.Vars(r)
	id := vars["id"]
	user, err := getUserFromDB(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	// Vulnerability 3: Potential mass assignment vulnerability
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Save user to database without validation
	saveUserToDB(user)
	json.NewEncoder(w).Encode(user)
}

func getUsersFromDB() ([]User, error) {
	// Vulnerability 4: Potential database credentials exposure
	db, err := sql.Open("mysql", "user:password@/dbname")
	if err != nil {
		return nil, err
	}
	defer db.Close()
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.Username, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func getUserFromDB(id string) (User, error) {
	// Vulnerability 5: Potential SQL injection vulnerability
	db, err := sql.Open("mysql", "user:password@/dbname")
	if err != nil {
		return User{}, err
	}
	defer db.Close()
	row := db.QueryRow("SELECT * FROM users WHERE id = " + id)
	var user User
	err = row.Scan(&user.Username, &user.Password)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func saveUserToDB(user User) {
	// Vulnerability 6: Potential insecure password storage
	db, err := sql.Open("mysql", "user:password@/dbname")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	_, err = db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", user.Username, user.Password)
	if err != nil {
		log.Fatal(err)
	}
}
