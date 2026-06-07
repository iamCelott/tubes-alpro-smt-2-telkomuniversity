package suppliers

import (
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

	fmt.Printf("Alamat Baru [%s]: ", current.Alamat)
	scanner.Scan()
	inputAlamat := scanner.Text()
	if strings.TrimSpace(inputAlamat) != "" {
		suppliers[indexFound].Alamat = inputAlamat
	}

	fmt.Printf("Kota [%s]: ", current.Kota)
	scanner.Scan()
	inputKota := scanner.Text()
	if strings.TrimSpace(inputKota) != "" {
		suppliers[indexFound].Kota = inputKota
	}

	fmt.Printf("Kabupaten [%s]: ", current.Kabupaten)
	scanner.Scan()
	inputKabupaten := scanner.Text()
	if strings.TrimSpace(inputKabupaten) != "" {
		suppliers[indexFound].Kabupaten = inputKabupaten
	}

	fmt.Printf("Provinsi [%s]: ", current.Provinsi)
	scanner.Scan()
	inputProvinsi := scanner.Text()
	if strings.TrimSpace(inputProvinsi) != "" {
		suppliers[indexFound].Provinsi = inputProvinsi
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

	fmt.Printf("Jumlah Pelayanan [%d]: ", current.JumlahPelayanan)
	scanner.Scan()
	textJumlahPelayanan := scanner.Text()
	if textJumlahPelayanan != "" {
		inputJumlahPelayanan, err := strconv.Atoi(textJumlahPelayanan)
		if err != nil {
			fmt.Println("Error: Input yang dimasukkan bukan angka!")
			return
		}
		suppliers[indexFound].JumlahPelayanan = inputJumlahPelayanan
	}

	fmt.Printf("Rating [%.1f]: ", current.Rating)
	scanner.Scan()
	textRating := scanner.Text()
	if textRating != "" {
		inputRating, err := strconv.ParseFloat(textRating, 64)
		if err != nil {
			fmt.Println("Error: Input yang dimasukkan bukan angka!")
			return
		}
		suppliers[indexFound].Rating = inputRating
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
