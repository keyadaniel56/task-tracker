# Task Tracker CLI

A simple command-line task manager built in Go.

Task Tracker allows you to track and manage your tasks directly from the terminal. You can add, update, delete, and manage tasks stored in a JSON file.

This project helps practice:
- Working with the filesystem
- Handling CLI arguments
- Structs and JSON encoding/decoding
- Project structure in Go
- Error handling
- Writing modular code

## Features

- Add tasks
- Update tasks
- Delete tasks
- Store tasks in `tasks.json`
- UUID-based task IDs
- Simple CLI interface

## Usage

Run the app using:
```bash
./run.sh add "Buy milk"
./run.sh update <id> "New description"
./run.sh delete <id>



bash

go build -o task
./task add "Buy milk"

Project Structure
text

.
├── app/
├── go.mod
├── go.sum
├── internal/
│   ├── app/
│   │   └── app.go
│   ├── handlers/
│   ├── modals/
│   │   └── task.go
│   └── utils/
│       ├── file.go
│       ├── task.go
│       └── type.go
├── main.go
├── README.md
├── run.sh
├── sample.txt
└── tasks.json

How It Works

Tasks are stored in tasks.json. Each task has:

    UUID ID

    Description

    Completion status

    CreatedAt timestamp

    UpdatedAt timestamp

The CLI reads user input using os.Args. Tasks are loaded, modified, then saved back to JSON.
Example Task Format (JSON)
json

[
  {
    "id": "c4f1c2e8-8a8e-4e9a-9e0e-6b9b8a7c1d3f",
    "description": "Buy milk",
    "completed": false,
    "created_at": "2026-02-16T10:00:00Z",
    "updated_at": "2026-02-16T10:00:00Z"
  }
]

Future Improvements

    Mark task as completed

    List all tasks

    Filter by status

    Add due dates

    Use Cobra for advanced CLI commands

    Add tests

text

