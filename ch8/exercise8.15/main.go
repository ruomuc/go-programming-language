package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

// 练习 8.15： 如果一个客户端没有及时地读取数据可能会导致所有的客户端被阻塞。
// 修改broadcaster来 跳过一条消息，而不是等待这个客户端一直到其准备好写。或者
// 为每一个客户端的消息发出channel建立 缓冲区，这样大部分的消息便不会被丢掉；
// broadcaster应该用一个非阻塞的send向这个channel中发消 息。
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
				select {
				case cli.ch <- msg:
				case <-time.After(10 * time.Second): // 超过10s，跳过给该客户端发送消息
				}
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
