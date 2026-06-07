package suppliers

import (
	"bangunin/models"
	"bangunin/services"
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

func Create() {
	fmt.Println("\n======= TAMBAH SUPPLIER =======")
	scanner := bufio.NewScanner(os.Stdin)

	suppliers, err := services.ReadSuppliers()
	if err != nil {
		fmt.Println("Gagal membaca data lama:", err)
		fmt.Println()
		return
	}

	var newSupplier models.Supplier

	fmt.Print("Masukkan ID Supplier: ")
	scanner.Scan()
	newSupplier.ID = scanner.Text()

	fmt.Print("Masukkan Nama: ")
	scanner.Scan()
	newSupplier.Nama = scanner.Text()

	fmt.Print("Masukkan Lokasi: ")
	scanner.Scan()
	newSupplier.Lokasi = scanner.Text()

	fmt.Print("Masukkan No Telepon: ")
	scanner.Scan()
	newSupplier.Telepon = scanner.Text()

	fmt.Print("Masukkan Email: ")
	scanner.Scan()
	newSupplier.Email = scanner.Text()

	fmt.Print("Masukkan Rating: ")
	scanner.Scan()

	ratingNilai, err := strconv.ParseFloat(scanner.Text(), 64)
	if err != nil {
		fmt.Printf("Error: Input harus berupa angka desimal (contoh: 4.5)!\n\n")
		return
	}

	newSupplier.Rating = ratingNilai
	suppliers = append(suppliers, newSupplier)

	updatedData, err := json.MarshalIndent(suppliers, "", "    ")
	if err != nil {
		fmt.Println("Gagal memproses JSON:", err)
		fmt.Println()
		return
	}

	err = os.WriteFile("storages/supplier.json", updatedData, 0644)
	if err != nil {
		fmt.Println("Gagal menyimpan ke file:", err)
		fmt.Println()
		return
	}

	fmt.Printf("\nSupplier berhasil ditambahkan!\n\n")
}
