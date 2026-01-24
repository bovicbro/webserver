package controller

import (
	"testing"
	. "webserver/server/http"
)

// TestControllerTypeSignature tests that Controller type has correct signature
func TestControllerTypeSignature(t *testing.T) {
	// Arrange
	controller := func(req Request, res Response) Response {
		return Response{Body: "Test", Status: OK, Content: HTML}
	}

	// Act
	req := Request{Url: "/test", HttpMethod: GET, Version: "HTTP/1.1"}
	res := Response{}
	result := controller(req, res)

	// Assert
	if result.Body != "Test" {
		t.Errorf("Expected body 'Test', got %s", result.Body)
	}
	if result.Status != OK {
		t.Errorf("Expected status OK, got %v", result.Status)
	}
}

// TestControllerWithDifferentResponses tests controller returning different responses
func TestControllerWithDifferentResponses(t *testing.T) {
	tests := []struct {
		name       string
		controller Controller
		expected   string
	}{
		{
			name: "Home controller",
			controller: func(req Request, res Response) Response {
				return Response{Body: "Home Page", Status: OK, Content: HTML}
			},
			expected: "Home Page",
		},
		{
			name: "JSON controller",
			controller: func(req Request, res Response) Response {
				return Response{Body: `{"status":"ok"}`, Status: OK, Content: JSON}
			},
			expected: `{"status":"ok"}`,
		},
		{
			name: "Error controller",
			controller: func(req Request, res Response) Response {
				return Response{Body: "Not Found", Status: NOT_FOUND, Content: PLAIN}
			},
			expected: "Not Found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			req := Request{Url: "/test", HttpMethod: GET, Version: "HTTP/1.1"}
			res := Response{}

			// Act
			result := tt.controller(req, res)

			// Assert
			if result.Body != tt.expected {
				t.Errorf("Expected body %s, got %s", tt.expected, result.Body)
			}
		})
	}
}

// TestControllerCanAccessRequest tests that controller can access request data
func TestControllerCanAccessRequest(t *testing.T) {
	// Arrange
	controller := func(req Request, res Response) Response {
		// Controller can access all request fields
		body := string(req.Url)
		return Response{Body: body, Status: OK, Content: PLAIN}
	}

	// Act
	req := Request{Url: "/api/users", HttpMethod: POST, Version: "HTTP/1.1"}
	res := Response{}
	result := controller(req, res)

	// Assert
	if result.Body != "/api/users" {
		t.Errorf("Expected body '/api/users', got %s", result.Body)
	}
}

// TestControllerCanAccessResponse tests that controller can access response data
func TestControllerCanAccessResponse(t *testing.T) {
	// Arrange
	controller := func(req Request, res Response) Response {
		// Modify input response
		res.Body = "Modified"
		res.Status = OK
		res.Content = HTML
		return res
	}

	// Act
	req := Request{Url: "/test", HttpMethod: GET, Version: "HTTP/1.1"}
	inputRes := Response{}
	result := controller(req, inputRes)

	// Assert
	if result.Body != "Modified" {
		t.Errorf("Expected modified body, got %s", result.Body)
	}
}

// TestControllerWithConditionalLogic tests controller with conditional logic
func TestControllerWithConditionalLogic(t *testing.T) {
	// Arrange
	controller := func(req Request, res Response) Response {
		if req.HttpMethod == GET {
			return Response{Body: "GET response", Status: OK, Content: PLAIN}
		}
		if req.HttpMethod == POST {
			return Response{Body: "POST response", Status: CREATED, Content: JSON}
		}
		return Response{Body: "Unknown method", Status: METHOD_NOT_ALLOWED, Content: PLAIN}
	}

	// Act & Assert
	tests := []struct {
		method       METHOD
		expectedBody string
		expectedCode STATUS
	}{
		{GET, "GET response", OK},
		{POST, "POST response", CREATED},
		{PUT, "Unknown method", METHOD_NOT_ALLOWED},
	}

	for _, tt := range tests {
		req := Request{Url: "/test", HttpMethod: tt.method, Version: "HTTP/1.1"}
		res := Response{}
		result := controller(req, res)

		if result.Body != tt.expectedBody {
			t.Errorf("Method %v: Expected body %s, got %s", tt.method, tt.expectedBody, result.Body)
		}
		if result.Status != tt.expectedCode {
			t.Errorf("Method %v: Expected status %v, got %v", tt.method, tt.expectedCode, result.Status)
		}
	}
}

// TestControllerWithStatelessBehavior tests that controllers are stateless
func TestControllerWithStatelessBehavior(t *testing.T) {
	// Arrange
	controller := func(req Request, res Response) Response {
		return Response{Body: "Constant", Status: OK, Content: PLAIN}
	}

	// Act & Assert - Same controller should always return same result
	for i := 0; i < 3; i++ {
		req := Request{Url: "/test", HttpMethod: GET, Version: "HTTP/1.1"}
		res := Response{}
		result := controller(req, res)

		if result.Body != "Constant" {
			t.Errorf("Iteration %d: Expected constant body, got %s", i, result.Body)
		}
	}
}

// TestMultipleControllers tests multiple different controllers
func TestMultipleControllers(t *testing.T) {
	// Arrange
	homeController := func(req Request, res Response) Response {
		return Response{Body: "Home", Status: OK, Content: HTML}
	}

	aboutController := func(req Request, res Response) Response {
		return Response{Body: "About", Status: OK, Content: HTML}
	}

	// Act
	homeReq := Request{Url: "/", HttpMethod: GET, Version: "HTTP/1.1"}
	aboutReq := Request{Url: "/about", HttpMethod: GET, Version: "HTTP/1.1"}

	homeRes := homeController(homeReq, Response{})
	aboutRes := aboutController(aboutReq, Response{})

	// Assert
	if homeRes.Body != "Home" {
		t.Errorf("Expected home body, got %s", homeRes.Body)
	}
	if aboutRes.Body != "About" {
		t.Errorf("Expected about body, got %s", aboutRes.Body)
	}
}

// TestControllerIgnoresInputResponse tests that controller can ignore input response
func TestControllerIgnoresInputResponse(t *testing.T) {
	// Arrange
	controller := func(req Request, res Response) Response {
		// Ignores the input res parameter
		return Response{Body: "New", Status: OK, Content: PLAIN}
	}

	// Act
	req := Request{Url: "/test", HttpMethod: GET, Version: "HTTP/1.1"}
	inputRes := Response{Body: "Old", Status: NOT_FOUND, Content: HTML}
	result := controller(req, inputRes)

	// Assert
	if result.Body != "New" {
		t.Errorf("Expected new body, got %s", result.Body)
	}
	if result.Status != OK {
		t.Errorf("Expected new status, got %v", result.Status)
	}
}

// TestControllerCanReadRequestURL tests controller reading URL from request
func TestControllerCanReadRequestURL(t *testing.T) {
	// Arrange
	controller := func(req Request, res Response) Response {
		url := string(req.Url)
		return Response{Body: "URL is " + url, Status: OK, Content: PLAIN}
	}

	// Act
	req := Request{Url: "/api/test", HttpMethod: GET, Version: "HTTP/1.1"}
	result := controller(req, Response{})

	// Assert
	if result.Body != "URL is /api/test" {
		t.Errorf("Expected body with URL, got %s", result.Body)
	}
}

// TestControllerCanReadRequestMethod tests controller reading method from request
func TestControllerCanReadRequestMethod(t *testing.T) {
	// Arrange
	controller := func(req Request, res Response) Response {
		method := string(req.HttpMethod)
		return Response{Body: "Method is " + method, Status: OK, Content: PLAIN}
	}

	// Act
	req := Request{Url: "/test", HttpMethod: DELETE, Version: "HTTP/1.1"}
	result := controller(req, Response{})

	// Assert
	if result.Body != "Method is DELETE" {
		t.Errorf("Expected body with method, got %s", result.Body)
	}
}
