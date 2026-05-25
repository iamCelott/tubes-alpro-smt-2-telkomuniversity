package main

import (
	"bangunin/utils"
	"fmt"
)

func main() {
	utils.Migrate()
	var action int
	fmt.Println("---------- BangunIn ---------- ")
	utils.PrintCustom(&action, []string{"Users", "Bahan Bangunan", "Logout"})
	utils.PrintCRUD(&action)
}
