package main

import (
	"flag"
	"fmt"
	"os"
	"seryn/src/internal/engine"
)

// Parse CLI flags (--workflow, --repo etc.)
// validate user input
// invoke engine to apply workflow
// call engine.ApplyWorkflow(...)

func main() {
	workflow := flag.String("workflow", "", "Git workflow to apply")
	repo := flag.String("repo", ".", "Path to the Git repository")
	flag.Parse()

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
