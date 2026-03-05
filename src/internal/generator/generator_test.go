package generator_test

import (
	"os"
	"path/filepath"
	"seryn/src/internal/generator"
	"strings"
	"testing"
)

func TestGenerateFiles_CreateAllFiles(t *testing.T) {
	dir := t.TempDir()

	err := generator.GenerateFiles(dir, "gitflow", "Open a PR against develop.")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := []string{"README.md", ".gitignore", "CONTRIBUTING.md"}
	for _, f := range expected {
		path := filepath.Join(dir, f)
		if _, err := os.Stat(path); os.IsNotExist(err) {
			t.Errorf("expected file '%s' to exist, but it does not", f)
		}
	}
}

func TestGenerateFiles_ReadmeContainsWorkflow(t *testing.T) {
	dir := t.TempDir()

	err := generator.GenerateFiles(dir, "gitflow", "Open a PR against develop.")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	data, err := os.ReadFile(filepath.Join(dir, "README.md"))
	if err != nil {
		t.Fatalf("failed to read README.md: %v", err)
	}

	if !strings.Contains(string(data), "gitflow") {
		t.Errorf("expected README.md to contain workflow name 'gitflow'")
	}
}

func TestGenerateFiles_ContributingContainsGuidelines(t *testing.T) {
	dir := t.TempDir()
	guidelines := "Open a PR against develop."

	err := generator.GenerateFiles(dir, "gitflow", guidelines)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	data, err := os.ReadFile(filepath.Join(dir, "CONTRIBUTING.md"))
	if err != nil {
		t.Fatalf("failed to read CONTRIBUTING.md: %v", err)
	}

	if !strings.Contains(string(data), guidelines) {
		t.Errorf("expected CONTRIBUTING.md to contain guidelines text")
	}
}

func TestGenerateFiles_Idempotent(t *testing.T) {
	dir := t.TempDir()

	if err := generator.GenerateFiles(dir, "gitflow", "Some guidelines."); err != nil {
		t.Fatalf("first call failed: %v", err)
	}

	readmePath := filepath.Join(dir, "README.md")
	original, _ := os.ReadFile(readmePath)

	if err := generator.GenerateFiles(dir, "feature", "Different guidelines."); err != nil {
		t.Fatalf("second call failed: %v", err)
	}

	after, _ := os.ReadFile(readmePath)
	if string(original) != string(after) {
		t.Errorf("README.md was overwritten on second call, expected idempotent behaviour")
	}
}

func TestGenerateCI_CreatesFile(t *testing.T) {
	dir := t.TempDir()

	err := generator.GenerateCI(dir, "pr-to-develop")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	ciPath := filepath.Join(dir, ".github", "workflows", "ci.yml")
	if _, err := os.Stat(ciPath); os.IsNotExist(err) {
		t.Errorf("expected ci.yml to exist as %s", ciPath)
	}
}

func TestGenerateCI_ContainsTrigger(t *testing.T) {
	cases := []struct {
		trigger  string
		expected string
	}{
		{"push-to-main", `branches: ["main"]`},
		{"pr-to-main", `branches: ["main"]`},
		{"pr-to-develop", `branches: ["develop"]`},
	}

	for _, tc := range cases {
		dir := t.TempDir()

		if err := generator.GenerateCI(dir, tc.trigger); err != nil {
			t.Fatalf("trigger '%s': unexpected error: %v", tc.trigger, err)
		}

		data, err := os.ReadFile(filepath.Join(dir, ".github", "workflows", "ci.yml"))
		if err != nil {
			t.Fatalf("trigger '%s': failed to read ci.yml: %v", tc.trigger, err)
		}

		if !strings.Contains(string(data), tc.expected) {
			t.Errorf("trigger '%s': expected ci.yml to contain '%s'", tc.trigger, tc.expected)
		}
	}
}

func TestGenerateCI_Idempotent(t *testing.T) {
	dir := t.TempDir()

	if err := generator.GenerateCI(dir, "push-to-main"); err != nil {
		t.Fatalf("first call failed: %v", err)
	}

	ciPath := filepath.Join(dir, ".github", "workflows", "ci.yml")
	original, _ := os.ReadFile(ciPath)

	if err := generator.GenerateCI(dir, "pr-to-develop"); err != nil {
		t.Fatalf("second call failed: %v", err)
	}

	after, _ := os.ReadFile(ciPath)
	if string(original) != string(after) {
		t.Errorf("ci.yml was overwritten on second call, expected idempotent behaviour")
	}
}
