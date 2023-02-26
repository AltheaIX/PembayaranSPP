package utils

import "github.com/jmoiron/sqlx"

type SQL struct {
	Db *sqlx.DB
}
