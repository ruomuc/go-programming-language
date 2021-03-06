package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

// 练习 8.13： 使聊天服务器能够断开空闲的客户端连接，
// 比如最近五分钟之后没有发送任何消息的那些 客户端。
// 提示：可以在其它goroutine中调用conn.Close()来解除
// Read调用，就像input.Scanner()所做 的那样。
func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConn(conn)
	}
}

type client struct {
	ch   chan<- string
	name string
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
)

func broadcaster() {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messages:
			for cli := range clients {
				cli.ch <- msg
			}
		case cli := <-entering:
			clients[cli] = true

			var names []string
			for c := range clients {
				names = append(names, c.name)
			}

			cli.ch <- fmt.Sprintf("%d arrival: %v\n", len(names), names)

		case cli := <-leaving:
			delete(clients, cli)
			close(cli.ch)
		}
	}
}

const timeout = 5 * time.Minute

func handleConn(conn net.Conn) {
	ch := make(chan string) // 对外发送客户消息的通道
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "You are" + who
	messages <- who + " has arrived"
	entering <- client{ch, who}

	// 定时器，超时会自动断开连接
	timer := time.NewTicker(timeout)
	go func() {
		<-timer.C
		conn.Close()
	}()

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
		timer.Reset(timeout)
	}

	leaving <- client{ch, who}
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // 注意，忽略网络层面的错误
	}
}
