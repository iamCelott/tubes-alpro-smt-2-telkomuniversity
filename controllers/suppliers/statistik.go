package suppliers

import (
	"bangunin/models"
	"bangunin/services"
	"fmt"
)

func Statistik() {
	fmt.Printf("\n===== STATISTIK SUPPLIER =====\n")

	suppliers, err := services.ReadSuppliers()
	if err != nil {
		fmt.Println("Belum ada data supllier:", err)
		fmt.Println()
		return
	}

	totalSupplier := len(suppliers)
	if totalSupplier == 0 {
		fmt.Println("\nBelum ada data supplier.")
		return
	}

	wilayahMap := make(map[string]int)
	var tertinggi models.Supplier
	var terendah models.Supplier
	var totalRating float64

	for i, s := range suppliers {
		wilayahMap[s.Lokasi]++
		totalRating += s.Rating

		if i == 0 {
			tertinggi = s
			terendah = s
			continue
		}

		if s.Rating > tertinggi.Rating {
			tertinggi = s
		}
		if s.Rating < terendah.Rating {
			terendah = s
		}
	}

	rataRataRating := totalRating / float64(totalSupplier)

	fmt.Println("Jumlah Supplier per Wilayah:")
	for kota, jumlah := range wilayahMap {
		fmt.Printf("%-12s : %d supplier\n", kota, jumlah)
	}

	fmt.Println("-----------------------------------------")
	fmt.Printf("%-35s : %d\n", "Total Supplier", totalSupplier)
	fmt.Printf("%-35s : %.2f\n", "Rata-rata Rating Mitra", rataRataRating)

	teksTertinggi := fmt.Sprintf("%s (%.1f)", tertinggi.Nama, tertinggi.Rating)
	fmt.Printf("%-35s : %s\n", "Supplier dengan Rating Tertinggi", teksTertinggi)

	teksTerendah := fmt.Sprintf("%s (%.1f)", terendah.Nama, terendah.Rating)
	fmt.Printf("%-35s : %s", "Supplier dengan Rating Terendah", teksTerendah)
	fmt.Printf("\n-----------------------------------------\n\n")
}
