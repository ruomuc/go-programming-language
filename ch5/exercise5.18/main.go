package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"path"
)

// 练习5.18：不改变原本的行为，重写fetch函数以使用defer语句关闭打开的可写文件。
func main() {
	_, _, err := fetch("https://www.zhihu.com")
	if err != nil {
		log.Fatalln(err)
	}
}

func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}

	f, err := os.Create(local + ".html")
	defer func() {
		closeErr := f.Close()
		if err == nil {
			err = closeErr
		}
	}()
	n, err = io.Copy(f, resp.Body)
	return
}
