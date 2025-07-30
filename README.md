# Tago - A TODO List Tool in Go

Tago is a command-line TODO list management tool built with Go. It allows users to create, manage, and track tasks efficiently from the terminal. This README provides detailed instructions on how to download, install, and use Tago, including setting up bash autocompletion for an enhanced user experience.

## Table of Contents
- [Features](#features)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Commands and Flags](#commands-and-flags)
- [Bash Autocompletion](#bash-autocompletion)
- [Project Structure](#project-structure)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)

## Features
- Create, update, delete, and list TODO items from the command line.
- Mark tasks as done or clean the entire task list.
- Simple and lightweight implementation in Go.
- Supports task prioritization and filtering.
- Built-in bash autocompletion for seamless command-line interaction.

## Preview
An example of adding a command and running it:
![Example of adding a command](/src/demo.gif)

Another example on how completed tasks are removed from the table
![Example of completing a command](/src/demo2.gif)

## Prerequisites
To use Tago, ensure you have the following installed:
- **Go** (version 1.18 or higher): [Install Go](https://go.dev/doc/install)
- **Git**: For cloning the repository.
- **Bash**: For autocompletion setup (Linux/macOS or WSL on Windows).

## Installation
Follow these steps to download and install Tago:

1. **Install Tago Using `go install`**:
   Use Go's `install` command to download and install the `tago` binary directly from the repository:
   ```bash
   go install github.com/ronthesoul/tago@latest
   ```
   This downloads the source code, builds the binary, and places it in `$GOPATH/bin` (or `$HOME/go/bin` if `GOPATH` is not set).

2. **Ensure `$GOPATH/bin` is in Your `PATH`**:
   To run `tago` from anywhere, add `$GOPATH/bin` to your `PATH`. Add the following line to your `~/.bashrc` or `~/.zshrc`:
   ```bash
   export PATH=$PATH:$(go env GOPATH)/bin
   ```
   Reload your shell configuration:
   ```bash
   source ~/.bashrc
   ```

3. **Verify Installation**:
   Check that Tago is installed correctly:
   ```bash
   tago --version
   ```

4. **(Optional) Clone the Repository for Development**:
   If you want to modify the source code, clone the repository:
   ```bash
   git clone https://github.com/ronthesoul/tago.git
   cd tago
   go build -o tago .
   ```
   This creates a `tago` binary in the project directory, which you can move to `/usr/local/bin` for global access:
   ```bash
   sudo mv tago /usr/local/bin/
   ```

## Commands and Flags
Tago provides a set of commands for managing TODO lists. Below is a table of available commands, their descriptions, and supported flags. Flags are inferred based on typical TODO list tool functionality, as specific flag details were not provided.

| Command            | Description                                    | Flags/Options                              |
|--------------------|------------------------------------------------|--------------------------------------------|
| `tago add`         | Adds a new task to the list                   | `-t, --title <string>`: Task title <br> `-p, --priority <int>`: Priority (1-5) <br> `-d, --due <date>`: Due date (YYYY-MM-DD) |
| `tago clean`       | Cleans the entire task list                   | None                                       |
| `tago completion`  | Generates autocompletion script for a shell   | `<shell>`: Shell type (e.g., `bash`, `zsh`) |
| `tago done`        | Marks a task as done                          | `-id <int>`: Task ID to mark as done       |
| `tago help`        | Shows help for any command                    | `<command>`: Specific command to get help for |
| `tago list`        | Lists specified tasks                         | `-s, --status <string>`: Filter by status (e.g., `done`, `pending`) <br> `--all`: Show all tasks |
| `tago remove`      | Removes a task by its number                  | `-id <int>`: Task ID to remove             |
| `tago run`         | Runs a specific task                          | `-id <int>`: Task ID to run                |

**Example Usage**:
```bash
# Add a task with priority and due date
tago add -t "Write README" -p 2 -d 2025-08-01

# List all pending tasks
tago list -s pending

# Mark a task as done
tago done -id 1

# Remove a task
tago remove -id 1

# Clean the entire task list
tago clean

# Get help for the add command
tago help add
```

## Usefull alises
I've written a number of aliases that would help you memorize the commands and establish a smarter and faster workfow

```bash
alias tals='tago list'
alias talsa='tago list -a'
alias talsc='tago list -c'
alias taru='tago run'
alias tadd='tago add'
alias trem='tago remove'
alias tacl='tago clean'
alias tado='tago done'
```

## Bash Autocompletion
Tago includes a `completion` command to generate a bash autocompletion script, making it easier to use the CLI. Follow these steps to set it up:

1. **Generate the Autocompletion Script**:
   Run the following command to generate the bash completion script:
   ```bash
   tago completion bash > tago-completion.bash
   ```

2. **Move the Script to a System Location**:
   Move the script to a standard bash completion directory:
   ```bash
   sudo mv tago-completion.bash /etc/bash_completion.d/
   ```

3. **Source the Script in `.bashrc`**:
   Add the following line to your `~/.bashrc` to enable autocompletion:
   ```bash
   source /etc/bash_completion.d/tago-completion.bash
   ```
   Alternatively, if you prefer not to move the script, source it directly from your project directory:
   ```bash
   echo "source $(pwd)/tago-completion.bash" >> ~/.bashrc
   ```

4. **Apply Changes**:
   Reload your bash configuration:
   ```bash
   source ~/.bashrc
   ```

---

## Makefile Commands

| Command         | Description                                      |
|-----------------|--------------------------------------------------|
| `make all`      | Runs build and tests                             |
| `make build`    | Builds the application                           |
| `make run`      | Runs the application                             |
| `make watch`    | Live reloads the application                     |
| `make test`     | Runs the test suite                              |
| `make clean`    | Cleans up binary from the last build             |

**Example Usage**:
```bash
make all
make build
make run
make watch
make test
make clean
```

---

## Project Structure
Below is the directory structure of the Tago project, based on a typical Go project layout:

```
tago/
├── cmd/                 # Main application entry points
│   └── tago/            # Main package for the tago CLI
│       └── main.go      # Entry point for the tago command
├── internal/            # Internal packages (e.g., task management logic)
│   ├── task/            # Task-related logic (e.g., add, list, done, remove)
│   │   ├── task.go      # Task struct and methods
│   │   └── task_test.go # Unit tests for task package
│   └── storage/         # Storage-related logic (e.g., file or database storage)
│       ├── storage.go   # Storage implementation
│       └── storage_test.go # Unit tests for storage
├── go.mod               # Go module file
├── go.sum               # Go module checksums
├── tago-completion.bash # Bash autocompletion script (generated or manual)
└── README.md            # Project documentation (this file)
```

This structure follows Go conventions, with `cmd/tago/main.go` as the main entry point and `internal/` for private packages. The actual structure may vary slightly; check the repository for specifics.

## Usage
Once installed, you can use Tago to manage your TODO list. Here are some example commands:
```bash
# Add a new task
tago add -t "Write README" -p 2 -d 2025-08-01

# List all tasks
tago list --all

# Mark a task as done
tago done -id 1

# Remove a task
tago remove -id 1

# Run a specific task
tago run -id 1

# Clean the task list
tago clean
```

Tago likely stores tasks in a local file or database (check the `internal/storage` package for details). Ensure you have write permissions in the directory where Tago runs.

## Contributing
Contributions are welcome! To contribute:
1. Fork the repository.
2. Create a new branch (`git checkout -b feature/your-feature`).
3. Make changes and commit (`git commit -m "Add your feature"`).
4. Push to your fork (`git push origin feature/your-feature`).
5. Open a pull request on the [ronthesoul/tago](https://github.com/ronthesoul/tago) repository.

Please include tests for new features and follow Go coding conventions.

## License
The license for Tago is not specified in this README. Check the `LICENSE` file in the repository for details.