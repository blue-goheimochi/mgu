# mgu

mgu - Manage git local users.

## Install

```
$ go get github.com/blue-goheimochi/mgu
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

## License

MIT

## Author

blue-goheimochi
