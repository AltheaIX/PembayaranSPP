package user_test

import (
	"fmt"
	"github.com/AltheaIX/PembayaranSPP/config"
	"testing"

	"github.com/AltheaIX/PembayaranSPP/user"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func TestGetAllPetugas(t *testing.T) {
	configModel := &config.Config{}
	configModel.LoadDatabaseEnv()

	dataSource := fmt.Sprintf("host=%v user=%v password=%v dbname=%v sslmode=disable", configModel.Database.Host, configModel.Database.Username, configModel.Database.Password, configModel.Database.Name)
	db, err := sqlx.Connect("postgres", dataSource)
	if err != nil {
		t.Fatal(err)
	}

	petugas, err := user.GetAllPetugas(db)
	t.Log(petugas)
}

func TestGetPetugasById(t *testing.T) {
	configModel := &config.Config{}
	configModel.LoadDatabaseEnv()

	dataSource := fmt.Sprintf("host=%v user=%v password=%v dbname=%v sslmode=disable", configModel.Database.Host, configModel.Database.Username, configModel.Database.Password, configModel.Database.Name)
	db, err := sqlx.Connect("postgres", dataSource)
	if err != nil {
		t.Fatal(err)
	}

	petugas, err := user.GetPetugasById(db, 15)
	t.Log(petugas)
}

func TestCreatePetugas(t *testing.T) {
	configModel := &config.Config{}
	configModel.LoadDatabaseEnv()

	dataSource := fmt.Sprintf("host=%v user=%v password=%v dbname=%v sslmode=disable", configModel.Database.Host, configModel.Database.Username, configModel.Database.Password, configModel.Database.Name)
	db, err := sqlx.Connect("postgres", dataSource)
	if err != nil {
		t.Fatal(err)
	}

	petugas := user.Petugas{
		Username:    "cTest",
		Password:    "cTest",
		NamaPetugas: "cTest",
		Level:       "admin",
	}

	newPetugas, err := user.CreatePetugas(db, petugas)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(newPetugas)
}

func TestUpdatePetugasById(t *testing.T) {
	configModel := &config.Config{}
	configModel.LoadDatabaseEnv()

	dataSource := fmt.Sprintf("host=%v user=%v password=%v dbname=%v sslmode=disable", configModel.Database.Host, configModel.Database.Username, configModel.Database.Password, configModel.Database.Name)
	db, err := sqlx.Connect("postgres", dataSource)
	if err != nil {
		t.Fatal(err)
	}

	petugas := user.Petugas{
		Username:    "uTest",
		Password:    "uTest",
		NamaPetugas: "uTest",
		Level:       "admin",
	}

	petugas.Id = 18
	t.Log(petugas)

	updatePetugas, err := user.UpdatePetugasById(db, petugas)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(updatePetugas)
}

func TestDeletePetugasById(t *testing.T) {
	configModel := &config.Config{}
	configModel.LoadDatabaseEnv()

	dataSource := fmt.Sprintf("host=%v user=%v password=%v dbname=%v sslmode=disable", configModel.Database.Host, configModel.Database.Username, configModel.Database.Password, configModel.Database.Name)
	db, err := sqlx.Connect("postgres", dataSource)
	if err != nil {
		t.Fatal(err)
	}

	err = user.DeletePetugasById(db, 23)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(err)
}
