package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func Migrate() {
	folderPath := "storages"

	fileNames := []string{
		"bahan_bangunan.json",
		"supplier.json",
	}

	err := os.MkdirAll(folderPath, 0755)
	if err != nil {
		fmt.Println("Gagal membuat folder storages:", err)
		return
	}

	for _, fileName := range fileNames {
		filePath := filepath.Join(folderPath, fileName)

		_, err := os.Stat(filePath)
		if os.IsNotExist(err) {
			initialData := []byte("[]")
			err = os.WriteFile(filePath, initialData, 0644)
			if err != nil {
				fmt.Printf("Gagal membuat file %s: %v\n", fileName, err)
				continue
			}
		}
	}
}
