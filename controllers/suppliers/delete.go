package suppliers

import (
	"bangunin/models"
	"bangunin/services"
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func Delete() {
	fmt.Println("\n======= HAPUS SUPPLIER =======")
	scanner := bufio.NewScanner(os.Stdin)

	suppliers, err := services.ReadSuppliers()
	if err != nil {
		fmt.Println("Gagal membaca data:", err)
		fmt.Println()
		return
	}

	fmt.Print("Masukkan ID Supplier yang ingin dihapus: ")
	scanner.Scan()
	targetID := strings.TrimSpace(scanner.Text())

	var updatedSuppliers []models.Supplier
	found := false

	for _, s := range suppliers {
		if s.ID == targetID {
			found = true
			continue
		}
		updatedSuppliers = append(updatedSuppliers, s)
	}

	if !found {
		fmt.Printf("Supplier dengan ID tersebut tidak ditemukan.\n\n")
		return
	}

	updatedData, err := json.MarshalIndent(updatedSuppliers, "", "    ")
	if err != nil {
		fmt.Println("Gagal memproses JSON:", err)
		fmt.Println()
		return
	}

	err = os.WriteFile("storages/supplier.json", updatedData, 0644)
	if err != nil {
		fmt.Println("Gagal memperbarui file:", err)
		fmt.Println()
		return
	}

	fmt.Printf("\nSupplier berhasil dihapus!\n\n")
}
