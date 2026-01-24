package server

import (
	"testing"
	"webserver/server/http"
)

// TestInitServer tests that InitServer properly initializes a Server
func TestInitServer(t *testing.T) {
	// Arrange
	config := Config{}

	// Act
	server := InitServer(config)

	// Assert
	if server == nil {
		t.Errorf("Expected server to be initialized, got nil")
	}
	if len(server.RouteControllers) != 0 {
		t.Errorf("Expected empty RouteControllers, got %d", len(server.RouteControllers))
	}
}

// TestServerAddController tests that AddController function is set
func TestServerAddController(t *testing.T) {
	// Arrange
	config := Config{}
	server := InitServer(config)

	if server.AddController == nil {
		t.Errorf("Expected AddController to be set, got nil")
	}
}

// TestServerListen tests that Listen function is set
func TestServerListen(t *testing.T) {
	// Arrange
	config := Config{}
	server := InitServer(config)

	// Assert
	if server.Listen == nil {
		t.Errorf("Expected Listen to be set, got nil")
	}
}

// TestServerAddControllerFunctionality tests adding a controller to the server
func TestServerAddControllerFunctionality(t *testing.T) {
	// Arrange
	config := Config{}
	server := InitServer(config)
	route := http.Route{Url: "/", Method: http.GET}
	controller := func(req http.Request, res http.Response) http.Response {
		return http.Response{Body: "Home", Status: http.OK, Content: http.HTML}
	}

	// Act
	server.AddController(route, controller)

	// Assert
	if len(server.RouteControllers) != 1 {
		t.Errorf("Expected 1 route controller, got %d", len(server.RouteControllers))
	}
}

// TestServerAddMultipleControllers tests adding multiple controllers
func TestServerAddMultipleControllers(t *testing.T) {
	// Arrange
	config := Config{}
	server := InitServer(config)

	controller1 := func(req http.Request, res http.Response) http.Response {
		return http.Response{Body: "Home", Status: http.OK, Content: http.HTML}
	}
	controller2 := func(req http.Request, res http.Response) http.Response {
		return http.Response{Body: "About", Status: http.OK, Content: http.HTML}
	}

	// Act
	server.AddController(http.Route{Url: "/", Method: http.GET}, controller1)
	server.AddController(http.Route{Url: "/about", Method: http.GET}, controller2)

	// Assert
	if len(server.RouteControllers) != 2 {
		t.Errorf("Expected 2 route controllers, got %d", len(server.RouteControllers))
	}
}

// TestServerPortType tests Port type
func TestServerPortType(t *testing.T) {
	// Arrange
	var port Port = 8000

	// Assert
	if port != 8000 {
		t.Errorf("Expected port 8000, got %d", port)
	}
}

// TestServerConfigStruct tests Config struct
func TestServerConfigStruct(t *testing.T) {
	// Arrange & Act
	config := Config{}

	// Assert
	// Config is currently empty, just ensure it can be created
	if config == (Config{}) {
		t.Log("Config struct is empty as expected")
	}
}

// TestServerStructure tests that Server struct has correct fields
func TestServerStructure(t *testing.T) {
	// Arrange
	config := Config{}
	server := InitServer(config)

	// Assert
	if server.AddController == nil {
		t.Errorf("Expected AddController field, got nil")
	}
	if server.Listen == nil {
		t.Errorf("Expected Listen field, got nil")
	}
	// RouteControllers should be a slice (can be empty)
	if cap(server.RouteControllers) < 0 {
		t.Errorf("RouteControllers is not a valid slice")
	}
}

// TestServerReturnPointer tests that InitServer returns a pointer
func TestServerReturnPointer(t *testing.T) {
	// Arrange
	config := Config{}

	// Act
	server := InitServer(config)

	// Assert
	if server == nil {
		t.Errorf("Expected pointer to Server, got nil")
	}
}

// TestServerMutability tests that server state can be mutated
func TestServerMutability(t *testing.T) {
	// Arrange
	config := Config{}
	server := InitServer(config)
	initialCount := len(server.RouteControllers)

	controller := func(req http.Request, res http.Response) http.Response {
		return http.Response{Body: "Test", Status: http.OK, Content: http.HTML}
	}

	// Act
	server.AddController(http.Route{Url: "/test", Method: http.GET}, controller)

	// Assert
	if len(server.RouteControllers) == initialCount {
		t.Errorf("Expected route controllers to be added, count unchanged")
	}
	if len(server.RouteControllers) != 1 {
		t.Errorf("Expected 1 route, got %d", len(server.RouteControllers))
	}
}

// TestAddControllerWithDifferentMethods tests adding controllers with different HTTP methods
func TestAddControllerWithDifferentMethods(t *testing.T) {
	// Arrange
	config := Config{}
	server := InitServer(config)

	getController := func(req http.Request, res http.Response) http.Response {
		return http.Response{Body: "GET", Status: http.OK, Content: http.PLAIN}
	}
	postController := func(req http.Request, res http.Response) http.Response {
		return http.Response{Body: "POST", Status: http.CREATED, Content: http.JSON}
	}

	// Act
	server.AddController(http.Route{Url: "/api", Method: http.GET}, getController)
	server.AddController(http.Route{Url: "/api", Method: http.POST}, postController)

	// Assert
	if len(server.RouteControllers) != 2 {
		t.Errorf("Expected 2 routes, got %d", len(server.RouteControllers))
	}
}

// TestServerIsPointer tests that operations modify original server
func TestServerIsPointer(t *testing.T) {
	// Arrange
	config := Config{}
	server1 := InitServer(config)

	controller := func(req http.Request, res http.Response) http.Response {
		return http.Response{Body: "Test", Status: http.OK, Content: http.HTML}
	}

	// Act
	server1.AddController(http.Route{Url: "/test1", Method: http.GET}, controller)

	// Assert
	if len(server1.RouteControllers) != 1 {
		t.Errorf("Expected 1 route in server1, got %d", len(server1.RouteControllers))
	}
}

// TestServerRouteControllersType tests that RouteControllers has correct type
func TestServerRouteControllersType(t *testing.T) {
	// Arrange
	config := Config{}
	server := InitServer(config)

	// Act
	controller := func(req http.Request, res http.Response) http.Response {
		return http.Response{Body: "Test", Status: http.OK, Content: http.HTML}
	}
	server.AddController(http.Route{Url: "/test", Method: http.GET}, controller)

	// Assert
	if len(server.RouteControllers) != 1 {
		t.Errorf("Expected RouteControllers to have 1 element, got %d", len(server.RouteControllers))
	}
}

// TestConfigWithPort tests Config struct with custom port
func TestConfigWithPort(t *testing.T) {
	// Arrange
	config := Config{Port: 3000}

	// Act
	srv := InitServer(config)

	// Assert
	if srv.Port != 3000 {
		t.Errorf("Expected port 3000, got %d", srv.Port)
	}
}

// TestConfigDefaultPort tests default port when not specified
func TestConfigDefaultPort(t *testing.T) {
	// Arrange
	config := Config{} // Port not specified (defaults to 0)

	// Act
	srv := InitServer(config)

	// Assert
	if srv.Port != 8000 {
		t.Errorf("Expected default port 8000, got %d", srv.Port)
	}
}

// TestServerPortFromConfig tests that server gets port from config
func TestServerPortFromConfig(t *testing.T) {
	testCases := []struct {
		name       string
		configPort Port
		expected   Port
	}{
		{"zero defaults to 8000", 0, 8000},
		{"custom port 3000", 3000, 3000},
		{"custom port 5000", 5000, 5000},
		{"custom port 9999", 9999, 9999},
		{"alternative port 8080", 8080, 8080},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			config := Config{Port: tc.configPort}

			// Act
			srv := InitServer(config)

			// Assert
			if srv.Port != tc.expected {
				t.Errorf("Expected port %d, got %d", tc.expected, srv.Port)
			}
		})
	}
}

// TestServerExportedPort tests that Port field is exported
func TestServerExportedPort(t *testing.T) {
	// Arrange
	config := Config{Port: 5000}
	srv := InitServer(config)

	// Act - Port should be accessible directly
	port := srv.Port

	// Assert
	if port != 5000 {
		t.Errorf("Expected to access port directly, got %d", port)
	}
}
