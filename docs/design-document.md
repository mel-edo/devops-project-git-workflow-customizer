# Seryn Design Document

## Problem Statement

Git workflows are applied inconsistently across repositories in most teams.
Manual setup of branches, repository files, and CI pipelines is error-prone
and time-consuming. Seryn solves this by providing a repeatable, automated,
configurable CLI tool that enforces standardized Git collaboration workflows.

---

## Architecture Overview

Seryn follows a layered architecture with strict package separation:
```
main.go
  └── engine (ApplyWorkflow / ApplyWorkflowBatch)
        ├── gitops    (repo init, branch enforcement)
        ├── generator (file generation, CI generation)
        ├── workflow  (workflow resolution)
        ├── monitoring (webhook alerts)
        └── logger   (structured output)
```

Each package has a single responsibility and no circular dependencies.

---

## Package Responsibilities

### `main`
Parses CLI flags (`--workflow`, `--repo`, `--config`). Delegates entirely
to the engine layer. Contains no business logic.

### `engine`
Orchestrates the full workflow application sequence. Coordinates between
gitops, generator, workflow, monitoring, and logger packages. Exposes both
single-repo and batch execution modes.

### `internal/workflow`
Defines the `WorkflowSpec` struct and resolves workflow names to their
specifications. Pure logic — no filesystem or network access.

### `internal/gitops`
Handles all Git operations via `os/exec`. Ensures repositories are
initialized and branches exist without performing destructive operations.
All operations are idempotent.

### `internal/generator`
Generates standard repository files and CI configuration. Never overwrites
existing files. Idempotent by design.

### `internal/config`
Parses and validates YAML configuration files. Enables config-driven and
batch execution modes.

### `internal/monitoring`
Sends webhook POST notifications on workflow completion.
Webhook URL can be provided via config file or the --webhook CLI flag.
If both are absent, alerts are silently skipped.

### `internal/logger`
Provides structured, consistent terminal output across all packages.
Replaces raw `fmt.Println` calls with typed functions: `Success`,
`Warning`, `Info`, and `Summary`.

---

## Key Design Decisions

### Non-destructive branch enforcement
Seryn never deletes or renames branches. It only creates missing required
branches and warns about unexpected ones. This is intentional — destructive
operations on shared branches would be unsafe in a team environment.

### Idempotency
Every operation in Seryn is safe to run multiple times. Files are only
created if missing. Branches are only created if absent. Git init is only
called if `.git/` does not exist. This makes Seryn safe to run repeatedly
as a policy enforcement tool.

### Forking workflow limitations
The Forking workflow cannot be fully enforced at the local Git level because
it depends on platform-level permissions (GitHub fork restrictions, protected
branch rules). Seryn handles this by generating appropriate CONTRIBUTING.md
documentation and CI configuration, while clearly documenting that full
enforcement requires platform configuration.

### Concurrency model
Batch mode uses one goroutine per repository with a `sync.WaitGroup` for
coordination and a `sync.Mutex` for safe concurrent writes to the results
slice. Output per repository is buffered through the logger to prevent
interleaved terminal output.

### Containerization approach
Seryn uses a multi-stage Docker build. The builder stage compiles the binary
using the full Go toolchain. The runtime stage uses Alpine Linux with only
Git installed, producing a minimal and secure image. The non-root user
(`seryn`) follows the principle of least privilege.

### CI pipeline design
The generated CI pipeline matches the workflow's trigger semantics exactly.
Centralized workflow triggers on push to main. Feature and Forking workflows
trigger on pull requests to main. Gitflow triggers on pull requests to
develop. This ensures the generated CI is immediately useful rather than
generic boilerplate.

---

## Infrastructure as Code

Terraform provisions an S3 bucket simulating enterprise policy template
storage. The bucket has versioning, AES256 server-side encryption, and
public access fully blocked, demonstrating security-conscious IaC practices.

---

## Security Decisions

- No hard-coded secrets anywhere in the codebase
- Webhook URL is config-driven and optional
- Docker image runs as non-root user
- Container image scanned with Trivy in CI (fails on HIGH/CRITICAL)
- Terraform bucket blocks all public access
- Git identity injected via environment variables at runtime