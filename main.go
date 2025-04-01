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
			content, err := os.ReadFile("./static/index.html")

			if err != nil {
				res = http.Response{Status: http.NOT_FOUND}
			} else {
				res = http.Response{Body: string(content), Status: http.OK}
			}
			return res
		})

	server.AddController(
		http.Route{Url: "/about", Method: http.GET},
		func(req http.Request, res http.Response) http.Response {
			content, err := os.ReadFile("./static/about.html")
			if err != nil {
				res = http.Response{Status: http.NOT_FOUND}
			} else {
				res = http.Response{Body: string(content), Status: http.OK}
			}
			return res
		})

	server.AddController(
		http.Route{Url: "/styles.css", Method: http.GET},
		func(req http.Request, res http.Response) http.Response {
			content, err := os.ReadFile("./static/styles.css")
			if err != nil {
				res = http.Response{Status: http.NOT_FOUND}
			} else {
				res = http.Response{Body: string(content), Status: http.OK}
			}
			return res
		})

	// End user should not have to pass rouce controllers here. Lets fix later
	// The port is not wired up yet so should change that as well.
	server.Listen(0000, server.RouteControllers)
}
