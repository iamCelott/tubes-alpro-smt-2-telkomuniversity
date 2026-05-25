package bahanBangunan

import (
	"bangunin/controllers/bahanBangunan/actions"
	"bangunin/utils"
	"fmt"
)

func Index() {
	for {
		var actionBangunan int
		fmt.Println("---------- Bahan Bangunan ---------- ")
		utils.PrintCRUD(&actionBangunan)

		switch actionBangunan {
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
