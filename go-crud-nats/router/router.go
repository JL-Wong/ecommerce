package router

import (
	"database/sql"
	"go-crud-nats/handler"
	"go-crud-nats/middleware"
	"net/http"

	"github.com/gorilla/mux"
	//"github.com/jmoiron/sqlx"
)

func NewRouter(db *sql.DB) *mux.Router {
    router := mux.NewRouter()
    router.Use(middleware.CORSMiddleware)
    router.Use(middleware.JsonContentTypeMiddleware)

    // Users endpoints
    router.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
        handler.GetUsers(w, r, db)
    }).Methods("GET")

    router.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
        handler.GetUser(w, r, db)
    }).Methods("GET")

    router.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
        handler.CreateUser(w, r, db)
    }).Methods("POST")

    router.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
        handler.UpdateUser(w, r, db)
    }).Methods("PUT")

    router.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
        handler.DeleteUser(w, r, db)
    }).Methods("DELETE")

    // Health check endpoint
    router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("Server is running"))
    }).Methods("GET")

    return router
}
