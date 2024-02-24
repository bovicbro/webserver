package http

type URL string

type Header string

type METHOD string

const (
	GET   METHOD = "GET"
	POST  METHOD = "POST"
	PUT   METHOD = "PUT"
	DELET METHOD = "DELETE"
)

type STATUS int

const (
	OK        STATUS = 200
	NOT_FOUND STATUS = 404
)

type Response string

type Request struct {
	Url        URL
	HttpMethod METHOD
	Headers    []Header
}

type Route struct {
	Url    URL
	Method METHOD
}
