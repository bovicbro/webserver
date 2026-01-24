# Data Structures Documentation

This directory contains comprehensive documentation about all data structures used in the webserver project.

## Documents

### 1. **DATA_STRUCTURES.md** (Main Reference)
The comprehensive guide covering all data structures with detailed explanations:
- All type aliases (URL, METHOD, CONTENT, STATUS, Port)
- All constants (HTTP methods, status codes, MIME types)
- All structs (Header, Request, Response, Route, ControlledRoutes, Server, Config)
- All function types (Controller, ListenerType, RouterType)
- Data structure relationships and composition
- Memory and performance considerations

**Best for:** Understanding the complete picture of every type in the project

### 2. **DATA_STRUCTURES_DIAGRAMS.md** (Visual Guide)
Visual representations of data structures and their relationships:
- HTTP request/response flow diagram
- Server structure composition
- Request structure details with field descriptions
- Response structure details with field descriptions
- Route matching process
- Type alias hierarchy
- Memory layout and flow
- Struct field visibility (exported/unexported)
- Type constants organization
- Data flow in main.go

**Best for:** Visual learners who want to understand how pieces fit together

### 3. **DATA_STRUCTURES_EXAMPLES.md** (Practical Guide)
Real code examples showing how to use each data structure:
- Type alias usage examples
- Struct creation and usage examples
- Function type alias examples
- Complete real-world example
- Testing examples
- Common usage patterns

**Best for:** Learning by doing, copy-paste starting points

### 4. **DATA_STRUCTURES_QUICK_REFERENCE.md** (Cheat Sheet)
Quick lookup reference for common tasks:
- Type aliases at a glance (table)
- Structs at a glance (table)
- Function types at a glance (table)
- HTTP methods, status codes, MIME types lists
- Common patterns and code snippets
- File locations of each type
- Field visibility matrix
- Type conversion
- Testing checklist
- Common errors to avoid

**Best for:** Quick lookups while coding

---

## Quick Navigation

### "I want to understand what types exist"
→ Read **DATA_STRUCTURES_QUICK_REFERENCE.md** first, then **DATA_STRUCTURES.md**

### "I want to see how things connect"
→ Read **DATA_STRUCTURES_DIAGRAMS.md**

### "I want to write code using these types"
→ Read **DATA_STRUCTURES_EXAMPLES.md**

### "I need to remember something specific"
→ Check **DATA_STRUCTURES_QUICK_REFERENCE.md** for tables and lists

### "I need complete details"
→ Read **DATA_STRUCTURES.md** (main reference)

---

## Type Summary

The project uses **minimal data structures**:

### Type Aliases (5)
- `URL` - Request paths
- `METHOD` - HTTP verbs  
- `CONTENT` - MIME types
- `STATUS` - HTTP status codes
- `Port` - Network port numbers

### Structs (7)
- `Header` - HTTP header pair
- `Request` - Parsed HTTP request
- `Response` - HTTP response
- `Route` - Route pattern definition
- `ControlledRoutes` - Route + handler pair
- `Server` - Main server object
- `Config` - Server configuration

### Function Types (3)
- `Controller` - Request handler function
- `ListenerType` - Server listener function
- `RouterType` - Route matcher function

### Constants (73 total)
- 9 HTTP methods (GET, POST, PUT, DELETE, PATCH, HEAD, OPTIONS, TRACE, CONNECT)
- 29 HTTP status codes (2xx, 3xx, 4xx, 5xx)
- 35 MIME type constants

---

## Data Flow

```
Raw TCP Bytes
    ↓
ParseRequest() → Request struct
    ↓
Router() → finds ControlledRoutes
    ↓
Controller() → creates Response struct
    ↓
Serialize() → HTTP response string
    ↓
Send to client
```

---

## Key Design Decisions

1. **String-based type aliases** for semantic clarity without runtime overhead
2. **Immutable structs** following functional programming style
3. **Function fields** (composition) instead of methods (inheritance)
4. **Unexported fields** for encapsulation where needed (ControlledRoutes)
5. **Concrete types** throughout, no interfaces
6. **Simple constants** instead of complex enums

---

## Common Patterns

### Register a route
```go
server.AddController(
    http.Route{Url: "/path", Method: http.GET},
    func(req http.Request, res http.Response) http.Response {
        return http.Response{Body: "...", Status: http.OK, Content: http.HTML}
    },
)
```

### Handle a request
```go
func(req http.Request, res http.Response) http.Response {
    // Access request
    if req.HttpMethod == http.GET {
        // Return response
        return http.Response{
            Body:    "Response body",
            Status:  http.OK,
            Content: http.JSON,
        }
    }
    return http.Response{Status: http.METHOD_NOT_ALLOWED}
}
```

---

## File Locations

```
server/http/http.go           - URL, METHOD, CONTENT, STATUS, Header, Request, Response, Route
server/router/router.go       - ControlledRoutes, RouterType, Router(), AddController()
server/controller/controller.go - Controller type
server/server.go              - Server, Config, Port
server/networking/networking.go - Port, ListenerType, Listen()
utility/utility.go            - SliceIndexOf() (uses generics)
```

---

## Learning Path

1. **Start here:** DATA_STRUCTURES_QUICK_REFERENCE.md
2. **Understand flow:** DATA_STRUCTURES_DIAGRAMS.md
3. **See examples:** DATA_STRUCTURES_EXAMPLES.md
4. **Deep dive:** DATA_STRUCTURES.md
5. **Reference while coding:** Keep QUICK_REFERENCE open

---

## Testing Data Structures

All data structures have comprehensive unit tests:
- `utility/utility_test.go` - SliceIndexOf tests
- `server/http/http_test.go` - Request/Response tests (20 tests)
- `server/router/router_test.go` - Route/Router tests (13 tests)
- `server/controller/controller_test.go` - Controller tests (12 tests)
- `server/server_test.go` - Server tests (12 tests)
- `server/networking/networking_test.go` - Networking tests (14 tests)

Run tests:
```bash
go test ./... -v          # All tests
go test ./... -cover      # With coverage
go test webserver/server/http -v  # Single package
```

---

## Frequently Asked Questions

### Q: Can I modify a Request/Response after creating it?
A: Technically yes, but not recommended. Structs are designed to be immutable.

### Q: Why are some struct fields unexported (lowercase)?
A: For encapsulation. `ControlledRoutes.route` and `ControlledRoutes.controller` are unexported to prevent direct access.

### Q: How do I add a custom HTTP header to a response?
A: The `Response.Headers` field exists but isn't currently serialized. You'd need to modify the `Serialize()` method.

### Q: Can the router match by HTTP method?
A: Not currently. The router only matches by URL. Routes with different methods on the same URL will collide.

### Q: What's the difference between ParseRequest() and creating a Request directly?
A: `ParseRequest()` parses raw HTTP request bytes and extracts method, URL, version, headers. Direct creation is for testing.

### Q: Why use function fields instead of methods on Server?
A: This allows different server instances to have different behaviors (composition over inheritance).

---

## Additional Resources

- **AGENTS.md** - Guidelines for AI agents working in this codebase (includes build/test commands)
- **README.md** - Project overview
- Unit tests - See actual usage in `*_test.go` files

---

## Summary

This webserver project demonstrates **simplicity through minimal types**:
- No external dependencies
- No interfaces (concrete types throughout)
- No complex generics (except utility.SliceIndexOf)
- Simple immutable data structures
- Functional composition style

The data structures are **purpose-built for HTTP servers**:
- URL, METHOD, STATUS, CONTENT capture HTTP concepts
- Request/Response handle protocol parsing and serialization
- Route/ControlledRoutes implement basic routing
- Server/Config provide the main application interface

Everything is **documented with examples, diagrams, and tests**.

Happy coding! 🚀
