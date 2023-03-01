package main

import (
	"fmt"
	"log"
	"os"

	"github.com/AltheaIX/PembayaranSPP/utils"
	"github.com/gofiber/fiber/v2"
)

// TODO: Initialize Fiber and Routing
func route() {
	utils.InitializeEnv(".env")
	app := fiber.New(fiber.Config{
		AppName: "Pembayaran SPP",
	})

	app.Get("/petugas/:param", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("param")
		if err != nil {
			username := c.Params("param")
			return c.SendString(fmt.Sprintf("%T", username))
		}
		fmt.Println(err)
		return c.SendString(fmt.Sprintf("%T", id))
	})

	fmt.Println(os.Getenv("PORT"))
	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}

// TODO: Initialize
func main() {
	route()
}
