package main

import (
	"webserver/server"
	"webserver/server/http"
)

func main() {
	server := server.InitServer(server.Config{})

	server.AddController(
		http.Route{Url: "/", Method: http.GET},
		func(req http.Request, res http.Response) http.Response {
			res = http.Response{Body: "smth", Status: 200}
			return res
		})

	// End user should not have to pass rouce controllers here. Lets fix later
	server.Listen(3333, server.RouteControllers)
}
