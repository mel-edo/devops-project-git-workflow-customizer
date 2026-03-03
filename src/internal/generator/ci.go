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
		return `name: CI
		
on:
	push:
		branches: ["main"]
jobs:
	build:
		runs-on: ubuntu-latest
		steps:
			- uses: actions/checkout@v4
			- uses: actions/setup-go@v5
			  with:
			  	go-version: "1.25"
			- run: go mod tidy && git diff --exit-code
			- run: go test ./...
			- run: go build ./...
`
	case "pr-to-main":
		return `name: CI
on:
	pull-request:
		branches: ["main"]
jobs:
	build:
		runs-on: ubuntu-latest
		steps:
			- uses: actions/checkout@v4
			- uses: actions/setup-go@v5
			  with:
			  	go-version: "1.25"
			- run: go mod tidy && git diff --exit-code
			- run: go test ./...
			- run: go build ./...
`
	case "pr-to-develop":
		return `name: CI
on:
	pull_request:
		branches: ["develop"]
jobs:
	build:
		runs-on: ubuntu-latest
		steps:
			- uses: actions/checkout@v4
			- uses: actions/setup-go@v5
			  with:
			  	go-version: "1.25"
			- run: go mod tidy && git diff --exit-code
			- run: go test ./...
			- run: go build ./...
`
	default:
		return `name: CI
on:
	push:
		branches: ["main"]
	pull_request:
jobs:
	build:
		runs-on: ubuntu-latest
		steps:
			- uses: actions/checkout@v4
			- uses: actions/setup-go@v5
			  with:
			  	go-version: "1.25"
			- run: go mod tidy && git diff --exit-code
			- run: go test ./...
			- run: go build ./...
`
	}
}
