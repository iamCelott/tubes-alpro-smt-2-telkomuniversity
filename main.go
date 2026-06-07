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
		utils.PrintCustom(&action, []string{"Tampilkan data suppliers", "Tambah data supplier", "Ubah data supplier", "Hapus data supplier", "Keluar"})
		switch action {
		case 1:
			suppliers.Read()
		case 2:
			suppliers.Create()
		case 3:
			suppliers.Update()
		case 4:
			suppliers.Delete()
		case 5:
			fmt.Println("Keluar!")
			return
		}
	}
}
