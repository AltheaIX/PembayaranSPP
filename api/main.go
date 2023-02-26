package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

// TODO: Initialize Fiber and Routing
func route() {
	app := fiber.New(fiber.Config{
		AppName: "Pembayaran SPP",
	})
	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}

// TODO: Initialize
func main() {
	route()
}
