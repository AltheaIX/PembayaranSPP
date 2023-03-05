package user_test

import (
	"fmt"
	"github.com/AltheaIX/PembayaranSPP/config"
	"github.com/AltheaIX/PembayaranSPP/user"
	"github.com/jmoiron/sqlx"
	"testing"
)

func TestGetAllSiswa(t *testing.T) {
	configModel := &config.Config{}
	configModel.LoadDatabaseEnv()

	dataSource := fmt.Sprintf("user=%v password=%v dbname=%v sslmode=disable", configModel.Database.Username, configModel.Database.Password, configModel.Database.Name)
	db, err := sqlx.Connect("postgres", dataSource)
	if err != nil {
		t.Fatal(err)
	}

	siswa, err := user.GetAllSiswa(db)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(siswa)
	/*
		_siswa := user.Siswa{}
		err = db.Get(&_siswa, "SELECT * FROM siswa WHERE nisn=$1", siswa[0].Nisn)
		t.Log(_siswa)
	*/
}

func TestGetSiswaByNisn(t *testing.T) {
	configModel := &config.Config{}
	configModel.LoadDatabaseEnv()

	dataSource := fmt.Sprintf("user=%v password=%v dbname=%v sslmode=disable", configModel.Database.Username, configModel.Database.Password, configModel.Database.Name)
	db, err := sqlx.Connect("postgres", dataSource)
	if err != nil {
		t.Fatal(err)
	}

	siswa, err := user.GetSiswaByNisn(db, "1234567890")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(siswa)
}

func TestCreateSiswa(t *testing.T) {
	configModel := &config.Config{}
	configModel.LoadDatabaseEnv()

	dataSource := fmt.Sprintf("user=%v password=%v dbname=%v sslmode=disable", configModel.Database.Username, configModel.Database.Password, configModel.Database.Name)
	db, err := sqlx.Connect("postgres", dataSource)
	if err != nil {
		t.Fatal(err)
	}

	siswa := user.Siswa{Nisn: "123456", Nis: "12345678", Nama: "Malik Zaky Zahroni", IdKelas: 1, Alamat: "Japan", NoTelp: "085156974848", IdSpp: 1}

	newSiswa, err := user.CreateSiswa(db, siswa)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(newSiswa)
}

func TestUpdateSiswa(t *testing.T) {
	configModel := &config.Config{}
	configModel.LoadDatabaseEnv()

	dataSource := fmt.Sprintf("user=%v password=%v dbname=%v sslmode=disable", configModel.Database.Username, configModel.Database.Password, configModel.Database.Name)
	db, err := sqlx.Connect("postgres", dataSource)
	if err != nil {
		t.Fatal(err)
	}

	siswa := user.Siswa{Nisn: "1234567890", Nis: "12345678", Nama: "Malik Zaky Zahroni", IdKelas: 1, Alamat: "Tokyo, Japan", NoTelp: "085156974842", IdSpp: 1}

	updateSiswa, err := user.UpdateSiswaByNisn(db, siswa)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(updateSiswa)
}

func TestDeleteSiswaByNisn(t *testing.T) {
	configModel := &config.Config{}
	configModel.LoadDatabaseEnv()

	dataSource := fmt.Sprintf("user=%v password=%v dbname=%v sslmode=disable", configModel.Database.Username, configModel.Database.Password, configModel.Database.Name)
	db, err := sqlx.Connect("postgres", dataSource)
	if err != nil {
		t.Fatal(err)
	}

	err = user.DeleteSiswaByNisn(db, "123456")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(err)
}
