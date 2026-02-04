package engine

import (
	"fmt"
	"seryn/src/internal/workflow"
)

// TODO:
// define ApplyWorkflow(repoPath, workflowName)
// repo exists or init it
// resolve rules via workflow resolver
// apply branch rules using gitops
// generate files and CI config
// log summary of actions performed

func ApplyWorkflow(workflowName string) error {
	spec, err := workflow.ResolveWorkflow(workflowName)
	if err != nil {
		return err
	}
	fmt.Println("Resolved workflow:")
	fmt.Println("Name:", spec.Name)
	fmt.Println("Required branches:", spec.RequiredBranches)
	fmt.Println("CI trigger:", spec.CITrigger)

	return nil
}
