package supplier

import "bangunin/controllers/supplier/actions"

func Index(action_supplier int) {
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
	}
}
