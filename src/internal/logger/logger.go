package logger

// print progress, warnings, final summary of applied workflow

import (
	"fmt"
	"strings"
)

func Success(msg string) {
	fmt.Printf("✔ %s\n", msg)
}

func Warning(msg string) {
	fmt.Printf("⚠ %s\n", msg)
}

func Info(msg string) {
	fmt.Printf(" %s \n", msg)
}

func Summary(workflow string, branches []string, files []string, unexpected []string) {
	fmt.Println(strings.Repeat("-", 40))
	Success("Repository ready")
	Success(fmt.Sprintf("Workflow applied: %s", workflow))
	Success(fmt.Sprintf("Branches ensured: %v", branches))
	Success(fmt.Sprintf("Files handled: %v", files))

	if len(unexpected) > 0 {
		Warning(fmt.Sprintf("Extra branches detected: %v", unexpected))
	}

	fmt.Println(strings.Repeat("-", 40))
}
