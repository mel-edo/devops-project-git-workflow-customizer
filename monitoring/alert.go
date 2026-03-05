package monitoring

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type AlertPayload struct {
	Repository string `json:"repository"`
	Workflow   string `json:"workflow"`
	Status     string `json:"status"`
	Timestamp  string `json:"timestamp"`
}

func SendAlert(webhookURL string, repo string, workflowName string, status string) error {
	if webhookURL == "" {
		return nil
	}

	payload := AlertPayload{
		Repository: repo,
		Workflow:   workflowName,
		Status:     status,
		Timestamp:  time.Now().UTC().Format(time.RFC3339),
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal alert payload: %w", err)
	}

	client := &http.Client{Timeout: 5 * time.Second}

	resp, err := client.Post(webhookURL, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("failed to send alert: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("alert webhook returned non-2xx status: %d", resp.StatusCode)
	}

	return nil
}
