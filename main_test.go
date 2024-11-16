package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"myproject/api"
	"myproject/storage"
	"myproject/worker"
)

func TestSubmitJob(t *testing.T) {
	// Initialize the store master data
	err := store.LoadStores("store_master.csv")
	if err != nil {
		t.Fatalf("Failed to load stores: %v", err)
	}

	// Initialize a WorkerPool
	workerPool := worker.NewWorkerPool()
	workerPool.Start()

	// Test request payload
	body := `{
		"count": 1,
		"visits": [
			{
				"store_id": "RP00001",
				"image_url": ["https://www.gstatic.com/webp/gallery/2.jpg"],
				"visit_time": "2024-11-15T10:00:00Z"
			}
		]
	}`

	req, err := http.NewRequest("POST", "/api/submit", strings.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	// Pass the initialized WorkerPool to the handler
	handler := api.SubmitJobHandler(workerPool)
	handler.ServeHTTP(rr, req)

	// Validate the response status code
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}
}

func TestStoreLoad(t *testing.T) {
	// Test loading the store master CSV file
	err := store.LoadStores("store_master.csv")
	if err != nil {
		t.Fatalf("Failed to load stores: %v", err)
	}

	// Validate a known store
	store, exists := store.GetStoreByID("RP00001")
	if !exists || store.StoreName != "B P STORE" {
		t.Errorf("store not loaded correctly: got %v", store)
	}
}
