package main

import (
	"log"

	"github.com/giftxtrade/api/src/app"
	"github.com/giftxtrade/api/src/utils"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Attempt connection with DB
	conn, err := utils.NewDbConnection()
	if err != nil {
		log.Fatal("Could not connect to database.\n", err)
		return
	}

	server := fiber.New(fiber.Config{
		ServerHeader: "giftxtrade api v2",
	})
	app.New(conn, server, false)

	const port = "8080"
	if err := server.Listen(":" + port); err != nil {
		log.Fatalf("❌ port %s already in use. could not start server\n\n", port)
		log.Fatal(err)
	}
}
