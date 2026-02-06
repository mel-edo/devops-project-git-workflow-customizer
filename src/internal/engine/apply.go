package engine

import (
	"fmt"
	"seryn/src/internal/generator"
	"seryn/src/internal/gitops"
	"seryn/src/internal/workflow"
)

// define ApplyWorkflow(repoPath, workflowName)
// repo exists or init it
// resolve rules via workflow resolver
// apply branch rules using gitops
// generate files and CI config
// log summary of actions performed

func ApplyWorkflow(repoPath string, workflowName string) error {
	err := gitops.EnsureRepo(repoPath)
	if err != nil {
		return err
	}

	spec, err := workflow.ResolveWorkflow(workflowName)
	if err != nil {
		return err
	}

	unexpected, err := gitops.EnsureBranches(repoPath, spec.RequiredBranches)
	if err != nil {
		return err
	}

	for _, b := range unexpected {
		fmt.Printf(
			"Warning: unexpected branch '%s' exists; workflow '%s' expects %v\n",
			b,
			spec.Name,
			spec.RequiredBranches,
		)
	}

	if err := generator.GenerateFiles(
		repoPath,
		spec.Name,
		spec.ContributionGuidelines,
	); err != nil {
		return err
	}

	filesGenerated := []string{
		"README.md",
		".gitignore",
		"CONTRIBUTING.md",
	}

	fmt.Println("✔ Repository ready")
	fmt.Printf("✔ Workflow applied: %s\n", spec.Name)
	fmt.Printf("✔ Branches ensured: %v\n", spec.RequiredBranches)

	if len(unexpected) > 0 {
		fmt.Printf("Extra branches detected: %v\n", unexpected)
	}

	fmt.Printf("✔ Files handled: %v\n", filesGenerated)

	return nil
}
