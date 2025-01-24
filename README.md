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