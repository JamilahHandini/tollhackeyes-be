package dbscan

import (
	"time"
)

//* Data User
type User struct {
	Nama 		string	`json:"nama"`
	TanggalLahir time.Time	`json:"tanggal_lahir"`
	Email		string	`json:"email"`
	NoHP		string	`json:"no_hp"`
	NoKTP		string	`json:"no_ktp"`
	NoSIM		string	`json:"no_sim"`
	Password	string	`json:"password"`
	PlatNomor	string	`json:"plat_nomor"`
	EToll       EToll	`json:"e_tolls"`
	InfoPerjalanan []InfoPerjalanan	`json:"info_perjalanans"`
}

type InfoPerjalanan struct {
	Tanggal				time.Time	`json:"tanggal"`
	Diskon				int			`json:"diskon"`
	Jarak				int			`json:"jarak"`
	JumlahPenumpang		int			`json:"jumlah_penumpang"`
	StatusPerjalanan	bool		`json:"status_perjalanan"`
}

type EToll struct {
	BerlakuSampai		time.Time	`json:"berlaku_sampai"`
	Nomor				string		`json:"nomor"`
	Saldo				int			`json:"saldo"`
}