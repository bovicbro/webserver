package main

import (
	"fmt"
)

// A HTTP servers function is to allow the user to receive http requests and send http responses
// It does this by using some underlying communication protocoll e.g. TCP.
// A core concept in this domain is the URL - Unified Resource locator. The requests specify a URL which specifies a resource.
// Another core conecpt is the HTTP method which indiciates intent.
// The combination of the URL and the method is used to select different behaviours which can modify state of the server as well as response.
// Another concept is Headers which contains meta-data about the request and responses

type URL string

type Header string

type HttpMethod string

const (
	GET   HttpMethod = "GET"
	POST  HttpMethod = "POST"
	PUT   HttpMethod = "PUT"
	DELET HttpMethod = "DELETE"
)

type Port int

type Response string

type Request struct {
	url        URL
	httpMethod HttpMethod
	headers    []Header
}

type Route struct {
	url    URL
	method HttpMethod
}
type Controller = func(req Request, res Response) Response

type Config struct {
}

type routeController struct {
	route      Route
	controller Controller
}

type Server struct {
	port             Port
	routeControllers []routeController
	addController    func(route Route, controller Controller)
	listen           func(port Port)
	handleRequest    func(res Request) Response
}

func initServer(config Config) Server {
	var server Server

	server.addController = func(route Route, controller Controller) {
		server.routeControllers = append(
			server.routeControllers,
			routeController{
				route:      route,
				controller: controller,
			})
	}

	server.handleRequest = func(req Request) Response {
		index := sliceIndexOf(server.routeControllers, func(rc routeController) bool {
			return rc.route.url == req.url
		})
		if index == -1 {
			return "404"
		}
		rc := server.routeControllers[index]
		return rc.controller(req, createBaseResponse(req))
	}

	server.listen = func(port Port) {
		fmt.Printf("Listening on port: %d\n", port)
		var requestUrl URL
		for {
			fmt.Print("Make request: ")
			fmt.Scan(&requestUrl)
			fmt.Println(server.handleRequest(Request{url: requestUrl, httpMethod: GET}))
		}
	}
	return server
}

func main() {
	server := initServer(Config{})

	server.addController(
		Route{url: "/", method: GET},
		func(req Request, res Response) Response {
			res = "smh"
			return res
		})

	server.listen(3333)
}

// Returns index of the first element of the array for which f returns true.
// If there is no element for which f returns true, -1 will be returned.
func sliceIndexOf[E any](slice []E, f func(E) bool) int {
	for i, v := range slice {
		if f(v) {
			return i
		}
	}
	return -1
}

func createBaseResponse(req Request) Response {
	return "response"
}
