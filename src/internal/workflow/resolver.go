package workflow

import "fmt"

// TODO:
// define supported workflows (centralized, feature, gitflow, forking)
// define WorkflowSpec struct (branches, CI triggers, contribution rules)
// implement Resolve(workflowName) -> WorkflowSpec
// validate unsupported workflow names

type WorkflowSpec struct {
	Name                   string
	RequiredBranches       []string
	ContributionGuidelines string
	CITrigger              string
}

func ResolveWorkflow(name string) (WorkflowSpec, error) {
	switch name {
	case "centralized":
		return WorkflowSpec{
			Name:                   "centralized",
			RequiredBranches:       []string{"main"},
			ContributionGuidelines: "All developers push directly to the main branch.",
			CITrigger:              "push-to-main",
		}, nil
	case "feature":
		return WorkflowSpec{
			Name:                   "feature",
			RequiredBranches:       []string{"main"},
			ContributionGuidelines: "Create feature branches and open pull requests to the main branch.",
			CITrigger:              "pr-to-main",
		}, nil
	case "gitflow":
		return WorkflowSpec{
			Name:                   "gitflow",
			RequiredBranches:       []string{"main", "develop"},
			ContributionGuidelines: "All pull requests much be opened against the develop branch. Do not merge directly to main.",
			CITrigger:              "pr-to-develop",
		}, nil
	case "forking":
		return WorkflowSpec{
			Name:                   "forking",
			RequiredBranches:       []string{"main"},
			ContributionGuidelines: "Fork the repository and open pull requests from your fork.",
			CITrigger:              "pr-to-main",
		}, nil
	default:
		return WorkflowSpec{}, fmt.Errorf("unsupported workflow: %s", name)
	}
}
