package utils

import (
	"bangunin/models"
	"fmt"
)

func PrintCustom(action_var *int, labels []string) {
	for i, label := range labels {
		fmt.Printf("[%d] %s\n", i+1, label)
	}
	fmt.Print("Masukan aksi: ")
	fmt.Scanln(action_var)
	if *action_var < 1 || *action_var > len(labels) {
		fmt.Println("Aksi yang anda masukan tidak valid!")
		PrintCustom(action_var, labels)
	}
}

func PrintSupplier(s models.Supplier) {
	fmt.Println("--------------------------------")
	fmt.Println("ID         \t\t\t:", s.ID)
	fmt.Println("Nama       \t\t\t:", s.Nama)
	fmt.Println("Alamat     \t\t\t:", s.Alamat)
	fmt.Println("Kota       \t\t\t:", s.Kota)
	fmt.Println("Kabupaten  \t\t\t:", s.Kabupaten)
	fmt.Println("Provinsi   \t\t\t:", s.Provinsi)
	fmt.Println("No Telepon \t\t\t:", s.Telepon)
	fmt.Println("Email      \t\t\t:", s.Email)
	fmt.Println("Jumlah Pelayanan \t\t:", s.JumlahPelayanan)
	fmt.Println("Rating     \t\t\t:", fmt.Sprintf("%.1f", s.Rating))
	fmt.Println("--------------------------------")
}
