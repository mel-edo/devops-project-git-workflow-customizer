package engine

import (
	"fmt"
	"seryn/src/internal/logger"
	"sync"
)

type BatchResult struct {
	Repo    string
	Success bool
	Error   error
}

func ApplyWorkflowBatch(repos []string, workflowName string) []BatchResult {
	var wg sync.WaitGroup
	var mu sync.Mutex

	results := make([]BatchResult, 0, len(repos))

	for _, repo := range repos {
		wg.Add(1)

		go func(r string) {
			defer wg.Done()

			err := ApplyWorkflow(r, workflowName)

			mu.Lock()
			defer mu.Unlock()

			if err != nil {
				results = append(results, BatchResult{Repo: r, Success: false, Error: err})
			} else {
				results = append(results, BatchResult{Repo: r, Success: true})
			}
		}(repo)
	}

	wg.Wait()
	return results
}

func PrintBatchSummary(results []BatchResult) {
	fmt.Println()
	logger.Info("Batch processing complete")

	succeeded := 0
	failed := 0

	for _, r := range results {
		if r.Success {
			logger.Success(fmt.Sprintf("%-40s OK", r.Repo))
			succeeded++
		} else {
			logger.Warning(fmt.Sprintf("%-40s FAILED: %v", r.Repo, r.Error))
			failed++
		}
	}

	fmt.Println()
	logger.Info(fmt.Sprintf("Total: %d succeeded, %d failed", succeeded, failed))
}
