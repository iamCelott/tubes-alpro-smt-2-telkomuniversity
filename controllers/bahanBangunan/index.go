package bahanBangunan

import "bangunin/controllers/bahanBangunan/actions"

func Index(action_bangunan int) {
	switch action_bangunan {
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
	}
}
