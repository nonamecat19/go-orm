# GitHub CI Workflows

This directory contains GitHub Actions workflows for continuous integration of the go-orm project.

## Workflows

### 1. Code Formatting (`code-formatting.yml`)
- Checks that all Go files are properly formatted using `gofmt`
- Runs on push to main/master and on pull requests
- Fails if any files are not properly formatted

### 2. Linting (`linting.yml`)
- Runs [golangci-lint](https://github.com/golangci/golangci-lint) on all modules in the workspace
- Checks for code quality issues, bugs, and style violations
- Uses a 5-minute timeout to allow for thorough linting

### 3. Static Analysis (`static-analysis.yml`)
- Runs [staticcheck](https://staticcheck.io/) on all modules in the workspace
- Performs advanced static analysis to find bugs and performance issues

### 4. Build Check (`build.yml`)
- Builds all modules in the workspace
- Ensures that the code compiles successfully
- Fails if any module fails to build

## Workspace Support

All workflows are designed to work with Go workspaces. They parse the `go.work` file to identify all modules and run the checks on each module separately.

## Running Locally

You can run these checks locally before pushing your changes:

```bash
# Format code
gofmt -w .

# Lint code
golangci-lint run ./...

# Static analysis
staticcheck ./...

# Build
go build ./...
```