package actions

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Detail() {
	fmt.Println("\n======= DETAIL SUPPLIER =======")
	scanner := bufio.NewScanner(os.Stdin)

	suppliers, err := ReadSuppliers()
	if err != nil {
		fmt.Println("Gagal membaca data:", err)
		return
	}

	fmt.Print("Masukkan ID Supplier yang dicari: ")
	scanner.Scan()
	searchID := strings.TrimSpace(scanner.Text())

	found := false
	for _, s := range suppliers {
		if s.ID == searchID {
			fmt.Println("\n--- Data Ditemukan ---")
			fmt.Println("ID         :", s.ID)
			fmt.Println("Nama       :", s.Nama)
			fmt.Println("Alamat     :", s.Lokasi)
			fmt.Println("No Telepon :", s.Kontak)
			fmt.Println("----------------------")
			found = true
			break
		}
	}

	if !found {
		fmt.Println("Supplier dengan ID tersebut tidak ditemukan.")
	}
}
