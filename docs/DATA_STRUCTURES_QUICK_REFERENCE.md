# Data Structures - Quick Reference

A concise reference for all data structures in the webserver project.

---

## Type Aliases at a Glance

| Type Alias | Base Type | Purpose | Example |
|-----------|-----------|---------|---------|
| `URL` | `string` | Request path | `"/api/users"` |
| `METHOD` | `string` | HTTP verb | `GET`, `POST`, `PUT` |
| `CONTENT` | `string` | MIME type | `"application/json"` |
| `STATUS` | `string` | HTTP status | `"200 OK"` |
| `Port` | `int` | Port number | `8000` |

---

## Structs at a Glance

| Struct | Fields | Purpose |
|--------|--------|---------|
| `Header` | `key`, `value` | HTTP header pair |
| `Request` | `Url`, `HttpMethod`, `Headers`, `Version` | Parsed HTTP request |
| `Response` | `Body`, `Status`, `Content`, `Headers` | HTTP response |
| `Route` | `Url`, `Method` | Route pattern |
| `ControlledRoutes` | `route`, `controller` | Route + handler |
| `Server` | `port`, `RouteControllers`, `AddController`, `Listen` | Main server object |
| `Config` | (empty) | Server configuration |

---

## Function Type Aliases at a Glance

| Type Alias | Signature | Purpose |
|-----------|-----------|---------|
| `Controller` | `func(Request, Response) Response` | Request handler |
| `ListenerType` | `func(string, []ControlledRoutes)` | Server listener |
| `RouterType` | `func(Request, []ControlledRoutes) Response` | Route matcher |

---

## HTTP Methods (9 constants)

```go
GET, POST, PUT, DELETE, PATCH, HEAD, OPTIONS, TRACE, CONNECT
```

---

## HTTP Status Codes (29 constants)

**2xx Success:** OK, CREATED, ACCEPTED, NON_AUTHORITATIVE, NO_CONTENT, RESET_CONTENT, PARTIAL_CONTENT

**3xx Redirect:** MULTIPLE_CHOICES, MOVED_PERMANENTLY, FOUND, SEE_OTHER, NOT_MODIFIED, TEMP_REDIRECT, PERM_REDIRECT

**4xx Client Error:** BAD_REQUEST, UNAUTHORIZED, PAYMENT_REQUIRED, FORBIDDEN, NOT_FOUND, METHOD_NOT_ALLOWED, NOT_ACCEPTABLE, PROXY_AUTH_REQUIRED, REQUEST_TIMEOUT, CONFLICT, GONE, LENGTH_REQUIRED, PRECONDITION_FAILED, PAYLOAD_TOO_LARGE, URI_TOO_LONG, UNSUPPORTED_MEDIA, IM_A_TEAPOT

**5xx Server Error:** INTERNAL_SERVER_ERROR, NOT_IMPLEMENTED, BAD_GATEWAY, SERVICE_UNAVAILABLE, GATEWAY_TIMEOUT, HTTP_VERSION_NOT_SUPPORTED

---

## MIME Types (35 constants)

**Text:** HTML, PLAIN, CSS, CSV, XML_TEXT, MARKDOWN

**Application:** JSON, JAVASCRIPT, XML_APP, FORM_URLENCODED, PDF, ZIP, GZIP, MSWORD, MS_EXCEL, MS_POWERPOINT

**Image:** PNG, JPEG, GIF, SVG, WEBP, ICO

**Audio/Video:** MP3, OGG_AUDIO, MP4, WEBM

**Font:** WOFF, WOFF2, TTF, OTF

---

## Common Patterns

### Create Server and Add Routes
```go
srv := server.InitServer(server.Config{})
srv.AddController(Route{Url: "/", Method: GET}, homeController)
srv.Listen("8000", srv.RouteControllers)
```

### Create Request
```go
req := Request{
    Url:        URL("/api/users"),
    HttpMethod: GET,
    Version:    "HTTP/1.1",
}
```

### Create Response
```go
res := Response{
    Body:    "Hello",
    Status:  OK,
    Content: JSON,
}
```

### Create Controller
```go
func(req Request, res Response) Response {
    return Response{
        Body:    "Data",
        Status:  OK,
        Content: JSON,
    }
}
```

### Parse Request
```go
req, err := ParseRequest("GET /test HTTP/1.1\n")
if err == nil {
    // Use req
}
```

### Serialize Response
```go
res := Response{Body: "Hello", Status: OK, Content: PLAIN}
httpString := res.Serialize()
// Send to client
```

---

## File Locations

| Type | Location |
|------|----------|
| URL, METHOD, CONTENT, STATUS, Request, Response, Route, Header | `server/http/http.go` |
| Controller | `server/controller/controller.go` |
| ControlledRoutes, RouterType, Router(), AddController() | `server/router/router.go` |
| Server, Config, Port | `server/server.go` |
| Port, ListenerType, Listen() | `server/networking/networking.go` |

---

## Field Visibility

| Type | Exported Fields | Unexported Fields |
|------|-----------------|-------------------|
| Header | None | `key`, `value` |
| Request | `Url`, `HttpMethod`, `Headers`, `Version` | None |
| Response | `Body`, `Status`, `Headers`, `Content` | None |
| Route | `Url`, `Method` | None |
| ControlledRoutes | None | `route`, `controller` |
| Server | `RouteControllers`, `AddController`, `Listen` | `port` |

---

## Type Conversion

```go
// URL to string
urlStr := string(url)

// METHOD to string
methodStr := string(method)

// CONTENT to string
mimeStr := string(content)

// STATUS to string
statusStr := string(status)

// String to URL
url := URL("path")

// String to METHOD
method := METHOD("GET")
```

---

## Zero Values

```go
var url URL                          // ""
var method METHOD                    // ""
var content CONTENT                  // ""
var status STATUS                    // ""
var req Request                      // All zero values
var res Response                     // All zero values
var route Route                      // {"", ""}
var srv *Server                      // nil
```

---

## Testing Checklist

- [ ] Create URL values
- [ ] Create METHOD constants
- [ ] Create CONTENT constants
- [ ] Create STATUS constants
- [ ] Parse Request
- [ ] Serialize Response
- [ ] Create routes
- [ ] Register controllers
- [ ] Match routes
- [ ] Handle 404s
- [ ] Call controllers with different methods
- [ ] Verify response body content
- [ ] Verify response status codes
- [ ] Verify content types

---

## Common Errors to Avoid

```go
// ❌ Cannot access unexported Header fields
header := Header{key: "Host", value: "localhost"}
println(header.key)  // Error!

// ❌ Cannot access unexported ControlledRoutes fields
cr := ControlledRoutes{...}
route := cr.route  // Error!

// ❌ Router matches by URL only, not method
// Both routes will match if URL is same:
Route{Url: "/api", Method: GET}
Route{Url: "/api", Method: POST}  // Second one overrides first

// ❌ Response with no Status or Content
res := Response{Body: "data"}
// Missing Status and Content - won't have proper headers

// ✓ Always provide Status and Content
res := Response{Body: "data", Status: OK, Content: JSON}

// ❌ Cannot modify Request or Response after creation
req := Request{Url: "/test", ...}
req.Url = "/new"  // Works (they're exported), but not idiomatic

// ✓ Create new instances instead
newReq := Request{Url: "/new", ...}
```

---

## Method Chaining Pattern

The project **does not use method chaining**. Instead, it uses:
- Functional composition
- Function fields
- Direct function calls

```go
// ✓ Functional style (used in project)
server.AddController(route, controller)
response := router.Router(request, routes)

// ✗ NOT used in project
server.AddController(route, controller).WithTimeout(30)
```

---

## Immutability

All data structures are **designed to be immutable** after creation:

```go
// Create once
res := Response{Body: "Hello", Status: OK}

// Use as-is
resString := res.Serialize()

// Don't modify
res.Body = "World"  // Technically allowed but not idiomatic
```

---

## Constants Reference

### HTTP Versions (defined in code)
- `HTTPVERSION = "HTTP/1.1"`

### Network (defined in networking.go)
- `HOST = "localhost"`
- `TYPE = "tcp"`

### All Other Constants
- 9 METHOD constants
- 29 STATUS constants
- 35 CONTENT constants

---

## Performance Notes

- **String comparison:** O(n) for URL matching
- **Route lookup:** O(n) linear scan through routes
- **Request parsing:** O(n) single pass through raw bytes
- **Response serialization:** O(n) string formatting
- **Memory:** No pre-allocation, grows as routes added

For small servers (< 1000 routes), performance is fine. Large servers would benefit from:
- Route trie/tree structure
- Binary search for method dispatch
- Pre-allocated buffer pools

---

## Future Enhancement Ideas

### Config Struct Could Include
```go
type Config struct {
    Port        Port
    Host        string
    Timeout     int
    MaxBodySize int64
    TLS         bool
    CertFile    string
    KeyFile     string
}
```

### Response Could Support Custom Headers
```go
res := Response{
    Body:    "Data",
    Status:  OK,
    Content: JSON,
    Headers: []Header{
        {key: "X-Custom", value: "Value"},
        {key: "Cache-Control", value: "no-cache"},
    },
}
// Would need to update Serialize() to include them
```

### Router Could Match by Method
```go
// Find route by both URL and Method
if route.Url == request.Url && route.Method == request.HttpMethod {
    // Execute controller
}
```

### Support Path Parameters
```go
Route{Url: "/users/{id}", Method: GET}
Request{Url: "/users/123"}  // Should match and extract "123"
```

---

## Architecture Decisions

1. **Type Aliases over Constants:** Semantic meaning with compile-time checking
2. **Function Fields over Methods:** Composition over inheritance
3. **Unexported Fields in ControlledRoutes:** Encapsulation, prevent direct access
4. **String-based Status/Content:** Simplicity, no enum overhead
5. **Immutable Structs:** Functional programming style
6. **No Interfaces:** Keep it simple, use concrete types

---

## When to Use Each Type

| Need | Use |
|------|-----|
| Represent a URL path | `URL` type alias |
| Represent HTTP method | `METHOD` type alias or constant |
| Represent response format | `CONTENT` type alias or constant |
| Represent response status | `STATUS` type alias or constant |
| Handle incoming HTTP request | `Request` struct |
| Send HTTP response | `Response` struct |
| Define a route pattern | `Route` struct |
| Register route + handler | Use `server.AddController()` |
| Create request handler | `Controller` function type |
| Customize server | `Config` struct (for future use) |

---

## Complete Type Hierarchy

```
string ─────────────────► URL, METHOD, CONTENT, STATUS
func ──────────────────► Controller, ListenerType, RouterType
int ───────────────────► Port

Struct ─────────────────► Header, Request, Response, Route, 
                         ControlledRoutes, Server, Config
```

---

## Quick Start

```go
// 1. Initialize server
server := server.InitServer(server.Config{})

// 2. Define controller
homeCtrl := func(req http.Request, res http.Response) http.Response {
    return http.Response{
        Body:    "<h1>Welcome</h1>",
        Status:  http.OK,
        Content: http.HTML,
    }
}

// 3. Register route
server.AddController(
    http.Route{Url: "/", Method: http.GET},
    homeCtrl,
)

// 4. Start server
server.Listen("8000", server.RouteControllers)
```

That's it! The server is now running and ready to handle requests.

