# Seryn User Guide

## Overview

Seryn is a CLI tool that standardizes Git workflows across repositories.
It initializes repositories, enforces branching strategies, generates
standard files, and produces CI configuration automatically.

---

## Installation

### Run locally
```bash
git clone https://github.com/mel-edo/devops-project-git-workflow-customizer.git
cd devops-project-git-workflow-customizer/src
go build -o seryn .
./seryn --workflow gitflow --repo /path/to/repo
```

### Run with Docker
```bash
make run ARGS="--workflow gitflow --repo /repo"
```

---

## Usage

### CLI Mode

Apply a workflow to a single repository:
```bash
seryn --workflow  --repo 
```

Supported workflows:

| Name         | Branches Created  | CI Trigger         |
|--------------|-------------------|--------------------|
| centralized  | main              | Push to main       |
| feature      | main              | PR to main         |
| gitflow      | main, develop     | PR to develop      |
| forking      | main              | PR to main         |

### Config Mode

Apply a workflow to multiple repositories using a YAML config file:
```bash
seryn --config src/config/config.yaml
```

Example config file:
```yaml
workflow: gitflow
default_branch: main
require_reviews: true
webhook_url: "https://discord.com/api/webhooks/your/webhook"

repositories:
  - /repo/project-a
  - /repo/project-b
  - /repo/project-c
```

When more than one repository is listed, Seryn processes them concurrently
using goroutines and prints a batch summary on completion.

---

## What Seryn generates

For every repository, Seryn creates the following files if they do not exist:

- `README.md` — identifies the applied workflow
- `.gitignore` — sensible defaults for Go projects
- `CONTRIBUTING.md` — workflow-specific contribution guidelines
- `.github/workflows/ci.yml` — GitHub Actions CI pipeline matching the workflow trigger

Seryn never overwrites existing files. All operations are idempotent and
safe to run multiple times.

---

## Webhook Monitoring

Seryn can send a POST notification to a webhook URL after each workflow
application. The payload looks like this:
```json
{
  "repository": "/repo/project-a",
  "workflow": "gitflow",
  "status": "success",
  "timestamp": "2025-01-01T12:00:00Z"
}
```

Set `webhook_url` in your config file to enable this. Leave it empty to
disable it. Compatible with Slack and Discord webhooks.

---

## Docker

Build and run Seryn inside Docker using the Makefile:
```bash
# Build image
make build

# Apply workflow via Docker
make run ARGS="--workflow gitflow"

# Clean up image
make clean
```

Git author identity is automatically injected from your global Git config.

---

## Running Tests
```bash
cd src
go test ./...

# Verbose output
go test -v ./...
```