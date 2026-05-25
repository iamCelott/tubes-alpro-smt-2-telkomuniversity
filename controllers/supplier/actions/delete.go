package actions

import (
	"bangunin/models"
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func Delete() {
	fmt.Println("\n======= HAPUS SUPPLIER =======")
	scanner := bufio.NewScanner(os.Stdin)

	suppliers, err := ReadSuppliers()
	if err != nil {
		fmt.Println("Gagal membaca data:", err)
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
		fmt.Println("Supplier dengan ID tersebut tidak ditemukan.")
		return
	}

	updatedData, err := json.MarshalIndent(updatedSuppliers, "", "    ")
	if err != nil {
		fmt.Println("Gagal memproses JSON:", err)
		return
	}

	err = os.WriteFile("storages/supplier.json", updatedData, 0644)
	if err != nil {
		fmt.Println("Gagal memperbarui file:", err)
		return
	}

	fmt.Println("\nSupplier berhasil dihapus!")
}
