package server

import (
	. "webserver/server/controller"
	. "webserver/server/http"
	"webserver/server/networking"
	"webserver/server/router"
)

type Config struct {
}

type routeController struct {
	route      Route
	controller Controller
}

type Port int

type Server struct {
	port             Port
	RouteControllers []router.ControlledRoutes
	AddController    func(route Route, controller Controller)
	Listen           networking.ListenerType
}

func InitServer(config Config) *Server {
	var server Server

	server.AddController = func(route Route, controller Controller) {
		server.RouteControllers = router.AddController(
			route,
			controller,
			server.RouteControllers,
		)
	}

	server.Listen = networking.Listen

	return &server
}
