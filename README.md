# Learn Go with Tests

This repository contains code for learning Go with tests, following the [Learn Go with Tests](https://quii.gitbook.io/learn-go-with-tests) book.

## Structure

Each package represents a lesson and contains corresponding tests.

## How to Run

```sh
go run src/main.go
```

## How to Test

```sh
go test ./...
```

The `./...` syntax in the `go test` command is a wildcard pattern that tells the Go tool to recursively find and run tests in the current directory and all its subdirectories. This is useful for running all tests in a multi-package project with a single command.

### Additional Test Commands

- **Run All Tests in the Current Directory and Subdirectories**:

  ```sh
  go test ./...
  ```

- **Run Tests in a Specific Package**:

  ```sh
  go test ./path/to/package
  ```

- **Run Tests with Verbose Output**:

  ```sh
  go test -v ./...
  ```

- **Run Tests with Coverage**:

  ```sh
  go test -cover ./...
  ```

- **Run Tests and Generate a Coverage Report**:

  ```sh
  go test -coverprofile=coverage.out ./...
  go tool cover -html=coverage.out
  ```

- **Run a Specific Test Function**:
  ```sh
  go test -run TestFunctionName ./path/to/package
  ```
