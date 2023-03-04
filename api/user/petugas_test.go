package user_test

import (
	"fmt"
	"testing"

	"github.com/AltheaIX/PembayaranSPP/user"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func TestAllPetugas(t *testing.T) {
	db, err := sqlx.Connect("postgres", "user=postgres password=althea dbname=PembayaranSPP sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}

	petugas, err := user.AllPetugas(db)
	t.Log(petugas)
}

func TestPetugasById(t *testing.T) {
	db, err := sqlx.Connect("postgres", "user=postgres password=althea dbname=PembayaranSPP sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}

	petugas, err := user.PetugasById(db, 15)
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
