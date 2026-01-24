# Data Structures - Usage Examples

This document provides practical code examples showing how to use each data structure.

---

## Type Aliases - Basic Usage

### URL

```go
package main

import "webserver/server/http"

// Creating a URL
var homePage http.URL = "/"
var aboutPage http.URL = "/about"
var apiEndpoint http.URL = "/api/v1/users"

// Using in Route
route := http.Route{
    Url:    "/api/users",
    Method: http.GET,
}

// Matching URLs
request := http.Request{Url: "/api/users"}
if request.Url == route.Url {
    println("Route matched!")
}

// Converting to string
urlStr := string(request.Url)  // Get as regular string
println(urlStr)  // Output: /api/users
```

---

### METHOD

```go
import "webserver/server/http"

// Using predefined constants
getMethod := http.GET      // "GET"
postMethod := http.POST    // "POST"
putMethod := http.PUT      // "PUT"

// Create custom method
customMethod := http.METHOD("CUSTOM")

// In request
request := http.Request{
    HttpMethod: http.GET,
}

// Checking method
if request.HttpMethod == http.POST {
    println("This is a POST request")
}

// All available constants
methods := []http.METHOD{
    http.GET,
    http.POST,
    http.PUT,
    http.DELETE,
    http.PATCH,
    http.HEAD,
    http.OPTIONS,
    http.TRACE,
    http.CONNECT,
}

for _, method := range methods {
    println(string(method))
}
```

---

### CONTENT (MIME Types)

```go
import "webserver/server/http"

// Text types
htmlContent := http.HTML                    // "text/html"
plainContent := http.PLAIN                  // "text/plain"
cssContent := http.CSS                      // "text/css"
markdownContent := http.MARKDOWN            // "text/markdown"

// Application types
jsonContent := http.JSON                    // "application/json"
xmlContent := http.XML_APP                  // "application/xml"
formContent := http.FORM_URLENCODED         // "application/x-www-form-urlencoded"
pdfContent := http.PDF                      // "application/pdf"

// Image types
pngContent := http.PNG                      // "image/png"
jpegContent := http.JPEG                    // "image/jpeg"
svgContent := http.SVG                      // "image/svg+xml"

// Font types
woffContent := http.WOFF                    // "font/woff"
woff2Content := http.WOFF2                  // "font/woff2"

// Audio/Video
mp3Content := http.MP3                      // "audio/mpeg"
mp4Content := http.MP4                      // "video/mp4"

// Using in response
response := http.Response{
    Body:    "<html>Hello</html>",
    Status:  http.OK,
    Content: htmlContent,
}

// Get MIME type as string
mimeType := string(response.Content)
println(mimeType)  // Output: text/html
```

---

### STATUS

```go
import "webserver/server/http"

// 2xx Success codes
okResponse := http.Response{Status: http.OK}                    // "200 OK"
createdResponse := http.Response{Status: http.CREATED}          // "201 Created"
noContentResponse := http.Response{Status: http.NO_CONTENT}     // "204 No Content"

// 3xx Redirection codes
movedResponse := http.Response{Status: http.MOVED_PERMANENTLY}  // "301 Moved Permanently"
foundResponse := http.Response{Status: http.FOUND}              // "302 Found"

// 4xx Client Error codes
badRequestResponse := http.Response{Status: http.BAD_REQUEST}   // "400 Bad Request"
unauthorizedResponse := http.Response{Status: http.UNAUTHORIZED} // "401 Unauthorized"
forbiddenResponse := http.Response{Status: http.FORBIDDEN}      // "403 Forbidden"
notFoundResponse := http.Response{Status: http.NOT_FOUND}       // "404 Not Found"
methodNotAllowedResponse := http.Response{Status: http.METHOD_NOT_ALLOWED} // "405 Method Not Allowed"

// 5xx Server Error codes
internalErrorResponse := http.Response{Status: http.INTERNAL_SERVER_ERROR} // "500 Internal Server Error"
notImplementedResponse := http.Response{Status: http.NOT_IMPLEMENTED}     // "501 Not Implemented"

// Get status as string
statusLine := string(okResponse.Status)
println(statusLine)  // Output: 200 OK
```

---

## Structs - Creation and Usage

### Header Struct

```go
import "webserver/server/http"

// Create individual headers
contentLengthHeader := http.Header{
    key:   "Content-Length",
    value: "1024",
}

hostHeader := http.Header{
    key:   "Host",
    value: "localhost:8000",
}

userAgentHeader := http.Header{
    key:   "User-Agent",
    value: "Mozilla/5.0",
}

// Create slice of headers
headers := []http.Header{
    contentLengthHeader,
    hostHeader,
    userAgentHeader,
}

// Note: Header fields (key, value) are unexported
// You cannot access them outside the http package
// Cannot do: println(header.key)

// They're used in Request and Response structs
request := http.Request{
    Headers: headers,
}
```

---

### Request Struct

```go
import "webserver/server/http"

// Create a minimal request
minimalRequest := http.Request{
    Url:        "/api/users",
    HttpMethod: http.GET,
    Version:    "HTTP/1.1",
    Headers:    []http.Header{},
}

// Create a more complete request
completeRequest := http.Request{
    Url:        "/api/users",
    HttpMethod: http.GET,
    Version:    "HTTP/1.1",
    Headers: []http.Header{
        {key: "Host", value: "example.com"},
        {key: "Accept", value: "application/json"},
    },
}

// Accessing request fields
println(string(minimalRequest.Url))         // Output: /api/users
println(string(minimalRequest.HttpMethod))  // Output: GET
println(minimalRequest.Version)             // Output: HTTP/1.1

// In real code, requests come from ParseRequest()
rawRequest := "GET /api/users HTTP/1.1\nHost: example.com\n"
parsedRequest, err := http.ParseRequest(rawRequest)
if err == nil {
    println(string(parsedRequest.Url))  // Output: /api/users
}
```

---

### Response Struct

```go
import "webserver/server/http"

// Create a simple HTML response
htmlResponse := http.Response{
    Body:    "<html><body>Hello World</body></html>",
    Status:  http.OK,
    Content: http.HTML,
    Headers: []http.Header{},
}

// Create a JSON response
jsonResponse := http.Response{
    Body:    `{"name": "John", "age": 30}`,
    Status:  http.OK,
    Content: http.JSON,
    Headers: []http.Header{},
}

// Create an error response
errorResponse := http.Response{
    Body:    "Page not found",
    Status:  http.NOT_FOUND,
    Content: http.PLAIN,
    Headers: []http.Header{},
}

// Create empty response (base response)
emptyResponse := http.Response{
    Body:    "",
    Status:  "",
    Content: "",
    Headers: []http.Header{},
}

// Serialize response to HTTP format
serialized := htmlResponse.Serialize()
println(serialized)
/* Output:
HTTP/1.1 200 OK
Content-length: 42
Content-Type: text/html; charset=utf-8

<html><body>Hello World</body></html>
*/

// Accessing response fields
println(string(jsonResponse.Status))   // Output: 200 OK
println(string(jsonResponse.Content))  // Output: application/json
println(len(jsonResponse.Body))        // Output: 31 (character count)
```

---

### Route Struct

```go
import "webserver/server/http"

// Create routes
homeRoute := http.Route{
    Url:    "/",
    Method: http.GET,
}

aboutRoute := http.Route{
    Url:    "/about",
    Method: http.GET,
}

apiRoute := http.Route{
    Url:    "/api/users",
    Method: http.GET,
}

apiCreateRoute := http.Route{
    Url:    "/api/users",
    Method: http.POST,
}

// Routes are typically stored in ControlledRoutes
// (covered in next section)

// Accessing route fields
println(string(homeRoute.Url))     // Output: /
println(string(homeRoute.Method))  // Output: GET
```

---

### ControlledRoutes Struct

```go
import (
    "webserver/server/http"
    "webserver/server/router"
)

// Create a controlled route (controller + route pair)
homeController := func(req http.Request, res http.Response) http.Response {
    return http.Response{
        Body:    "<html>Home</html>",
        Status:  http.OK,
        Content: http.HTML,
    }
}

controlledRoute := router.ControlledRoutes{
    route: http.Route{
        Url:    "/",
        Method: http.GET,
    },
    controller: homeController,
}

// Note: Fields are unexported (lowercase)
// Cannot access: controlledRoute.route
// Cannot access: controlledRoute.controller

// ControlledRoutes are used in slices
var routes []router.ControlledRoutes
routes = append(routes, controlledRoute)

// Typically created via router.AddController()
routes = router.AddController(
    http.Route{Url: "/about", Method: http.GET},
    func(req http.Request, res http.Response) http.Response {
        return http.Response{
            Body:    "About page",
            Status:  http.OK,
            Content: http.HTML,
        }
    },
    routes,
)
```

---

### Server Struct

```go
import (
    "webserver/server"
    "webserver/server/http"
)

// Initialize server
srv := server.InitServer(server.Config{})

// Add routes via AddController function field
srv.AddController(
    http.Route{Url: "/", Method: http.GET},
    func(req http.Request, res http.Response) http.Response {
        return http.Response{
            Body:    "<html>Home</html>",
            Status:  http.OK,
            Content: http.HTML,
        }
    },
)

srv.AddController(
    http.Route{Url: "/about", Method: http.GET},
    func(req http.Request, res http.Response) http.Response {
        return http.Response{
            Body:    "<html>About</html>",
            Status:  http.OK,
            Content: http.HTML,
        }
    },
)

// Access registered routes
routeCount := len(srv.RouteControllers)  // How many routes registered
println(routeCount)  // Output: 2

// Start listening
srv.Listen("8000", srv.RouteControllers)
// Note: This will block forever (infinite loop)
```

---

### Config Struct

```go
import "webserver/server"

// Create config (currently empty, but designed for expansion)
config := server.Config{}

// Pass to InitServer
srv := server.InitServer(config)

// In future, could use like:
// config := server.Config{
//     Port:    8000,
//     Host:    "0.0.0.0",
//     TLS:     true,
//     CertFile: "/path/to/cert.pem",
//     Timeout: 30,
// }
```

---

## Function Type Aliases - Usage

### Controller Function Type

```go
import (
    "webserver/server/controller"
    "webserver/server/http"
)

// Define a simple controller
var helloController controller.Controller = func(req http.Request, res http.Response) http.Response {
    return http.Response{
        Body:    "Hello, World!",
        Status:  http.OK,
        Content: http.PLAIN,
    }
}

// Call the controller
request := http.Request{
    Url:        "/",
    HttpMethod: http.GET,
}
baseResponse := http.Response{}
response := helloController(request, baseResponse)

println(response.Body)    // Output: Hello, World!
println(string(response.Status))  // Output: 200 OK

// More complex controller with logic
var conditionalController controller.Controller = func(req http.Request, res http.Response) http.Response {
    if req.HttpMethod == http.GET {
        return http.Response{
            Body:    "GET request received",
            Status:  http.OK,
            Content: http.PLAIN,
        }
    } else if req.HttpMethod == http.POST {
        return http.Response{
            Body:    "POST request received",
            Status:  http.CREATED,
            Content: http.JSON,
        }
    }
    return http.Response{
        Body:    "Method not allowed",
        Status:  http.METHOD_NOT_ALLOWED,
        Content: http.PLAIN,
    }
}

// Controller that reads file
var fileController controller.Controller = func(req http.Request, res http.Response) http.Response {
    content, err := os.ReadFile("./static/index.html")
    if err != nil {
        return http.Response{
            Body:    "File not found",
            Status:  http.NOT_FOUND,
            Content: http.PLAIN,
        }
    }
    return http.Response{
        Body:    string(content),
        Status:  http.OK,
        Content: http.HTML,
    }
}
```

---

### ListenerType Function Type

```go
import (
    "webserver/server/networking"
    "webserver/server/router"
)

// The actual networking.Listen function implements ListenerType
var listener networking.ListenerType = networking.Listen

// You could implement your own:
var customListener networking.ListenerType = func(port string, rcs []router.ControlledRoutes) {
    println("Starting server on port " + port)
    println("Registered " + string(rune(len(rcs))) + " routes")
    // Custom listener implementation
}

// Call the listener
listener("8000", []router.ControlledRoutes{})
```

---

### RouterType Function Type

```go
import (
    "webserver/server/http"
    "webserver/server/router"
)

// The actual router.Router function implements RouterType
var routerFunc router.RouterType = router.Router

// Use the router
request := http.Request{
    Url:        "/api/users",
    HttpMethod: http.GET,
}

response := routerFunc(request, []router.ControlledRoutes{})
println(string(response.Status))  // Output: 404 Not Found (no routes registered)
```

---

## Complete Real-World Example

```go
package main

import (
    "os"
    "webserver/server"
    "webserver/server/http"
)

func main() {
    // Initialize server
    srv := server.InitServer(server.Config{})

    // Home page route
    srv.AddController(
        http.Route{Url: "/", Method: http.GET},
        func(req http.Request, res http.Response) http.Response {
            content, err := os.ReadFile("./static/index.html")
            if err != nil {
                return http.Response{Status: http.NOT_FOUND}
            }
            return http.Response{
                Body:    string(content),
                Status:  http.OK,
                Content: http.HTML,
            }
        })

    // About page route
    srv.AddController(
        http.Route{Url: "/about", Method: http.GET},
        func(req http.Request, res http.Response) http.Response {
            return http.Response{
                Body:    "About page content",
                Status:  http.OK,
                Content: http.HTML,
            }
        })

    // API route returning JSON
    srv.AddController(
        http.Route{Url: "/api/data", Method: http.GET},
        func(req http.Request, res http.Response) http.Response {
            return http.Response{
                Body:    `{"status":"ok","data":[1,2,3]}`,
                Status:  http.OK,
                Content: http.JSON,
            }
        })

    // Start listening on port 8000
    srv.Listen("8000", srv.RouteControllers)
}
```

---

## Testing Data Structures

```go
package main

import (
    "testing"
    "webserver/server/http"
)

func TestRequestCreation(t *testing.T) {
    req := http.Request{
        Url:        "/api/users",
        HttpMethod: http.GET,
        Version:    "HTTP/1.1",
        Headers:    []http.Header{},
    }

    if string(req.Url) != "/api/users" {
        t.Errorf("Expected URL /api/users, got %v", req.Url)
    }
}

func TestResponseSerialization(t *testing.T) {
    res := http.Response{
        Body:    "Hello",
        Status:  http.OK,
        Content: http.PLAIN,
    }

    serialized := res.Serialize()
    if !strings.Contains(serialized, "200 OK") {
        t.Errorf("Expected OK status in response")
    }
}

func TestParseRequest(t *testing.T) {
    rawReq := "GET /test HTTP/1.1\n"
    req, err := http.ParseRequest(rawReq)

    if err != nil {
        t.Errorf("Expected no error, got %v", err)
    }
    if string(req.Url) != "/test" {
        t.Errorf("Expected URL /test, got %v", req.Url)
    }
}
```

---

## Summary of Usage Patterns

1. **Type Aliases** - Use constants for HTTP concepts (GET, POST, OK, NOT_FOUND, JSON, HTML)
2. **Structs** - Create Request/Response for network operations
3. **Routes** - Register with AddController using Route + controller function
4. **Controllers** - Return Response with appropriate Body, Status, Content
5. **Server** - Initialize, register routes, start listening

The design emphasizes **functional composition** over object-oriented patterns.
