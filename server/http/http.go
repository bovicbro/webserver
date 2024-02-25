package http

import (
	"fmt"
	"strings"
)

type URL string

type Header string

type METHOD string

const (
	GET   METHOD = "GET"
	POST  METHOD = "POST"
	PUT   METHOD = "PUT"
	DELET METHOD = "DELETE"
)

type CONTENT string

const (
	HTML CONTENT = "text/html"
	CSS  CONTENT = "text/css"
	JSON CONTENT = "application/json"
)

const HTTPVERSION = "HTTP/1.1"

type STATUS string

const (
	OK        STATUS = "200 OK"
	NOT_FOUND STATUS = "404 Not Found"
)

type Response struct {
	Body    string
	Status  STATUS
	Headers []Header
	Content CONTENT
}

func (res Response) Serialize() string {
	resString := fmt.Sprintf(`HTTP/1.1 %s 
Content-length: %d 
Content-Type: %s; charset=utf-8

%s`,
		res.Status,
		len(res.Body),
		res.Content,
		res.Body,
	)
	return resString
}

type Request struct {
	Url        URL
	HttpMethod METHOD
	Headers    []Header
	Version    string
}

type Route struct {
	Url    URL
	Method METHOD
}

func ParseRequest(reqRaw string) Request {
	var req = Request{}
	rows := strings.Split(reqRaw, "\n")
	method := strings.Split(rows[0], " ")[0]
	url := strings.Split(rows[0], " ")[1]
	version := strings.Split(rows[0], " ")[2]
	req.HttpMethod = METHOD(method)
	req.Url = URL(url)
	req.Version = version
	return req
}

func CreateBaseResponse(req Request) Response {
	res := Response{}
	return res
}
