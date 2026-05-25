package actions

import (
	"encoding/json"
	"fmt"
	"os"
)

func Delete() {
	if len(DB_BahanBangunan) == 0 {
		fmt.Println("\nData kosong, tidak ada yang bisa dihapus.")
		return
	}

	var targetID string
	fmt.Print("\nMasukkan ID Bahan yang ingin dihapus: ")
	fmt.Scanln(&targetID)

	idx := -1
	for i, bahan := range DB_BahanBangunan {
		if bahan.ID == targetID {
			idx = i
			break
		}
	}

	if idx == -1 {
		fmt.Println("Data dengan ID tersebut tidak ditemukan.")
		return
	}

	DB_BahanBangunan = append(DB_BahanBangunan[:idx], DB_BahanBangunan[idx+1:]...)
	fmt.Println("Data berhasil dihapus!")

	// --- KODE SIMPAN KE JSON (Taruh di paling bawah Delete) ---
	jsonData, err := json.MarshalIndent(DB_BahanBangunan, "", "  ")
	if err == nil {
		os.WriteFile("storages/bahan_bangunan.json", jsonData, 0644)
	} else {
		fmt.Println("Gagal mengubah data ke JSON:", err)
	}
}
