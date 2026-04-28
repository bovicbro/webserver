# Changelog

This document describes the changes made to the webserver project in the current development cycle.

---

## 1. HTTP Protocol Layer (`server/http/http.go`)

### Request Parsing
- **Headers**: `ParseRequest` now parses HTTP headers into `req.Headers` as `[]Header`.
- **Query Parameters**: URLs with query strings (e.g. `/search?q=test`) are now split. The path is stored in `req.Url` and parameters are stored in `req.Query` (`map[string]string`).
- **Request Body**: The body after the header section is now captured in `req.Body`.
- **Exported `Header` fields**: The `Header` struct fields were renamed from `key`/`value` (unexported) to `Key`/`Value` (exported) so controllers and tests can read and write them.
- Added `parseQueryString` helper to populate the query map.

### Request Struct
- Added `Query map[string]string`
- Added `Body string`
- Added `PathParams map[string]string`

### Response Serialization
- `Response.Serialize()` now **defaults `Content-Type` to `text/html`** if the controller does not set one. Previously it produced `Content-Type: ; charset=utf-8`.
- Custom headers stored in `res.Headers` are now actually included in the serialized output. Previously the field was defined but ignored.
- Fixed trailing spaces and removed an unwanted trailing newline pattern in the serialized response string.

---

## 2. Router (`server/router/router.go`)

### New Matching Logic
The router was rewritten to support more than exact string equality:

- **Exact match**: `/about` still matches `/about`.
- **Wildcard routes**: `/*` matches any path. This is used for the static file catch-all.
- **Path parameters**: Routes like `/api/users/:id` now extract segments into `req.PathParams`.
  - Example: `GET /api/users/42` sets `req.PathParams["id"] = "42"`.

### Route Priority
Routes are checked in registration order. Exact and parameterized routes should be registered **before** wildcard routes so they take priority.

### Implementation Detail
- Replaced the `SliceIndexOf` lookup with a simple `for` loop over all registered routes.
- Added `matchRoute(pattern, target)` helper that handles exact, wildcard (`/*`), and colon-prefixed (`:`) parameter segments.

---

## 3. Networking Layer (`server/networking/networking.go`)

### Listener Lifecycle
- **Fixed resource leak**: Removed the nested `for { initListener(...) }` pattern. `Listen` now creates a single `net.Listener`, defers its close, and loops on `Accept()` directly.
- **Graceful accept errors**: If `Accept()` fails, the server logs the error with `log.Printf` and **continues** instead of calling `log.Fatal` and crashing.

### Connection Handling
- **Buffer size**: Increased request read buffer from `1024` bytes to `4096` bytes.
- **Buffer safety**: The parser now receives only the bytes actually read (`buffer[:n]`) rather than the entire zero-padded buffer.
- **Connection cleanup**: `conn.Close()` is deferred in `handleRequest` so it always runs, even if parsing or routing panics.
- **Removed extra newline**: The response writer no longer appends an artificial `\n` after the serialized HTTP response. This was corrupting non-text responses (e.g. images).

---

## 4. Static File Helper (`server/controller/controller.go`)

### `Static(basePath string) Controller`
A new helper function was added so `main.go` no longer needs a dedicated controller for every static file.

**Behavior:**
- Maps the request URL directly to a file inside `basePath`.
- If the URL has no file extension (e.g. `/about`), it tries appending `.html` automatically.
- **Content-Type auto-detection**: `detectContentType` infers the correct MIME type from the file extension (`.css`, `.js`, `.png`, `.json`, `.html`, etc.).
- **Basic directory traversal guard**: Rejects URLs containing `..` with `403 Forbidden`.

### Note on Security
The `..` check is a first-line defense but is not exhaustive. URL-encoded traversal sequences could still bypass it. This is noted as a known limitation, not a production-grade sandbox.

---

## 5. Application Entry Point (`main.go`)

### Before
Four repetitive controllers manually read individual files from `./static`:
```go
AddController(http.Route{Url: "/"},       ... ReadFile("./static/index.html") ...)
AddController(http.Route{Url: "/about"},  ... ReadFile("./static/about.html") ...)
// etc.
```

### After
- A single wildcard route handles all static assets:
  ```go
  AddController(http.Route{Url: "/*", Method: http.GET}, controller.Static("./static"))
  ```
- **New demo API routes** were added to exercise the new features:
  - `GET /api/search?q=hello` → returns JSON containing the query value.
  - `GET /api/users/:id` → returns JSON with the extracted `user_id` path parameter.
  - `POST /api/echo` → returns JSON echoing back the request body.

---

## 6. Test Coverage

Tests were added or updated across all packages to verify the new behavior:

- `server/http/http_test.go`
  - Query parameter extraction
  - Header parsing
  - Body extraction
  - Custom header serialization
  - Default content-type fallback

- `server/router/router_test.go`
  - Wildcard route matching (`/*`)
  - Path parameter extraction (`:id`)
  - Route priority (exact routes over wildcards)

- `server/controller/controller_test.go`
  - Content type detection by extension
  - `Static` serving existing files
  - `Static` returning `404` for missing files
  - `Static` rejecting `..` traversal attempts

---

## 7. Tooling Verification

All changes were verified with:
- `go test ./...` — all tests pass
- `gofmt -w .` — all files formatted
- `go vet ./...` — no static analysis issues
- `go build` — binary compiles successfully

---

## Known Remaining Limitations (Not Addressed)

These are out of scope for this iteration but worth noting for future work:

- **No request body size limit** — large POST bodies can still exhaust memory.
- **Buffer truncation** — 4 KB is better than 1 KB, but requests larger than the buffer are still truncated.
- **No connection timeouts / rate limits** — the server will accept unlimited concurrent connections.
- **No TLS / HTTPS** — all traffic is plaintext.
- **No HTTP Keep-Alive** — every response closes the connection.
- **No chunked transfer encoding** support.
- **Path traversal filter is basic** — URL-encoded sequences are not normalized.
