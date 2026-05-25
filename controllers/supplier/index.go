package supplier

import (
	"bangunin/controllers/supplier/actions"
	"fmt"
)

func Index() {
	var action_supplier int

	for {
		fmt.Println("\n--------- Supplier ---------")
		fmt.Println("[1] Read")
		fmt.Println("[2] Detail")
		fmt.Println("[3] Create")
		fmt.Println("[4] Update")
		fmt.Println("[5] Delete")
		fmt.Println("[0] Kembali ke Menu Utama")
		fmt.Print("Masukan aksi: ")
		fmt.Scanln(&action_supplier)

		switch action_supplier {
		case 1:
			actions.Read()
		case 2:
			actions.Detail()
		case 3:
			actions.Create()
		case 4:
			actions.Update()
		case 5:
			actions.Delete()
		case 0:
			fmt.Println("Kembali ke menu utama...")
			return 
		default:
			fmt.Println("Aksi tidak valid! Silakan pilih menu yang tersedia.")
		}
	}
}
