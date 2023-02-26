package user

type Petugas struct {
	IdPetugas   int `db:"id_petugas"`
	Username    string
	Password    string
	NamaPetugas string `db:"nama_petugas"`
	Level       string
}

type payload struct {
	Username string
	Password string
}
