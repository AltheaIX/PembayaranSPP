package user_test

import (
	"fmt"
	"testing"

	"github.com/AltheaIX/PembayaranSPP/user"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func TestGetAllPetugas(t *testing.T) {
	db, err := sqlx.Connect("postgres", "user=postgres password=althea dbname=PembayaranSPP sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}

	petugas, err := user.GetAllPetugas(db)
	t.Log(petugas)
}

func TestGetPetugasById(t *testing.T) {
	db, err := sqlx.Connect("postgres", "user=postgres password=althea dbname=PembayaranSPP sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}

	petugas, err := user.GetPetugasById(db, 15)
	t.Log(petugas)
}

func TestCreatePetugas(t *testing.T) {
	db, err := sqlx.Connect("postgres", "user=postgres password=althea dbname=PembayaranSPP sslmode=disable")
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
	db, err := sqlx.Connect("postgres", "user=postgres password=althea dbname=PembayaranSPP sslmode=disable")
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
	db, err := sqlx.Connect("postgres", "user=postgres password=althea dbname=PembayaranSPP sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}

	err = user.DeletePetugasById(db, 23)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(err)
}
