package utils

import "fmt"

func PrintCRUD(action_var *int) {
	var labels = []string{"Read", "Detail", "Create", "Update", "Delete", "Kembali"}
	for i, label := range labels {
		fmt.Printf("[%d] %s\n", i+1, label)
	}
	fmt.Print("Masukan aksi: ")
	fmt.Scanln(action_var)
	if *action_var < 1 || *action_var > len(labels) {
		fmt.Println("Aksi yang anda masukan tidak valid!")
		PrintCRUD(action_var)
	}
}
func PrintCustom(action_var *int, labels []string) {
	for i, label := range labels {
		fmt.Printf("[%d] %s\n", i+1, label)
	}
	fmt.Print("Masukan aksi: ")
	fmt.Scanln(action_var)
	if *action_var < 1 || *action_var > len(labels) {
		fmt.Println("Aksi yang anda masukan tidak valid!")
		PrintCustom(action_var, labels)
	}
}
