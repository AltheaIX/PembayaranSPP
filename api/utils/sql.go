package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// TODO: Make a db connection that can be reuse
func DBConnection() (*SQL, error) {
	InitializeEnv()
	dst := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DBUSER"), os.Getenv("DBPASS"), os.Getenv("DBNAME"))
	fmt.Println(dst)
	db, err := sqlx.Connect("postgres", dst)
	if err != nil {
		log.Fatalln(err)
	}

	return &SQL{Db: db}, nil
}
