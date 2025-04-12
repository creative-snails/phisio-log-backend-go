# Setup Instructions

> **Note**: If you encounter a "command not found" or "not recognized" error after installing the `air` package via `go install github.com/air-verse/air@latest` and running `air` in the terminal, it indicates that Go's binary path is not properly configured in your system's PATH environment variable. Follow the instructions below to resolve this issue.

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

## Verifying Installation

After setting up the PATH, verify that `air` is accessible by running:

```bash
air -v
```

If you haven't installed `air` yet, you can install it using:

```bash
go install github.com/cosmtrek/air@latest
```

## PostgreSQL Installation and Setup on macOS

1. If you haven't installed Homebrew yet, you can install it using:

   ```bash
   /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
   ```

2. To install PostgreSQL, run:

   ```bash
   brew install postgresql
   ```

3. Start the PostgreSQL service with::

   ```bash
   brew services start postgresql
   ```

4. Verify the installation by running:

   ```bash
   psql --version
   ```

## Resolving "FATAL: role 'postgres' does not exist" Error in pgAdmin4

If you try to click save and create the server for PostgreSQL in pgAdmin4 and you get an error saying that the role does not exist, follow these steps to resolve it. Use the username returned in the terminal as the username.

**Switch to the PostgreSQL User**: Use the following command to switch to the PostgreSQL user. If you installed PostgreSQL using Homebrew, it might use your macOS username instead of `postgres`. You can find your username by running `whoami` in the terminal.

```bash
sudo -u $(whoami) psql
```

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
