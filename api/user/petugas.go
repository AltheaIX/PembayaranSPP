package user

import (
	"database/sql"
	"fmt"

	"github.com/AltheaIX/PembayaranSPP/utils"
	_ "github.com/lib/pq"
)

// TODO: Create petugas's data
func CreatePetugas(SQL *utils.SQL, Petugas *Petugas) {
	db := SQL.Db

	// Validate level for Enumerate
	if Petugas.Level != "admin" && Petugas.Level != "petugas" {
		fmt.Println("Invalid level.")
		return
	}

	// Check petugas's existance to avoid duplicate
	_, err := ReadPetugas(SQL, Petugas.Username)
	if err != sql.ErrNoRows {
		fmt.Println("Petugas already exist")
		return
	}

	query := "INSERT INTO petugas (username, password, nama_petugas, level) VALUES (:username, :password, :namaPetugas, :level)"
	_, err = db.NamedExec(query, map[string]interface{}{
		"username":    Petugas.Username,
		"password":    Petugas.Password,
		"namaPetugas": Petugas.NamaPetugas,
		"level":       Petugas.Level,
	})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Success inserting petugas!")
}

// TODO: Remove petugas's data
func ReadPetugas[V int | string](SQL *utils.SQL, input V) (Petugas Petugas, err error) {
	db := SQL.Db

	dataType := fmt.Sprintf("%T", input)
	switch dataType {
	// Search by ID
	case "int":
		err = db.Get(&Petugas, "SELECT * FROM petugas WHERE id_petugas=$1", input)
		if err != nil {
			// Any error goes here
			return Petugas, err
		}
		// Result found
		return Petugas, err

	// Search by Username
	case "string":
		err = db.Get(&Petugas, "SELECT * FROM petugas WHERE username=$1", input)
		if err != nil {
			// Any error goes here
			return Petugas, err
		}
		// Result found
		return Petugas, err
	}

	return
}

// TODO: Update petugas's data
func UpdatePetugas(SQL *utils.SQL) {
}

// TODO: Delete petugas's data
func DeletePetugas(SQL *utils.SQL) {
}
