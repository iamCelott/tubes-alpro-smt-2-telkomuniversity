package main

import (
	"bangunin/controllers/suppliers"
	"bangunin/utils"
	"fmt"
)

func main() {
	utils.Migrate()
	for {
		fmt.Println("---------- BangunIn ---------- ")
		var action int
		utils.PrintCustom(&action, []string{"Tampilkan semua data supplier", "Tampilkan detail data supplier", "Tambah data supplier", "Ubah data supplier", "Hapus data supplier", "Statistik supplier", "Keluar"})
		switch action {
		case 1:
			suppliers.Read()
		case 2:
			suppliers.Detail()
		case 3:
			suppliers.Create()
		case 4:
			suppliers.Update()
		case 5:
			suppliers.Delete()
		case 6:
			suppliers.Statistik()
		case 7:
			fmt.Println("Keluar!")
			return
		}
	}
}
