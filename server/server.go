package server

import (
	"webserver/server/controller"
	"webserver/server/http"
	"webserver/server/networking"
	"webserver/server/router"
)

type Config struct {
	Port Port // Server port (default: 8000)
}

type routeController struct {
	route      http.Route
	controller controller.Controller
}

type Port int

type Server struct {
	Port             Port // Server port
	RouteControllers []router.ControlledRoutes
	AddController    func(
		route http.Route,
		controller controller.Controller)
	Listen networking.ListenerType
}

func InitServer(config Config) *Server {
	var server Server

	// Set port from config, default to 8000
	server.Port = config.Port
	if server.Port == 0 {
		server.Port = 8000
	}

	server.AddController = func(
		route http.Route,
		controller controller.Controller) {
		server.RouteControllers = router.AddController(
			route,
			controller,
			server.RouteControllers,
		)
	}

	server.Listen = networking.Listen

	return &server
}
