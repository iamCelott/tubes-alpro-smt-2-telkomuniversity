package utils

import (
	"bangunin/models"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func sequentialSearch(suppliers []models.Supplier, keyword string) []models.Supplier {
	keyword = strings.ToLower(keyword)
	var results []models.Supplier
	for _, s := range suppliers {
		if strings.Contains(strings.ToLower(s.Nama), keyword) ||
			strings.Contains(strings.ToLower(s.Lokasi), keyword) {
			results = append(results, s)
		}
	}
	return results
}

func binarySearch(suppliers []models.Supplier, keyword string) []models.Supplier {
	keyword = strings.ToLower(strings.TrimSpace(keyword))
	low := 0
	high := len(suppliers) - 1
	for low <= high {
		mid := (low + high) / 2
		nama := strings.ToLower(suppliers[mid].Nama)
		if nama == keyword {
			return []models.Supplier{suppliers[mid]}
		}
		if nama < keyword {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return []models.Supplier{}
}

func SearchSuppliers(suppliers []models.Supplier) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("\nMetode pencarian:")
	var algoInputCode int
	var isSequential bool
	PrintCustom(&algoInputCode, []string{"Sequential Search (bisa kata sebagian)", "Binary Search (harus kata penuh", "Kembali"})
	switch algoInputCode {
	case 1:
		isSequential = true
	case 2:
		isSequential = false

	}

	fmt.Print("Masukkan kata kunci: ")
	keyword, _ := reader.ReadString('\n')
	keyword = strings.TrimSpace(keyword)
	if keyword == "" {
		fmt.Printf("Kata kunci tidak boleh kosong.\n\n")
		return
	}

	var results []models.Supplier
	if isSequential {
		results = sequentialSearch(suppliers, keyword)
	} else {
		results = binarySearch(suppliers, keyword)
	}

	fmt.Println("\n======= HASIL PENCARIAN =======")
	if len(results) == 0 {
		fmt.Print("Data tidak ditemukan.\n\n")
		return
	}

	fmt.Printf("Ditemukan %d data:\n", len(results))
	for _, s := range results {
		PrintSupplier(s)
	}
	fmt.Println()
}
