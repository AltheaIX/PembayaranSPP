package config

import (
	"github.com/jmoiron/sqlx"
	"os"
)

type Config struct {
	Database struct {
		Host     string `json:"DB_HOST"`
		Username string `json:"DB_USER"`
		Password string `json:"DB_PASS"`
		Name     string `json:"DB_NAME"`
	}
}

func (config *Config) LoadDatabaseEnv() {
	config.Database.Username = os.Getenv("DB_USER")
	config.Database.Password = os.Getenv("DB_PASS")
	config.Database.Name = os.Getenv("DB_NAME")
	config.Database.Host = os.Getenv("DB_HOST")
}

func InitializeDatabase(db *sqlx.DB) {
	schema := `
	create type level as enum ('admin', 'petugas');

	create table if not exists petugas
	(
		id_petugas   bigint generated always as identity
			primary key,
		username     varchar(25) not null,
		password     varchar(32) not null,
		nama_petugas varchar(35) not null,
		level        level       not null
	);

	create table if not exists spp
	(
		id_spp  bigint not null
			primary key,
		tahun   bigint not null,
		nominal bigint not null
	);

	create table if not exists kelas
	(
		id_kelas            bigint      not null
			primary key,
		nama_kelas          varchar(10) not null,
		kompetensi_keahlian varchar(50) not null
	);

	create table if not exists siswa
	(
		nisn     char(10)    not null
			primary key,
		nis      char(8)     not null,
		nama     varchar(35) not null,
		id_kelas bigint      not null
			references kelas,
		alamat   text        not null,
		no_telp  varchar(13) not null,
		id_spp   bigint      not null
			references spp,
		unique (nisn, id_spp)
	);

	create table if not exists pembayaran
	(
		id_pembayaran bigint      not null
			primary key,
		id_petugas    bigint      not null
			references petugas,
		nisn          varchar(10) not null,
		tgl_bayar     date        not null,
		bulan_dibayar varchar(8)  not null,
		tahun_dibayar varchar(4)  not null,
		id_spp        bigint      not null,
		jumlah_bayar  bigint      not null,
		unique (id_spp, nisn),
		foreign key (nisn, id_spp) references siswa (nisn, id_spp)
	)
`
	db.MustExec(schema)
}
