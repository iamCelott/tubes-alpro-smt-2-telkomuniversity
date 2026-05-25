package actions

import (
	"bangunin/models"
	"encoding/json"
	"fmt"
	"os"
)

func ReadSuppliers() ([]models.Supplier, error) {
	filePath := "storages/supplier.json"
	var suppliers []models.Supplier

	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &suppliers)
	if err != nil {
		return nil, err
	}

	return suppliers, nil
}

func Read() {
	fmt.Println("\n======= DAFTAR SUPPLIER =======")
	suppliers, err := ReadSuppliers()
	if err != nil {
		fmt.Println("Gagal membaca data:", err)
		return
	}

	if len(suppliers) == 0 {
		fmt.Println("Belum ada data supplier.")
		return
	}

	for _, s := range suppliers {
		fmt.Println("--------------------------------")
		fmt.Println("ID         :", s.ID)
		fmt.Println("Nama       :", s.Nama)
		fmt.Println("Alamat     :", s.Lokasi)
		fmt.Println("No Telepon :", s.Kontak)
	}
	fmt.Println("--------------------------------")

}
