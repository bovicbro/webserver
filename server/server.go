package server

import (
	"fmt"
	. "webserver/http"
	"webserver/utility"
)

type Controller = func(req Request, res Response) Response

type Config struct {
}

type routeController struct {
	route      Route
	controller Controller
}

type Port int

type Server struct {
	port             Port
	routeControllers []routeController
	AddController    func(route Route, controller Controller)
	Listen           func(port Port)
	handleRequest    func(res Request) Response
}

func InitServer(config Config) Server {
	var server Server

	server.AddController = func(route Route, controller Controller) {
		server.routeControllers = append(
			server.routeControllers,
			routeController{
				route:      route,
				controller: controller,
			})
	}

	server.handleRequest = func(req Request) Response {
		index := utility.SliceIndexOf(server.routeControllers, func(rc routeController) bool {
			return rc.route.Url == req.Url
		})
		if index == -1 {
			return "404"
		}
		rc := server.routeControllers[index]
		return rc.controller(req, createBaseResponse(req))
	}

	server.Listen = func(port Port) {
		fmt.Printf("Listening on port: %d\n", port)
		var requestUrl URL
		for {
			fmt.Print("Make request: ")
			fmt.Scan(&requestUrl)
			fmt.Println(server.handleRequest(Request{Url: requestUrl, HttpMethod: GET}))
		}
	}
	return server
}

func createBaseResponse(req Request) Response {
	return "response"
}
