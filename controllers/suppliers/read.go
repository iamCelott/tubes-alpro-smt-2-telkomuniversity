package suppliers

import (
	"bangunin/models"
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
)

func ReadSuppliers() ([]models.Supplier, error) {
	data, err := os.ReadFile("storages/supplier.json")
	if err != nil {
		return nil, err
	}
	var suppliers []models.Supplier
	err = json.Unmarshal(data, &suppliers)
	if err != nil {
		return nil, err
	}
	return suppliers, nil
}

func Read() {
	fmt.Println("\n======= DATA SUPPLIER =======")
	suppliers, err := ReadSuppliers()
	if err != nil {
		fmt.Println("Gagal membaca data:", err)
		return
	}
	if len(suppliers) == 0 {
		fmt.Println("Belum ada data supplier.")
		return
	}
	for _, s := range suppliers {
		printSupplier(s)
	}

	fmt.Println("\n[1] Cari data supplier")
	fmt.Println("[2] Kembali ke menu utama")
	fmt.Print("Masukan aksi: ")
	var subAction int
	fmt.Scanln(&subAction)
	switch subAction {
	case 1:
		Search()
	case 2:
		return
	default:
		fmt.Println("Aksi tidak valid!")
	}
}

func Search() {
	suppliers, err := ReadSuppliers()
	if err != nil {
		fmt.Println("Gagal membaca data:", err)
		return
	}

	if len(suppliers) == 0 {
		fmt.Println("Belum ada data supplier.")
		return
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("\n======= PENCARIAN SUPPLIER =======")
	fmt.Println("Cari berdasarkan:")
	fmt.Println("1. Nama")
	fmt.Println("2. Alamat")
	fmt.Println("3. Kota")
	fmt.Println("4. Kabupaten")
	fmt.Println("5. Provinsi")
	fmt.Print("Pilih (1-5): ")
	fieldInput, _ := reader.ReadString('\n')
	fieldInput = strings.TrimSpace(fieldInput)

	var field string
	switch fieldInput {
	case "1":
		field = "nama"
	case "2":
		field = "alamat"
	case "3":
		field = "kota"
	case "4":
		field = "kabupaten"
	case "5":
		field = "provinsi"
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
		printSupplier(s)
	}
}

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
	case "alamat":
		return s.Alamat
	case "kota":
		return s.Kota
	case "kabupaten":
		return s.Kabupaten
	case "provinsi":
		return s.Provinsi
	}
	return ""
}

func printSupplier(s models.Supplier) {
	fmt.Println("--------------------------------")
	fmt.Println("ID         \t\t\t:", s.ID)
	fmt.Println("Nama       \t\t\t:", s.Nama)
	fmt.Println("Alamat     \t\t\t:", s.Alamat)
	fmt.Println("Kota       \t\t\t:", s.Kota)
	fmt.Println("Kabupaten  \t\t\t:", s.Kabupaten)
	fmt.Println("Provinsi   \t\t\t:", s.Provinsi)
	fmt.Println("No Telepon \t\t\t:", s.Telepon)
	fmt.Println("Email      \t\t\t:", s.Email)
	fmt.Println("Jumlah Pelayanan \t:", s.JumlahPelayanan)
	fmt.Println("Rating     \t\t\t:", fmt.Sprintf("%.1f", s.Rating))
	fmt.Println("--------------------------------")
}
