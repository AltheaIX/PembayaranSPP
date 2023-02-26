package user

import (
	"fmt"

	"github.com/AltheaIX/PembayaranSPP/utils"
	_ "github.com/lib/pq"
)

// TODO: Make a db connection that can be reuse
func CheckLogin(SQL *utils.SQL, payload *payload) Petugas {
	db := SQL.Db
	Petugas := Petugas{}
	err := db.Get(&Petugas, "SELECT * FROM petugas WHERE username=$1 AND password=$2", payload.Username, payload.Password)
	if err != nil {
		fmt.Println(err.Error())
		return Petugas
	}

	return Petugas
}
