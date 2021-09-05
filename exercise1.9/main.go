package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

// 1.9 修改 fetch来输出 HTTP 的状态码，可以在 resp.Status 中找到它。
func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Http status: ", resp.Status)
	}
}
