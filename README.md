# Golang Starter Kit

A starter kit written in Golang + Ports and Adapters structure on top of <https://github.com/golang-standards/project-layout>

## Getting Started

  1. `git clone git@github.com/yuraxdrumz/golang-starter-kit`
  2. `go install golang.org/x/tools/cmd/goimports@latest`
  3. `go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.55.2`
  4. Run `brew install pre-commit` / `pip install pre-commit`
  5. Run `pre-commit`

If you want to check pre-commit hook on current folder run `pre-commit run --all-files`

### Prerequisites

Ports and Adapters divides your code to 3 parts:

- Business-logic - These are your business rules + types, implemented without any dependency on 3rd party modules (self-contained)
- Ports - Interfaces to speak with your business rules
- Adapters - Implementations of the ports, There are two kinds of adapters:
- In(Driver) - your external API to the world. For example - `internal/pkg/adapters/in/http.go`
- Out(Driven) - what your business logic uses. For example - `internal/pkg/adapters/out/reverser/in-memory.go`

Usually, you divide `ports` and `adapters` to separate directories, but the best practice in golang is to keep structs near implementations. That is why, I decided to add `ports.go` near each adapter.

Adding a new business logic:

  1. Create appropriate structs in `ports.go` file under `internal/app/your-use-case/ports.go`
  2. Create your use-case with your application specific logic under `internal/app/your-use-case/logic.go`
  3. Create your in/out adapter, for example - `Repository(out)` or `gRPC(in)`. `internal/pkg/adapters/*`
  4. Tests!!! `your-file-name_test.go` under same directory as file `internal/app/your-use-case/logic_test.go`

### Installing Packages

```golang
go get -u <repository-name>
```

### Running the service

```bash
go run main.go
```

### Environment Variables

To read environement variables, this service uses `envconfig` library, which allows defining a struct of environment variables. You can easily add new variables under `internal/pkg/adapters/out/env`

The default is:

```golang

type Specification struct {
  LogLevel    string `envconfig:"LOG_LEVEL" default:"info"`
  Port        string `envConfig:"PORT" default:"8080"`
}

```

By default this service looks for .env in root folder.

### Logging

The repo uses `logrus` logger to write to stdout.

To change log level change the `LOG_LEVEL` environment variable

Possible log levels:

- info - default
- debug
- error
- warn
- fatal
