package models

type Supplier struct {
	ID     string `json:"ID"`
	Nama   string `json:"Nama"`
	Lokasi string `json:"Lokasi"`
	Kontak string `json:"Kontak"`
}