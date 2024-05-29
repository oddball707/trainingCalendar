# trainingCalendar
Generate training calendars

# Getting Started

To run without docker, and see all logs, in two separate terminals:
Run `make dev`
    - runs backend locally

Run `make npm`
    - runs frontend locally

To run dockerized:
- `make build && make run`

## Command line
There is also a command line interface if you have go installed locally
You can run `go run main.go -cmd -h` to view available instructions, but running `go run main.go -cmd -date 10/07/2024` will prompt you with some questions and output a schedule, but not a .ics file (yet)


# Deployment Info

## Infrastructure
    - Terraform Cloud used for infrastructure deployments
    - see /main.tf

## Frontend
    - AWS Amplify used for frontend deployments 
    - tracks main branch, updates automatically on commit to master

## Backend
    - Run on lambda (see /lambda)
    - deployed via Update Lambda Github Action (triggered manually)

