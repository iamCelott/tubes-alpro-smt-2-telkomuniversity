package actions

import (
	"bangunin/models"
	"bufio"
	"encoding/json"
	"fmt"
	"os"
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
	newSupplier.Lokasi = scanner.Text()

	fmt.Print("Masukkan No Telepon: ")
	scanner.Scan()
	newSupplier.Kontak = scanner.Text()

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
