package models

type BahanBangunan struct {
    ID         string  `json:"id"`
    IDSupplier string  `json:"id_supplier"` // Penghubung ke data supplier
    Nama       string  `json:"nama"`
    Kategori   string  `json:"kategori"`
    Satuan     string  `json:"satuan"`      // "sak", "kg", "lembar", dll
	TipeSatuan string  `json:"tipe_satuan"` // "volume" atau "berat"
    Stock      int     `json:"stock"`
    Harga      float64 `json:"harga"`
}