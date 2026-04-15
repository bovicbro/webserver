package main

import (
	"os"
	"strconv"
	"webserver/server"
	"webserver/server/http"
)

func main() {
	// Initialize server with custom port (use 0 or omit for default 8000)
	webServer := server.InitServer(server.Config{
		Port: 8000,
	})

	webServer.AddController(
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

	webServer.AddController(
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

	webServer.AddController(
		http.Route{Url: "/contributors", Method: http.GET},
		func(req http.Request, res http.Response) http.Response {
			content, err := os.ReadFile("./static/contributors.html")
			if err != nil {
				res = http.Response{Status: http.NOT_FOUND}
			} else {
				res = http.Response{Body: string(content), Status: http.OK}
			}
			return res
		})

	webServer.AddController(
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

	webServer.Listen(strconv.Itoa(int(webServer.Port)), webServer.RouteControllers)
}
