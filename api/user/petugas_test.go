package user

import (
	"fmt"
	"testing"

	"github.com/AltheaIX/PembayaranSPP/utils"
)

func TestCreatePetugas(t *testing.T) {
	db, err := utils.DBConnection()
	if err != nil {
		t.Log(err)
	}

	petugas := &Petugas{
		Username:    "altheax",
		Password:    "althea",
		NamaPetugas: "Malik Zaky Zahroni",
		Level:       "petugas",
	}

	CreatePetugas(db, petugas)
}

func TestReadPetugas(t *testing.T) {
	db, err := utils.DBConnection()
	if err != nil {
		t.Log(err)
	}

	petugas, err := ReadPetugas(db, 15)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(petugas)
}
