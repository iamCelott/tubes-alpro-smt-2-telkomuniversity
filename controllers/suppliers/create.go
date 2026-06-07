package suppliers

import (
	"bangunin/models"
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

func Create() {
	fmt.Println("\n======= TAMBAH SUPPLIER =======")
	scanner := bufio.NewScanner(os.Stdin)

	suppliers, err := ReadSuppliers()
	if err != nil {
		fmt.Println("Gagal membaca data lama:", err)
		return
	}

	var newSupplier models.Supplier

	newSupplier.ID = fmt.Sprintf("%d", len(suppliers)+1)
	fmt.Printf("ID Supplier Otomatis: %s\n", newSupplier.ID)

	fmt.Print("Masukkan Nama Supplier: ")
	scanner.Scan()
	newSupplier.Nama = scanner.Text()

	fmt.Print("Masukkan Alamat: ")
	scanner.Scan()
	newSupplier.Alamat = scanner.Text()

	fmt.Print("Masukkan Kota: ")
	scanner.Scan()
	newSupplier.Kota = scanner.Text()

	fmt.Print("Masukkan Kabupaten: ")
	scanner.Scan()
	newSupplier.Kabupaten = scanner.Text()

	fmt.Print("Masukkan Provinsi: ")
	scanner.Scan()
	newSupplier.Provinsi = scanner.Text()

	fmt.Print("Masukkan No Telepon: ")
	scanner.Scan()
	newSupplier.Telepon = scanner.Text()

	fmt.Print("Masukkan Email: ")
	scanner.Scan()
	newSupplier.Email = scanner.Text()

	fmt.Print("Masukkan Jumlah Pelayanan: ")
	scanner.Scan()
	angka, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("Error: Input yang dimasukkan bukan angka!")
		return
	}
	newSupplier.JumlahPelayanan = angka

	fmt.Print("Masukkan Rating: ")
	scanner.Scan()

	ratingNilai, err := strconv.ParseFloat(scanner.Text(), 64)
	if err != nil {
		fmt.Println("Error: Input harus berupa angka desimal (contoh: 4.5)!")
		return
	}

	newSupplier.Rating = ratingNilai
	suppliers = append(suppliers, newSupplier)

	updatedData, err := json.MarshalIndent(suppliers, "", "    ")
	if err != nil {
		fmt.Println("Gagal memproses JSON:", err)
		return
	}

	err = os.WriteFile("storages/supplier.json", updatedData, 0644)
	if err != nil {
		fmt.Println("Gagal menyimpan ke file:", err)
		return
	}

	fmt.Println("\nSupplier berhasil ditambahkan!")
}
