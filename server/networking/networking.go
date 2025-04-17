package networking

import (
	"log"
	"net"
	"os"
	"webserver/server/http"
	"webserver/server/router"
)

// This should be set by config
// tcp could be hardcoded, unless implementing http3
const (
	HOST = "localhost"
	PORT = "8080"
	TYPE = "tcp"
)

type Port int

type ListenerType = func(port Port, rcs []router.ControlledRoutes)

func Listen(port Port, rcs []router.ControlledRoutes) {
	for {
		initListener(rcs)
	}
}

func initListener(rcs []router.ControlledRoutes) {
	listen, err := net.Listen(TYPE, HOST+":"+PORT)

	if err != nil {
		os.Exit(1)
	}

	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		go handleRequest(conn, rcs)
	}
}

func handleRequest(conn net.Conn, rcs []router.ControlledRoutes) {
	buffer := make([]byte, 1024)
	_, err := conn.Read(buffer)

	if err != nil {
		log.Fatal(err)
	}

	req, err := http.ParseRequest(string(buffer))
	var res http.Response
	if err != nil {
		res = http.Response{Status: http.BAD_REQUEST, Body: "400 Bad Request"}
	} else {
		res = router.Router(req, rcs)
	}

	conn.Write([]byte(res.Serialize() + "\n"))
	conn.Close()
}
