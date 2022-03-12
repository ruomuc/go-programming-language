package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

// 练习8.1：修改clock2来接收一个端口号，写一个程序 clockwall,作为
// 多个时钟服务器的客户端，读取每一个服务器的时间，类似于不同地区办
// 公室的时钟，然后显示在一个表中。如果可以访问不同地域的计算机，可以
// 远程运行示例程序；否则可以伪装不同的时区，在不同的端口上运行。
func main() {
	servers := parseArgs(os.Args[1:])

	for _, s := range servers {
		conn, err := net.Dial("tcp", s.address)
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()

		go func(s *server) {
			sc := bufio.NewScanner(conn)
			for sc.Scan() {
				s.message = sc.Text()
			}
		}(s)
	}

	for {
		fmt.Println()
		for _, s := range servers {
			fmt.Printf("%s : %s\n", s.name, s.message)
		}
		fmt.Print("---")
		time.Sleep(time.Second)
	}
}

type server struct {
	name    string
	address string
	message string
}

func parseArgs(args []string) []*server {
	var servers []*server
	for _, arg := range args {
		s := strings.SplitN(arg, "=", 2)
		servers = append(servers, &server{s[0], s[1], ""})
	}
	return servers
}
