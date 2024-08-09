package handler

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"go-crud-nats/model"
	"go-crud-nats/publisher"

	"github.com/gorilla/mux"
	//"github.com/jmoiron/sqlx"
)

func GetUsers(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
    users, err := publisher.PublishGetUsers(w)
    if err != nil {
        log.Println("Error getting users:", err)
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(users)
}

func GetUser(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	vars := mux.Vars(r)
	id := vars["id"]
	log.Println("Fetching user with ID:", id)

	user, err := publisher.PublishGetUser(id, w)
	if err != nil {
		return // Error already handled in publisher
	}

	json.NewEncoder(w).Encode(user)
}

func CreateUser(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
    var user model.User
    json.NewDecoder(r.Body).Decode(&user)

    // Call PublishUserCreation with both user and http.ResponseWriter
    response, err := publisher.PublishUserCreation(user, w)
    if err != nil {
        http.Error(w, "Failed to create user", http.StatusInternalServerError)
        log.Println("Error publishing user creation:", err)
        return
    }

    json.NewEncoder(w).Encode(response)
}

func UpdateUser(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var u model.User
	json.NewDecoder(r.Body).Decode(&u)

	vars := mux.Vars(r)
	id := vars["id"]

	response, err := publisher.PublishUpdateUser(u, id, w)
	if err != nil {
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(response)
}

func DeleteUser(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	vars := mux.Vars(r)
	id := vars["id"]

	response, err := publisher.PublishDeleteUser(id, w)
	if err != nil {
		http.Error(w, "Failed to delete user", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(response)
}