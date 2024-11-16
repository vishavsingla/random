package api

import (
	"encoding/json"
	"myproject/job"
	"myproject/storage"
	"myproject/worker"
	"net/http"
)

type SubmitRequest struct {
	Count  int         `json:"count"`
	Visits []job.Visit `json:"visits"`
}

func SubmitJobHandler(workerPool *worker.WorkerPool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req SubmitRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		if req.Count != len(req.Visits) {
			http.Error(w, "Count mismatch with visits", http.StatusBadRequest)
			return
		}

		// Validate Store IDs
		for _, visit := range req.Visits {
			if _, exists := store.GetStoreByID(visit.StoreID); !exists {
				http.Error(w, "Invalid store ID: "+visit.StoreID, http.StatusBadRequest)
				return
			}
		}

		jobID, err := job.CreateJob(req.Visits)
		if err != nil {
			http.Error(w, "Error creating job", http.StatusInternalServerError)
			return
		}

		workerPool.Enqueue(jobID)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]int{"job_id": jobID})
	}
}
