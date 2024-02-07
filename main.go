package main

import (
	"webserver/http"
	"webserver/server"
)

func main() {
	server := server.InitServer(server.Config{})

	server.AddController(
		http.Route{Url: "/", Method: http.GET},
		func(req http.Request, res http.Response) http.Response {
			res = "smh"
			return res
		})

	server.Listen(3333)
}
