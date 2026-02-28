# trainingCalendar
Generate training calendars

# Getting Started

## Requirements
- Go
- Docker (optional)
- npm
- react-scripts

To run without docker, and see all logs, in two separate terminals:
Run `make dev`
    - runs backend locally

Run `make npm`
    - runs frontend locally

To run dockerized:
- `make build && make run`

## Command line
There is also a command line interface if you have go installed locally
You can run `go run main.go -cmd -h` to view available instructions, but running `make cmd` will prompt you with some questions and output a schedule as an .ics file


# Deployment Info

## Infrastructure
    - Terraform Cloud used for infrastructure deployments
    - see /main.tf
    - Must log into terraform cloud and manually deploy any changes

## Frontend
    - AWS Amplify used for frontend deployments
    - tracks main branch, updates automatically on commit to master

## Backend
    - Run on lambda (see /lambda)
    - deployed via Update Lambda Github Action (triggered manually)
    - Must update /lambda/go.mod to reference newest tag version

## Releasing a New Version

To create a new release:

1. Ensure all changes are committed and pushed to main branch
2. Run the release command with the desired version:
   ```bash
   make release VERSION=v1.3.1
   ```
   This will:
   - Create and push a new git tag
   - Update lambda/go.mod to use the new version
   - Fetch the module from GitHub
   - Run go mod tidy

3. Test the lambda build:
   ```bash
   cd lambda && make build
   ```

4. Commit and push the updated lambda/go.mod:
   ```bash
   git add lambda/go.mod
   git commit -m "Update lambda to v1.3.1"
   git push
   ```

5. Deploy via the 'Update Lambda' GitHub Action (triggered manually)

### Development Mode

To switch lambda back to using local code for development:
```bash
make dev-mode
```

This adds a replace directive to use your local changes without creating a new release.
