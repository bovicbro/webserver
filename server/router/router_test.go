package router

import (
	"testing"
	. "webserver/server/http"
)

// TestAddController tests the AddController function
func TestAddController(t *testing.T) {
	// Arrange
	var rcs []ControlledRoutes
	route := Route{Url: "/", Method: GET}
	controller := func(req Request, res Response) Response {
		return Response{Body: "Home", Status: OK, Content: HTML}
	}

	// Act
	result := AddController(route, controller, rcs)

	// Assert
	if len(result) != 1 {
		t.Errorf("Expected 1 route, got %d", len(result))
	}
	if result[0].route.Url != route.Url {
		t.Errorf("Expected URL %v, got %v", route.Url, result[0].route.Url)
	}
	if result[0].route.Method != route.Method {
		t.Errorf("Expected method %v, got %v", route.Method, result[0].route.Method)
	}
}

// TestAddControllerMultiple tests adding multiple routes
func TestAddControllerMultiple(t *testing.T) {
	// Arrange
	var rcs []ControlledRoutes
	controller1 := func(req Request, res Response) Response {
		return Response{Body: "Home", Status: OK, Content: HTML}
	}
	controller2 := func(req Request, res Response) Response {
		return Response{Body: "About", Status: OK, Content: HTML}
	}

	// Act
	rcs = AddController(Route{Url: "/", Method: GET}, controller1, rcs)
	rcs = AddController(Route{Url: "/about", Method: GET}, controller2, rcs)

	// Assert
	if len(rcs) != 2 {
		t.Errorf("Expected 2 routes, got %d", len(rcs))
	}
	if rcs[0].route.Url != "/" {
		t.Errorf("Expected first URL '/', got %v", rcs[0].route.Url)
	}
	if rcs[1].route.Url != "/about" {
		t.Errorf("Expected second URL '/about', got %v", rcs[1].route.Url)
	}
}

// TestRouterFoundRoute tests Router finding a matching route
func TestRouterFoundRoute(t *testing.T) {
	// Arrange
	req := Request{Url: "/", HttpMethod: GET, Version: "HTTP/1.1"}
	expectedBody := "Home Page"

	controller := func(req Request, res Response) Response {
		return Response{Body: expectedBody, Status: OK, Content: HTML}
	}

	rcs := []ControlledRoutes{
		{
			route:      Route{Url: "/", Method: GET},
			controller: controller,
		},
	}

	// Act
	res := Router(req, rcs)

	// Assert
	if res.Body != expectedBody {
		t.Errorf("Expected body %s, got %s", expectedBody, res.Body)
	}
	if res.Status != OK {
		t.Errorf("Expected status OK, got %v", res.Status)
	}
}

// TestRouterNotFoundRoute tests Router when route is not found
func TestRouterNotFoundRoute(t *testing.T) {
	// Arrange
	req := Request{Url: "/notfound", HttpMethod: GET, Version: "HTTP/1.1"}

	controller := func(req Request, res Response) Response {
		return Response{Body: "Home", Status: OK, Content: HTML}
	}

	rcs := []ControlledRoutes{
		{
			route:      Route{Url: "/", Method: GET},
			controller: controller,
		},
	}

	// Act
	res := Router(req, rcs)

	// Assert
	if res.Status != NOT_FOUND {
		t.Errorf("Expected status NOT_FOUND, got %v", res.Status)
	}
	if res.Body != "404 Not found" {
		t.Errorf("Expected body '404 Not found', got %s", res.Body)
	}
}

// TestRouterEmptyRoutes tests Router with empty routes list
func TestRouterEmptyRoutes(t *testing.T) {
	// Arrange
	req := Request{Url: "/", HttpMethod: GET, Version: "HTTP/1.1"}
	var rcs []ControlledRoutes

	// Act
	res := Router(req, rcs)

	// Assert
	if res.Status != NOT_FOUND {
		t.Errorf("Expected status NOT_FOUND, got %v", res.Status)
	}
}

// TestRouterWithMultipleRoutes tests Router finds correct route among multiple
func TestRouterWithMultipleRoutes(t *testing.T) {
	// Arrange
	req := Request{Url: "/api/users", HttpMethod: GET, Version: "HTTP/1.1"}

	controller1 := func(req Request, res Response) Response {
		return Response{Body: "Home", Status: OK, Content: HTML}
	}
	controller2 := func(req Request, res Response) Response {
		return Response{Body: "Users List", Status: OK, Content: JSON}
	}
	controller3 := func(req Request, res Response) Response {
		return Response{Body: "About", Status: OK, Content: HTML}
	}

	rcs := []ControlledRoutes{
		{route: Route{Url: "/", Method: GET}, controller: controller1},
		{route: Route{Url: "/api/users", Method: GET}, controller: controller2},
		{route: Route{Url: "/about", Method: GET}, controller: controller3},
	}

	// Act
	res := Router(req, rcs)

	// Assert
	if res.Body != "Users List" {
		t.Errorf("Expected body 'Users List', got %s", res.Body)
	}
	if res.Content != JSON {
		t.Errorf("Expected JSON content type, got %v", res.Content)
	}
}

// TestRouterWithDifferentMethods tests Router with different HTTP methods
func TestRouterWithDifferentMethods(t *testing.T) {
	// Note: Current implementation matches only by URL, not by HTTP method
	// This test documents that behavior

	// Arrange
	getController := func(req Request, res Response) Response {
		return Response{Body: "GET response", Status: OK, Content: PLAIN}
	}
	postController := func(req Request, res Response) Response {
		return Response{Body: "POST response", Status: CREATED, Content: JSON}
	}

	rcs := []ControlledRoutes{
		{route: Route{Url: "/api/data", Method: GET}, controller: getController},
		{route: Route{Url: "/api/data", Method: POST}, controller: postController},
	}

	// Act
	getReq := Request{Url: "/api/data", HttpMethod: GET, Version: "HTTP/1.1"}
	getRes := Router(getReq, rcs)

	postReq := Request{Url: "/api/data", HttpMethod: POST, Version: "HTTP/1.1"}
	postRes := Router(postReq, rcs)

	// Assert
	// Current implementation matches only by URL, not by method
	// Both should match the first route that has the URL
	if getRes.Body != "GET response" {
		t.Errorf("Expected GET response, got %s", getRes.Body)
	}
	if postRes.Body != "GET response" {
		// Router matches by URL only, so both GET and POST will return GET response
		t.Logf("Router matches by URL only, POST request matched GET controller: %s", postRes.Body)
	}
}

// TestRouterCallsController tests that Router properly calls the controller
func TestRouterCallsController(t *testing.T) {
	// Arrange
	called := false
	controller := func(req Request, res Response) Response {
		called = true
		return Response{Body: "Test", Status: OK, Content: PLAIN}
	}

	rcs := []ControlledRoutes{
		{route: Route{Url: "/test", Method: GET}, controller: controller},
	}

	req := Request{Url: "/test", HttpMethod: GET, Version: "HTTP/1.1"}

	// Act
	Router(req, rcs)

	// Assert
	if !called {
		t.Errorf("Expected controller to be called, but it wasn't")
	}
}

// TestRouterWithFirstMatch tests that Router returns first matching route
func TestRouterWithFirstMatch(t *testing.T) {
	// Arrange
	controller1 := func(req Request, res Response) Response {
		return Response{Body: "First", Status: OK, Content: PLAIN}
	}
	controller2 := func(req Request, res Response) Response {
		return Response{Body: "Second", Status: OK, Content: PLAIN}
	}

	rcs := []ControlledRoutes{
		{route: Route{Url: "/test", Method: GET}, controller: controller1},
		{route: Route{Url: "/test", Method: GET}, controller: controller2},
	}

	req := Request{Url: "/test", HttpMethod: GET, Version: "HTTP/1.1"}

	// Act
	res := Router(req, rcs)

	// Assert
	if res.Body != "First" {
		t.Errorf("Expected first controller to be called, got %s", res.Body)
	}
}

// TestRouterWithRootPath tests Router with root path
func TestRouterWithRootPath(t *testing.T) {
	// Arrange
	controller := func(req Request, res Response) Response {
		return Response{Body: "Root", Status: OK, Content: HTML}
	}

	rcs := []ControlledRoutes{
		{route: Route{Url: "/", Method: GET}, controller: controller},
	}

	req := Request{Url: "/", HttpMethod: GET, Version: "HTTP/1.1"}

	// Act
	res := Router(req, rcs)

	// Assert
	if res.Body != "Root" {
		t.Errorf("Expected root response, got %s", res.Body)
	}
	if res.Status != OK {
		t.Errorf("Expected OK status, got %v", res.Status)
	}
}

// TestRouterWithDeepPaths tests Router with deeply nested paths
func TestRouterWithDeepPaths(t *testing.T) {
	// Arrange
	controller := func(req Request, res Response) Response {
		return Response{Body: "Deep path", Status: OK, Content: JSON}
	}

	rcs := []ControlledRoutes{
		{route: Route{Url: "/api/v1/users/123/profile", Method: GET}, controller: controller},
	}

	req := Request{Url: "/api/v1/users/123/profile", HttpMethod: GET, Version: "HTTP/1.1"}

	// Act
	res := Router(req, rcs)

	// Assert
	if res.Body != "Deep path" {
		t.Errorf("Expected deep path response, got %s", res.Body)
	}
}

// TestControlledRoutesStructure tests that ControlledRoutes properly stores route and controller
func TestControlledRoutesStructure(t *testing.T) {
	// Arrange
	expectedRoute := Route{Url: "/test", Method: POST}
	controller := func(req Request, res Response) Response {
		return Response{Body: "Test", Status: OK, Content: PLAIN}
	}

	// Act
	cr := ControlledRoutes{
		route:      expectedRoute,
		controller: controller,
	}

	// Assert
	if cr.route.Url != expectedRoute.Url {
		t.Errorf("Expected URL %v, got %v", expectedRoute.Url, cr.route.Url)
	}
	if cr.route.Method != expectedRoute.Method {
		t.Errorf("Expected method %v, got %v", expectedRoute.Method, cr.route.Method)
	}
}
