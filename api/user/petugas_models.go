package user

import "github.com/jmoiron/sqlx"

type Env struct {
	db *sqlx.DB
}

type Petugas struct {
	Id          int    `db:"id_petugas" json:",omitempty"`
	Username    string `db:"username" json:"username"`
	Password    string `db:"password" json:"password"`
	NamaPetugas string `db:"nama_petugas" json:"nama"`
	Level       string `db:"level" json:"level"`
}

type ResponseJson struct {
	Success bool
	Data    []Petugas
	Message string
}
