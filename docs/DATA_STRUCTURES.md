# Data Structures in WebServer

This document describes all data structures (types, type aliases, and structs) used throughout the webserver project.

## Overview

The project uses a minimal set of data structures focused on HTTP protocol handling and request routing. There are no complex nested structures or third-party types—only what's necessary for an HTTP server.

---

## Type Aliases

Type aliases provide semantic clarity by giving meaningful names to base types. They represent domain concepts in the HTTP server.

### `type URL string`
**Location:** `server/http/http.go:10`

**Purpose:** Represents the path component of an HTTP request (e.g., "/", "/about", "/api/users")

**Example:**
```go
url := URL("/api/users")
route := Route{Url: url, Method: GET}
```

**Usage in codebase:**
- `Request.Url` field
- `Route.Url` field
- Router matching by URL comparison

**Notes:**
- Stores only the path, not the full URI with query parameters
- No validation is performed on the URL value

---

### `type METHOD string`
**Location:** `server/http/http.go:17`

**Purpose:** Represents HTTP request methods

**Constants defined:**
```go
const (
    GET     METHOD = "GET"
    POST    METHOD = "POST"
    PUT     METHOD = "PUT"
    DELETE  METHOD = "DELETE"
    PATCH   METHOD = "PATCH"
    HEAD    METHOD = "HEAD"
    OPTIONS METHOD = "OPTIONS"
    TRACE   METHOD = "TRACE"
    CONNECT METHOD = "CONNECT"
)
```

**Example:**
```go
method := METHOD("GET")
route := Route{Method: method}
```

**Usage in codebase:**
- `Request.HttpMethod` field
- `Route.Method` field
- Controller condition checking (e.g., `if req.HttpMethod == GET`)

**Notes:**
- 9 standard HTTP methods are defined as constants
- Can store any string value, not just the predefined constants
- Router currently matches only by URL, not by method

---

### `type CONTENT string`
**Location:** `server/http/http.go:31`

**Purpose:** Represents MIME types for HTTP response Content-Type headers

**Constants defined (35 total):**

**Text types:**
```go
HTML            CONTENT = "text/html"
PLAIN           CONTENT = "text/plain"
CSS             CONTENT = "text/css"
CSV             CONTENT = "text/csv"
XML_TEXT        CONTENT = "text/xml"
MARKDOWN        CONTENT = "text/markdown"
```

**Application types:**
```go
JSON            CONTENT = "application/json"
JAVASCRIPT      CONTENT = "application/javascript"
XML_APP         CONTENT = "application/xml"
FORM_URLENCODED CONTENT = "application/x-www-form-urlencoded"
PDF             CONTENT = "application/pdf"
ZIP             CONTENT = "application/zip"
GZIP            CONTENT = "application/gzip"
MSWORD          CONTENT = "application/msword"
MS_EXCEL        CONTENT = "application/vnd.ms-excel"
MS_POWERPOINT   CONTENT = "application/vnd.ms-powerpoint"
```

**Image types:**
```go
PNG             CONTENT = "image/png"
JPEG            CONTENT = "image/jpeg"
GIF             CONTENT = "image/gif"
SVG             CONTENT = "image/svg+xml"
WEBP            CONTENT = "image/webp"
ICO             CONTENT = "image/x-icon"
```

**Audio/Video types:**
```go
MP3             CONTENT = "audio/mpeg"
OGG_AUDIO       CONTENT = "audio/ogg"
MP4             CONTENT = "video/mp4"
WEBM            CONTENT = "video/webm"
```

**Font types:**
```go
WOFF            CONTENT = "font/woff"
WOFF2           CONTENT = "font/woff2"
TTF             CONTENT = "font/ttf"
OTF             CONTENT = "font/otf"
```

**Example:**
```go
response := Response{
    Body: "<html>...</html>",
    Content: HTML,
}
```

**Usage in codebase:**
- `Response.Content` field
- Set by controllers to specify response MIME type

**Notes:**
- Comprehensive MIME type coverage for web assets
- Default value is empty string (no content type)

---

### `type STATUS string`
**Location:** `server/http/http.go:77`

**Purpose:** Represents HTTP status codes with their reason phrases

**Constants defined (29 total):**

**1xx Informational:**
```go
CONTINUE            STATUS = "100 Continue"
SWITCHING_PROTOCOLS STATUS = "101 Switching Protocols"
PROCESSING          STATUS = "102 Processing"
```

**2xx Success:**
```go
OK                STATUS = "200 OK"
CREATED           STATUS = "201 Created"
ACCEPTED          STATUS = "202 Accepted"
NON_AUTHORITATIVE STATUS = "203 Non-Authoritative Information"
NO_CONTENT        STATUS = "204 No Content"
RESET_CONTENT     STATUS = "205 Reset Content"
PARTIAL_CONTENT   STATUS = "206 Partial Content"
```

**3xx Redirection:**
```go
MULTIPLE_CHOICES  STATUS = "300 Multiple Choices"
MOVED_PERMANENTLY STATUS = "301 Moved Permanently"
FOUND             STATUS = "302 Found"
SEE_OTHER         STATUS = "303 See Other"
NOT_MODIFIED      STATUS = "304 Not Modified"
TEMP_REDIRECT     STATUS = "307 Temporary Redirect"
PERM_REDIRECT     STATUS = "308 Permanent Redirect"
```

**4xx Client Errors:**
```go
BAD_REQUEST         STATUS = "400 Bad Request"
UNAUTHORIZED        STATUS = "401 Unauthorized"
PAYMENT_REQUIRED    STATUS = "402 Payment Required"
FORBIDDEN           STATUS = "403 Forbidden"
NOT_FOUND           STATUS = "404 Not Found"
METHOD_NOT_ALLOWED  STATUS = "405 Method Not Allowed"
NOT_ACCEPTABLE      STATUS = "406 Not Acceptable"
PROXY_AUTH_REQUIRED STATUS = "407 Proxy Authentication Required"
REQUEST_TIMEOUT     STATUS = "408 Request Timeout"
CONFLICT            STATUS = "409 Conflict"
GONE                STATUS = "410 Gone"
LENGTH_REQUIRED     STATUS = "411 Length Required"
PRECONDITION_FAILED STATUS = "412 Precondition Failed"
PAYLOAD_TOO_LARGE   STATUS = "413 Payload Too Large"
URI_TOO_LONG        STATUS = "414 URI Too Long"
UNSUPPORTED_MEDIA   STATUS = "415 Unsupported Media Type"
IM_A_TEAPOT         STATUS = "418 I'm a teapot"
```

**5xx Server Errors:**
```go
INTERNAL_SERVER_ERROR      STATUS = "500 Internal Server Error"
NOT_IMPLEMENTED            STATUS = "501 Not Implemented"
BAD_GATEWAY                STATUS = "502 Bad Gateway"
SERVICE_UNAVAILABLE        STATUS = "503 Service Unavailable"
GATEWAY_TIMEOUT            STATUS = "504 Gateway Timeout"
HTTP_VERSION_NOT_SUPPORTED STATUS = "505 HTTP Version Not Supported"
```

**Example:**
```go
response := Response{
    Body: "Page not found",
    Status: NOT_FOUND,  // "404 Not Found"
}
```

**Usage in codebase:**
- `Response.Status` field
- Set by controllers to indicate response status
- Used in HTTP response serialization

**Notes:**
- Includes the full status line (code + reason phrase)
- Covers most common HTTP status codes
- Format: "CODE Reason Phrase"

---

## Structs

Structs are concrete types that hold multiple related fields.

### `struct Header`
**Location:** `server/http/http.go:12`

**Fields:**
```go
type Header struct {
    key   string  // Header name (e.g., "Content-Length")
    value string  // Header value (e.g., "1024")
}
```

**Purpose:** Represents an HTTP header key-value pair

**Example:**
```go
header := Header{
    key:   "Content-Type",
    value: "text/html",
}
```

**Usage in codebase:**
- `Request.Headers` field (slice of headers)
- `Response.Headers` field (slice of headers)

**Status:** Currently **not fully implemented** - headers are parsed but not fully utilized in request handling

**Notes:**
- Fields are unexported (lowercase), so headers cannot be directly accessed outside the package
- Headers are extracted from request but not actively used in routing
- Response currently does not send custom headers

---

### `struct Request`
**Location:** `server/http/http.go:152`

**Fields:**
```go
type Request struct {
    Url        URL       // The request path (e.g., "/api/users")
    HttpMethod METHOD    // The HTTP method (GET, POST, etc.)
    Headers    []Header  // Array of HTTP headers from the request
    Version    string    // HTTP version (e.g., "HTTP/1.1")
}
```

**Purpose:** Represents a parsed HTTP request

**Example:**
```go
req := Request{
    Url:        URL("/api/users"),
    HttpMethod: GET,
    Version:    "HTTP/1.1",
    Headers:    []Header{},
}
```

**Creation:** Via `ParseRequest(reqRaw string)` function which extracts:
1. First line of HTTP request: `GET /path HTTP/1.1`
2. Headers: (stored but not currently used)

**Usage flow:**
1. Raw HTTP request received by `handleRequest()` in networking
2. Parsed into `Request` struct using `ParseRequest()`
3. Passed to `Router()` to find matching route
4. Passed to matched controller function

**Examples in codebase:**
```go
// In networking.go
req, err := http.ParseRequest(string(buffer))
res := router.Router(req, rcs)

// In controllers
server.AddController(
    http.Route{Url: "/", Method: http.GET},
    func(req http.Request, res http.Response) http.Response {
        // Access request data
        println(req.Url)
        println(req.HttpMethod)
        return res
    })
```

**Fields detail:**

| Field | Type | Example | Notes |
|-------|------|---------|-------|
| `Url` | `URL` | `"/api/users"` | Path only, parsed from request line |
| `HttpMethod` | `METHOD` | `GET` | Extracted from first word of request line |
| `Headers` | `[]Header` | `[]Header{{...}}` | Currently parsed but not used |
| `Version` | `string` | `"HTTP/1.1"` | Extracted from third word of request line |

---

### `struct Response`
**Location:** `server/http/http.go:131`

**Fields:**
```go
type Response struct {
    Body    string   // The response body content
    Status  STATUS   // HTTP status code with reason phrase
    Headers []Header // Array of HTTP headers for the response
    Content CONTENT  // MIME type for Content-Type header
}
```

**Purpose:** Represents an HTTP response to be sent to the client

**Example:**
```go
res := Response{
    Body:    "<html><body>Hello</body></html>",
    Status:  OK,
    Content: HTML,
    Headers: []Header{},
}
```

**Creation Methods:**

1. **CreateBaseResponse()** - Creates empty response:
```go
res := CreateBaseResponse(req)
// Returns: Response{Body: "", Status: "", Headers: nil, Content: ""}
```

2. **Direct construction** - Most common in controllers:
```go
res := Response{
    Body:    "Data",
    Status:  http.OK,
    Content: http.JSON,
}
```

**Serialization:** `Response.Serialize()` method converts to HTTP response format:
```go
// Input:
Response{
    Body:    "Hello World",
    Status:  OK,
    Content: HTML,
}

// Output:
HTTP/1.1 200 OK
Content-length: 11
Content-Type: text/html; charset=utf-8

Hello World
```

**Usage flow:**
1. Controllers create or modify Response
2. Response is serialized to string format
3. Sent to client via TCP connection
4. Connection is closed

**Fields detail:**

| Field | Type | Purpose | Example |
|-------|------|---------|---------|
| `Body` | `string` | Response body content | `"Hello World"` |
| `Status` | `STATUS` | HTTP status line | `"200 OK"` |
| `Headers` | `[]Header` | Custom headers (unused) | `[]Header{}` |
| `Content` | `CONTENT` | MIME type | `"text/html"` |

**Notes:**
- Default empty Response has no status or content type
- Serialization adds `Content-Type: charset=utf-8` automatically
- Content-Length calculated from Body length
- Headers field exists but is not serialized in `Serialize()` method

---

### `struct Route`
**Location:** `server/http/http.go:159`

**Fields:**
```go
type Route struct {
    Url    URL    // The URL path to match
    Method METHOD // The HTTP method to match
}
```

**Purpose:** Defines a route pattern to match against incoming requests

**Example:**
```go
route := Route{
    Url:    URL("/api/users"),
    Method: GET,
}
```

**Usage in routing:**
```go
// Register a route with a controller
server.AddController(
    Route{Url: "/", Method: GET},
    func(req Request, res Response) Response {
        return Response{Body: "Home", Status: OK, Content: HTML}
    },
)
```

**Fields detail:**

| Field | Type | Matching | Notes |
|-------|------|----------|-------|
| `Url` | `URL` | Exact match only | `"/about"` matches only exactly `"/about"` |
| `Method` | `METHOD` | Exact match (GET, POST, PUT, DELETE, etc.) | Router now matches both URL and METHOD |

**Route Matching Behavior:**
Router matches routes based on **both URL and HTTP METHOD** exactly. This allows you to define different handlers for the same URL with different HTTP methods.

**Example with Multiple Methods:**
```go
// GET /api/users → Returns list of users
server.AddController(
    Route{Url: "/api/users", Method: GET},
    getUsersController,
)

// POST /api/users → Creates new user
server.AddController(
    Route{Url: "/api/users", Method: POST},
    createUserController,
)

// PUT /api/users → Updates user
server.AddController(
    Route{Url: "/api/users", Method: PUT},
    updateUserController,
)

// DELETE /api/users → Deletes user
server.AddController(
    Route{Url: "/api/users", Method: DELETE},
    deleteUserController,
)
```

**Limitations:**
- Routes like `/users/{id}` not supported (exact match only)
- No wildcards or regex patterns
- Dynamic URL parameters require separate implementation

---

### `struct ControlledRoutes`
**Location:** `server/router/router.go:10`

**Fields:**
```go
type ControlledRoutes struct {
    route      Route             // The route pattern to match
    controller Controller        // The function to handle requests
}
```

**Purpose:** Pairs a route with its handler function

**Example:**
```go
cr := ControlledRoutes{
    route: Route{Url: "/", Method: GET},
    controller: func(req Request, res Response) Response {
        return Response{Body: "Home", Status: OK, Content: HTML}
    },
}
```

**Field visibility:**
- Fields are unexported (lowercase)
- Cannot access directly outside router package
- Only accessible via router functions

**Usage:**
```go
// In server/server.go
rcs := []router.ControlledRoutes{}  // Slice of controlled routes

// Added via AddController
rcs = router.AddController(route, controller, rcs)

// Passed to router
response := router.Router(request, rcs)

// Passed to listener
server.Listen("8000", rcs)
```

**Related functions:**
```go
// Add a new controlled route
func AddController(route Route, controller Controller, rcs []ControlledRoutes) []ControlledRoutes

// Find and execute matching route
func Router(req Request, rcs []ControlledRoutes) Response
```

---

### `struct Server`
**Location:** `server/server.go:20`

**Fields:**
```go
type Server struct {
    port             Port                           // Server port (unexported)
    RouteControllers []router.ControlledRoutes      // List of registered routes (exported)
    AddController    func(...)                      // Function field for adding routes
    Listen           networking.ListenerType        // Function field for starting listener
}
```

**Full definition:**
```go
type Server struct {
    port             Port  // Unexported, not currently used
    RouteControllers []router.ControlledRoutes
    AddController    func(
        route http.Route,
        controller controller.Controller)
    Listen networking.ListenerType
}
```

**Purpose:** Main server object that manages routes and starts the HTTP server

**Creation:**
```go
server := server.InitServer(server.Config{})
```

**Field details:**

| Field | Type | Visibility | Purpose |
|-------|------|------------|---------|
| `port` | `Port` | Unexported | Intended for port config (unused) |
| `RouteControllers` | `[]router.ControlledRoutes` | Exported | List of all registered routes |
| `AddController` | `func(...)` | Exported | Adds a route-controller pair |
| `Listen` | `networking.ListenerType` | Exported | Starts the TCP listener |

**Usage example:**
```go
// Initialize server
server := server.InitServer(server.Config{})

// Add routes
server.AddController(
    http.Route{Url: "/", Method: http.GET},
    func(req http.Request, res http.Response) http.Response {
        return http.Response{Body: "Hello", Status: http.OK, Content: http.HTML}
    },
)

server.AddController(
    http.Route{Url: "/about", Method: http.GET},
    func(req http.Request, res http.Response) http.Response {
        return http.Response{Body: "About", Status: http.OK, Content: http.HTML}
    },
)

// Start listening
server.Listen("8000", server.RouteControllers)
```

**Design notes:**
- Uses function fields for composition instead of interfaces
- `AddController` and `Listen` are function fields, not methods
- `RouteControllers` is the only exported data field
- Follows functional programming patterns

---

### `struct Config`
**Location:** `server/server.go:10`

**Fields:**
```go
type Config struct {
    // Currently empty
}
```

**Purpose:** Placeholder for server configuration

**Current usage:**
```go
server := server.InitServer(server.Config{})
```

**Future potential:**
- Server port configuration
- Host configuration
- TLS/SSL settings
- Request timeouts
- Buffer sizes

**Notes:**
- Exists but is completely empty
- Passed to `InitServer()` but not used
- Designed for future expansion

---

## Type Aliases (Advanced)

### `type Controller`
**Location:** `server/controller/controller.go:5`

**Definition:**
```go
type Controller = func(req Request, res Response) Response
```

**Purpose:** Function type alias for HTTP request handlers

**Example:**
```go
controller := func(req http.Request, res http.Response) http.Response {
    return http.Response{Body: "Test", Status: http.OK, Content: http.HTML}
}
```

**Parameters:**
- `req Request` - Parsed HTTP request
- `res Response` - Base response (usually empty)

**Return:**
- `Response` - HTTP response to send to client

**Usage:**
```go
// Pass controller to AddController
server.AddController(route, controller)

// Or call directly in tests
result := controller(req, res)
```

**Typical pattern:**
```go
func(req http.Request, res http.Response) http.Response {
    if req.HttpMethod == http.GET {
        return http.Response{Body: "GET", Status: http.OK, Content: http.PLAIN}
    }
    return http.Response{Body: "Method not allowed", Status: http.METHOD_NOT_ALLOWED, Content: http.PLAIN}
}
```

---

### `type ListenerType`
**Location:** `server/networking/networking.go:20`

**Definition:**
```go
type ListenerType = func(port string, rcs []router.ControlledRoutes)
```

**Purpose:** Function type for starting the TCP listener

**Parameters:**
- `port string` - Port number as string (e.g., "8000")
- `rcs []router.ControlledRoutes` - Slice of registered routes

**Implementation:**
```go
func Listen(port string, rcs []router.ControlledRoutes) {
    for {
        initListener(port, rcs)
    }
}
```

**Usage:**
```go
// Assigned to server during initialization
server.Listen = networking.Listen

// Called when starting server
server.Listen("8000", server.RouteControllers)
```

---

### `type RouterType`
**Location:** `server/router/router.go:15`

**Definition:**
```go
type RouterType = func(req Request, rcs []ControlledRoutes) Response
```

**Purpose:** Function type for matching requests to routes

**Parameters:**
- `req Request` - Parsed HTTP request
- `rcs []ControlledRoutes` - Slice of registered routes

**Return:**
- `Response` - Response from matched controller or 404

**Implementation:**
```go
func Router(req Request, rcs []ControlledRoutes) Response {
    // Find matching route by URL
    // Call controller and return response
}
```

---

### `type Port`
**Location:** `server/server.go:18` and `server/networking.go:18`

**Definition:**
```go
type Port int
```

**Purpose:** Represents a network port number

**Example:**
```go
var port Port = 8000
```

**Notes:**
- Defined in two places (server and networking packages)
- Port values: 0-65535
- Currently not actively used in server initialization
- Intended for type safety and clarity

---

### `type AddControllerType`
**Location:** `server/router/router.go:28`

**Definition:**
```go
type AddControllerType = func(route Route, controller Controller, rcs []ControlledRoutes)
```

**Purpose:** Function type for registering new routes

**Not actively used** - kept for reference

---

## Data Structure Relationships

### Flow Diagram: HTTP Request Processing

```
Raw TCP Data (bytes)
    ↓
ParseRequest() → Request struct
    ↓
Router(req, routes) → searches ControlledRoutes slice
    ↓
Found ControlledRoutes → calls Controller function
    ↓
Controller → creates/returns Response struct
    ↓
Response.Serialize() → string (HTTP format)
    ↓
Send to client via TCP
```

### Data Structure Composition

```
Server
├── RouteControllers: []ControlledRoutes
│   ├── route: Route
│   │   ├── Url: URL (string)
│   │   └── Method: METHOD (string)
│   └── controller: Controller (function type)
│       ├── Receives: Request
│       │   ├── Url: URL (string)
│       │   ├── HttpMethod: METHOD (string)
│       │   ├── Headers: []Header
│       │   └── Version: string
│       └── Returns: Response
│           ├── Body: string
│           ├── Status: STATUS (string)
│           ├── Headers: []Header
│           └── Content: CONTENT (string)
├── AddController: function
└── Listen: function
```

---

## Memory and Performance Considerations

### String-based Types
- `URL`, `METHOD`, `CONTENT`, `STATUS` are all strings
- No validation on string values
- Searching is O(n) linear scan through routes

### Slices
- `RouteControllers` - grows dynamically as routes added
- `Headers` - stored in Request/Response but mostly unused
- No pre-allocated capacity specified

### Function Fields
- Server uses function fields instead of interfaces
- Each Server instance can have different functions assigned
- No virtual dispatch overhead

### Parsing
- Single-pass parsing of HTTP request
- Splits on newlines and spaces
- No buffering of extra request data

---

## Summary Table

| Name | Type | Location | Purpose |
|------|------|----------|---------|
| `URL` | Type Alias | http.go | Request path |
| `METHOD` | Type Alias | http.go | HTTP verb |
| `CONTENT` | Type Alias | http.go | MIME type |
| `STATUS` | Type Alias | http.go | HTTP status |
| `Header` | Struct | http.go | HTTP header |
| `Request` | Struct | http.go | Parsed HTTP request |
| `Response` | Struct | http.go | HTTP response |
| `Route` | Struct | http.go | Route pattern |
| `ControlledRoutes` | Struct | router.go | Route + handler |
| `Server` | Struct | server.go | Main server object |
| `Config` | Struct | server.go | Server config |
| `Controller` | Function Type | controller.go | Request handler |
| `ListenerType` | Function Type | networking.go | Listener function |
| `RouterType` | Function Type | router.go | Router function |
| `Port` | Type Alias | server.go | Port number |
