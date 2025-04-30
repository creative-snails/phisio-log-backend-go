# Project Setup Instructions

## Prerequisites

Ensure you have the following installed on your system:

- Docker
- Docker Compose

## Environment Variables

Create a `.env` file in the root of your project and add the following environment variables:

```env
SERVER_PORT=5000
SERVER_HOST=0.0.0.0
DB_PORT=5432
DB_HOST=db
DB_USER=postgres
DB_PASSWORD=phisio-log-dev
DB_NAME=phisiolog
DB_SSLMODE=disable
```

## Docker Setup

The project uses Docker for containerization. Follow these steps to set up and run the application:

1. **Build and Start Containers**
      `sh
   docker-compose up --build
   `

2. **Access the Application**
      - The application will be accessible at `http://localhost:5000`.

3. **Database Access**
      - The PostgreSQL database will be accessible at `localhost:5433`.

## Running the Application

To start developing, ensure the containers are running and then you can begin coding. The application will automatically reload with changes due to the `air` tool.

## Generate SQL queries

1. Install sqlc using go, run:

   ```bash
   go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
   ```

2. Using homebrew:

   ```bash
   brew install sqlc
   ```

3. Generate queries
   ```bash
   sqlc generate
   ```

**Note**: If you encounter a "command not found" or "not recognized" error after installing the `sqlc` package via `go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest` and running `sqlc generate` in the terminal, it indicates that Go's binary path is not properly configured in your system's PATH environment variable. Follow the instructions below to resolve this issue.

## Adding Go binary to PATH [[1]](https://stackoverflow.com/questions/28162577)

### Unix-based Systems (macOS/Linux)

1. Open your .zshrc file (for Zsh) or .bashrc (for Bash):

   ```bash
   nano ~/.zshrc
   ```

2. Add the following line at the end of the file:

   ```bash
   export PATH=$PATH:$(go env GOPATH)/bin
   ```

3. Save the file and reload your shell configuration:
   ```bash
   source ~/.zshrc
   ```

### Windows

1. Open System Properties (Win + Pause)
2. Click on "Advanced system settings" [[2]](https://legiondev.hashnode.dev/go-setup)
3. Click on "Environment Variables"
4. Under "User variables", find and select "Path"
5. Click "Edit"
6. Click "New"
7. Add the Go binary path:
   ```
   %USERPROFILE%\go\bin
   ```
8. Click "OK" to save all changes
