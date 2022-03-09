package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	var port int
	var timeZone string
	flag.IntVar(&port, "port", 8000, "run port")
	flag.StringVar(&timeZone, "tz", "UTC", "time zone")
	flag.Parse()

	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn, timeZone)
	}
}

func handleConn(c net.Conn, timeZone string) {
	defer c.Close()
	for {
		var now string
		switch timeZone {
		case "UTC":
			now = time.Now().UTC().Format("15:04:05\n")
		case "LOCAL":
			now = time.Now().Local().Format("15:04:05\n")
		}

		_, err := io.WriteString(c, now)
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
