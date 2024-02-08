package server

import (
	"fmt"
	. "webserver/server/controller"
	. "webserver/server/http"
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
	Listen           func(port Port)
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

	server.Listen = func(port Port) {

		var requestUrl URL
		for {
			fmt.Print("Make request: ")
			fmt.Scan(&requestUrl)
			fmt.Println(server.router(Request{Url: requestUrl, HttpMethod: GET}, server.routeControllers))
		}
	}
	return server
}

func createBaseResponse(req Request) Response {
	return "response"
}
