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
