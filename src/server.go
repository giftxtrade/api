package main

import (
	"log"
	"net/http"

	"github.com/giftxtrade/api/src/app"
	"github.com/giftxtrade/api/src/utils"
	"github.com/gorilla/mux"
)

func main() {
	// Attempt connection with DB
	conn, err := utils.NewDbConnection()
	if err != nil {
		log.Fatal("Could not connect to database.\n", err)
		return
	}
	
	// Create router instance
	router := mux.NewRouter()
	// Create server base with DB connection
	server := app.New(conn)
	server.CreateRoutes(router)

	const port = "8080"
	log.Printf("🚀 server starting on port %s\n", port)
	if err := http.ListenAndServe(":" + port, router); err != nil {
		log.Fatalf("❌ port %s already in use. could not start server\n\n", port)
		log.Fatal(err)
	}
}
