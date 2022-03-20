package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"

	"exercise8.7/links"
)

// 练习 8.7： 完成一个并发程序来创建一个线上网站的本地镜像，
// 把该站点的所有可达的页面都抓取到本 地硬盘。为了省事，我们
// 这里可以只取出现在该域下的所有页面(比如golang.org结尾，译注
// ：外链的应 该就不算了。)当然了，出现在页面里的链接你也需要进行
// 一些处理，使其能够在你的镜像站点上进行跳 转，而不是指向原始的链接。
var (
	wg      sync.WaitGroup
	baseUrl = flag.String("base", "https://go.dev/", "base url")
)

func main() {
	flag.Parse()
	for _, url := range crawl(*baseUrl) {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			download(url)
		}(url)
	}

	done := make(chan struct{})
	go func() {
		wg.Wait()
		done <- struct{}{}
	}()
	<-done
}

func crawl(url string) []string {
	list, err := links.Extract(url)
	if err != nil {
		log.Fatal(err)
	}
	return list
}

func download(url string) {
	if !strings.HasPrefix(url, *baseUrl) {
		return
	}

	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	dir := strings.TrimPrefix(strings.TrimPrefix(url, "http://"), "https://")
	if err := os.MkdirAll(dir, 0755); err != nil {
		log.Fatalln(err)
	}

	filename := dir + "index.html"
	f, err := os.Create(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	_, err = io.Copy(f, resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
}
