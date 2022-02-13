package main

import (
	"exercise5.13/links"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
)

// 练习5.13：修改crawl函数保存找到的页面，根据需要创建目录。
// 不要保存不同域名下的页面。比如，如果本来的页面来自 golang.org，
// 那么就把它们保存下来但是不要保存 vimeo.com 下的页面。

func main() {

	baseUrl := flag.String("-u", "https://www.zhihu.com/", "the url to crawl")
	flag.Parse()

	for _, u := range crawl(*baseUrl) {
		fmt.Println(u)
		download(*baseUrl, u)
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Fatalln(err)
	}
	return list
}

func download(baseUrl, url string) {
	// 如果域名不同，那么不保存
	if !strings.HasPrefix(url, baseUrl) {
		return
	}

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	var dir string
	dir = strings.TrimPrefix(url, "http://")
	dir = strings.TrimPrefix(url, "https://")
	if err := os.MkdirAll(dir, 0755); err != nil {
		log.Fatalln(err)
	}

	filename := path.Join(dir, "index.html")
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
