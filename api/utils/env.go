package utils

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

//TODO: Load .env file from current directory
func InitializeEnv() {
	fmt.Println("Initializing .env file ...")
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error on Initializing .env file")
		return
	}
}
