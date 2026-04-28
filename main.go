package main

import (
	"encoding/json"
	"strconv"
	"webserver/server"
	"webserver/server/controller"
	"webserver/server/http"
)

func main() {
	webServer := server.InitServer(server.Config{
		Port: 8000,
	})

	// API: search with query params
	webServer.AddController(
		http.Route{Url: "/api/search", Method: http.GET},
		func(req http.Request, res http.Response) http.Response {
			query := req.Query["q"]
			if query == "" {
				query = "nothing"
			}
			return http.Response{
				Status:  http.OK,
				Content: http.JSON,
				Body:    `{"query":"` + query + `"}`,
			}
		})

	// API: user by ID (path params)
	webServer.AddController(
		http.Route{Url: "/api/users/:id", Method: http.GET},
		func(req http.Request, res http.Response) http.Response {
			userID := req.PathParams["id"]
			body, _ := json.Marshal(map[string]string{"user_id": userID})
			return http.Response{
				Status:  http.OK,
				Content: http.JSON,
				Body:    string(body),
			}
		})

	// API: echo body
	webServer.AddController(
		http.Route{Url: "/api/echo", Method: http.POST},
		func(req http.Request, res http.Response) http.Response {
			return http.Response{
				Status:  http.OK,
				Content: http.JSON,
				Body:    `{"received":"` + req.Body + `"}`,
			}
		})

	// Static file serving (catch-all)
	webServer.AddController(
		http.Route{Url: "/*", Method: http.GET},
		controller.Static("./static"))

	webServer.Listen(strconv.Itoa(int(webServer.Port)), webServer.RouteControllers)
}
