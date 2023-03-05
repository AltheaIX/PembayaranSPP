package main

import (
	"fmt"
	"github.com/AltheaIX/PembayaranSPP/config"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Env struct {
	db *sqlx.DB
}

// Initialize Fiber and Routing
func route() {
	configModel := &config.Config{}
	configModel.LoadDatabaseEnv()

	dataSource := fmt.Sprintf("user=%v password=%v dbname=%v sslmode=disable", configModel.Database.Username, configModel.Database.Password, configModel.Database.Name)
	db, err := sqlx.Connect("postgres", dataSource)
	if err != nil {
		fmt.Println(err.Error())
	}

	env := &Env{db: db}

	app := fiber.New(fiber.Config{
		AppName: "Pembayaran SPP",
	})

	//GET Method
	app.Get("/petugas", env.HandlerGetPetugas)
	app.Get("/petugas/id/:id", env.HandlerGetPetugasById)

	//POST Method
	app.Post("/petugas", env.HandlerCreatePetugas)

	//PUT Method
	app.Put("/petugas/id/:id", env.HandlerUpdatePetugas)

	//DELETE Method
	app.Delete("/petugas/id/:id", env.HandlerDeletePetugas)

	fmt.Println(os.Getenv("PORT"))
	log.Fatal(app.Listen(":3000"))
}

// Initialize
func main() {

	route()
}
