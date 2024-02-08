package http

type URL string

type Header string

type HttpMethod string

const (
	GET   HttpMethod = "GET"
	POST  HttpMethod = "POST"
	PUT   HttpMethod = "PUT"
	DELET HttpMethod = "DELETE"
)

type Response string

type Request struct {
	Url        URL
	HttpMethod HttpMethod
	Headers    []Header
}

type Route struct {
	Url    URL
	Method HttpMethod
}
