package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8001")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	go mustCopy(os.Stdout, conn)

	sc := bufio.NewScanner(os.Stdin)

CLOSE:
	for sc.Scan() {
		args := strings.Fields(sc.Text())
		cmd := args[0]
		switch cmd {
		case "close":
			fmt.Fprintln(conn, sc.Text())
			break CLOSE
		case "ls", "cd", "get":
			fmt.Fprintln(conn, sc.Text())
		default:
			fmt.Printf("unexpected command %s\n", cmd)
		}
	}
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
