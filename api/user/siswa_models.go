package user

type Siswa struct {
	Nisn    string `json:"nisn" db:"nisn"`
	Nis     string `json:"nis" db:"nis"`
	Nama    string `json:"nama" db:"nama"`
	IdKelas int    `json:"id_kelas" db:"id_kelas"`
	Alamat  string `json:"alamat" db:"alamat"`
	NoTelp  string `json:"no_telp" db:"no_telp"`
	IdSpp   int    `json:"id_spp" db:"id_spp"`
}
