package supplier

import (
	"bangunin/controllers/supplier/actions"
	"bangunin/utils"
	"fmt"
)

func Index() {
	for {
		var action_supplier int
		fmt.Println("---------- Supplier ---------- ")
		utils.PrintCRUD(&action_supplier)

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
		case 6:
			return
		}
	}
}
