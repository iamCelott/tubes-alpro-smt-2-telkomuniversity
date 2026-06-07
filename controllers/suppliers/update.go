package suppliers

import (
	"bangunin/services"
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Update() {
	fmt.Println("\n======= UPDATE SUPPLIER =======")
	scanner := bufio.NewScanner(os.Stdin)

	suppliers, err := services.ReadSuppliers()
	if err != nil {
		fmt.Println("Gagal membaca data:", err)
		fmt.Println()
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
		fmt.Printf("Supplier tidak ditemukan.\n\n")
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

	fmt.Printf("Lokasi Baru [%s]: ", current.Lokasi)
	scanner.Scan()
	inputLokasi := scanner.Text()
	if strings.TrimSpace(inputLokasi) != "" {
		suppliers[indexFound].Lokasi = inputLokasi
	}

	fmt.Printf("No telpon baru [%s]: ", current.Telepon)
	scanner.Scan()
	inputTelepon := scanner.Text()
	if strings.TrimSpace(inputTelepon) != "" {
		suppliers[indexFound].Telepon = inputTelepon
	}

	fmt.Printf("Email baru [%s]: ", current.Email)
	scanner.Scan()
	inputEmail := scanner.Text()
	if strings.TrimSpace(inputEmail) != "" {
		suppliers[indexFound].Email = inputEmail
	}

	fmt.Printf("Rating [%.1f]: ", current.Rating)
	scanner.Scan()
	textRating := scanner.Text()
	if textRating != "" {
		inputRating, err := strconv.ParseFloat(textRating, 64)
		if err != nil {
			fmt.Printf("Error: Input yang dimasukkan bukan angka!\n\n")
			return
		}
		suppliers[indexFound].Rating = inputRating
	}

	// Simpan perubahan ke JSON
	updatedData, err := json.MarshalIndent(suppliers, "", "    ")
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

	fmt.Printf("\nData Supplier berhasil diperbarui!\n\n")
}
