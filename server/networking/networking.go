package networking

import (
	"log"
	"net"
	"webserver/server/http"
	"webserver/server/router"
)

// This should be set by config
// tcp could be hardcoded, unless implementing http3
const (
	HOST = "localhost"
	TYPE = "tcp"
)

type Port int

type ListenerType = func(
	port string,
	rcs []router.ControlledRoutes)

func Listen(port string, rcs []router.ControlledRoutes) {
	listener, err := net.Listen(TYPE, HOST+":"+port)
	if err != nil {
		log.Fatalf("Failed to listen on %s:%s: %v", HOST, port, err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Accept error: %v", err)
			continue
		}
		go handleRequest(conn, rcs)
	}
}

func handleRequest(conn net.Conn, rcs []router.ControlledRoutes) {
	defer conn.Close()

	buffer := make([]byte, 4096)
	n, err := conn.Read(buffer)
	if err != nil {
		log.Printf("Read error: %v", err)
		return
	}

	req, err := http.ParseRequest(string(buffer[:n]))
	var res http.Response
	if err != nil {
		res = http.Response{Status: http.BAD_REQUEST, Body: "400 Bad Request", Content: http.PLAIN}
	} else {
		res = router.Router(req, rcs)
	}

	_, writeErr := conn.Write([]byte(res.Serialize()))
	if writeErr != nil {
		log.Printf("Write error: %v", writeErr)
	}
}
