package utils

import (
	"bangunin/models"
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func sequentialSearch(suppliers []models.Supplier, field, keyword string) []models.Supplier {
	var results []models.Supplier
	key := strings.ToLower(keyword)

	for _, s := range suppliers {
		value := strings.ToLower(getField(s, field))
		if strings.Contains(value, key) {
			results = append(results, s)
		}
	}
	return results
}

func binarySearch(suppliers []models.Supplier, field, keyword string) []models.Supplier {
	data := make([]models.Supplier, len(suppliers))
	copy(data, suppliers)

	sort.Slice(data, func(i, j int) bool {
		return strings.ToLower(getField(data[i], field)) < strings.ToLower(getField(data[j], field))
	})

	key := strings.ToLower(keyword)
	low, high := 0, len(data)-1
	found := -1

	for low <= high {
		mid := (low + high) / 2
		value := strings.ToLower(getField(data[mid], field))

		if value == key {
			found = mid
			break
		} else if value < key {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	if found == -1 {
		return nil
	}

	start := found
	for start > 0 && strings.ToLower(getField(data[start-1], field)) == key {
		start--
	}

	var results []models.Supplier
	for i := start; i < len(data) && strings.ToLower(getField(data[i], field)) == key; i++ {
		results = append(results, data[i])
	}
	return results
}

func getField(s models.Supplier, field string) string {
	switch field {
	case "nama":
		return s.Nama
	case "lokasi":
		return s.Lokasi
	}
	return ""
}

func SearchSuppliers(suppliers []models.Supplier) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("\n======= PENCARIAN SUPPLIER =======")
	fmt.Println("Cari berdasarkan:")
	fmt.Println("1. Nama")
	fmt.Println("2. Lokasi")
	fmt.Print("Pilih (1-5): ")
	fieldInput, _ := reader.ReadString('\n')
	fieldInput = strings.TrimSpace(fieldInput)

	var field string
	switch fieldInput {
	case "1":
		field = "nama"
	case "2":
		field = "lokasi"
	default:
		fmt.Println("Pilihan field tidak valid.")
		return
	}

	fmt.Println("\nMetode pencarian:")
	fmt.Println("1. Sequential Search (bisa kata sebagian)")
	fmt.Println("2. Binary Search (harus kata penuh)")
	fmt.Print("Pilih (1-2): ")
	algoInput, _ := reader.ReadString('\n')
	algoInput = strings.TrimSpace(algoInput)

	fmt.Print("Masukkan kata kunci: ")
	keyword, _ := reader.ReadString('\n')
	keyword = strings.TrimSpace(keyword)
	if keyword == "" {
		fmt.Println("Kata kunci tidak boleh kosong.")
		return
	}

	var results []models.Supplier
	switch algoInput {
	case "1":
		results = sequentialSearch(suppliers, field, keyword)
	case "2":
		results = binarySearch(suppliers, field, keyword)
	default:
		fmt.Println("Pilihan metode tidak valid.")
		return
	}

	fmt.Println("\n======= HASIL PENCARIAN =======")
	if len(results) == 0 {
		fmt.Println("Data tidak ditemukan.")
		return
	}

	fmt.Printf("Ditemukan %d data:\n", len(results))
	for _, s := range results {
		PrintSupplier(s)
	}
}
