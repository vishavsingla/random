package main

import (
	"log"
	"myproject/api"
	"myproject/storage"
	"myproject/worker"
	"net/http"
)

var workerPool *worker.WorkerPool

func init() {
	workerPool = worker.NewWorkerPool()
	workerPool.Start()

	// Load the Store Master file
	err := store.LoadStores("store_master.csv")
	if err != nil {
		log.Fatalf("Failed to load Store Master: %v", err)
	}
}

func main() {
	http.HandleFunc("/api/submit", api.SubmitJobHandler(workerPool))
	http.HandleFunc("/api/status", api.StatusHandler)

	log.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
