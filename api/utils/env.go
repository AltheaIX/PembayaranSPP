package utils

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

// TODO: Load .env file from current directory
func InitializeEnv(envLoc string) {
	fmt.Println("Initializing .env file ...")
	err := godotenv.Load(envLoc)
	if err != nil {
		log.Fatal("Error on Initializing .env file")
		return
	}
}
