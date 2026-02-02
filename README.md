# Seryn

**Seryn** is a DevOps-focused CLI tool that standardizes Git workflows across repositories based on team structure and organizational rules.  
It enables developers and platform teams to quickly initialize, enforce, and upgrade Git repositories to predefined workflows such as Gitflow or Trunk-Based Development.

---

## Student Details

**Student Name:** 
**Registration Number:** 
**Course:** CSE3253 – DevOps [PE6]  
**Semester:** VI (2025–2026)  
**Project Type:** Git & DevOps Automation  
**Difficulty Level:** Intermediate  

---

## Project Overview

### Problem Statement

In many teams, Git workflows are applied inconsistently across repositories.  
Setting up branches, repository rules, and CI pipelines manually for every new project is error-prone and time-consuming.

There is a need for a repeatable, automated, and configurable tool that can apply standardized Git workflows reliably across single or multiple repositories.

---

### Objectives

- Automate Git repository initialization and workflow setup
- Enforce consistent branching strategies across teams
- Generate standard repository files and CI pipelines automatically
- Support batch processing of multiple repositories
- Demonstrate practical DevOps concepts such as CI/CD, containerization, IaC, and monitoring

---

### Key Features

- CLI-based Git workflow enforcement
- Supports Gitflow and Trunk-Based workflows
- Automatic generation of README, `.gitignore`, and CI configuration
- Batch mode using Go concurrency
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
docker build -t seryn .
docker run -v $(pwd):/repo seryn apply --workflow trunk
```

### Project Structure

```bash
devops-project-git-workflow-customizer/
│
├── README.md
├── LICENSE
├── .gitignore
│
├── src/
│   ├── main.go
│   ├── go.mod
│   ├── config/
│   │   └── config.yaml
│   └── internal/
│       ├── gitops/
│       ├── generator/
│       ├── monitoring/
│       └── utils/
│
├── infrastructure/
│   ├── docker/
│   │   └── Dockerfile
│   │   └── docker-compose.yml
│   └── terraform/
│       └── main.tf
│
├── pipelines/
│   └── .github/workflows/
│       └── ci-cd.yml
│
├── tests/
│   ├── unit/
│   └── test-data/
│
├── monitoring/
│   └── alerts/
│       └── alert-config.json
│
├── docs/
│   ├── project-plan.md
│   ├── user-guide.md
│   └── design-document.md
│
└── deliverables/
```

### Configuration

Seryn can be configured using a YAML file:

```yaml
default_branch: main
workflow: gitflow
require_reviews: true
repositories:
  - /repo/project-a
  - /repo/project-b
```

### CI/CD Pipeline

The GitHub Actions pipeline performs:

- Code linting and formatting
- Go build and test
- Docker image build
- Security scan using Trivy

Pipeline definition:

```bash
pipelines/.github/workflows/ci-cd.yml
```

### Testing

- Unit tests verify file generation and workflow application
- Test data stored in tests/test-data/
- Tests run automatically in CI

```bash
go test ./...
```

### Monitoring & Alerts

After successful workflow application, Seryn sends a webhook notification containing:

- Repository name
- Applied workflow
- Status (success/failure)

Webhook configuration:

```bash
monitoring/alerts/alert-config.json
```

### Docker & Infrastructure

#### Docker

- Multi-stage build for minimal image size
- Alpine Linux base image

#### Terraform

- Provisions an S3 bucket to store workflow templates
- Simulates enterprise policy storage

### Development Workflow

#### Git Branching Strategy

```bash
main
└── develop
    ├── feature/*
    ├── hotfix/*
```

### Commit Convention

- feat: New feature
- fix: Bug fix
- docs: Documentation
- refactor: Code refactoring
- chore: Maintenance tasks

### Security

- No hard-coded secrets
- Environment-based configuration
- Container image scanning with Trivy
- Principle of least privilege (Terraform)

### Demo

A demo video demonstrating:

- Repository initialization
- Workflow application
- Batch processing
- CI/CD pipeline execution