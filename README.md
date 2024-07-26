# Project Structure

This project follows the [golang-standards/project-layout](https://github.com/golang-standards/project-layout) guidelines. Below is the structure of the project and an explanation of each folder and file.

## Directory Structure
go-employee-web-server/\
├── cmd/\
│ └── employee/\
│ └── main.go\
├── internal/\
│ ├── configs/\
│ │ └── db.go\
│ ├── data/\
│ │ └── db.go\
│ ├── handlers/\
│ │ ├── employees.go\
│ │ ├── view.go\
│ │ └── edit.go\
│ ├── models/\
│ │ └── employee.go\
│ └── utils/\
│ └── utils.go\
├── web/\
│ ├── static/\
│ │ └── style.css\
│ └── templates/\
│ ├── employees.html\
│ ├── view.html\
│ └── edit.html\
├── go.mod\
└── go.sum


## Explanation

### `cmd/`

- **cmd/employee/main.go**: The entry point for your application. This is where you set up and start your application.

### `internal/`

- **internal/configs/db.go**: Contains the database initialization and connection logic.
- **internal/data/db.go**: Contains functions to interact with the database (fetching, inserting, updating records).
- **internal/handlers/**: Contains HTTP handlers for different routes (`employees`, `view`, `edit`).
- **internal/models/employee.go**: Defines the `Employee` struct.
- **internal/utils/utils.go**: Contains utility functions used across the project.

### `web/`

- **web/static/**: Contains static files such as CSS, JavaScript, and images.
- **web/templates/**: Contains HTML templates used by the application.

### Project Files

- **go.mod**: The Go module file that defines the module path and dependencies.
- **go.sum**: The Go sum file that contains the checksums for the module's dependencies.
