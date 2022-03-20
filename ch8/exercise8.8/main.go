package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

// 练习 8.8： 使用select来改造8.3节中的echo服务器，为其增加超时，
// 这样服务器可以在客户端10秒中 没有任何喊话时自动断开连接。
func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	defer c.Close()
	input := bufio.NewScanner(c)

	ch := make(chan struct{})
	go func() {
		for input.Scan() {
			ch <- struct{}{}
		}
	}()

	for {
		select {
		case <-time.After(10 * time.Second):
			return
		case <-ch:
			go echo(c, input.Text(), 1*time.Second)
		}
	}
}
