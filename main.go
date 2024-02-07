package main

import (
	. "webserver/http"
	. "webserver/server"
)

// A HTTP servers function is to allow the user to receive http requests and send http responses
// It does this by using some underlying communication protocoll e.g. TCP.
// A core concept in this domain is the URL - Unified Resource locator. The requests specify a URL which specifies a resource.
// Another core conecpt is the HTTP method which indiciates intent.
// The combination of the URL and the method is used to select different behaviours which can modify state of the server as well as response.
// Another concept is Headers which contains meta-data about the request and responses

func main() {
	server := InitServer(Config{})

	server.AddController(
		Route{Url: "/", Method: GET},
		func(req Request, res Response) Response {
			res = "smh"
			return res
		})

	server.Listen(3333)
}
