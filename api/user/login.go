package user

import (
	"github.com/AltheaIX/PembayaranSPP/utils"
	_ "github.com/lib/pq"
)

// TODO: Make an authentication checker for API
func CheckAuthentication(SQL *utils.SQL, payload *payload) (Petugas Petugas, err error) {
	db := SQL.Db
	defer db.Close()

	err = db.Get(&Petugas, "SELECT * FROM petugas WHERE username=$1 AND password=$2", payload.Username, payload.Password)
	if err != nil {
		// Any error goes here
		return Petugas, err
	}
	// Result found
	return Petugas, err
}
