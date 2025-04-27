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

### Running Tests

```
$ go test ./...
```

### Project Structure

- `cmd/mgu`: Main executable and CLI definition
- `pkg/config`: Configuration management
- `pkg/git`: Git repository operations

## License

MIT

## Author

blue-goheimochi
