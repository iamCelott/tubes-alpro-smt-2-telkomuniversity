package actions

import (
	"bangunin/models"
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

var DB_BahanBangunan []models.BahanBangunan

func Create() {
	var baru models.BahanBangunan
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("\n=== TAMBAH BAHAN BANGUNAN ===")

	fmt.Print("ID Bahan       : ")
	fmt.Scanln(&baru.ID)

	fmt.Print("ID Supplier    : ")
	fmt.Scanln(&baru.IDSupplier)

	fmt.Print("Nama Bahan     : ")
	nama, _ := reader.ReadString('\n')
	baru.Nama = strings.TrimSpace(nama)

	fmt.Print("Kategori       : ")
	kategori, _ := reader.ReadString('\n')
	baru.Kategori = strings.TrimSpace(kategori)

	fmt.Print("Satuan         : ")
	satuan, _ := reader.ReadString('\n')
	baru.Satuan = strings.TrimSpace(satuan)

	fmt.Print("Tipe Satuan    : ")
	tipe, _ := reader.ReadString('\n')
	baru.TipeSatuan = strings.TrimSpace(tipe)

	fmt.Print("Jumlah Stok    : ")
	fmt.Scanln(&baru.Stock)

	fmt.Print("Harga          : ")
	fmt.Scanln(&baru.Harga)

	DB_BahanBangunan = append(DB_BahanBangunan, baru)
	fmt.Println(" Data bahan bangunan berhasil ditambahkan!")

	// --- KODE SIMPAN KE JSON LANGSUNG DI SINI ---
	jsonData, err := json.MarshalIndent(DB_BahanBangunan, "", "  ")
	if err == nil {
		// Pastikan path foldernya benar sesuai strukturmu
		os.WriteFile("storages/bahan_bangunan.json", jsonData, 0644)
	} else {
		fmt.Println("Gagal mengubah data ke JSON:", err)
	}
}
