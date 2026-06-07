package services

import (
	"bangunin/models"
	"encoding/json"
	"os"
)

func ReadRiwayatPelayanan() ([]models.RiwayatPelayanan, error) {
	data, err := os.ReadFile("storages/riwayatPelayanan.json")
	if err != nil {
		return nil, err
	}
	var riwayat []models.RiwayatPelayanan
	err = json.Unmarshal(data, &riwayat)
	if err != nil {
		return nil, err
	}
	return riwayat, nil
}
