# Project Structure

This project follows the [golang-standards/project-layout](https://github.com/golang-standards/project-layout) guidelines. Below is the structure of the project and an explanation of each folder and file.

## Directory Structure
go-employee-web-server/\
├── cmd/\
│ └── employee/\
│ └── main.go\
├── internal/\
│ ├── api/\
│ │ └── client.go\
│ ├── configs/\
│ │ └── db.go\
│ ├── data/\
│ │ ├── file.go\
│ │ ├── mysql.go\
│ │ └── storage.go\
│ ├── factory/\
│ │ └── handler_factory.go\
│ ├── handlers/\
│ │ ├── add.go\
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
│ ├── edit.html\
│ └── add.html\
├── go.mod\
├── docker-compose.yml\
├── .env\
└── go.sum

### cmd/
This directory contains the main application entry point.

- **employee/**: The main application directory.
  - **main.go**: The main file that initializes and runs the application.

### internal/
Contains application code that is not intended to be used outside of this project.

- **api/**: Contains the API client.
  - **client.go**: Implements the HTTP client and interface for fetching employees from the API.

- **configs/**: Contains configuration-related code.
  - **db.go**: Handles database configuration.

- **data/**: Contains storage-related code.
  - **file_storage.go**: Implements file-based storage operations.
  - **mysql_storage.go**: Implements MySQL-based storage operations.
  - **storage.go**: Defines the storage interface.

- **factory/**: Contains factory functions for creating handlers.
  - **handler_factory.go**: Implements the factory for creating handlers.

- **handlers/**: Contains the HTTP handlers for different routes.
  - **add.go**: Handles adding a new employee.
  - **employees.go**: Handles listing employees with pagination and search.
  - **view.go**: Handles viewing an employee's details.
  - **edit.go**: Handles editing an employee's details.

- **models/**: Contains the data models.
  - **employee.go**: Defines the Employee struct used within the internal application logic.

- **utils/**: Contains utility functions.
  - **utils.go**: Implements helper functions used throughout the project.

### web/
Contains web-related resources.

- **static/**: Contains static files (e.g., CSS, JavaScript).
  - **style.css**: Stylesheet for the web application.
- **templates/**: Contains HTML templates.
  - **employees.html**: Template for listing employees with pagination and search.
  - **view.html**: Template for viewing an employee's details.
  - **edit.html**: Template for editing an employee's details.
  - **add.html**: Template for adding a new employee.

### go.mod
The Go module file, which defines the module path and its dependencies.

### go.sum
The file that contains the expected cryptographic checksums of the content of specific module versions.

## Getting Started

To run the project, follow these steps:

1. Clone the repository.
2. Navigate to the project directory.
3. Run `go mod tidy` to install the dependencies.
4. Run the application using `go run cmd/employee/main.go`.

Ensure that you have Go installed on your machine. For more information on installing Go, visit [the official Go documentation](https://golang.org/doc/install).

## Using Docker for MySQL

This project uses Docker to manage the MySQL database. Ensure Docker is installed and running on your machine.

1. Create a `.env` file in the root directory of the project with the following content:
    ```
    MYSQL_ROOT_PASSWORD=your_root_password
    MYSQL_PASSWORD=your_password
    MYSQL_DATABASE=employee
    ```

2. Run the Docker Compose command to start the MySQL container:
    ```
    docker-compose up -d
    ```

3. Check the status of the Docker container:
    ```
    docker ps
    ```

This will start a MySQL database instance in a Docker container, which your Go application can connect to.

### Important Note

- Make sure to replace `your_root_password` and `your_password` with your desired passwords.
- The `.env` file is included in the `.gitignore` to prevent sensitive information from being committed to version control.