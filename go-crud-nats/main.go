package main

import (
	"database/sql"
	"go-crud-nats/middleware"
	"go-crud-nats/router"
	"go-crud-nats/subscriber"
	"log"
	"net/http"
	"os"

	//"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/nats-io/nats.go"
)

func main() {
	log.Println("Starting the application...")

	// connection to db
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	log.Println("Successfully connected to the database")

	nc, err := nats.Connect(os.Getenv("NATS_URL"))
    if err != nil {
        log.Fatal(err)
    }
    defer nc.Close()

	go subscriber.CreateUserSubscriber(db, nc)
	go subscriber.GetAllUserSubscriber(db, nc)
	go subscriber.GetUserSubscriber(db, nc)
	go subscriber.UpdateUserSubscriber(db,nc)
	go subscriber.DeleteUserSubscriber(db,nc)

	// create router
	r := router.NewRouter(db)

    // Apply middleware
    corsRouter := middleware.CORSMiddleware(r)
    jsonRouter := middleware.JsonContentTypeMiddleware(corsRouter)

    log.Println("Router setup complete. Starting server on port 8000")
    log.Fatal(http.ListenAndServe(":8000", jsonRouter))
}

