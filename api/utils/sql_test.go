package utils

import (
	"fmt"
	"testing"
)

func TestDBConnection(t *testing.T) {
	db, err := DBConnection()
	if err != nil {
		t.Log(err)
	}

	fmt.Println("DB: ", db)
}
