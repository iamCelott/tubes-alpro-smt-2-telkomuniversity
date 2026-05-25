package main

import (
	"bangunin/controllers/bahanBangunan"
	"bangunin/controllers/supplier"
	"bangunin/utils"
	"fmt"
)

func main() {
	utils.Migrate()
	var action int
	fmt.Println("---------- BangunIn ---------- ")
	utils.PrintCustom(&action, []string{"Bahan Bangunan", "Supplier", "Keluar"})
	switch action {
	case 1:
		var action_bangunan int
		fmt.Println("---------- Bahan Bangunan ---------- ")
		utils.PrintCRUD(&action_bangunan)
		bahanBangunan.Index(action_bangunan)
	case 2:
		var action_supplier int
		fmt.Println("---------- Supplier ---------- ")
		utils.PrintCRUD(&action_supplier)
		supplier.Index()
	case 3:
		fmt.Println("Keluar!")
		return
	}
}
