package gitops

import (
	"fmt"
	"os/exec"
	"strings"
)

// list existing branches, create branch if it doesen't exist
// log warnings

func EnsureBranches(repoPath string, branches []string) ([]string, error) {
	// List existing branches
	listCmd := exec.Command("git", "branch", "--list")
	listCmd.Dir = repoPath

	out, err := listCmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf(
			"git branch --list failed: %s",
			strings.TrimSpace(string(out)),
		)
	}

	existing := make(map[string]bool)
	lines := strings.Split(string(out), "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		// Remove leading '*' from current branch
		line = strings.TrimPrefix(line, "* ")

		existing[line] = true
	}

	// Create required branches if missing
	for _, branch := range branches {
		if existing[branch] {
			continue
		}

		createCmd := exec.Command("git", "branch", branch)
		createCmd.Dir = repoPath

		out, err := createCmd.CombinedOutput()
		if err != nil {
			return nil, fmt.Errorf(
				"git branch %s failed: %s",
				branch,
				strings.TrimSpace(string(out)),
			)
		}

		existing[branch] = true
	}

	// identify unexpected branches
	required := make(map[string]bool)
	for _, b := range branches {
		required[b] = true
	}

	var unexpected []string
	for b := range existing {
		if !required[b] {
			unexpected = append(unexpected, b)
		}
	}

	return unexpected, nil
}
