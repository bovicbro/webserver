# Agent Guidelines for WebServer Project

This document provides essential information for AI coding agents working in this codebase.

## Project Overview

- **Language**: Go 1.21.6+
- **Type**: Custom HTTP web server built from scratch using standard library only
- **Module**: `webserver`
- **Status**: Work in progress educational project
- **No external dependencies** - pure standard library implementation

## Build, Test, and Lint Commands

### Build Commands
```bash
go build              # Build executable (outputs: webserver)
go build -v           # Build with verbose output
go run main.go        # Run server directly without building
go clean              # Remove build artifacts
```

### Test Commands
```bash
# Run all tests
go test ./...

# Run all tests with verbose output
go test ./... -v

# Run tests for a single package
go test webserver/utility
go test webserver/server/networking

# Run a specific test function
go test webserver/utility -run TestSliceIndexOf
go test webserver/utility -run TestSliceIndexOfHappy

# Run tests with coverage
go test ./... -cover
go test ./... -coverprofile=coverage.out
```

### Linting and Formatting
```bash
# Format code (always run before committing)
gofmt -w .

# Check which files need formatting
gofmt -l .

# Run static analysis
go vet ./...

# Tidy dependencies
go mod tidy
```

## Code Style Guidelines

### Package Structure
```
webserver/
├── main.go                    # Entry point only
├── server/                    # Core server implementation
│   ├── server.go             # Server initialization
│   ├── controller/           # Controller types
│   ├── http/                 # HTTP protocol
│   ├── networking/           # TCP listener
│   └── router/               # Route matching
├── utility/                   # Generic utilities
└── static/                    # Static assets
```

### Import Conventions
```go
import (
    // Standard library imports first (alphabetically)
    "errors"
    "fmt"
    "log"
    "net"
    "strings"
    
    // Then internal packages (alphabetically)
    "webserver/server/http"
    "webserver/server/router"
    "webserver/utility"
)
```

**Dot imports** are used sparingly for convenience in specific files:
```go
import . "webserver/server/http"  // Only in controller.go and router.go
```

### Naming Conventions

**Packages**: lowercase, single word
```go
package server
package networking
```

**Types**: PascalCase (exported), camelCase (unexported)
```go
type Server struct { ... }        // Exported
type routeController struct { ... } // Unexported
```

**Type Aliases**: PascalCase for domain concepts
```go
type URL string
type METHOD string
type STATUS string
type Port int
```

**Functions**: PascalCase (exported), camelCase (unexported)
```go
func InitServer(port Port) Server { ... }  // Exported
func handleRequest(conn net.Conn) { ... }  // Unexported
```

**Constants**: SCREAMING_SNAKE_CASE or UPPERCASE
```go
const (
    GET     METHOD = "GET"
    POST    METHOD = "POST"
    NOT_FOUND STATUS = "404"
    FORM_URLENCODED CONTENT = "application/x-www-form-urlencoded"
)
```

**Variables**: camelCase, prefer short names for common types
```go
var server Server
var req Request
var res Response
var rcs []ControlledRoutes  // rc = route controller
```

**Method receivers**: single letter or short abbreviation
```go
func (res Response) Serialize() string { ... }
```

### Error Handling

Follow these patterns based on context:

**1. Fatal exit for critical initialization errors:**
```go
listener, err := net.Listen("tcp", HOST+":"+port)
if err != nil {
    log.Fatal(err)
    os.Exit(1)
}
```

**2. Return errors to caller for recoverable issues:**
```go
func ParseRequest(reqRaw string) (Request, error) {
    if len(rows) < 1 {
        return req, errors.New("Malformed request")
    }
    return req, nil
}
```

**3. Convert errors to HTTP responses in handlers:**
```go
content, err := os.ReadFile("./static/index.html")
if err != nil {
    res = http.Response{Status: http.NOT_FOUND}
} else {
    res = http.Response{Body: string(content), Status: http.OK}
}
```

**General principles:**
- Always check errors explicitly with `if err != nil`
- Use `errors.New("message")` for simple errors
- Don't ignore errors silently
- Log errors before fatal exit

### Type Usage

**Function types** for flexibility:
```go
type Controller = func(req Request, res Response) Response
type ListenerType = func(port string, rcs []router.ControlledRoutes)
```

**Generics** for type-safe utilities:
```go
func SliceIndexOf[E any](slice []E, f func(E) bool) int { ... }
```

**Struct composition** for configuration:
```go
type Server struct {
    port             Port  // Unexported config
    RouteControllers []router.ControlledRoutes  // Exported for access
    AddController    func(...)  // Function fields for composition
    Listen           networking.ListenerType
}
```

### Documentation

Add comments for all exported identifiers:
```go
// SliceIndexOf returns the index of the first element in the slice that
// satisfies the predicate f. Returns -1 if no element is found.
func SliceIndexOf[E any](slice []E, f func(E) bool) int { ... }
```

Package-level comments at the top of files:
```go
// Package http handles HTTP protocol parsing and response serialization.
package http
```

### Testing

**File naming**: `*_test.go` in the same package
```go
// utility/utility_test.go
package utility
```

**Test naming**:
```go
func TestSliceIndexOfHappy(t *testing.T) { ... }          // Exported function
func Test_initListener(t *testing.T) { ... }              // Unexported function
```

**Test structure**:
```go
func TestFunctionName(t *testing.T) {
    // Arrange
    input := "test"
    expected := "result"
    
    // Act
    result := Function(input)
    
    // Assert
    if result != expected {
        t.Errorf("Expected %v, got %v", expected, result)
    }
}
```

## Development Workflow

1. **Before writing code**: Run `go test ./...` to ensure tests pass
2. **Write code**: Follow style guidelines above
3. **Test your changes**: Add/update tests, run `go test ./... -v`
4. **Format code**: Run `gofmt -w .` before committing
5. **Static analysis**: Run `go vet ./...` to catch issues
6. **Build**: Run `go build` to ensure compilation succeeds

## Common Patterns in This Codebase

- **No interfaces**: Concrete types throughout, function types for abstraction
- **Composition over inheritance**: Server struct composes functions
- **Minimal dependencies**: Standard library only
- **Simple error handling**: No custom error types
- **Hardcoded config**: Currently uses constants (see networking.go:12)
- **Single method receivers**: Most logic in standalone functions

## Notes for Agents

- This is an educational project building an HTTP server from scratch
- Prefer simplicity over frameworks
- Don't add external dependencies without discussion
- Test coverage is limited - add tests when modifying code
- Server runs on localhost:8000 by default
- Static files are in `./static/` directory
