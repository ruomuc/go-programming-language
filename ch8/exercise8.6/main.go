package main

import (
	"crawl3/links"
	"fmt"
	"log"
	"os"
)

//练习8.6：对并发的爬虫添加深度限制。如果用户设置-depth=3，
// 那么仅最多通过三个链接可达的URL能被找到。

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

	// 创建20个爬虫goroutine来获取每个不可见链接
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenWorks {
				foundLinks := crawl(link)
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

func crawl(w work) []work {
	fmt.Printf("url: %s, depth: %d\n", w.link, w.depth)
	if w.depth > 3 {
		return nil
	}

	list, err := links.Extract(w.link)
	if err != nil {
		log.Print(err)
	}

	var works []work
	for _, l := range list {
		works = append(works, work{l, w.depth + 1})
	}
	return works
}
