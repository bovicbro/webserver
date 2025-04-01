// Package that handles http
package http

import (
	"errors"
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
	MALFORMED STATUS = "400 Bad Request"
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

func ParseRequest(reqRaw string) (Request, error) {
	var req = Request{}
	rows := strings.Split(reqRaw, "\n")
	if len(rows) < 1 {
		return req, errors.New("Malformed request")
	}
	firstRow := strings.Split(rows[0], " ")
	if len(firstRow) < 3 {
		return req, errors.New("Malformed request")
	}
	method := firstRow[0]
	url := firstRow[1]
	version := firstRow[2]
	req.HttpMethod = METHOD(method)
	req.Url = URL(url)
	req.Version = version
	return req, nil
}

func CreateBaseResponse(req Request) Response {
	res := Response{}
	return res
}
