package main

import (
	"bangunin/utils"
	"fmt"
)

func main() {
	var action int
	fmt.Println("---------- BangunIn ---------- ")
	utils.PrintCustom(&action, []string{"Profile", "Settings", "Logout"})
}
