package models

type RiwayatPelayanan struct {
    IDTransaksi   string `json:"id_transaksi"`
    IDSupplier    string `json:"id_supplier"`
    Tanggal       string `json:"tanggal"`
    SkorKepuasan  int    `json:"skor_kepuasan"` // Nilai dari 1-5, nanti dirata-rata ke RatingLayanan
    Catatan       string `json:"catatan"`
}