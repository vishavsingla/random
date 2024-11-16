package store

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Store struct {
	StoreID   string `json:"store_id"`
	StoreName string `json:"store_name"`
	AreaCode  string `json:"area_code"`
}

var stores map[string]Store

func LoadStores(filePath string) error {
	stores = make(map[string]Store)

	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open store master file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	_, err = reader.Read() // Skip the header row
	if err != nil {
		return fmt.Errorf("failed to read CSV header: %v", err)
	}

	records, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("failed to read CSV records: %v", err)
	}

	for _, record := range records {
		if len(record) < 3 {
			continue
		}
		store := Store{
			StoreID:   record[2],
			StoreName: record[1],
			AreaCode:  record[0],
		}
		stores[store.StoreID] = store
	}

	return nil
}

func GetStoreByID(storeID string) (Store, bool) {
	store, exists := stores[storeID]
	return store, exists
}
