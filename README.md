# Chars

Developed and tested og Golang 1.20

This is a Web app, JSON API, and CLI tool for generating random characters.

## The CLI tool

The CLI tool has the following options:

| Flag | Type   | Description                      | Default |
| ---- | ------ | -------------------------------- | ------- |
| -c   | `int`  | Number of characters to generate | `1`     |
| -l   | `int`  | Length of each character         | `32`    |
| -lc  | `bool` | Include lower case characters    | `true`  |
| -uc  | `bool` | Include upper case characters    | `false` |
| -num | `bool` | Include numbers                  | `false` |
| -sc  | `bool` | Include special characters       | `true`  |

### Examples

Generate a single 32 character string with lower case characters and special characters:

```bash
go run ./cmd/cli
```

Generate a single 64 character string with lower case characters, upper case characters, numbers, and special characters:

```bash
go run ./cmd/cli -l 64 -uc -num -sc
```

Generate 10 64 character strings with lower case characters, upper case characters, numbers, and special characters:

```bash
go run ./cmd/cli -c 10 -l 64 -uc -num -sc
```

## The JSON API

The JSON API has the following endpoints:

| Endpoint              | Method | Description                   |
| --------------------- | ------ | ----------------------------- |
| /characters           | `GET`  | Generate random characters    |
| /characters?count=10  | `GET`  | Generate 10 random characters |
| /characters?length=64 | `GET`  | Generate 64 character strings |

The Web app is available at <http://localhost:4444/> when you run `make run/live` in the root of the project.
