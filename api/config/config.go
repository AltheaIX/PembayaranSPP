package config

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"regexp"
)

type Config struct {
	Database struct {
		Username string `json:"DB_USER"`
		Password string `json:"DB_PASS"`
		Name     string `json:"DB_NAME"`
	}
}

func (config *Config) LoadDatabaseEnv() {
	projectName := regexp.MustCompile(`^(.*` + "api" + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	mapConfig, err := godotenv.Read(string(rootPath) + `/.env`)

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	jsonData, err := json.Marshal(mapConfig)
	if err != nil {
		return
	}

	if err = json.Unmarshal(jsonData, &config.Database); err != nil {
		return
	}

	fmt.Println(config.Database)
}
