package jobs

import "encoding/json"

type Jobs struct {
	Type       string          `json:"type"`
	Payload    json.RawMessage `json:"payload"`
	MaxRetries int             `json:"max_retries"`
}

type JobResponse struct {
	Status string `json:"status"`
}
