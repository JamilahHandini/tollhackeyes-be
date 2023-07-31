package dbscan


//* Data User
type User struct {
	Nama 		string	`json:"nama"`
	TanggalLahir string	`json:"tanggal_lahir"`
	Email		string	`json:"email"`
	NoHP		string	`json:"no_hp"`
	NoKTP		string	`json:"no_ktp"`
	NoSIM		string	`json:"no_sim"`
	Password	string	`json:"password"`
	PlatNomor	string	`json:"plat_nomor"`
	EToll       EToll	`json:"e_tolls"`
	InfoPerjalanan InfoPerjalanan	`json:"info_perjalanans"`
}

type Perjalanan struct {
    Diskon           int    `json:"diskon"`
    Jarak            int    `json:"jarak"`
    JumlahPenumpang  int    `json:"jumlah_penumpang"`
    StatusPerjalanan bool   `json:"status_perjalanan"`
    Tanggal          string `json:"tanggal"`
}

type InfoPerjalanan struct {
	Perjalanan1 Perjalanan `json:"perjalanan1"`
    Perjalanan2 Perjalanan `json:"perjalanan2"`
    Perjalanan3 Perjalanan `json:"perjalanan3"`
}

type EToll struct {
	BerlakuSampai		string	`json:"berlaku_sampai"`
	Nomor				string		`json:"nomor"`
	Saldo				int			`json:"saldo"`
}