package suppliers

import (
	"bangunin/controllers/riwayatPelayanan"
	"bangunin/services"
	"bangunin/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Detail() {
	fmt.Println("\n======= DETAIL SUPPLIER =======")
	scanner := bufio.NewScanner(os.Stdin)

	suppliers, err := services.ReadSuppliers()
	if err != nil {
		fmt.Println("Gagal membaca data:", err)
		fmt.Println()
		return
	}

	fmt.Print("Masukkan ID Supplier yang dicari: ")
	scanner.Scan()
	searchID := strings.TrimSpace(scanner.Text())

	found := false
	for _, s := range suppliers {
		if s.ID == searchID {
			fmt.Println("\n--- Data Ditemukan ---")
			utils.PrintSupplier(s)
			fmt.Println()
			found = true
			break
		}
	}

	if !found {
		fmt.Printf("Supplier dengan ID tersebut tidak ditemukan.\n\n")
		return
	}

	var action int
	utils.PrintCustom(&action, []string{"Tambah riwayat pelayanan", "Ubah riwayat pelayanan", "Hapus riwayat pelayanan", "Kembali"})
	switch action {
	case 1:
		riwayatPelayanan.Create(searchID)
	case 2:
		riwayatPelayanan.Update(searchID)
	case 3:
		riwayatPelayanan.Delete(searchID)
	case 4:
		fmt.Printf("Kembali!\n\n")
		return
	}
}
