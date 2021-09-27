package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strings"
	"time"
)

// 1.10 找一个数据量比较大的网站，用本小节中的程序调研网站的缓存策略，
// 对每个URL执行两遍 请求，查看两次时间是否有较大的差别，并且每次获取
// 到的响应内容是否一致，修改本节中的程序，将 响应结果输出，以便于进行对比。
func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	if !strings.HasPrefix(url, "http://") {
		url = "http://" + url
	}
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	fileName := path.Join("./", strings.Split(url, ".")[1])
	f, _ := os.Create(fileName)
	nbytes, err := io.Copy(f, resp.Body)

	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s : %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
