# Getting started with Go 1.22

This is a small approach to getting started with Go 1.22. This walkthrough is for MacOS.

The aim is to create a basic "Hello, World" API with the new mux router option from the Go 1.22 release.

## Install Go

You can install Go from [the official website](https://go.dev/dl/). I used [Homebrew](https://brew.sh/) in my case.

```bash
# Install with Homebrew
$ brew install go
```

Once installed, you can confirm the version with `go version`.

## Initialize a Go project

You can use the `go mod init` function in order to initialize a new `go.mod` file.

```bash
# Init the project
$ go mod init github.com/okeeffed/go-1-22-mux
```

`go mod` is a helper for manager modules. `go mod help` can display more information about this.

```bash
$ go mod help
Go mod provides access to operations on modules.

Note that support for modules is built into all the go commands,
not just 'go mod'. For example, day-to-day adding, removing, upgrading,
and downgrading of dependencies should be done using 'go get'.
See 'go help modules' for an overview of module functionality.

Usage:

        go mod <command> [arguments]

The commands are:

        download    download modules to local cache
        edit        edit go.mod from tools or scripts
        graph       print module requirement graph
        init        initialize new module in current directory
        tidy        add missing and remove unused modules
        vendor      make vendored copy of dependencies
        verify      verify dependencies have expected content
        why         explain why packages or modules are needed

Use "go help mod <command>" for more information about a command.
```

## Configuring the Makefile

A Makefile is a special file used in software development projects, particularly in Unix-like operating systems, to automate the process of compiling and building software. It contains a set of instructions (written in a format called "Makefile syntax") that define how to build the project and its various components.

Here's a simple explanation of a Makefile for someone who's not familiar with it:

```makefile
# Makefile for a Go project
# Assumes entry point at project/cmd/server/server.go

# Binary output name
BINARY_NAME=hello_world

# Default make command
all: build

# Build the binary
build:
	@echo "Building..."
	go build -o $(BINARY_NAME) ./cmd/server

# Run the server
run: build
	@echo "Running..."
	./$(BINARY_NAME)

# Test your application
test:
	@echo "Testing..."
	go test ./...

# Clean up binaries
clean:
	@echo "Cleaning..."
	go clean
	rm -f $(BINARY_NAME)

# Help command to display available commands
help:
	@echo "Makefile commands:"
	@echo "all    - Build the application"
	@echo "build  - Build the binary"
	@echo "run    - Build and run the application"
	@echo "test   - Run tests"
	@echo "clean  - Remove binaries"
	@echo "help   - Display this help"

# Mark commands that don't correspond to files as .PHONY
.PHONY: all build run test clean help

```

> Warning: Make sure you use the correct indentation.

The following explains the different parts of this file.

### Variables

- `BINARY_NAME`: Represents the name of the binary executable.

### Default Target (`all`)

- When you run `make` without specifying a target, it executes the `build` target.

### Build Target

- Compiles the Go source files and generates the binary executable specified by `BINARY_NAME`.

### Clean Target

- Removes any generated binary executable and cleans up any temporary files created during the build process.

### Test Target

- Runs all the tests in the project.

### Dependencies Target

- Installs any dependencies the project might have.

### Run Target

- Builds and runs the executable.

### Help Target

- Provides a helpful message listing available targets and their descriptions.

To use this Makefile, place it in your Go project directory and run `make`. You can also use `make clean` to remove the generated binary, `make test` to run tests, `make deps` to install dependencies, and `make run` to build and run the executable.

This **Makefile** is a little more complex than it needs to be (for this simple project), but understanding it will also help you when you see other Makefiles. It is not just used for Golang applications.

Add the example code above into a `Makefile`.

## Setting up our remaining files and folders

The [Go project layout](https://github.com/golang-standards/project-layout) project on GitHub is a great place to help understand a layout for Go projects. That being, a lot of it can be overkill. We will follow these standards for the sake of today's example, but please note that this is not necessary for such a small project.

It may help in future if you are build up a large application.

For now, we will use the following folder structure:

```txt
project/
├── cmd/server/          # Application entry points
├── api/v1/              # Route handlers
└── internal/            # Private application and library code
    └── model/           # Domain models
        └── user.go      # User struct and methods
```

> Please note: this is very much overkill for the size of the project.

For this, add a few folders:

```s
# Make the folders and add the files
$ mkdir -p cmd/server internal/model api/v1
$ touch cmd/server/server.go internal/model/user.go api/v1/goodbye.go api/v1/hello.go
```

We are ready to add some code.

## Adding our User model

We will be writing the code in this order:

1. Adding our **User** model.
2. Adding our **GET** routes for `/v1/hello` and `/v1/goodbye`.
3. Adding our `main` function (using the `http.NewServeMux` export from Go 1.22).

Even though I say **User** model, really it will be a barebones struct to demonstrate using it across files. The actually implementation of it will be a bit rogue and not-for-a-user.

That being said, our `internal/model/user.go` file will look like this:

```go
package model

// User is a simple model for a user.
type User struct {
	// Name is the name of the user.
	Name string
}
```

He, we are create a package `model` that we can use later in our imports as the following:

```go
import (
	model "github.com/okeeffed/go-1-22-mux/internal/model"
)
```

Naming the struct with a capital letter (`User`) makes the struct public and will be importable from other files.

We can then use this in our application.

## Configuring our GET route handlers

Inside of `api/v1/hello.go`, add the following:

```go
package api

import (
	"net/http"

	model "github.com/okeeffed/go-1-22-mux/internal/model"
)

// HelloHandler is a simple HTTP handler that writes a response.
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	// Create a new user.
	user := model.User{
		Name: "World",
	}

	// Write a response to the client.
	w.Write([]byte("Hello, " + user.Name + "!"))
}
```

Here we are using the `http` import helping us type our arguments for the response writer and the request, while the `model` import is used so we can create our new user with the name "World" (just to stick to hello world convention for now).

> In practice, you may see variables denoted with shorthand e.g. `user` would be `u`.

Similar to before, we've used our keyword `package` to define our `api` package and `HelloHandler` with a capital to make it public.

We can do something similar for `api/v1/goodbye.go`:

```go
package api

import (
	"net/http"

	model "github.com/okeeffed/go-1-22-mux/internal/model"
)

// HelloHandler is a simple HTTP handler that writes a response.
func GoodbyeHandler(w http.ResponseWriter, r *http.Request) {
	// Create a new user.
	user := model.User{
		Name: "World",
	}

	// Write a response to the client.
	w.Write([]byte("Goodbye, " + user.Name + "!"))
}
```

## Configuring our main application

As seen in our **Makefile**, our main entry point is `cmd/server/server.go`.

The convention `cmd/<app-name>/<app-name>.go` is common across larger projects, although you would normally just see a simple `main.go` at the root directory for small applications.

Within that file, we will add our server code:

```go
package main

import (
	"net/http"

	v1 "github.com/okeeffed/go-1-22-mux/api/v1"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /v1/hello", v1.HelloHandler)
	mux.HandleFunc("GET /v1/goodbye", v1.GoodbyeHandler)

	http.ListenAndServe(":8080", mux)
}
```

In our code, we are using Go 1.22's [NewServeMux](https://pkg.go.dev/net/http@master#NewServeMux) function to serve our routers. We also import `v1` to use both our `HelloHandler` and `GoodbyeHandler` functions that we wrote.

## Running our router

With out Makefile sorted, we can run either `make run` or simple `make` to build out our binary `hello_world`.

For now, run `make run`.

Once built, our server will be running on port 8080.

To test, we can use `curl` on another terminal window:

```s
# Test /v1/hello
$ curl http://localhost:8080/v1/hello
Hello, World!%
# Test /v1/goodbye
curl http://localhost:8080/v1/goodbye
Goodbye, World!%
```

Success!

## Conclusion

In today's post, we used the Golang project layout convention to help test drive out the Go 1.22 built-in mux router.

We did this from the process of installing Go on Mac until we had a running server.

## Resources and further reading

- [Download Go](https://go.dev/dl/)
- [Homebrew](https://brew.sh/)
- [Go project layout](https://github.com/golang-standards/project-layout)
- [NewServeMux | Go Docs](https://pkg.go.dev/net/http@master#NewServeMux)
