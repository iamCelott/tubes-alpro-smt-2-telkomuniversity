package actions

import (
	"encoding/json"
	"fmt"
	"os"
)

func Read() {
	fmt.Println("\n=== DAFTAR BAHAN BANGUNAN ===")

	// 1. Membaca langsung dari file JSON
	fileData, err := os.ReadFile("storages/bahan_bangunan.json")
	if err != nil {
		fmt.Println("Gagal membaca file JSON. Pastikan file ada di folder storages.")
		return
	}

	// 2. Menerjemahkan teks JSON ke dalam variabel global
	// Ini akan otomatis menimpa/memperbarui data lama di memori dengan data terbaru dari JSON
	errUnmarshal := json.Unmarshal(fileData, &DB_BahanBangunan)
	if errUnmarshal != nil {
		fmt.Println("Gagal membaca format JSON:", errUnmarshal)
		return
	}

	// 3. Menampilkan data yang sudah berhasil dibaca
	if len(DB_BahanBangunan) == 0 {
		fmt.Println("Belum ada data bahan bangunan.")
		return
	}

	for i, bahan := range DB_BahanBangunan {
		fmt.Printf("%d. ID: %s | Nama: %s | Kategori: %s | Satuan: %s (%s) | Stok: %d | Harga: %.2f\n",
			i+1, bahan.ID, bahan.Nama, bahan.Kategori, bahan.Satuan, bahan.TipeSatuan, bahan.Stock, bahan.Harga)
	}
}