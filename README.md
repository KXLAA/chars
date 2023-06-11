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

| Endpoint                                 | Method | Description                                |
| ---------------------------------------- | ------ | ------------------------------------------ |
| /api/v1/generate                         | `GET`  | Generate random characters of length 32    |
| /api/v1/generate?length=10               | `GET`  | Generate random character of length 10     |
| /api/v1/generate?lowercase=true          | `GET`  | Include lower case characters              |
| /api/v1/generate?uppercase=true          | `GET`  | Include upper case characters              |
| /api/v1/generate?numbers=true            | `GET`  | Include numbers                            |
| /api/v1/generate?symbols=true            | `GET`  | Include special characters                 |
| /api/v1/generate/bulk?count=10?length=32 | `GET`  | Generate 10 random characters of length 32 |

The Web app is available at <http://localhost:4444/> when you run `make run/live` in the root of the project.

## TODS's

- [ ] Add integration & E2E tests for CLI tool
- [ ] CI/CD for repo
- [ ] Distribute CLI tool as a binary for Linux, Windows, and Mac
- [ ] Deploy web app
- [ ] Migrate to HTTPS for web app
- [ ] Add copied flashed message to web app
- [ ] Proper Error Handling for the API endpoints
- [ ] Add proper footer to home page
- [ ] update readme
- [ ] Add a favicon
- [ ] remove unused helper functions & general clean ups
- [ ] Clean up css styles
- [ ] Give Credit for starter template
