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

const HTTPVERSION = "HTTP/1.1"

type STATUS int

const (
	OK        STATUS = 200
	NOT_FOUND STATUS = 404
)

type Response struct {
	Body    string
	Status  STATUS
	Headers []Header
}

func (res Response) Serialize() string {
	resString := fmt.Sprintf(`HTTP/1.1 200 OK 
Content-length: %d 
Content-Type: text/html; charset=utf-8

%s`,
		len(res.Body),
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
