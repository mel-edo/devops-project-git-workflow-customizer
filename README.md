# Seryn

**Seryn** is a DevOps-focused CLI tool that standardizes Git workflows across repositories based on team structure and organizational rules.  
It enables developers and platform teams to quickly initialize, enforce, and upgrade Git repositories to canonical Git collaboration workflows such as Centralized, Feature Branch, Gitflow, and Forking workflows.

---

## Student Details

**Student Name:** Saurabh P

**Registration Number:** 23FE10CSE00264

**Course:** CSE3253 – DevOps [PE6]  

**Semester:** VI (2025–2026)  

**Project Type:** Git & DevOps Automation  

**Difficulty Level:** Intermediate  

---

## Project Overview

### Problem Statement

In many teams, Git workflows are applied inconsistently across repositories.  
Setting up branches, repository rules, and CI pipelines manually for every new project is error-prone and time-consuming.

There is a need for a repeatable, automated, and configurable tool that can apply standardized
Git collaboration workflows reliably across single or multiple repositories, while respecting
real-world constraints of hosting platforms such as GitHub.

---

### Objectives

- Automate Git repository initialization and workflow setup
- Enforce consistent branching strategies across teams
- Generate standard repository files and CI pipelines automatically
- Support batch processing of multiple repositories
- Demonstrate practical DevOps concepts such as CI/CD, containerization, IaC, and monitoring

---

### Key Features

- CLI-based Git workflow standardization
- Supports four canonical Git workflows:
  - Centralized workflow
  - Feature Branch workflow
  - Gitflow workflow
  - Forking workflow (documentation and policy support)
- Automatic generation of README, `.gitignore`, CONTRIBUTING.md, and CI configuration
- Workflow-specific branch creation and enforcement rules
- Batch mode using Go concurrency for multi-repository setup
- Dockerized execution for portability
- Cloud-backed template storage (S3 simulation)
- Webhook notifications on completion

---

## Technology Stack

### Core Technologies

- **Programming Language:** Go (Golang)
- **CLI Framework:** Standard Go CLI (`flag`, `os`, `exec`)
- **Configuration:** YAML

### DevOps Tools

- **Version Control:** Git
- **CI/CD:** GitHub Actions
- **Containerization:** Docker
- **Infrastructure as Code:** Terraform
- **Security Scanning:** Trivy
- **Monitoring / Alerts:** Webhooks (Slack / Discord)

---

## Getting Started

### Prerequisites

- Git 2.30+
- Go 1.21+
- Docker 20.10+
- GitHub account (for CI/CD)

---

### Installation

#### Clone the repository

```bash
git clone https://github.com/mel-edo/devops-project-git-workflow-customizer.git
cd devops-project-git-workflow-customizer
```

### Usage

#### Running Seryn Locally

```bash
go run src/main.go --workflow gitflow --repo /path/to/project
```

#### Running with Docker

```bash
make build
make run ARGS="--workflow gitflow"
```

### Project Structure

```bash
devops-project-git-workflow-customizer/
│
├── README.md
├── LICENSE
├── .gitignore
├── Makefile
│
├── src/
│   ├── main.go
│   ├── go.mod
│   ├── config/
│   │   └── config.yaml
│   └── internal/
│       ├── engine/
│       │   ├── apply.go
│       │   └── batch.go
│       ├── gitops/
│       │   ├── repo.go
│       │   └── branches.go
│       ├── generator/
│       │   ├── files.go
│       │   ├── ci.go
│       │   └── generator_test.go
│       ├── workflow/
│       │   ├── resolver.go
│       │   └── workflow_test.go
│       ├── monitoring/
│       │   └── alert.go
│       ├── config/
│       │   └── loader.go
│       └── logger/
│           └── logger.go
│
├── infrastructure/
│   ├── docker/
│   │   └── Dockerfile
│   └── terraform/
│       └── main.tf
│
├── .github/
│   └── workflows/
│       └── ci.yml
│
└── docs/
    ├── user-guide.md
    └── design-document.md
```

### Configuration

Seryn can be configured using a YAML file:

```yaml
default_branch: main
workflow: gitflow   # centralized | feature | gitflow | forking
webhook_url: "https://discord.com/api/webhooks/your/webhook"

repositories:
  - /repo/project-a
  - /repo/project-b
```

Seryn enforces workflows differently based on their nature.
Centralized, Feature Branch, and Gitflow workflows are fully enforced at the repository level.
The Forking workflow is supported through repository policies and documentation scaffolding,
as it depends on platform-level permissions outside the scope of a local Git CLI.

### Supported Workflows

#### 1. Centralized Workflow
- Single shared repository
- Developers push directly to `main`
- CI runs on every push to `main`

#### 2. Feature Branch Workflow
- `main` is protected
- Feature branches (`feature/*`, `fix/*`)
- Pull requests required for merging
- CI triggered on pull requests

#### 3. Gitflow Workflow
- Long-lived `main` and `develop` branches
- Feature branches merged into `develop`
- CI triggered on pull requests to `develop`

#### 4. Forking Workflow
- Intended for open-source or external contributors
- Direct pushes to `main` are restricted
- CONTRIBUTING.md documents fork-based contribution flow
- Enforcement is partial due to reliance on hosting-platform permissions

### CI/CD Pipeline

The GitHub Actions pipeline performs:

- Code linting and formatting
- Go build and test
- Docker image build
- Security scan using Trivy

Pipeline definition:

```bash
.github/workflows/ci.yml
```

### Testing

- Unit tests verify file generation and workflow application
- Tests run automatically in CI

```bash
cd src && go test ./...
```

### Monitoring & Alerts

After successful workflow application, Seryn sends a webhook notification containing:

- Repository name
- Applied workflow
- Status (success/failure)

Webhook configuration is done via the config YAML file:

```yaml
webhook_url: "https://discord.com/api/webhooks/your/webhook"
```

Compatible with Slack and Discord webhooks. Leave empty to disable.

### Docker & Infrastructure

#### Docker

- Multi-stage build for minimal image size
- Alpine Linux base image

#### Terraform

- Provisions an S3 bucket to store workflow templates
- Simulates enterprise policy storage

## Contributing

Suggestions, fixes and improvments are welcome. Feel free to open an issue or a PR.

---

## License

This project is licensed under the GNU General Public License v3.0 - see the [LICENSE](LICENSE) file for details.