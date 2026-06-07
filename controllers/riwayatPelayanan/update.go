package riwayatPelayanan

import (
	"bangunin/services"
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func Update(supplierID string) {
	fmt.Println("\n======= UPDATE RIWAYAT PELAYANAN =======")
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

	fmt.Print("\nMasukkan ID Riwayat Pelayanan yang ingin diubah: ")
	scanner.Scan()
	targetID := strings.TrimSpace(scanner.Text())

	indexFound := -1
	for i, r := range riwayat {
		if r.ID == targetID && r.SupplierID == supplierID {
			indexFound = i
			break
		}
	}

	if indexFound == -1 {
		fmt.Printf("Riwayat Pelayanan dengan ID tersebut tidak ditemukan.\n\n")
		return
	}

	current := riwayat[indexFound]
	fmt.Println("\n*Kosongkan (langsung Enter) jika tidak ingin mengubah data*")

	fmt.Printf("Pelayanan Baru [%s]: ", current.Pelayanan)
	scanner.Scan()
	inputPelayanan := scanner.Text()
	if strings.TrimSpace(inputPelayanan) != "" {
		riwayat[indexFound].Pelayanan = inputPelayanan
	}

	updatedData, err := json.MarshalIndent(riwayat, "", "    ")
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

	fmt.Printf("\nRiwayat Pelayanan berhasil diperbarui!\n\n")
}
