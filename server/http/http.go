// Package that handles http
package http

import (
	"errors"
	"fmt"
	"strings"
)

type URL string

type Header struct {
	key   string
	value string
}

type METHOD string

const (
	GET     METHOD = "GET"
	POST    METHOD = "POST"
	PUT     METHOD = "PUT"
	DELETE  METHOD = "DELETE"
	PATCH   METHOD = "PATCH"
	HEAD    METHOD = "HEAD"
	OPTIONS METHOD = "OPTIONS"
	TRACE   METHOD = "TRACE"
	CONNECT METHOD = "CONNECT"
)

type CONTENT string

const (
	// Text types
	HTML     CONTENT = "text/html"
	PLAIN    CONTENT = "text/plain"
	CSS      CONTENT = "text/css"
	CSV      CONTENT = "text/csv"
	XML_TEXT CONTENT = "text/xml"
	MARKDOWN CONTENT = "text/markdown"

	// Application types
	JSON            CONTENT = "application/json"
	JAVASCRIPT      CONTENT = "application/javascript"
	XML_APP         CONTENT = "application/xml"
	FORM_URLENCODED CONTENT = "application/x-www-form-urlencoded"
	PDF             CONTENT = "application/pdf"
	ZIP             CONTENT = "application/zip"
	GZIP            CONTENT = "application/gzip"
	MSWORD          CONTENT = "application/msword"
	MS_EXCEL        CONTENT = "application/vnd.ms-excel"
	MS_POWERPOINT   CONTENT = "application/vnd.ms-powerpoint"

	// Image types
	PNG  CONTENT = "image/png"
	JPEG CONTENT = "image/jpeg"
	GIF  CONTENT = "image/gif"
	SVG  CONTENT = "image/svg+xml"
	WEBP CONTENT = "image/webp"
	ICO  CONTENT = "image/x-icon"

	// Audio/Video types
	MP3       CONTENT = "audio/mpeg"
	OGG_AUDIO CONTENT = "audio/ogg"
	MP4       CONTENT = "video/mp4"
	WEBM      CONTENT = "video/webm"

	// Font types
	WOFF  CONTENT = "font/woff"
	WOFF2 CONTENT = "font/woff2"
	TTF   CONTENT = "font/ttf"
	OTF   CONTENT = "font/otf"
)

const HTTPVERSION = "HTTP/1.1"

type STATUS string

const (
	// 1xx Informational
	CONTINUE            STATUS = "100 Continue"
	SWITCHING_PROTOCOLS STATUS = "101 Switching Protocols"
	PROCESSING          STATUS = "102 Processing"

	// 2xx Success
	OK                STATUS = "200 OK"
	CREATED           STATUS = "201 Created"
	ACCEPTED          STATUS = "202 Accepted"
	NON_AUTHORITATIVE STATUS = "203 Non-Authoritative Information"
	NO_CONTENT        STATUS = "204 No Content"
	RESET_CONTENT     STATUS = "205 Reset Content"
	PARTIAL_CONTENT   STATUS = "206 Partial Content"

	// 3xx Redirection
	MULTIPLE_CHOICES  STATUS = "300 Multiple Choices"
	MOVED_PERMANENTLY STATUS = "301 Moved Permanently"
	FOUND             STATUS = "302 Found"
	SEE_OTHER         STATUS = "303 See Other"
	NOT_MODIFIED      STATUS = "304 Not Modified"
	TEMP_REDIRECT     STATUS = "307 Temporary Redirect"
	PERM_REDIRECT     STATUS = "308 Permanent Redirect"

	// 4xx Client Errors
	BAD_REQUEST         STATUS = "400 Bad Request"
	UNAUTHORIZED        STATUS = "401 Unauthorized"
	PAYMENT_REQUIRED    STATUS = "402 Payment Required"
	FORBIDDEN           STATUS = "403 Forbidden"
	NOT_FOUND           STATUS = "404 Not Found"
	METHOD_NOT_ALLOWED  STATUS = "405 Method Not Allowed"
	NOT_ACCEPTABLE      STATUS = "406 Not Acceptable"
	PROXY_AUTH_REQUIRED STATUS = "407 Proxy Authentication Required"
	REQUEST_TIMEOUT     STATUS = "408 Request Timeout"
	CONFLICT            STATUS = "409 Conflict"
	GONE                STATUS = "410 Gone"
	LENGTH_REQUIRED     STATUS = "411 Length Required"
	PRECONDITION_FAILED STATUS = "412 Precondition Failed"
	PAYLOAD_TOO_LARGE   STATUS = "413 Payload Too Large"
	URI_TOO_LONG        STATUS = "414 URI Too Long"
	UNSUPPORTED_MEDIA   STATUS = "415 Unsupported Media Type"
	IM_A_TEAPOT         STATUS = "418 I'm a teapot"

	// 5xx Server Errors
	INTERNAL_SERVER_ERROR      STATUS = "500 Internal Server Error"
	NOT_IMPLEMENTED            STATUS = "501 Not Implemented"
	BAD_GATEWAY                STATUS = "502 Bad Gateway"
	SERVICE_UNAVAILABLE        STATUS = "503 Service Unavailable"
	GATEWAY_TIMEOUT            STATUS = "504 Gateway Timeout"
	HTTP_VERSION_NOT_SUPPORTED STATUS = "505 HTTP Version Not Supported"
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
