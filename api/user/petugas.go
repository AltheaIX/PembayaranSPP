package user

import (
	"github.com/jmoiron/sqlx"
)

func GetAllPetugas(db *sqlx.DB) ([]Petugas, error) {
	model := Petugas{}
	petugas := []Petugas{}

	rows, err := db.Queryx("SELECT * FROM petugas")
	for rows.Next() {
		err = rows.StructScan(&model)
		if err != nil {
			return petugas, err
		}

		petugas = append(petugas, Petugas{Id: model.Id, Username: model.Username, Password: model.Password, NamaPetugas: model.NamaPetugas, Level: model.Level})
	}
	return petugas, err
}

func GetPetugasById(db *sqlx.DB, id int) (Petugas, error) {
	model := Petugas{}

	err := db.Get(&model, "SELECT * FROM petugas WHERE id_petugas=$1", id)
	if err != nil {
		return model, err
	}

	return model, err
}

func CreatePetugas(db *sqlx.DB, petugas Petugas) (Petugas, error) {
	newPetugas := Petugas{}

	rows, err := db.NamedQuery("INSERT INTO petugas (username, password, nama_petugas, level) VALUES (:username, :password, :nama_petugas, :level) RETURNING *", petugas)
	if err != nil {
		return petugas, err
	}
	defer rows.Close()

	if !rows.Next() {
		return newPetugas, rows.Err()
	}

	err = rows.StructScan(&newPetugas)
	return newPetugas, err
}

func UpdatePetugasById(db *sqlx.DB, petugas Petugas) (Petugas, error) {
	updatePetugas := Petugas{}

	rows, err := db.NamedQuery("UPDATE petugas set username=:username, password=:password, nama_petugas=:nama_petugas, level=:level where id_petugas=:id_petugas RETURNING *", petugas)
	if err != nil {
		return updatePetugas, err
	}
	defer rows.Close()

	if !rows.Next() {
		return updatePetugas, err
	}

	err = rows.StructScan(&updatePetugas)
	return updatePetugas, err
}

func DeletePetugasById(db *sqlx.DB, id int) error {
	_, err := db.Exec("DELETE from petugas where id_petugas=$1", id)
	if err != nil {
		return err
	}

	return err
}
