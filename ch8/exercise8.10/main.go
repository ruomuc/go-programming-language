package main

import (
	"fmt"
	"log"
	"os"

	"exercise8.10/links"
)

// 练习 8.10： HTTP请求可能会因http.Request结构体中Cancel channel的关闭而取消。
// 修改8.6节中的 web crawler来支持取消http请求。（提示：http.Get并没有提供方便
// 地定制一个请求的方法。你可以用 http.NewRequest来取而代之，设置它的Cancel字段，
// 然后用http.DefaultClient.Do(req)来进行这个 http请求。）

type work struct {
	link  string
	depth int
}

func main() {
	worklist := make(chan []work)
	unseenWorks := make(chan work)

	// 向任务列表添加命令行参数
	go func() {
		var works []work
		for _, url := range os.Args[1:] {
			works = append(works, work{url, 1})
		}
		worklist <- works
	}()

	canceled := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(canceled)
	}()

	// 创建20个爬虫goroutine来获取每个不可见链接
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenWorks {
				foundLinks := crawl(link, canceled)
				go func() { worklist <- foundLinks }()
			}
		}()
	}

	seen := make(map[string]bool)
	for works := range worklist {
		for _, w := range works {
			if !seen[w.link] {
				seen[w.link] = true
				unseenWorks <- w
			}
		}
	}
}

func crawl(w work, canceled chan struct{}) []work {
	fmt.Printf("url: %s, depth: %d\n", w.link, w.depth)
	if w.depth > 3 {
		return nil
	}

	list, err := links.Extract(w.link, canceled)
	if err != nil {
		log.Print(err)
	}

	var works []work
	for _, l := range list {
		works = append(works, work{l, w.depth + 1})
	}
	return works
}
