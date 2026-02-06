package generator

import (
	"fmt"
	"os"
	"path/filepath"
)

// create template content for each workflow
// write files if they don't exist (warn if they do)
// keep file generation idempotent

func GenerateFiles(repoPath string, workflowName string, contributionText string) error {

	// README.md
	readmePath := filepath.Join(repoPath, "README.md")
	if _, err := os.Stat(readmePath); os.IsNotExist(err) {
		content := fmt.Sprintf(
			"# Repository\n\nThis repository uses the **%s** Git workflow.\n",
			workflowName,
		)

		if err := os.WriteFile(readmePath, []byte(content), 0644); err != nil {
			return err
		}
	} else {
		fmt.Println("README.md already exists, skipping")
	}

	// .gitignore
	gitignorePath := filepath.Join(repoPath, ".gitignore")
	if _, err := os.Stat(gitignorePath); os.IsNotExist(err) {
		content := "bin/\n*.log\n"

		if err := os.WriteFile(gitignorePath, []byte(content), 0644); err != nil {
			return err
		}
	} else {
		fmt.Println(".gitignore already exists, skipping")
	}

	// CONTRIBUTING.md
	contributingPath := filepath.Join(repoPath, "CONTRIBUTING.md")
	if _, err := os.Stat(contributingPath); os.IsNotExist(err) {
		content := fmt.Sprintf(
			"# Contributing\n\n%s\n",
			contributionText,
		)

		if err := os.WriteFile(contributingPath, []byte(content), 0644); err != nil {
			return err
		}
	} else {
		fmt.Println("CONTRIBUTING.md already exists, skipping")
	}

	return nil
}
