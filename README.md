# Go Agency CLI

A command-line interface (CLI) tool written in Go to manage company branches across different regions.
This project demonstrates file handling (CSV), flag parsing, and CRUD operations without using an external database.

## Features
- **Create:** Add new agency details.
- **List:** View all agencies in a specific region.
- **Get:** Retrieve details of a specific agency by ID.
- **Edit:** Update agency information.
- **Status:** View statistics (count of agencies and employees).

## Usage

Run the program using the following flags:

```bash
# List all agencies in Tehran
go run main.go -command list -region tehran

# Create a new agency
go run main.go -command create -region isfahan
# (Follow the interactive prompts)

# Check status
go run main.go -command status -region shiraz