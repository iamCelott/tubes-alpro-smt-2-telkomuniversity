package main

import (
	"bangunin/controllers/bahanBangunan"
	"bangunin/controllers/supplier"
	"bangunin/utils"
	"fmt"
)

func main() {
	utils.Migrate()
	for {
		fmt.Println("---------- BangunIn ---------- ")
		var action int
		utils.PrintCustom(&action, []string{"Bahan Bangunan", "Supplier", "Keluar"})
		switch action {
		case 1:
			bahanBangunan.Index()
		case 2:
			supplier.Index()
		case 3:
			fmt.Println("Keluar!")
			return
		}
	}
}
