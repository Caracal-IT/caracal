# Go Best Practices

## Code Style and Formatting

### General Guidelines
- Use `gofmt` to format all code automatically
- Keep lines under 80 characters when possible
- Use meaningful variable and function names
- Use camelCase for variable and function names
- Use PascalCase for exported types and functions

### Naming Conventions

**Packages:**
- Use short, lowercase names
- Avoid generic names like "utils" or "common"
- Example: `http`, `strings`, `fmt`

**Variables and Functions:**
- Use camelCase for unexported names
- Use PascalCase for exported names
- Be descriptive but concise
- Avoid single letter variables except for loop counters and parameters

**Constants:**
- Use PascalCase for exported constants
- Use camelCase for unexported constants
- Example: `MaxRetries`, `defaultTimeout`

## Error Handling

```go
// Always check errors
if err != nil {
    return err
}

// Don't ignore errors with blank identifier
_ = doSomething()  // Avoid this

// Wrap errors with context (Go 1.13+)
if err != nil {
    return fmt.Errorf("failed to process: %w", err)
}
```

## Documentation

- Write godoc comments for all exported functions, types, and packages
- godoc comments start with the name of the item being documented
- Keep comments clear and concise
- Example:

```go
// Package main is the entry point for the application.
package main

// Add returns the sum of two integers.
func Add(a, b int) int {
    return a + b
}
```

## Code Organization

### Project Structure
```
project/
├── main.go           # Entry point
├── go.mod
├── go.sum
├── internal/         # Unexported packages
│   ├── auth/
│   └── database/
├── pkg/              # Exported packages
│   ├── api/
│   └── models/
├── cmd/              # Command-line tools
│   └── cli/
└── docs/             # Documentation
```

### Package Organization
- Keep related functionality in the same package
- Use `internal/` for packages not meant to be imported elsewhere
- Use `pkg/` for packages that can be imported by external projects

## Concurrency

### Goroutines
- Use goroutines for concurrent operations
- Always provide a way to stop goroutines gracefully
- Use channels for communication between goroutines

```go
// Good pattern with context
func process(ctx context.Context) {
    select {
    case <-ctx.Done():
        return
    default:
        // Process
    }
}
```

## Testing

- Write tests for all public functions
- Use `*testing.T` for unit tests
- Use `*testing.B` for benchmarks
- Place tests in the same package with `_test.go` suffix
- Example: `main_test.go`

```go
func TestAdd(t *testing.T) {
    result := Add(2, 3)
    if result != 5 {
        t.Errorf("expected 5, got %d", result)
    }
}
```

## Performance

- Avoid unnecessary allocations
- Use `sync.Pool` for object reuse
- Profile your code with `pprof`
- Benchmark critical sections
- Use `strings.Builder` instead of `+` for string concatenation

## Dependencies

- Keep dependencies minimal
- Use Go modules for dependency management
- Run `go mod tidy` regularly
- Keep Go version updated

## Common Patterns

### Interface-based Design
```go
type Reader interface {
    Read(p []byte) (n int, err error)
}
```

### Functional Options
```go
type Config struct {
    timeout time.Duration
    retries int
}

type Option func(*Config)

func WithTimeout(d time.Duration) Option {
    return func(c *Config) {
        c.timeout = d
    }
}
```

## Tools and Commands

- `go fmt ./...` - Format code
- `go vet ./...` - Run static analysis
- `go test ./...` - Run tests
- `go test -cover ./...` - Check test coverage
- `golint ./...` - Run linter
- `go mod tidy` - Clean up dependencies

## Resources

- [Effective Go](https://golang.org/doc/effective_go)
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- [Standard Go Project Layout](https://github.com/golang-standards/project-layout)

