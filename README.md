# mgu

mgu - Manage git local users.

![Go Tests](https://github.com/blue-goheimochi/mgu/workflows/Go%20Tests/badge.svg)

## Install

```
$ go install github.com/blue-goheimochi/mgu/cmd/mgu@latest
```

For development version:

```
$ git clone https://github.com/blue-goheimochi/mgu.git
$ cd mgu
$ go build -o mgu ./cmd/mgu
```

## Usage

* Create setting file
    ```
    $ mgu init
    ~/.config/mgu/setting.json has been created
    ```

* Confirm Current User
    ```
    $ mgu
    blue-goheimochi <blue-goheimochi@example.com>
    ```

* Add User
    You can add users interactively.
    ````
    $ mgu add
    ? user.name blue-goheimochi
    ? user.email blue-goheimochi@example.com
    blue-goheimochi <blue-goheimochi@example.com> is added.
    ````

* Confirm user list
    ```
    $ mgu list
    blue-goheimochi <blue-goheimochi@example.com>
    pink-goheimochi <pink-goheimochi@example.com>
    ````

* Remove user
    ```
    $ mgu remove
    ? Please select a user
      blue-goheimochi <blue-goheimochi@example.com>
    > pink-goheimochi <pink-goheimochi@example.com>
    ? Do you want to remove? [y/N] y
    pink-goheimochi <pink-goheimochi@example.com> is removed.
    ````

* Set User
    ```
    $ mgu set
    ? Please select a user: (current: blue-goheimochi <blue-goheimochi@example.com>) 
    > blue-goheimochi <blue-goheimochi@example.com>
      pink-goheimochi <pink-goheimochi@example.com>
    blue-goheimochi <blue-goheimochi@example.com> has been set as a Git' local user.
    ````

## Development

### Requirements

- Go 1.18 or later (tested up to 1.20)
- Git (for running integration tests)

### Running Tests

To run all tests:

```
$ go test ./...
```

To run tests with race detection and coverage reporting:

```
$ go test -race -coverprofile=coverage.txt -covermode=atomic ./...
```

To view test coverage in your browser:

```
$ go tool cover -html=coverage.txt
```

### Project Structure

- `cmd/mgu`: Main executable and CLI definition
  - `cmd/mgu/commands/`: Implementation of the CLI commands
  - `cmd/mgu/commands/test_helpers.go`: Testing utilities for CLI commands
- `pkg/config`: Configuration management
  - `pkg/config/manager.go`: User configuration handling
  - `pkg/config/files.go`: File operations
  - `pkg/config/types.go`: Core type definitions
- `pkg/git`: Git repository operations
  - `pkg/git/interface.go`: Git repository interface
  - `pkg/git/repository.go`: Local Git repository implementation
  - `pkg/git/mock.go`: Mock implementation for testing

### Testing Architecture

The project uses dependency injection for testability:

1. **Interfaces**: Core functionality is defined through interfaces in `pkg/git/interface.go`
2. **Mocks**: Mock implementations are provided in `pkg/git/mock.go` and used in tests
3. **Factory Functions**: Components are created via factory functions that can be replaced in tests
4. **Test Helpers**: Helper utilities for testing are in `*_test_helpers.go` files
5. **Table-Driven Tests**: Tests use table-driven approach for comprehensive test cases

### GitHub Actions

This project uses GitHub Actions for continuous integration testing:

- Tests run on Go versions 1.18, 1.19, and 1.20
- Includes race detection
- Verifies build succeeds

## License

MIT

## Author

blue-goheimochi
