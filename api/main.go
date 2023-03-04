package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Env struct {
	db *sqlx.DB
}

// TODO: Initialize Fiber and Routing
func route() {
	db, err := sqlx.Connect("postgres", "user=postgres password=althea dbname=PembayaranSPP sslmode=disable")
	if err != nil {
		fmt.Println(err.Error())
	}

	env := &Env{db: db}

	app := fiber.New(fiber.Config{
		AppName: "Pembayaran SPP",
	})

	app.Get("/petugas", env.HandlerGetPetugas)
	app.Get("/petugas/id/:id", env.HandlerGetPetugasById)
	app.Post("/petugas", env.HandlerCreatePetugas)

	fmt.Println(os.Getenv("PORT"))
	log.Fatal(app.Listen(":3000"))
}

// TODO: Initialize
func main() {

	route()
}
