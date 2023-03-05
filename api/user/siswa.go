package user

import "github.com/jmoiron/sqlx"

func GetAllSiswa(db *sqlx.DB) ([]Siswa, error) {
	model := Siswa{}
	listSiswa := []Siswa{}

	rows, err := db.Queryx("SELECT * FROM siswa")
	for rows.Next() {
		err = rows.StructScan(&model)
		if err != nil {
			return listSiswa, err
		}

		listSiswa = append(listSiswa, model)
	}
	return listSiswa, err
}

func GetSiswaByNisn(db *sqlx.DB, nisn string) (Siswa, error) {
	model := Siswa{}

	err := db.Get(&model, "SELECT * FROM siswa WHERE nisn=$1", nisn)
	if err != nil {
		return model, err
	}

	return model, err
}

func CreateSiswa(db *sqlx.DB, siswa Siswa) (Siswa, error) {
	newSiswa := Siswa{}

	rows, err := db.NamedQuery("INSERT INTO siswa (nisn, nis, nama, id_kelas, alamat, no_telp, id_spp) VALUES (:nisn, :nis, :nama, :id_kelas, :alamat, :no_telp, :id_spp) RETURNING *", siswa)
	if err != nil {
		return newSiswa, err
	}
	defer rows.Close()

	if !rows.Next() {
		return newSiswa, rows.Err()
	}

	err = rows.StructScan(&newSiswa)
	return newSiswa, err
}

// TODO: Update Siswa Data
func UpdateSiswaByNisn(db *sqlx.DB, siswa Siswa) (Siswa, error) {
	updateSiswa := Siswa{}

	rows, err := db.NamedQuery("UPDATE siswa set nis=:nis, nama=:nama, id_kelas=:id_kelas, alamat=:alamat, no_telp=:no_telp, id_spp=:id_spp where nisn=:nisn RETURNING *", siswa)
	if err != nil {
		return updateSiswa, err
	}
	defer rows.Close()

	if !rows.Next() {
		return updateSiswa, err
	}

	err = rows.StructScan(&updateSiswa)
	return updateSiswa, err
}

// TODO: Delete Siswa Data
func DeleteSiswaByNisn(db *sqlx.DB, nisn string) error {
	_, err := db.Exec("DELETE from siswa where nisn=$1", nisn)
	if err != nil {
		return err
	}

	return err
}
