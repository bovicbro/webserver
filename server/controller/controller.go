package controller

import (
	"os"
	"strings"

	. "webserver/server/http"
)

type Controller = func(req Request, res Response) Response

// Static returns a Controller that serves files from the given directory.
// The request URL is mapped directly to files within basePath.
// If the URL has no file extension, a .html extension is tried automatically.
func Static(basePath string) Controller {
	return func(req Request, res Response) Response {
		path := strings.TrimPrefix(string(req.Url), "/")
		if path == "" {
			path = "index.html"
		}

		// Prevent directory traversal
		if strings.Contains(path, "..") {
			return Response{Status: FORBIDDEN, Content: PLAIN, Body: "403 Forbidden"}
		}

		filePath := basePath + "/" + path

		// If no extension, try .html fallback
		if !strings.Contains(path, ".") {
			htmlPath := filePath + ".html"
			if _, err := os.Stat(htmlPath); err == nil {
				filePath = htmlPath
			}
		}

		content, err := os.ReadFile(filePath)
		if err != nil {
			return Response{Status: NOT_FOUND, Content: HTML, Body: "<h1>404 Not Found</h1>"}
		}

		return Response{
			Body:    string(content),
			Status:  OK,
			Content: detectContentType(filePath),
		}
	}
}

func detectContentType(path string) CONTENT {
	switch {
	case strings.HasSuffix(path, ".css"):
		return CSS
	case strings.HasSuffix(path, ".js"):
		return JAVASCRIPT
	case strings.HasSuffix(path, ".json"):
		return JSON
	case strings.HasSuffix(path, ".png"):
		return PNG
	case strings.HasSuffix(path, ".jpg"), strings.HasSuffix(path, ".jpeg"):
		return JPEG
	case strings.HasSuffix(path, ".gif"):
		return GIF
	case strings.HasSuffix(path, ".svg"):
		return SVG
	case strings.HasSuffix(path, ".html"), strings.HasSuffix(path, ".htm"):
		return HTML
	default:
		return PLAIN
	}
}
