package engine

import (
	"seryn/src/internal/generator"
	"seryn/src/internal/gitops"
	"seryn/src/internal/logger"
	"seryn/src/internal/monitoring"
	"seryn/src/internal/workflow"
)

// define ApplyWorkflow(repoPath, workflowName)
// repo exists or init it
// resolve rules via workflow resolver
// apply branch rules using gitops
// generate files and CI config
// log summary of actions performed

func ApplyWorkflow(repoPath string, workflowName string) error {
	return applyWorkflow(repoPath, workflowName, "")
}

func ApplyWorkflowWithAlert(repoPath string, workflowName string, webhookURL string) error {
	return applyWorkflow(repoPath, workflowName, webhookURL)
}

func applyWorkflow(repoPath string, workflowName string, webhookURL string) error {
	// validate workflow first before touching filesystem
	spec, err := workflow.ResolveWorkflow(workflowName)
	if err != nil {
		monitoring.SendAlert(webhookURL, repoPath, workflowName, "failure")
		return err
	}

	logger.Info("Initializing repository at: " + repoPath)

	if err := gitops.EnsureRepo(repoPath); err != nil {
		monitoring.SendAlert(webhookURL, repoPath, workflowName, "failure")
		return err
	}

	logger.Info("Resolving workflow: " + workflowName)

	unexpected, err := gitops.EnsureBranches(repoPath, spec.RequiredBranches)
	if err != nil {
		monitoring.SendAlert(webhookURL, repoPath, workflowName, "failure")
		return err
	}

	for _, b := range unexpected {
		logger.Warning("Unexpected branch '" + b + "' exists; not part of workflow '" + spec.Name + "'")
	}

	if err := generator.GenerateFiles(
		repoPath,
		spec.Name,
		spec.ContributionGuidelines,
	); err != nil {
		monitoring.SendAlert(webhookURL, repoPath, workflowName, "failure")
		return err
	}

	if err := generator.GenerateCI(repoPath, spec.CITrigger); err != nil {
		monitoring.SendAlert(webhookURL, repoPath, workflowName, "failure")
		return err
	}

	filesGenerated := []string{
		"README.md",
		".gitignore",
		"CONTRIBUTING.md",
		".github/workflows/ci.yml",
	}

	logger.Summary(spec.Name, spec.RequiredBranches, filesGenerated, unexpected)

	monitoring.SendAlert(webhookURL, repoPath, workflowName, "success")

	return nil
}
