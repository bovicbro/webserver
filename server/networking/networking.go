package networking

import (
	"fmt"
	"webserver/server/http"
)

type Port int

type Listener interface {
	greet() string
}

func Listen(port Port, callback func(res http.URL)) {
	var requestUrl http.URL
	for {
		fmt.Print("Make request: ")
		fmt.Scan(&requestUrl)
		callback(requestUrl)
	}
}
