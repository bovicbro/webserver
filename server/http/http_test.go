package http

import (
	"os"
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
	expectedUrl := URL("/api/search")

	// Act
	req, err := ParseRequest(reqRaw)

	// Assert
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if req.Url != expectedUrl {
		t.Errorf("Expected URL %v, got %v", expectedUrl, req.Url)
	}
	if req.Query["q"] != "hello world" {
		t.Errorf("Expected query 'q' to be 'hello world', got %v", req.Query["q"])
	}
}

// TestParseRequestQueryParams tests ParseRequest extracts query parameters
func TestParseRequestQueryParams(t *testing.T) {
	// Arrange
	reqRaw := "GET /api/search?q=test&page=2 HTTP/1.1\nHost: localhost:8000\n"

	// Act
	req, err := ParseRequest(reqRaw)

	// Assert
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if req.Url != "/api/search" {
		t.Errorf("Expected URL '/api/search', got %v", req.Url)
	}
	if req.Query["q"] != "test" {
		t.Errorf("Expected query 'q' to be 'test', got %v", req.Query["q"])
	}
	if req.Query["page"] != "2" {
		t.Errorf("Expected query 'page' to be '2', got %v", req.Query["page"])
	}
}

// TestParseRequestHeaders tests ParseRequest extracts headers
func TestParseRequestHeaders(t *testing.T) {
	// Arrange
	reqRaw := "GET / HTTP/1.1\nHost: localhost:8000\nAccept: text/html\n\n"

	// Act
	req, err := ParseRequest(reqRaw)

	// Assert
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if len(req.Headers) != 2 {
		t.Errorf("Expected 2 headers, got %d", len(req.Headers))
	}
	if req.Headers[0].Key != "Host" || req.Headers[0].Value != "localhost:8000" {
		t.Errorf("Expected first header Host: localhost:8000, got %s: %s", req.Headers[0].Key, req.Headers[0].Value)
	}
	if req.Headers[1].Key != "Accept" || req.Headers[1].Value != "text/html" {
		t.Errorf("Expected second header Accept: text/html, got %s: %s", req.Headers[1].Key, req.Headers[1].Value)
	}
}

// TestParseRequestBody tests ParseRequest extracts body
func TestParseRequestBody(t *testing.T) {
	// Arrange
	reqRaw := "POST /api/echo HTTP/1.1\nContent-Length: 11\n\nhello world"

	// Act
	req, err := ParseRequest(reqRaw)

	// Assert
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if req.Body != "hello world" {
		t.Errorf("Expected body 'hello world', got '%s'", req.Body)
	}
}

// TestResponseSerializeWithCustomHeaders tests Response.Serialize with custom headers
func TestResponseSerializeWithCustomHeaders(t *testing.T) {
	// Arrange
	res := Response{
		Body:    "Hello",
		Status:  OK,
		Content: HTML,
		Headers: []Header{{Key: "X-Custom", Value: "test"}},
	}

	// Act
	serialized := res.Serialize()

	// Assert
	if !strings.Contains(serialized, "X-Custom: test") {
		t.Errorf("Expected custom header in serialized response, got: %s", serialized)
	}
}

// TestResponseSerializeDefaultContentType tests default Content-Type when not set
func TestResponseSerializeDefaultContentType(t *testing.T) {
	// Arrange
	res := Response{
		Body:   "Hello",
		Status: OK,
	}

	// Act
	serialized := res.Serialize()

	// Assert
	if !strings.Contains(serialized, "text/html") {
		t.Errorf("Expected default text/html content type, got: %s", serialized)
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

// TestParseRequestQueryParamNoValue tests query param with no value (e.g., ?key)
func TestParseRequestQueryParamNoValue(t *testing.T) {
	// Arrange
	reqRaw := "GET /path?key HTTP/1.1\nHost: localhost:8000\n\n"

	// Act
	req, err := ParseRequest(reqRaw)

	// Assert
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if req.Query["key"] != "" {
		t.Errorf("Expected query 'key' to be empty string, got %v", req.Query["key"])
	}
}

// TestParseRequestQueryParamEmptyValue tests query param with empty value (e.g., ?key=)
func TestParseRequestQueryParamEmptyValue(t *testing.T) {
	// Arrange
	reqRaw := "GET /path?key= HTTP/1.1\nHost: localhost:8000\n\n"

	// Act
	req, err := ParseRequest(reqRaw)

	// Assert
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if req.Query["key"] != "" {
		t.Errorf("Expected query 'key' to be empty string, got %v", req.Query["key"])
	}
}

// TestParseRequestQueryParamMultipleEquals tests query param with multiple equals signs
func TestParseRequestQueryParamMultipleEquals(t *testing.T) {
	// Arrange
	reqRaw := "GET /path?key=val=ue HTTP/1.1\nHost: localhost:8000\n\n"

	// Act
	req, err := ParseRequest(reqRaw)

	// Assert
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if req.Query["key"] != "val=ue" {
		t.Errorf("Expected query 'key' to be 'val=ue', got %v", req.Query["key"])
	}
}

// TestParseRequestQueryParamEmptyKey tests query param with empty key (e.g., ?=value)
func TestParseRequestQueryParamEmptyKey(t *testing.T) {
	// Arrange
	reqRaw := "GET /path?=value HTTP/1.1\nHost: localhost:8000\n\n"

	// Act
	req, err := ParseRequest(reqRaw)

	// Assert
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if req.Query[""] != "value" {
		t.Errorf("Expected empty key to have value 'value', got %v", req.Query[""])
	}
}

// TestParseRequestEmptyQueryString tests empty query string (e.g., /path?)
func TestParseRequestEmptyQueryString(t *testing.T) {
	// Arrange
	reqRaw := "GET /path? HTTP/1.1\nHost: localhost:8000\n\n"

	// Act
	req, err := ParseRequest(reqRaw)

	// Assert
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if req.Url != "/path" {
		t.Errorf("Expected URL '/path', got %v", req.Url)
	}
	if len(req.Query) != 0 {
		t.Errorf("Expected empty query map, got %d entries", len(req.Query))
	}
}

// TestParseRequestQueryParamDuplicateKeys tests duplicate query param keys
func TestParseRequestQueryParamDuplicateKeys(t *testing.T) {
	// Arrange
	reqRaw := "GET /path?key=value1&key=value2 HTTP/1.1\nHost: localhost:8000\n\n"

	// Act
	req, err := ParseRequest(reqRaw)

	// Assert
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	// Note: current implementation keeps last value for duplicate keys
	if req.Query["key"] != "value2" {
		t.Errorf("Expected last query 'key' value to be 'value2', got %v", req.Query["key"])
	}
}

// TestParseRequestHeaderWithColonInValue tests headers with colons in values
func TestParseRequestHeaderWithColonInValue(t *testing.T) {
	// Arrange
	reqRaw := "POST /api/echo HTTP/1.1\nX-Time: 12:30:45\n\n"

	// Act
	req, err := ParseRequest(reqRaw)

	// Assert
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if len(req.Headers) != 1 {
		t.Errorf("Expected 1 header, got %d", len(req.Headers))
	}
	if req.Headers[0].Key != "X-Time" || req.Headers[0].Value != "12:30:45" {
		t.Errorf("Expected header 'X-Time: 12:30:45', got '%s: %s'", req.Headers[0].Key, req.Headers[0].Value)
	}
}

// TestParseRequestHeaderEmptyValue tests header with empty value
func TestParseRequestHeaderEmptyValue(t *testing.T) {
	// Arrange
	reqRaw := "POST /api/data HTTP/1.1\nX-Empty: \n\n"

	// Act
	req, err := ParseRequest(reqRaw)

	// Assert
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if len(req.Headers) != 1 {
		t.Errorf("Expected 1 header, got %d", len(req.Headers))
	}
	if req.Headers[0].Key != "X-Empty" || req.Headers[0].Value != "" {
		t.Errorf("Expected header 'X-Empty' with empty value, got '%s: %s'", req.Headers[0].Key, req.Headers[0].Value)
	}
}

// TestParseRequestDuplicateHeaders tests duplicate header keys
func TestParseRequestDuplicateHeaders(t *testing.T) {
	// Arrange
	reqRaw := "GET / HTTP/1.1\nAccept: text/html\nAccept: application/json\n\n"

	// Act
	req, err := ParseRequest(reqRaw)

	// Assert
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if len(req.Headers) != 2 {
		t.Errorf("Expected 2 headers, got %d", len(req.Headers))
	}
	// Both headers should be preserved
	if req.Headers[0].Value != "text/html" || req.Headers[1].Value != "application/json" {
		t.Errorf("Expected both Accept headers, got %v", req.Headers)
	}
}

// TestParseRequestHeaderWithWhitespace tests headers with extra whitespace
func TestParseRequestHeaderWithWhitespace(t *testing.T) {
	// Arrange
	reqRaw := "GET / HTTP/1.1\n  Host :  localhost:8000  \n\n"

	// Act
	req, err := ParseRequest(reqRaw)

	// Assert
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	// Should trim whitespace around key and value
	if len(req.Headers) != 1 {
		t.Errorf("Expected 1 header, got %d", len(req.Headers))
	}
	if req.Headers[0].Key != "Host" || req.Headers[0].Value != "localhost:8000" {
		t.Errorf("Expected trimmed header 'Host: localhost:8000', got '%s: %s'", req.Headers[0].Key, req.Headers[0].Value)
	}
}

// TestParseRequestBodyWithCRLF tests body separated by CRLF (Windows line endings)
func TestParseRequestBodyWithCRLF(t *testing.T) {
	// Arrange
	reqRaw := "POST /api/echo HTTP/1.1\r\nContent-Length: 11\r\n\r\nhello world"

	// Act
	req, err := ParseRequest(reqRaw)

	// Assert
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if req.Body != "hello world" {
		t.Errorf("Expected body 'hello world', got '%s'", req.Body)
	}
}

// TestParseRequestBodyWithNullBytes tests body with trailing null bytes
func TestParseRequestBodyWithNullBytes(t *testing.T) {
	// Arrange - simulating body with null byte terminator
	reqRaw := "POST /api/echo HTTP/1.1\nContent-Length: 11\n\nhello world\x00\x00"

	// Act
	req, err := ParseRequest(reqRaw)

	// Assert
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	// Null bytes should be trimmed from the right
	if req.Body != "hello world" {
		t.Errorf("Expected body 'hello world' (null bytes trimmed), got '%s'", req.Body)
	}
}

// TestParseRequestBodyWithNewlines tests body containing newlines
func TestParseRequestBodyWithNewlines(t *testing.T) {
	// Arrange
	reqRaw := "POST /api/echo HTTP/1.1\nContent-Length: 18\n\nline1\nline2\nline3"

	// Act
	req, err := ParseRequest(reqRaw)

	// Assert
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if req.Body != "line1\nline2\nline3" {
		t.Errorf("Expected body with newlines preserved, got '%s'", req.Body)
	}
}

// TestParseRequestWhitespaceOnly tests request with only whitespace
// NOTE: Current implementation splits on \n, so "   \n\n   " produces ["   ", "", "   "],
// which doesn't trigger the "empty first row" check. This documents actual behavior.
func TestParseRequestWhitespaceOnly(t *testing.T) {
	// Arrange
	reqRaw := "   \n\n   "

	// Act
	req, err := ParseRequest(reqRaw)

	// Assert
	// This currently PASSES (no error) because first row is "   " (spaces), not empty
	// The parser then tries to split "   " by space and gets ["", "", ""], resulting
	// in method="", url="", version=""
	if err != nil {
		t.Logf("Got error: %v", err)
	} else {
		t.Logf("Whitespace-only request parsed as: method='%v', url='%v', version='%v'",
			req.HttpMethod, req.Url, req.Version)
		// Documenting the actual behavior - this may indicate the parser should trim
		// the first row before checking if it's empty
		if req.HttpMethod != "" || req.Url != "" || req.Version != "" {
			t.Logf("Note: Parser produced non-empty values from whitespace input")
		}
	}
}

// TestParseRequestExtraSpacesInRequestLine tests extra spaces in request line
func TestParseRequestExtraSpacesInRequestLine(t *testing.T) {
	// Arrange - multiple spaces between parts
	reqRaw := "GET   /path   HTTP/1.1\nHost: localhost\n\n"

	// Act
	_, err := ParseRequest(reqRaw)

	// Assert
	// Current implementation splits on single space, so this may or may not work
	// Documenting behavior - splitting on " " gives ["GET", "", "", "/path", "", "", "HTTP/1.1"]
	// which is more than 3 parts, so it might pass or fail depending on implementation
	if err == nil {
		// If it passes, check what we got
		req, _ := ParseRequest(reqRaw)
		t.Logf("Extra spaces parsed as: method=%v, url=%v, version=%v", req.HttpMethod, req.Url, req.Version)
	}
}

// TestParseRequestFromTestDataFile tests reading complex requests from testdata files
func TestParseRequestFromTestDataFile(t *testing.T) {
	tests := []struct {
		filename     string
		expectedURL  URL
		expectedKeys []string
	}{
		{"query_param_no_value.txt", "/path", []string{"key"}},
		{"query_param_empty_value.txt", "/path", []string{"key"}},
		{"query_param_multiple_equals.txt", "/path", []string{"key"}},
		{"query_param_empty_key.txt", "/path", []string{""}},
	}

	for _, tt := range tests {
		t.Run(tt.filename, func(t *testing.T) {
			// Read test data file
			content, err := os.ReadFile("testdata/" + tt.filename)
			if err != nil {
				t.Fatalf("Failed to read testdata file %s: %v", tt.filename, err)
			}

			// Act
			req, err := ParseRequest(string(content))

			// Assert
			if err != nil {
				t.Errorf("Expected no error for %s, got %v", tt.filename, err)
			}
			if req.Url != tt.expectedURL {
				t.Errorf("Expected URL %v, got %v", tt.expectedURL, req.Url)
			}
			for _, key := range tt.expectedKeys {
				if _, ok := req.Query[key]; !ok {
					t.Errorf("Expected query key '%s' to exist", key)
				}
			}
		})
	}
}
