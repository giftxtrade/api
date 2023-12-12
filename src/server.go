package main

import (
	"log"

	"github.com/giftxtrade/api/src/app"
	"github.com/giftxtrade/api/src/utils"
	"github.com/gofiber/fiber/v2"

	_ "github.com/lib/pq"
)

func main() {
	// Attempt connection with DB
	conn, err := utils.NewDbConnection()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	server := fiber.New(fiber.Config{
		ServerHeader: "giftxtrade api v2",
	})
	app.New(conn, server)

	const port = "8080"
	if err := server.Listen(":" + port); err != nil {
		log.Fatalf("❌ port %s already in use. could not start server\n\n", port)
		log.Fatal(err)
	}
}
