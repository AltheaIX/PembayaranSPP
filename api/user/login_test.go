package user

import (
	"fmt"
	"testing"

	"github.com/AltheaIX/PembayaranSPP/utils"
)

func TestCheckAuthentication(t *testing.T) {
	db, err := utils.DBConnection()
	if err != nil {
		t.Log(err)
	}

	payload := &payload{Username: "althea", Password: "althxea"}

	petugas, err := CheckAuthentication(db, payload)
	fmt.Println(petugas, err)
}
