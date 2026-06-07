package models

type Supplier struct {
	ID              string  `json:"ID"`
	Nama            string  `json:"Nama"`
	Alamat          string  `json:"Alamat"`
	Kota            string  `json:"Kota"`
	Kabupaten       string  `json:"Kabupaten"`
	Provinsi        string  `json:"Provinsi"`
	Telepon         string  `json:"Telepon"`
	Email           string  `json:"Email"`
	JumlahPelayanan int     `json:"JumlahPelayanan"`
	Rating          float64 `json:"Rating"`
}
