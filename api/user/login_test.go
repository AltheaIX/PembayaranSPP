package user

import (
	"fmt"
	"testing"

	"github.com/AltheaIX/PembayaranSPP/utils"
)

func TestCheckLogin(t *testing.T) {
	db, err := utils.DBConnection()
	if err != nil {
		t.Log(err)
	}

	payload := &payload{Username: "althea", Password: "althea"}

	res := CheckLogin(db, payload)
	fmt.Println(res.NamaPetugas)
}
