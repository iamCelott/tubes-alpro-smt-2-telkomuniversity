package actions

import (
	"bufio"
	"encoding/json" // Wajib ditambahkan untuk membaca JSON
	"fmt"
	"os"            // Wajib ditambahkan untuk membuka file
	"strings"
)

func Detail() {
	// 1. Membaca langsung dari file JSON sebelum melakukan pencarian
	fileData, err := os.ReadFile("storages/bahan_bangunan.json")
	if err != nil {
		fmt.Println("Gagal membaca file JSON. Pastikan file ada di folder storages.")
		return
	}

	// 2. Menerjemahkan isi file JSON ke dalam variabel global DB_BahanBangunan
	errUnmarshal := json.Unmarshal(fileData, &DB_BahanBangunan)
	if errUnmarshal != nil {
		fmt.Println("Gagal membaca format JSON:", errUnmarshal)
		return
	}

	// 3. Validasi setelah data berhasil dimuat
	if len(DB_BahanBangunan) == 0 {
		fmt.Println("\nData bahan bangunan masih kosong.")
		return
	}

	// 4. Proses input kata kunci dari user
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\nMasukkan ID atau sebagian Nama Bahan yang ingin dicari: ")
	input, _ := reader.ReadString('\n')

	// Mengubah input menjadi huruf kecil agar pencarian tidak sensitif huruf besar/kecil
	keyword := strings.TrimSpace(strings.ToLower(input))

	ditemukan := 0

	fmt.Println("\n=== DETAIL BAHAN BANGUNAN ===")

	// 5. Melakukan pencarian menggunakan Sequential Search
	for _, bahan := range DB_BahanBangunan {

		namaBahan := strings.ToLower(bahan.Nama)
		idBahan := strings.ToLower(bahan.ID)

		// Menampilkan jika ID cocok persis, ATAU keyword ada di dalam nama bahan
		if idBahan == keyword || strings.Contains(namaBahan, keyword) {
			ditemukan++
			fmt.Printf("\n--- Hasil ke-%d ---\n", ditemukan)
			fmt.Printf("ID Bahan    : %s\n", bahan.ID)
			fmt.Printf("Nama Bahan  : %s\n", bahan.Nama)
			fmt.Printf("Kategori    : %s\n", bahan.Kategori)
			fmt.Printf("Jumlah Stok : %d %s\n", bahan.Stock, bahan.Satuan)
			fmt.Printf("Tipe Satuan : %s\n", bahan.TipeSatuan)
			fmt.Printf("Harga       : Rp %.2f\n", bahan.Harga)
			fmt.Printf("ID Supplier : %s\n", bahan.IDSupplier)
		}
	}

	if ditemukan == 0 {
		fmt.Println("Data tidak ditemukan.")
	} else {
		fmt.Printf("\nTotal ditemukan: %d bahan\n", ditemukan)
	}
}