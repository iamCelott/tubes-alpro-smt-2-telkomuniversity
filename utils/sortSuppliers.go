package utils

import (
	"bangunin/models"
	"fmt"
)

func insertionSort(data *[]models.Supplier, isAsc bool) {
	pData := *data
	for i := 1; i < len(pData); i++ {
		for j := i; j > 0; j-- {
			if isAsc {
				if pData[j].Rating < pData[j-1].Rating {
					pData[j], pData[j-1] = pData[j-1], pData[j]
				}
			} else {
				if pData[j].Rating > pData[j-1].Rating {
					pData[j], pData[j-1] = pData[j-1], pData[j]
				}
			}
		}
	}
}

func selectionSort(data *[]models.Supplier, isAsc bool) {
	pData := *data
	for i := range pData {
		var smallest int = i
		for j := i + 1; j < len(pData); j++ {
			if isAsc {
				if pData[j].Rating < pData[smallest].Rating {
					smallest = j
				}
			} else {
				if pData[j].Rating > pData[smallest].Rating {
					smallest = j
				}
			}
		}
		pData[i], pData[smallest] = pData[smallest], pData[i]
	}
}

func SortSuppliers(suppliers []models.Supplier) {
	fmt.Println("\n======= PENGURUTAN SUPPLIER =======")
	fmt.Println("Urutkan berdasarkan:")
	var ascCode int
	var isAsc bool
	PrintCustom(&ascCode, []string{"Terendah", "Tertinggi"})
	switch ascCode {
	case 1:
		isAsc = true
	case 2:
		isAsc = false
	}

	var methodCode int
	var isInsertionSort bool
	fmt.Println("\nMetode pengurutan:")
	PrintCustom(&methodCode, []string{"Insertion Sort", "Selection Sort"})
	switch methodCode {
	case 1:
		isInsertionSort = true
	case 2:
		isInsertionSort = false
	}

	if isInsertionSort {
		insertionSort(&suppliers, isAsc)
	} else {
		selectionSort(&suppliers, isAsc)
	}

	for _, s := range suppliers {
		PrintSupplier(s)
	}
}
