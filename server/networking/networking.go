package networking

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"
	"webserver/server/http"
)

// This should be set by config
// tcp could be hardcoded, unless implementing http3
const (
	HOST = "localhost"
	PORT = "8080"
	TYPE = "tcp"
)

type Port int

type ListenerType = func(port Port)

func Listen(port Port) {
	var requestUrl http.URL
	for {
		fmt.Print("Make request: ")
		fmt.Scan(&requestUrl)
	}
}

func init() {
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
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	buffer := make([]byte, 1024)
	_, err := conn.Read(buffer)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(buffer))

	time := time.Now()
	responseStr := fmt.Sprintf("Your message was received at: %v\n", time)
	conn.Write([]byte(responseStr))
	conn.Close()
}
