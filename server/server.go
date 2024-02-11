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
	routeControllers []router.ControlledRoutes
	AddController    func(route Route, controller Controller)
	Listen           networking.ListenerType
	router           router.RouterType
}

func InitServer(config Config) Server {
	var server Server

	server.AddController = func(route Route, controller Controller) {
		server.routeControllers = router.AddController(
			route,
			controller,
			server.routeControllers,
		)
	}

	server.router = router.Router

	server.Listen = networking.Listen

	return server
}

func createBaseResponse(req Request) Response {
	return "response"
}
