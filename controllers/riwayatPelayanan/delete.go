package riwayatPelayanan

import (
	"bangunin/models"
	"bangunin/services"
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func Delete(supplierID string) {
	fmt.Println("\n======= HAPUS RIWAYAT PELAYANAN =======")
	scanner := bufio.NewScanner(os.Stdin)

	riwayat, err := services.ReadRiwayatPelayanan()
	if err != nil {
		fmt.Println("Gagal membaca data:", err)
		fmt.Println()
		return
	}

	fmt.Printf("Riwayat Pelayanan untuk Supplier ID %s:\n", supplierID)
	count := 0
	for _, r := range riwayat {
		if r.SupplierID == supplierID {
			fmt.Printf("  ID: %s | Pelayanan: %s\n", r.ID, r.Pelayanan)
			count++
		}
	}

	if count == 0 {
		fmt.Printf("Tidak ada riwayat pelayanan untuk supplier ini.\n\n")
		return
	}

	fmt.Print("\nMasukkan ID Riwayat Pelayanan yang ingin dihapus: ")
	scanner.Scan()
	targetID := strings.TrimSpace(scanner.Text())

	var updatedRiwayat []models.RiwayatPelayanan
	found := false

	for _, r := range riwayat {
		if r.ID == targetID && r.SupplierID == supplierID {
			found = true
			continue
		}
		updatedRiwayat = append(updatedRiwayat, r)
	}

	if !found {
		fmt.Printf("Riwayat Pelayanan dengan ID tersebut tidak ditemukan.\n\n")
		return
	}

	updatedData, err := json.MarshalIndent(updatedRiwayat, "", "    ")
	if err != nil {
		fmt.Println("Gagal memproses JSON:", err)
		fmt.Println()
		return
	}

	err = os.WriteFile("storages/riwayatPelayanan.json", updatedData, 0644)
	if err != nil {
		fmt.Println("Gagal memperbarui file:", err)
		fmt.Println()
		return
	}

	fmt.Printf("\nRiwayat Pelayanan berhasil dihapus!\n\n")
}
