package suppliers

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
		fmt.Printf("%-22s : %s\n", "ID", s.ID)
		fmt.Printf("%-22s : %s\n", "Nama", s.Nama)
		fmt.Printf("%-22s : %s\n", "Alamat", s.Alamat)
		fmt.Printf("%-22s : %s\n", "Kota", s.Kota)
		fmt.Printf("%-22s : %s\n", "Kabupaten", s.Kabupaten)
		fmt.Printf("%-22s : %s\n", "Provinsi", s.Provinsi)
		fmt.Printf("%-22s : %s\n", "No Telepon", s.Telepon)
		fmt.Printf("%-22s : %s\n", "Email", s.Email)
		fmt.Printf("%-22s : %d\n", "Jumlah Pelayanan", s.JumlahPelayanan)
		fmt.Printf("%-22s : %.1f\n", "Rating", s.Rating)
	}
	fmt.Println("--------------------------------")

}
