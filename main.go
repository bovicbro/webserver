package main

import (
	"os"
	"webserver/server"
	"webserver/server/http"
)

func main() {
	server := server.InitServer(server.Config{})

	server.AddController(
		http.Route{Url: "/", Method: http.GET},
		func(req http.Request, res http.Response) http.Response {
			content, err := os.ReadFile("./index.html")
			if err != nil {
				res = http.Response{Body: "Hello, World!", Status: 200}
			} else {
				res = http.Response{Body: string(content), Status: 200}
			}
			return res
		})

	// End user should not have to pass rouce controllers here. Lets fix later
	server.Listen(3333, server.RouteControllers)
}
