package generator

import (
	"fmt"
	"os"
	"path/filepath"
)

// generate CI YAML based on workflow
// support diff ci triggers (push, pull)
// ensure dir exists before writing

func GenerateCI(repoPath string, ciTrigger string) error {
	dirPath := filepath.Join(repoPath, ".github", "workflows")

	if err := os.MkdirAll(dirPath, 0755); err != nil {
		return fmt.Errorf("failed to create workflows directory: %w", err)
	}

	ciPath := filepath.Join(dirPath, "ci.yml")

	if _, err := os.Stat(ciPath); err == nil {
		fmt.Println("ci.yml already exists, skipping")
		return nil
	}

	content := buildCIContent(ciTrigger)

	if err := os.WriteFile(ciPath, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to write ci.yml: %w", err)
	}

	return nil
}

func buildCIContent(ciTrigger string) string {
	switch ciTrigger {
	case "push-to-main":
		return "name: CI\n\non:\n  push:\n    branches: [\"main\"]\n\njobs:\n  build:\n    runs-on: ubuntu-latest\n    steps:\n      - uses: actions/checkout@v4\n      - run: echo \"CI running\"\n"
	case "pr-to-main":
		return "name: CI\n\non:\n  pull_request:\n    branches: [\"main\"]\n\njobs:\n  build:\n    runs-on: ubuntu-latest\n    steps:\n      - uses: actions/checkout@v4\n      - run: echo \"CI running\"\n"
	case "pr-to-develop":
		return "name: CI\n\non:\n  pull_request:\n    branches: [\"develop\"]\n\njobs:\n  build:\n    runs-on: ubuntu-latest\n    steps:\n      - uses: actions/checkout@v4\n      - run: echo \"CI running\"\n"
	default:
		return "name: CI\n\non:\n  push:\n    branches: [\"main\"]\n  pull_request:\n\njobs:\n  build:\n    runs-on: ubuntu-latest\n    steps:\n      - uses: actions/checkout@v4\n      - run: echo \"CI running\"\n"
	}
}
