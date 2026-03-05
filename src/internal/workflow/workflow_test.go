package workflow_test

import (
	"seryn/src/internal/workflow"
	"testing"
)

func TestResolveWorkflow(t *testing.T) {
	spec, err := workflow.ResolveWorkflow("centralized")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if spec.Name != "centralized" {
		t.Errorf("expected name 'centralized', got '%s'", spec.Name)
	}

	if len(spec.RequiredBranches) != 1 || spec.RequiredBranches[0] != "main" {
		t.Errorf("expected branches [main], got %v", spec.RequiredBranches)
	}

	if spec.CITrigger != "push-to-main" {
		t.Errorf("expected CITrigger 'push-to-main', got '%s'", spec.CITrigger)
	}
}

func TestResolveWorkflow_Feature(t *testing.T) {
	spec, err := workflow.ResolveWorkflow("feature")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if spec.Name != "feature" {
		t.Errorf("expected name 'feature', got '%s'", spec.Name)
	}

	if spec.CITrigger != "pr-to-main" {
		t.Errorf("expected CITrigger 'pr-to-main', got '%s'", spec.CITrigger)
	}
}

func TestResolveWorkflow_Gitflow(t *testing.T) {
	spec, err := workflow.ResolveWorkflow("gitflow")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if spec.Name != "gitflow" {
		t.Errorf("expected name 'gitflow', got '%s'", spec.Name)
	}

	if len(spec.RequiredBranches) != 2 {
		t.Fatalf("expected name 'gitflow', got '%s'", spec.Name)
	}

	if spec.RequiredBranches[0] != "main" || spec.RequiredBranches[1] != "develop" {
		t.Errorf("expected branches [main develop], got %v", spec.RequiredBranches)
	}

	if spec.CITrigger != "pr-to-develop" {
		t.Errorf("expected CITrigger 'pr-to-develop', got '%s'", spec.CITrigger)
	}
}

func TestResolveWorkflow_Forking(t *testing.T) {
	spec, err := workflow.ResolveWorkflow("forking")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if spec.Name != "forking" {
		t.Errorf("expected name 'forking', got '%s'", spec.Name)
	}

	if spec.CITrigger != "pr-to-main" {
		t.Errorf("expected CITrigger 'pr-to-main', got '%s'", spec.CITrigger)
	}
}

func TestResolveWorkflow_Unsupported(t *testing.T) {
	_, err := workflow.ResolveWorkflow("trunk")
	if err == nil {
		t.Fatal("expected error for unsupported workflow, got nil")
	}

	expected := "unsupported workflow: trunk"
	if err.Error() != expected {
		t.Errorf("expected error '%s', got '%s'", expected, err.Error())
	}
}

func TestResolveWorkflow_ContributionGuidelines(t *testing.T) {
	workflows := []string{"centralized", "feature", "gitflow", "forking"}
	for _, w := range workflows {
		spec, err := workflow.ResolveWorkflow(w)
		if err != nil {
			t.Fatalf("unexpected error for workflow '%s': %v", w, err)
		}

		if spec.ContributionGuidelines == "" {
			t.Errorf("expected non-empty ContributionGuidelines for workflow '%s'", w)
		}
	}
}
