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
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT env is required.")
	}

	dataSource := fmt.Sprintf("host=%v user=%v password=%v dbname=%v sslmode=disable", configModel.Database.Host, configModel.Database.Username, configModel.Database.Password, configModel.Database.Name)
	db, err := sqlx.Connect("postgres", dataSource)
	if err != nil {
		log.Fatal(err.Error())
	}

	env := &Env{db: db}
	config.InitializeDatabase(db)
	db.Exec("SELECT * FROM petugas")

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

	log.Fatal(app.Listen(":" + port))
}

// Initialize
func main() {
	route()
}
