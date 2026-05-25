package actions

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func Update() {
	if len(DB_BahanBangunan) == 0 {
		fmt.Println("\nData kosong, tidak ada yang bisa diubah.")
		return
	}

	var targetID string
	fmt.Print("\nMasukkan ID Bahan yang ingin diubah: ")
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

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Masukkan data baru (kosongkan/tekan enter jika ingin mempertahankan data teks lama):")

	fmt.Printf("Nama Bahan baru (Lama: %s): ", DB_BahanBangunan[idx].Nama)
	nama, _ := reader.ReadString('\n')
	if strings.TrimSpace(nama) != "" {
		DB_BahanBangunan[idx].Nama = strings.TrimSpace(nama)
	}

	fmt.Printf("Stock baru (Lama: %d): ", DB_BahanBangunan[idx].Stock)
	fmt.Scanln(&DB_BahanBangunan[idx].Stock)

	fmt.Printf("Harga baru (Lama: %.2f): ", DB_BahanBangunan[idx].Harga)
	fmt.Scanln(&DB_BahanBangunan[idx].Harga)

	fmt.Println("Data berhasil diperbarui!")

	// --- KODE SIMPAN KE JSON (Taruh di paling bawah Update) ---
	jsonData, err := json.MarshalIndent(DB_BahanBangunan, "", "  ")
	if err == nil {
		os.WriteFile("storages/bahan_bangunan.json", jsonData, 0644)
	} else {
		fmt.Println("Gagal mengubah data ke JSON:", err)
	}
}
