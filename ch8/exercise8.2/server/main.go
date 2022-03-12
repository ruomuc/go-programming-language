package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"path"
	"strings"
)

func main() {
	conn, err := net.Listen("tcp", "localhost:8001")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := conn.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	sc := bufio.NewScanner(conn)
	cwd := "."

CLOSE:

	for sc.Scan() {
		args := strings.Fields(sc.Text())
		cmd := args[0]
		switch cmd {
		case "close":
			break CLOSE
		case "ls":
			if len(args) < 2 {
				ls(conn, cwd)
			} else {
				path := args[1]
				if err := ls(conn, path); err != nil {
					fmt.Fprint(conn, err)
				}
			}
		case "cd":
			if len(args) < 2 {
				fmt.Fprintln(conn, "not enough argument")
			} else {
				cwd += "/" + args[1]
			}
		case "get":
			if len(args) < 2 {
				fmt.Fprintln(conn, "not enough argument")
			} else {
				filename := args[1]
				data, err := ioutil.ReadFile(path.Join(cwd, filename))
				if err != nil {
					fmt.Fprint(conn, err)
				}
				fmt.Fprintf(conn, "%s\n", data)
			}
		}
	}
}

func ls(w io.Writer, dir string) error {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}
	for _, file := range files {
		fmt.Fprintf(w, "%s\n", file.Name())
	}
	return nil
}
