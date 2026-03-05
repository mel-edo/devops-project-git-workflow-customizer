package main

import (
	"flag"
	"fmt"
	"os"
	"seryn/src/internal/config"
	"seryn/src/internal/engine"
)

// Parse CLI flags (--workflow, --repo etc.)
// validate user input
// invoke engine to apply workflow
// call engine.ApplyWorkflow(...)

func main() {
	workflow := flag.String("workflow", "", "Git workflow to apply")
	repo := flag.String("repo", ".", "Path to the Git repository")
	configPath := flag.String("config", "", "Path to a YAML config file")
	flag.Parse()

	if *configPath != "" {
		cfg, err := config.LoadConfig(*configPath)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		if len(cfg.Repositories) == 1 {
			if err := engine.ApplyWorkflowWithAlert(cfg.Repositories[0], cfg.Workflow, cfg.WebhookURL); err != nil {
				fmt.Println("Error:", err)
				os.Exit(1)
			}
			return
		}

		results := engine.ApplyWorkflowBatch(cfg.Repositories, cfg.Workflow)
		engine.PrintBatchSummary(results)
		return
	}

	if *workflow == "" {
		fmt.Println("Error: --workflow flag is required (or use --config)")
		os.Exit(1)
	}

	err := engine.ApplyWorkflow(*repo, *workflow)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
