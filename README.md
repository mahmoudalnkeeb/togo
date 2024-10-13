# ToGo

A simple CLI todo application built with Golang.

## Table of Contents

- [Features](#features)
- [Known Issues](#known-issues)
- [Usage](#usage)
- [Setup](#setup)
- [Cli](#cli)
- [License](#license)

## Features

- Create new todo items
- List all todos
- Retrieve a todo by ID
- Update the title of a todo
- Mark a todo as complete
- > **TODO:** Filters for list command like --filters "completed=0&page=2&limit=10&before=2006-01-02 15:04:05"
- > **[not working for now]** Automatically mark todos as complete based on a specified time

## Usage

### Prerequisites

Ensure you have the following installed on your machine:

- [Go](https://go.dev/dl/)
- [Git](https://git-scm.com/downloads)

### Setup

1. Clone the repository and navigate to the directory:

    ```bash
    git clone https://github.com/mahmoudalnkeeb/togo.git && cd togo
    ```

2. Install the necessary modules:

    ```bash
    go mod tidy
    ```

3. Add `DB_FILE` to .env file - start in the root directory

    ```dotenv
       DB_FILE="/path/to/togo.db"
    ```

    or using export command

    ```bash
      export DB_FILE="/path/to/togo.db"&&go run .
    ```

4. Run the application:

    ```bash
    go run .
    ```

## Cli

### Available Commands

- **add**
  Add a new todo item.

  Usage:

  ```bash
  togo add --title "Your Todo Title" [--complete_by "YYYY-MM-DD HH:MM:SS"]
  ```

  - `--title` (required): The title of the todo item.
  - `--complete_by` (optional): The time when the todo should be marked as complete.

- **complete**
  Mark a todo item as complete.

  Usage:

  ```bash
  togo complete <ID>
  ```

  - `<ID>`: The ID of the todo item to mark as complete.

- **completion**
  Generate the autocompletion script for the specified shell.

- **delete**
  Delete a todo item.

  Usage:

  ```bash
  togo delete <ID>
  ```

  - `<ID>`: The ID of the todo item to delete.

- **help**
  Help about any command.

- **list**
  List all todo items.

  Usage:

  ```bash
  togo list
  ```

- **update**
  Update a todo item title.

  Usage:

  ```bash
  togo update <ID> --title "New Title"
  ```

  - `<ID>`: The ID of the todo item to update.
  - `--title` (required): The new title for the todo item.

### Flags

- `-h, --help`: Help for togo.

Use `"togo [command] --help"` for more information about a command.

## Known Issues

- The auto-complete feature may not work as intended due to time-related problems. (**Working with dates can be tricky!**)

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
