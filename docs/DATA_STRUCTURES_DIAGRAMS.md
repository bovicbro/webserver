# Data Structures - Visual Diagrams

This document provides visual representations of the data structures and their relationships.

---

## 1. HTTP Request/Response Flow

```
┌─────────────────────────────────────────────────────────────────┐
│                    Client TCP Connection                        │
└─────────────────────────────────────────────────────────────────┘
                             ↓
                    Raw HTTP Request Data
                    "GET /api/users HTTP/1.1\n
                     Host: localhost:8000\n\n"
                             ↓
                    ┌────────────────────┐
                    │  ParseRequest()    │
                    └────────────────────┘
                             ↓
        ┌────────────────────────────────────────────┐
        │              Request Struct                │
        ├────────────────────────────────────────────┤
        │ Url: "/api/users"          (URL)          │
        │ HttpMethod: "GET"          (METHOD)       │
        │ Version: "HTTP/1.1"        (string)       │
        │ Headers: [...]             ([]Header)     │
        └────────────────────────────────────────────┘
                             ↓
                  ┌──────────────────────┐
                  │   Router(req, rcs)   │
                  └──────────────────────┘
                             ↓
                        /    |    \
            Found    /        |        \  Not Found
                /             |             \
        ┌─────────────┐       ↓       ┌──────────────┐
        │  Execute    │       ×       │Return 404    │
        │ Controller  │               │Response      │
        └─────────────┘               └──────────────┘
                ↓
        ┌────────────────────────────────────┐
        │        Response Struct             │
        ├────────────────────────────────────┤
        │ Body: "<html>..."      (string)    │
        │ Status: "200 OK"       (STATUS)    │
        │ Content: "text/html"   (CONTENT)   │
        │ Headers: []            ([]Header)  │
        └────────────────────────────────────┘
                ↓
        ┌────────────────────────┐
        │ Response.Serialize()   │
        └────────────────────────┘
                ↓
        Raw HTTP Response Data
        "HTTP/1.1 200 OK\n
         Content-length: 16\n
         Content-Type: text/html; charset=utf-8\n\n
         <html>...</html>"
                ↓
        ┌──────────────────────────────┐
        │ Send to Client & Close Conn  │
        └──────────────────────────────┘
```

---

## 2. Server Structure Composition

```
┌──────────────────────────────────────────────────────────────┐
│                       Server Struct                          │
├──────────────────────────────────────────────────────────────┤
│                                                              │
│  port: Port (unexported, unused)                            │
│                                                              │
│  RouteControllers: []ControlledRoutes                       │
│  ┌───────────────────────────────────────────────────────┐  │
│  │ [0] ControlledRoutes                                  │  │
│  │ ├── route: Route                                      │  │
│  │ │   ├── Url: "/"                                      │  │
│  │ │   └── Method: "GET"                                 │  │
│  │ └── controller: func(req, res) Response              │  │
│  │     └── Returns: Response{...}                        │  │
│  │                                                       │  │
│  │ [1] ControlledRoutes                                  │  │
│  │ ├── route: Route                                      │  │
│  │ │   ├── Url: "/about"                                 │  │
│  │ │   └── Method: "GET"                                 │  │
│  │ └── controller: func(req, res) Response              │  │
│  │     └── Returns: Response{...}                        │  │
│  │                                                       │  │
│  │ [2] ControlledRoutes                                  │  │
│  │ ├── route: Route                                      │  │
│  │ │   ├── Url: "/api/data"                              │  │
│  │ │   └── Method: "POST"                                │  │
│  │ └── controller: func(req, res) Response              │  │
│  │     └── Returns: Response{...}                        │  │
│  └───────────────────────────────────────────────────────┘  │
│                                                              │
│  AddController: func(route, controller) { ... }            │
│                                                              │
│  Listen: networking.ListenerType                           │
│         func(port string, rcs []ControlledRoutes) { ... }  │
│                                                              │
└──────────────────────────────────────────────────────────────┘
```

---

## 3. Request Structure Details

```
┌────────────────────────────────────────────────────────────┐
│                   Request Struct                          │
├────────────────────────────────────────────────────────────┤
│                                                            │
│  ┌──────────────────────────────────────────────────────┐ │
│  │ Url: URL (string type alias)                         │ │
│  │ Value: "/api/users"                                  │ │
│  │ Purpose: The request path to match against routes    │ │
│  │ Example: "/", "/about", "/api/v1/products"          │ │
│  └──────────────────────────────────────────────────────┘ │
│                                                            │
│  ┌──────────────────────────────────────────────────────┐ │
│  │ HttpMethod: METHOD (string type alias)               │ │
│  │ Value: "GET" | "POST" | "PUT" | "DELETE" | ...     │ │
│  │ Purpose: HTTP verb for the request                   │ │
│  │ Note: Router currently ignores this field            │ │
│  └──────────────────────────────────────────────────────┘ │
│                                                            │
│  ┌──────────────────────────────────────────────────────┐ │
│  │ Version: string                                      │ │
│  │ Value: "HTTP/1.1"                                    │ │
│  │ Purpose: HTTP protocol version from request line     │ │
│  │ Note: Not currently used in response handling        │ │
│  └──────────────────────────────────────────────────────┘ │
│                                                            │
│  ┌──────────────────────────────────────────────────────┐ │
│  │ Headers: []Header (slice of Header structs)          │ │
│  │                                                       │ │
│  │ Header[0]                                             │ │
│  │ ├── key: "Host"                                       │ │
│  │ └── value: "localhost:8000"                           │ │
│  │                                                       │ │
│  │ Header[1]                                             │ │
│  │ ├── key: "User-Agent"                                │ │
│  │ └── value: "Mozilla/5.0..."                           │ │
│  │                                                       │ │
│  │ Header[2]                                             │ │
│  │ ├── key: "Accept"                                    │ │
│  │ └── value: "text/html"                                │ │
│  │                                                       │ │
│  │ Purpose: HTTP headers parsed from request            │ │
│  │ Note: Parsed but not actively used                   │ │
│  └──────────────────────────────────────────────────────┘ │
│                                                            │
└────────────────────────────────────────────────────────────┘
```

---

## 4. Response Structure Details

```
┌────────────────────────────────────────────────────────────┐
│                   Response Struct                         │
├────────────────────────────────────────────────────────────┤
│                                                            │
│  ┌──────────────────────────────────────────────────────┐ │
│  │ Body: string                                         │ │
│  │ Value: "<html><body>Hello</body></html>"            │ │
│  │ Purpose: Response body content                       │ │
│  │ Usage: Can be HTML, JSON, plain text, etc.          │ │
│  │ Max Size: Limited by buffer size (tested to 10MB)   │ │
│  └──────────────────────────────────────────────────────┘ │
│                                                            │
│  ┌──────────────────────────────────────────────────────┐ │
│  │ Status: STATUS (string type alias)                   │ │
│  │ Value: "200 OK" | "404 Not Found" | ...             │ │
│  │ Purpose: HTTP status code with reason phrase         │ │
│  │ Format: "CODE Reason"                                │ │
│  │ Examples:                                             │ │
│  │   "200 OK"                                            │ │
│  │   "201 Created"                                       │ │
│  │   "404 Not Found"                                     │ │
│  │   "500 Internal Server Error"                         │ │
│  └──────────────────────────────────────────────────────┘ │
│                                                            │
│  ┌──────────────────────────────────────────────────────┐ │
│  │ Content: CONTENT (string type alias)                 │ │
│  │ Value: "text/html" | "application/json" | ...       │ │
│  │ Purpose: MIME type for Content-Type header           │ │
│  │ Available types: 35+ MIME types defined              │ │
│  │ Examples:                                             │ │
│  │   "text/html"          (HTML)                         │ │
│  │   "application/json"   (JSON)                         │ │
│  │   "text/plain"         (Plain text)                   │ │
│  │   "text/css"           (CSS)                          │ │
│  │   "image/png"          (PNG image)                    │ │
│  └──────────────────────────────────────────────────────┘ │
│                                                            │
│  ┌──────────────────────────────────────────────────────┐ │
│  │ Headers: []Header (unexported, not serialized)       │ │
│  │                                                       │ │
│  │ Purpose: Custom HTTP response headers                │ │
│  │ Status: Field exists but not currently used          │ │
│  │ Note: Serialization doesn't include custom headers   │ │
│  └──────────────────────────────────────────────────────┘ │
│                                                            │
└────────────────────────────────────────────────────────────┘
```

---

## 5. Route Matching Process

```
┌──────────────────────────────────────┐
│  Incoming Request                    │
│  Request{Url: "/api/users", ...}     │
└──────────────────────────────────────┘
         ↓
┌──────────────────────────────────────┐
│  Server.RouteControllers[]           │
│  ┌────────────────────────────────┐  │
│  │ [0] Route: "/"                 │  │ ← Url "/api/users" ≠ "/" → SKIP
│  └────────────────────────────────┘  │
│  ┌────────────────────────────────┐  │
│  │ [1] Route: "/about"            │  │ ← Url "/api/users" ≠ "/about" → SKIP
│  └────────────────────────────────┘  │
│  ┌────────────────────────────────┐  │
│  │ [2] Route: "/api/users"        │  │ ← Url "/api/users" = "/api/users" ✓ MATCH!
│  │ Controller: func(req, res)     │  │
│  └────────────────────────────────┘  │
└──────────────────────────────────────┘
         ↓
┌──────────────────────────────────────┐
│  Execute Matched Controller          │
│  controller(request, baseResponse)   │
│  Returns: Response{...}              │
└──────────────────────────────────────┘
         ↓
    Response sent to client
```

---

## 6. Type Alias Hierarchy

```
Go Base Type                Domain Type Alias              Constants/Values
═════════════════════════════════════════════════════════════════════════

string ─────────────────────→ URL                    "/", "/about", "/api/users"
        ─────────────────────→ METHOD                GET, POST, PUT, DELETE...
        ─────────────────────→ CONTENT               text/html, application/json...
        ─────────────────────→ STATUS                200 OK, 404 Not Found...

int ────────────────────────→ Port                   8000, 3000, 5000...

func ───────────────────────→ Controller             func(req, res) Response
    ───────────────────────→ ListenerType           func(port, routes)
    ───────────────────────→ RouterType             func(req, routes) Response
```

---

## 7. Memory Layout: Request/Response in Network Flow

```
TCP Packet received:
┌──────────────────────────────────────────────────┐
│ Raw bytes: 71 bytes                              │
│ "GET /api/users HTTP/1.1\r\n                    │
│ Host: localhost:8000\r\n\r\n"                   │
└──────────────────────────────────────────────────┘
         ↓ ParseRequest()
         ↓ (strings.Split on "\n" and " ")
         ↓
┌──────────────────────────────────────────────────┐
│ Request struct: ~160 bytes (estimated)           │
│ ┌────────────────────────────────────────────┐  │
│ │ Url: "/api/users" (16 bytes string)        │  │
│ │ HttpMethod: "GET" (3 bytes string)         │  │
│ │ Version: "HTTP/1.1" (8 bytes string)       │  │
│ │ Headers: []Header (empty or with values)   │  │
│ └────────────────────────────────────────────┘  │
└──────────────────────────────────────────────────┘
         ↓ Controller processes
         ↓
┌──────────────────────────────────────────────────┐
│ Response struct: ~200 bytes (estimated)          │
│ ┌────────────────────────────────────────────┐  │
│ │ Body: "<html>..." (variable length)        │  │
│ │ Status: "200 OK" (6 bytes string)          │  │
│ │ Content: "text/html" (9 bytes string)      │  │
│ │ Headers: []Header (empty)                  │  │
│ └────────────────────────────────────────────┘  │
└──────────────────────────────────────────────────┘
         ↓ Serialize()
         ↓ (fmt.Sprintf creates response string)
         ↓
┌──────────────────────────────────────────────────┐
│ HTTP Response string:                            │
│ "HTTP/1.1 200 OK\r\n                           │
│ Content-length: 18\r\n                         │
│ Content-Type: text/html; charset=utf-8\r\n    │
│ \r\n                                            │
│ <html>...</html>"                              │
│ (~150+ bytes depending on body)                 │
└──────────────────────────────────────────────────┘
         ↓ TCP Write
         ↓
┌──────────────────────────────────────────────────┐
│ Raw bytes sent to client                         │
└──────────────────────────────────────────────────┘
```

---

## 8. Struct Field Visibility

```
Package Boundary: http package
╔════════════════════════════════════════════════════════════════╗
║                                                                ║
║  type Request struct {                                         ║
║      Url        URL       ← EXPORTED (visible outside package) ║
║      HttpMethod METHOD    ← EXPORTED                           ║
║      Headers    []Header  ← EXPORTED                           ║
║      Version    string    ← EXPORTED                           ║
║  }                                                             ║
║                                                                ║
║  type Header struct {                                          ║
║      key   string  ← UNEXPORTED (lowercase)                   ║
║      value string  ← UNEXPORTED                               ║
║  }                                                             ║
║                                                                ║
╚════════════════════════════════════════════════════════════════╝

Package Boundary: router package
╔════════════════════════════════════════════════════════════════╗
║                                                                ║
║  type ControlledRoutes struct {                                ║
║      route      Route        ← UNEXPORTED (lowercase)         ║
║      controller Controller   ← UNEXPORTED                      ║
║  }                                                             ║
║                                                                ║
║  Cannot be accessed directly outside router package:           ║
║  ✗ cr.route.Url                                               ║
║  ✗ cr.controller                                              ║
║                                                                ║
╚════════════════════════════════════════════════════════════════╝

Package Boundary: server package
╔════════════════════════════════════════════════════════════════╗
║                                                                ║
║  type Server struct {                                          ║
║      port             Port                ← UNEXPORTED         ║
║      RouteControllers []ControlledRoutes  ← EXPORTED           ║
║      AddController    func(...)           ← EXPORTED           ║
║      Listen           ListenerType        ← EXPORTED           ║
║  }                                                             ║
║                                                                ║
║  Can access: server.RouteControllers, server.AddController    ║
║  Cannot access: server.port                                   ║
║                                                                ║
╚════════════════════════════════════════════════════════════════╝
```

---

## 9. Type Constants Organization

### STATUS Constants (29 total)
```
1xx Informational (3 codes)
├── CONTINUE
├── SWITCHING_PROTOCOLS
└── PROCESSING

2xx Success (7 codes)
├── OK
├── CREATED
├── ACCEPTED
├── NON_AUTHORITATIVE
├── NO_CONTENT
├── RESET_CONTENT
└── PARTIAL_CONTENT

3xx Redirection (7 codes)
├── MULTIPLE_CHOICES
├── MOVED_PERMANENTLY
├── FOUND
├── SEE_OTHER
├── NOT_MODIFIED
├── TEMP_REDIRECT
└── PERM_REDIRECT

4xx Client Errors (17 codes)
├── BAD_REQUEST
├── UNAUTHORIZED
├── PAYMENT_REQUIRED
├── FORBIDDEN
├── NOT_FOUND
├── METHOD_NOT_ALLOWED
├── NOT_ACCEPTABLE
├── PROXY_AUTH_REQUIRED
├── REQUEST_TIMEOUT
├── CONFLICT
├── GONE
├── LENGTH_REQUIRED
├── PRECONDITION_FAILED
├── PAYLOAD_TOO_LARGE
├── URI_TOO_LONG
├── UNSUPPORTED_MEDIA
└── IM_A_TEAPOT

5xx Server Errors (6 codes)
├── INTERNAL_SERVER_ERROR
├── NOT_IMPLEMENTED
├── BAD_GATEWAY
├── SERVICE_UNAVAILABLE
├── GATEWAY_TIMEOUT
└── HTTP_VERSION_NOT_SUPPORTED
```

### CONTENT Constants (35 total)
```
Text types (6)
├── HTML
├── PLAIN
├── CSS
├── CSV
├── XML_TEXT
└── MARKDOWN

Application types (10)
├── JSON
├── JAVASCRIPT
├── XML_APP
├── FORM_URLENCODED
├── PDF
├── ZIP
├── GZIP
├── MSWORD
├── MS_EXCEL
└── MS_POWERPOINT

Image types (6)
├── PNG
├── JPEG
├── GIF
├── SVG
├── WEBP
└── ICO

Audio/Video types (4)
├── MP3
├── OGG_AUDIO
├── MP4
└── WEBM

Font types (4)
├── WOFF
├── WOFF2
├── TTF
└── OTF
```

### METHOD Constants (9 total)
```
GET       ← Most common
POST      ← Form submissions
PUT       ← Resource updates
DELETE    ← Resource deletion
PATCH     ← Partial updates
HEAD      ← Like GET but no body
OPTIONS   ← CORS preflight
TRACE     ← Request tracing
CONNECT   ← Proxy tunneling
```

---

## 10. Data Flow in main.go

```
main()
  │
  ├─ InitServer()
  │  └─ Creates empty Server{RouteControllers: []}
  │
  ├─ server.AddController("/", GET handler)
  │  └─ router.AddController() appends to RouteControllers
  │
  ├─ server.AddController("/about", GET handler)
  │  └─ router.AddController() appends to RouteControllers
  │
  ├─ server.AddController("/styles.css", GET handler)
  │  └─ router.AddController() appends to RouteControllers
  │
  └─ server.Listen("8000", server.RouteControllers)
     └─ networking.Listen()
        └─ Infinite loop calling initListener()
           └─ Infinite loop accepting connections
              └─ For each connection:
                 ├─ Read raw bytes
                 ├─ ParseRequest() → Request
                 ├─ Router() → Response
                 ├─ Serialize() → string
                 ├─ Write to connection
                 └─ Close connection
```

---

## Summary

The project uses a **minimal, focused set of data structures**:

- **4 string type aliases** (URL, METHOD, CONTENT, STATUS) for semantic clarity
- **3 structs** for HTTP (Header, Request, Response)
- **3 structs** for routing (Route, ControlledRoutes, Server)
- **4 function type aliases** for composition (Controller, ListenerType, RouterType, AddControllerType)
- **35 MIME type constants** and **29 HTTP status constants**

All structures are **immutable** after creation (no methods to modify), following a **functional programming style** with function fields for composition rather than methods or interfaces.
