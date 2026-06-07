package services

import (
	"bangunin/models"
	"encoding/json"
	"os"
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
