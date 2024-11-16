package job

import (
	"errors"
	"fmt"
	"image"
	"math/rand"
	"net/http"
	"sync"
	"time"
	_ "image/jpeg"
	_ "image/png"
)

type ImageResult struct {
	URL       string `json:"url"`
	Perimeter int    `json:"perimeter"`
	Status    string `json:"status"`
	Error     string `json:"error,omitempty"`
}

type Visit struct {
	StoreID   string        `json:"store_id"`
	ImageURLs []string      `json:"image_url"`
	VisitTime string        `json:"visit_time"`
	Results   []ImageResult `json:"results"`
}

type Job struct {
	ID     int
	Visits []Visit
	Status string // ongoing, completed, failed
}

var jobs = make(map[int]*Job)
var mu sync.Mutex

func CreateJob(visits []Visit) (int, error) {
	mu.Lock()
	defer mu.Unlock()

	jobID := rand.Intn(1000)
	job := &Job{ID: jobID, Visits: visits, Status: "ongoing"}
	jobs[jobID] = job
	return jobID, nil
}

func ProcessJob(jobID int) {
	job := jobs[jobID]
	for vi, visit := range job.Visits {
		for _, url := range visit.ImageURLs {
			result := processImage(url)
			job.Visits[vi].Results = append(job.Visits[vi].Results, result)
		}
	}

	job.Status = "completed"
	for _, visit := range job.Visits {
		for _, result := range visit.Results {
			if result.Status == "failed" {
				job.Status = "failed"
				return
			}
		}
	}
}

func processImage(url string) ImageResult {
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != 200 {
		return ImageResult{URL: url, Status: "failed", Error: fmt.Sprintf("Failed to download: %v", err)}
	}
	defer resp.Body.Close()

	img, _, err := image.Decode(resp.Body)
	if err != nil {
		return ImageResult{URL: url, Status: "failed", Error: fmt.Sprintf("Decode error: %v", err)}
	}

	bounds := img.Bounds()
	width, height := bounds.Dx(), bounds.Dy()
	perimeter := 2 * (width + height)

	time.Sleep(time.Duration(100+rand.Intn(300)) * time.Millisecond)

	return ImageResult{URL: url, Perimeter: perimeter, Status: "completed"}
}

func GetJobStatus(jobID int) (map[string]interface{}, error) {
	mu.Lock()
	defer mu.Unlock()

	job, exists := jobs[jobID]
	if !exists {
		return nil, errors.New("job not found")
	}

	return map[string]interface{}{
		"job_id": job.ID,
		"status": job.Status,
		"visits": job.Visits,
	}, nil
}
