package suppliers

import (
	"bangunin/services"
	"bangunin/utils"
	"fmt"
)

func Read() {
	fmt.Println("\n======= DATA SUPPLIER =======")
	suppliers, err := services.ReadSuppliers()
	if err != nil {
		fmt.Println("Gagal membaca data:", err)
		fmt.Println()
		return
	}
	if len(suppliers) == 0 {
		fmt.Printf("Belum ada data supplier.\n\n")
		return
	}
	for _, s := range suppliers {
		utils.PrintSupplier(s)
	}

	var subAction int
	utils.PrintCustom(&subAction, []string{"Cari data supplier", "Urutkan data supplier berdasarkan rating", "Kembali"})
	switch subAction {
	case 1:
		utils.SearchSuppliers(suppliers)
	case 2:
		utils.SortSuppliers(suppliers)
	case 3:
		fmt.Printf("Kembali!\n\n")
		return
	}
}
