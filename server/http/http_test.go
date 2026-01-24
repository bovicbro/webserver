package http

import (
	"strings"
	"testing"
)

// TestParseRequestHappy tests ParseRequest with a valid HTTP request
func TestParseRequestHappy(t *testing.T) {
	// Arrange
	reqRaw := "GET /index.html HTTP/1.1\nHost: localhost:8000\n"
	expectedMethod := METHOD("GET")
	expectedUrl := URL("/index.html")
	expectedVersion := "HTTP/1.1"

	// Act
	req, err := ParseRequest(reqRaw)

	// Assert
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if req.HttpMethod != expectedMethod {
		t.Errorf("Expected method %v, got %v", expectedMethod, req.HttpMethod)
	}
	if req.Url != expectedUrl {
		t.Errorf("Expected URL %v, got %v", expectedUrl, req.Url)
	}
	if req.Version != expectedVersion {
		t.Errorf("Expected version %v, got %v", expectedVersion, req.Version)
	}
}

// TestParseRequestWithPostMethod tests ParseRequest with POST method
func TestParseRequestWithPostMethod(t *testing.T) {
	// Arrange
	reqRaw := "POST /api/user HTTP/1.1\nHost: localhost:8000\n"
	expectedMethod := METHOD("POST")
	expectedUrl := URL("/api/user")

	// Act
	req, err := ParseRequest(reqRaw)

	// Assert
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if req.HttpMethod != expectedMethod {
		t.Errorf("Expected method %v, got %v", expectedMethod, req.HttpMethod)
	}
	if req.Url != expectedUrl {
		t.Errorf("Expected URL %v, got %v", expectedUrl, req.Url)
	}
}

// TestParseRequestEmptyString tests ParseRequest with empty string
func TestParseRequestEmptyString(t *testing.T) {
	// Arrange
	reqRaw := ""

	// Act
	_, err := ParseRequest(reqRaw)

	// Assert
	if err == nil {
		t.Errorf("Expected error for empty request, got nil")
	}
}

// TestParseRequestMalformedFirstLine tests ParseRequest with malformed first line
func TestParseRequestMalformedFirstLine(t *testing.T) {
	// Arrange
	reqRaw := "INVALID\nHost: localhost:8000\n"

	// Act
	_, err := ParseRequest(reqRaw)

	// Assert
	if err == nil {
		t.Errorf("Expected error for malformed request, got nil")
	}
}

// TestParseRequestMissingVersion tests ParseRequest with missing version
func TestParseRequestMissingVersion(t *testing.T) {
	// Arrange
	reqRaw := "GET /index.html\n"

	// Act
	_, err := ParseRequest(reqRaw)

	// Assert
	if err == nil {
		t.Errorf("Expected error for malformed request, got nil")
	}
}

// TestParseRequestAllMethods tests ParseRequest with all HTTP methods
func TestParseRequestAllMethods(t *testing.T) {
	methods := []string{"GET", "POST", "PUT", "DELETE", "PATCH", "HEAD", "OPTIONS"}

	for _, method := range methods {
		// Arrange
		reqRaw := method + " /test HTTP/1.1\nHost: localhost:8000\n"

		// Act
		req, err := ParseRequest(reqRaw)

		// Assert
		if err != nil {
			t.Errorf("Expected no error for method %s, got %v", method, err)
		}
		if string(req.HttpMethod) != method {
			t.Errorf("Expected method %s, got %v", method, req.HttpMethod)
		}
	}
}

// TestResponseSerialize tests Response.Serialize method
func TestResponseSerialize(t *testing.T) {
	// Arrange
	res := Response{
		Body:    "Hello World",
		Status:  OK,
		Content: HTML,
	}

	// Act
	serialized := res.Serialize()

	// Assert
	if !strings.Contains(serialized, "200 OK") {
		t.Errorf("Expected '200 OK' in serialized response, got: %s", serialized)
	}
	if !strings.Contains(serialized, "Hello World") {
		t.Errorf("Expected 'Hello World' in serialized response, got: %s", serialized)
	}
	if !strings.Contains(serialized, "text/html") {
		t.Errorf("Expected 'text/html' in serialized response, got: %s", serialized)
	}
	if !strings.Contains(serialized, "Content-length: 11") {
		t.Errorf("Expected correct content length, got: %s", serialized)
	}
}

// TestResponseSerializeWithEmptyBody tests Response.Serialize with empty body
func TestResponseSerializeWithEmptyBody(t *testing.T) {
	// Arrange
	res := Response{
		Body:    "",
		Status:  NO_CONTENT,
		Content: PLAIN,
	}

	// Act
	serialized := res.Serialize()

	// Assert
	if !strings.Contains(serialized, "204 No Content") {
		t.Errorf("Expected '204 No Content' in response, got: %s", serialized)
	}
	if !strings.Contains(serialized, "Content-length: 0") {
		t.Errorf("Expected content length 0, got: %s", serialized)
	}
}

// TestResponseSerializeWithJSON tests Response.Serialize with JSON content
func TestResponseSerializeWithJSON(t *testing.T) {
	// Arrange
	jsonBody := `{"name":"test","value":123}`
	res := Response{
		Body:    jsonBody,
		Status:  OK,
		Content: JSON,
	}

	// Act
	serialized := res.Serialize()

	// Assert
	if !strings.Contains(serialized, "application/json") {
		t.Errorf("Expected 'application/json' in response, got: %s", serialized)
	}
	if !strings.Contains(serialized, jsonBody) {
		t.Errorf("Expected JSON body in response, got: %s", serialized)
	}
}

// TestResponseSerializeWithError tests Response.Serialize with error status
func TestResponseSerializeWithError(t *testing.T) {
	// Arrange
	res := Response{
		Body:    "404 Not Found",
		Status:  NOT_FOUND,
		Content: PLAIN,
	}

	// Act
	serialized := res.Serialize()

	// Assert
	if !strings.Contains(serialized, "404 Not Found") {
		t.Errorf("Expected '404 Not Found' in response, got: %s", serialized)
	}
}

// TestCreateBaseResponse tests CreateBaseResponse function
func TestCreateBaseResponse(t *testing.T) {
	// Arrange
	req := Request{
		Url:        URL("/test"),
		HttpMethod: GET,
		Version:    "HTTP/1.1",
	}

	// Act
	res := CreateBaseResponse(req)

	// Assert
	if res.Body != "" {
		t.Errorf("Expected empty body, got %s", res.Body)
	}
	if res.Status != "" {
		t.Errorf("Expected empty status, got %s", res.Status)
	}
	if res.Content != "" {
		t.Errorf("Expected empty content type, got %s", res.Content)
	}
}

// TestParseRequestWithSpecialCharacters tests ParseRequest with special characters in URL
func TestParseRequestWithSpecialCharacters(t *testing.T) {
	// Arrange
	reqRaw := "GET /api/search?q=hello%20world HTTP/1.1\nHost: localhost:8000\n"
	expectedUrl := URL("/api/search?q=hello%20world")

	// Act
	req, err := ParseRequest(reqRaw)

	// Assert
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if req.Url != expectedUrl {
		t.Errorf("Expected URL %v, got %v", expectedUrl, req.Url)
	}
}

// TestResponseSerializeWithLargeBody tests Response.Serialize with large body
func TestResponseSerializeWithLargeBody(t *testing.T) {
	// Arrange
	largeBody := strings.Repeat("x", 10000)
	res := Response{
		Body:    largeBody,
		Status:  OK,
		Content: HTML,
	}

	// Act
	serialized := res.Serialize()

	// Assert
	if !strings.Contains(serialized, "Content-length: 10000") {
		t.Errorf("Expected correct large content length, got: %s", serialized[:100])
	}
	if !strings.Contains(serialized, largeBody) {
		t.Errorf("Expected large body in response")
	}
}

// TestContentTypeConstants tests that content type constants are correctly defined
func TestContentTypeConstants(t *testing.T) {
	tests := []struct {
		name     string
		content  CONTENT
		expected string
	}{
		{"HTML", HTML, "text/html"},
		{"JSON", JSON, "application/json"},
		{"CSS", CSS, "text/css"},
		{"PLAIN", PLAIN, "text/plain"},
		{"PNG", PNG, "image/png"},
	}

	for _, tt := range tests {
		if string(tt.content) != tt.expected {
			t.Errorf("Test %s: expected %s, got %s", tt.name, tt.expected, tt.content)
		}
	}
}

// TestStatusCodeConstants tests that status code constants are correctly defined
func TestStatusCodeConstants(t *testing.T) {
	tests := []struct {
		name     string
		status   STATUS
		expected string
	}{
		{"OK", OK, "200 OK"},
		{"CREATED", CREATED, "201 Created"},
		{"BAD_REQUEST", BAD_REQUEST, "400 Bad Request"},
		{"NOT_FOUND", NOT_FOUND, "404 Not Found"},
		{"INTERNAL_SERVER_ERROR", INTERNAL_SERVER_ERROR, "500 Internal Server Error"},
	}

	for _, tt := range tests {
		if string(tt.status) != tt.expected {
			t.Errorf("Test %s: expected %s, got %s", tt.name, tt.expected, tt.status)
		}
	}
}

// TestMethodConstants tests that HTTP method constants are correctly defined
func TestMethodConstants(t *testing.T) {
	tests := []struct {
		name     string
		method   METHOD
		expected string
	}{
		{"GET", GET, "GET"},
		{"POST", POST, "POST"},
		{"PUT", PUT, "PUT"},
		{"DELETE", DELETE, "DELETE"},
		{"PATCH", PATCH, "PATCH"},
	}

	for _, tt := range tests {
		if string(tt.method) != tt.expected {
			t.Errorf("Test %s: expected %s, got %s", tt.name, tt.expected, tt.method)
		}
	}
}
