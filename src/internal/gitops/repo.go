package gitops

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// check if .git exists
// init repo using `git init`
// ret clear errors for invalid repo states

func EnsureRepo(path string) error {
	gitDir := filepath.Join(path, ".git")

	if _, err := os.Stat(gitDir); os.IsNotExist(err) {
		initCmd := exec.Command("git", "init", "-b", "main")
		initCmd.Dir = path

		out, err := initCmd.CombinedOutput()
		if err != nil {
			return fmt.Errorf("git init failed: %s", strings.TrimSpace(string(out)))
		}
	}

	return ensureInitialCommit(path)
}

func ensureInitialCommit(path string) error {
	var out []byte
	var err error

	checkCmd := exec.Command("git", "rev-parse", "--verify", "HEAD")
	checkCmd.Dir = path

	if err = checkCmd.Run(); err == nil {
		// repo already has at least one commit; nothing to do
		return nil
	}

	// Use a hidden placeholder instead of README.md so GenerateFiles
	// can create README.md cleanly
	placeholderPath := filepath.Join(path, ".seryn-init")
	if err := os.WriteFile(placeholderPath, []byte(""), 0644); err != nil {
		return err
	}

	addCmd := exec.Command("git", "add", ".")
	addCmd.Dir = path

	out, err = addCmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("git add failed: %s", strings.TrimSpace(string(out)))
	}

	commitCmd := exec.Command(
		"git",
		"commit",
		"-m",
		"Initial commit",
	)
	commitCmd.Dir = path

	out, err = commitCmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("git commit failed: %s", strings.TrimSpace(string(out)))
	}

	// Remove placeholder now that the initial commit exists
	if err := os.Remove(placeholderPath); err != nil {
		return fmt.Errorf("failed to remove placeholder: %w", err)
	}

	return nil
}
