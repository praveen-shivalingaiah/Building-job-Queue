package main

import (
	"encoding/json"
	"fmt"
	"jobqueue/jobs"
	"log"
	"net/http"
)

const (
	port = ":8080"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /createjob", createJob)

	log.Println("Listening on", port)
	log.Fatal(http.ListenAndServe(port, mux))
}

func createJob(w http.ResponseWriter, r *http.Request) {
	var job jobs.Jobs

	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)
	if err := json.NewDecoder(r.Body).Decode(&job); err != nil {
		http.Error(w, "Unable to parse json", http.StatusBadRequest)
		return
	}

	fmt.Printf("type=%s payload=%s retries=%d\n", job.Type, job.Payload, job.MaxRetries)

	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(jobs.JobResponse{Status: "queued"})
}
