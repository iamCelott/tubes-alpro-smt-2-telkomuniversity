package riwayatPelayanan

import (
	"bangunin/models"
	"bangunin/services"
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func Create(supplierID string) {
	fmt.Println("\n======= TAMBAH RIWAYAT PELAYANAN =======")
	scanner := bufio.NewScanner(os.Stdin)

	riwayat, err := services.ReadRiwayatPelayanan()
	if err != nil {
		fmt.Println("Gagal membaca data lama:", err)
		fmt.Println()
		return
	}

	var newRiwayat models.RiwayatPelayanan

	fmt.Print("Masukkan ID: ")
	scanner.Scan()
	newRiwayat.ID = scanner.Text()

	newRiwayat.SupplierID = supplierID

	fmt.Print("Masukkan Pelayanan: ")
	scanner.Scan()
	newRiwayat.Pelayanan = scanner.Text()

	if newRiwayat.Pelayanan == "" {
		fmt.Printf("Error: Pelayanan tidak boleh kosong!\n\n")
		return
	}

	riwayat = append(riwayat, newRiwayat)

	updatedData, err := json.MarshalIndent(riwayat, "", "    ")
	if err != nil {
		fmt.Println("Gagal memproses JSON:", err)
		fmt.Println()
		return
	}

	err = os.WriteFile("storages/riwayatPelayanan.json", updatedData, 0644)
	if err != nil {
		fmt.Println("Gagal menyimpan ke file:", err)
		fmt.Println()
		return
	}

	fmt.Printf("\nRiwayat Pelayanan berhasil ditambahkan!\n\n")
}
