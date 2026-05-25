package actions

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func Update() {
	fmt.Println("\n======= UPDATE SUPPLIER =======")
	scanner := bufio.NewScanner(os.Stdin)

	suppliers, err := ReadSuppliers()
	if err != nil {
		fmt.Println("Gagal membaca data:", err)
		return
	}

	fmt.Print("Masukkan ID Supplier yang ingin diubah: ")
	scanner.Scan()
	targetID := strings.TrimSpace(scanner.Text())

	indexFound := -1
	for i, s := range suppliers {
		if s.ID == targetID {
			indexFound = i
			break
		}
	}

	if indexFound == -1 {
		fmt.Println("Supplier tidak ditemukan.")
		return
	}

	current := suppliers[indexFound]
	fmt.Println("\n*Kosongkan (langsung Enter) jika tidak ingin mengubah data*")

	fmt.Printf("Nama Baru [%s]: ", current.Nama)
	scanner.Scan()
	inputNama := scanner.Text()
	if strings.TrimSpace(inputNama) != "" {
		suppliers[indexFound].Nama = inputNama
	}

	fmt.Printf("Alamat Baru [%s]: ", current.Lokasi)
	scanner.Scan()
	inputAlamat := scanner.Text()
	if strings.TrimSpace(inputAlamat) != "" {
		suppliers[indexFound].Lokasi = inputAlamat
	}

	fmt.Printf("No Telepon Baru [%s]: ", current.Kontak)
	scanner.Scan()
	inputTelp := scanner.Text()
	if strings.TrimSpace(inputTelp) != "" {
		suppliers[indexFound].Kontak = inputTelp
	}

	// Simpan perubahan ke JSON
	updatedData, err := json.MarshalIndent(suppliers, "", "    ")
	if err != nil {
		fmt.Println("Gagal memproses JSON:", err)
		return
	}

	err = os.WriteFile("storages/supplier.json", updatedData, 0644)
	if err != nil {
		fmt.Println("Gagal memperbarui file:", err)
		return
	}

	fmt.Println("\nData Supplier berhasil diperbarui!")
}
