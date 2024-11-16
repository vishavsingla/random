package api

import (
	"encoding/json"
	"myproject/job"
	"net/http"
	"strconv"
)

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	jobIDStr := r.URL.Query().Get("jobid")
	jobID, err := strconv.Atoi(jobIDStr)
	if err != nil || jobID == 0 {
		http.Error(w, "Invalid job ID", http.StatusBadRequest)
		return
	}

	jobStatus, err := job.GetJobStatus(jobID)
	if err != nil {
		http.Error(w, "Job not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(jobStatus)
}
