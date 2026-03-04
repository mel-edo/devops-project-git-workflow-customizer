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

		for _, r := range cfg.Repositories {
			fmt.Printf("Applying workflow '%s' to %s\n", cfg.Workflow, r)
			if err := engine.ApplyWorkflow(r, cfg.Workflow); err != nil {
				fmt.Printf("Error processing %s: %v\n", r, err)
			}
		}

		return
	}

	if *workflow == "" {
		fmt.Println("Error: --workflow flag is required")
		os.Exit(1)
	}

	err := engine.ApplyWorkflow(*repo, *workflow)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
