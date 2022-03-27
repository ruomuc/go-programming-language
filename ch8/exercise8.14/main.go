package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

// 练习 8.14： 修改聊天服务器的网络协议这样每一个客户端就可以在entering时可以提供它们的名字。 将消息前缀由之前的网络地址改为这个名字。
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

			cli.ch <- fmt.Sprintf("%d arrival: [%s]\n", len(names), strings.Join(names, ","))

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

	fmt.Fprint(conn, "You are: ___\b\b\b")
	sc := bufio.NewScanner(conn)
	sc.Scan()
	who := sc.Text()

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
