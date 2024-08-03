# Project Structure

This project follows the [golang-standards/project-layout](https://github.com/golang-standards/project-layout) guidelines. Below is the structure of the project and an explanation of each folder and file.

## Directory Structure

```plaintext
go-employee-web-server/
├── cmd/
│   └── employee/
│       └── main.go
├── internal/
│   ├── api/
│   │   ├── models/
│   │   │   └── employee.go
│   │   ├── client.go
│   │   └── httpClient.go
│   ├── configs/
│   │   ├── db.go
│   │   └── redis.go
│   ├── data/
│   │   ├── file.go
│   │   ├── mysql.go
│   │   └── storage.go
│   ├── factory/
│   │   └── factory.go
│   ├── handlers/
│   │   ├── add.go
│   │   ├── cleanup.go
│   │   ├── count.go
│   │   ├── edit.go
│   │   ├── employees.go
│   │   ├── login.go
│   │   ├── ping.go
│   │   ├── templates.go
│   │   └── top.go
│   │   └── view.go
│   ├── models/
│   │   └── employee.go
│   │   └── userCallCount.go
│   ├── utils/
│   │   └── utils.go
├── web/
│   ├── data/
│   │   └── employees.txt
│   ├── static/
│   │   └── style.css
│   └── templates/
│       ├── 404.html
│       ├── add.html
│       ├── edit.html
│       ├── employees.html
│       ├── login.html
│       └── view.html
├── docker-compose.yml
├── Dockerfile
├── go.mod
└── go.sum
```
## Directory Details

### `cmd/`
Contains the main application entry point.
- **`employee/`**
    - **`main.go`**: Initializes and runs the application.

### `internal/`
Contains application code that is not intended to be used outside of this project.
- **`api/`**
    - **`models/`**
        - **`employee.go`**: Defines the Employee model.
    - **`client.go`**: Implements the HTTP client.
    - **`httpClient.go`**: Provides additional HTTP client utilities.
- **`configs/`**
    - **`db.go`**: Database configuration.
    - **`redis.go`**: Redis configuration.
- **`data/`**
    - **`file.go`**: File-based storage operations.
    - **`mysql.go`**: MySQL-specific storage operations.
    - **`storage.go`**: Defines the storage interface.
- **`factory/`**
    - **`factory.go`**: Implements handler creation.
- **`handlers/`**
    - **`add.go`**: Handles adding a new employee.
    - **`cleanup.go`**: Handles cleanup operations.
    - **`count.go`**: Handles API call count statistics.
    - **`edit.go`**: Handles editing employee details.
    - **`employees.go`**: Handles listing employees.
    - **`login.go`**: Handles user login.
    - **`ping.go`**: Handles the ping API with rate limiting and locking.
    - **`templates.go`**: Provides utilities for templates.
    - **`top.go`**: Handles top API for retrieving top callers.
    - **`view.go`**: Handles viewing employee details.
- **`models/`**
    - **`employee.go`**: Defines the Employee model.
    - **`userCallCount.go`**: Manages user call count statistics.
- **`utils/`**
    - **`utils.go`**: Contains helper functions.

### `web/`
Contains web-related resources.
- **`data/`**
    - **`employees.txt`**: File for storing employee data.
- **`static/`**
    - **`style.css`**: CSS file for styling.
- **`templates/`**
    - **`404.html`**: Template for 404 error page.
    - **`add.html`**: Template for adding an employee.
    - **`edit.html`**: Template for editing an employee.
    - **`employees.html`**: Template for listing employees.
    - **`login.html`**: Template for user login.
    - **`view.html`**: Template for viewing an employee.

### Root Directory
- **`docker-compose.yml`**: Defines Docker services for the project.
- **`Dockerfile`**: Dockerfile for building the application image.
- **`go.mod`**: Go module file specifying dependencies.
- **`go.sum`**: Go sum file with checksums for dependencies.
- **`.env`**: Environment variables file for configuring sensitive settings and secrets.

## Getting Started

### Prerequisites

Before you begin, ensure you have the following installed on your system:

- [Docker](https://docs.docker.com/get-docker/) - For containerization and managing dependencies.
- [Docker Compose](https://docs.docker.com/compose/install/) - For defining and running multi-container Docker applications.
- [Go](https://golang.org/doc/install) - Go programming language (version specified in `go.mod`).

### Setting Up

1. **Clone the Repository**

   Clone this repository to your local machine using:

   ```bash
   git clone https://github.com/bendoppler/go-employee-web-server.git
   cd go-employee-web-server
   ```
### Create a `.env` File

Create a `.env` file in the root directory of the project. This file should contain the environment variables required for the application to run. Here's a sample `.env` file:

```env
DB_USER=root
DB_PASSWORD=yourpassword
DB_NAME=yourdatabase
DB_ROOT_PASSWORD=yourrootpassword
REDIS_PASSWORD=yourredispassword
```
Update the values with your own configuration:

- **`DB_USER`**: MySQL username.
- **`DB_PASSWORD`**: MySQL user password.
- **`DB_NAME`**: MySQL database name.
- **`DB_ROOT_PASSWORD`**: MySQL root password.
- **`REDIS_PASSWORD`**: Redis password.

## Build and Run the Docker Application

Use Docker Compose to build and run the containers. This command will:

1. Build the Docker images for your application based on the `Dockerfile`.
2. Start the containers as defined in the `docker-compose.yml` file.

Run the following command in the root of your project directory:

```bash
docker-compose up --build
```
- `--build`: Forces Docker Compose to rebuild the images even if they are up-to-date.
- `--detach` or `-d`: Runs the containers in the background and prints the container IDs.
- `--remove-orphans`: Removes containers for services not defined in the `docker-compose.yml` file.

To build and run the Docker application, execute the following command:

```bash
docker-compose up --build
```
This command will build the Docker images as defined in the `Dockerfile` and `docker-compose.yml` file, then start the containers in the background. If you make changes to the Dockerfile or dependencies, you can re-run this command to rebuild the images and restart the containers.

After running the application, you can access the application at `http://localhost:8080`.

To stop and remove the application, networks, and volumes created by `docker-compose up`, use the following command:

```bash
docker-compose down
```
This command will stop the running application and remove them, along with the networks and volumes that were created. It is useful for cleaning up after development or when you want to ensure a fresh start.

If you need to stop the application without removing them, you can use:

```bash
docker-compose stop
```
This will stop the running application but leave them in place, so they can be restarted later.

To restart the application, you can use:

```bash
docker-compose start
```

This command starts the stopped application without rebuilding the images. It is a quick way to resume your application if you need to pause and resume development or testing.

To remove all application, networks, and volumes defined in your docker-compose.yml, you can use:

```bash
docker-compose down --volumes
```
This command stops and removes the application, networks, and volumes associated with it. It is useful for cleaning up your environment or ensuring a fresh start.

For more detailed information on Docker Compose commands and options, refer to the [official Docker Compose documentation](https://docs.docker.com/compose/).
