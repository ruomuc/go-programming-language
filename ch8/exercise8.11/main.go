package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// 练习 8.11： 紧接着8.4.4中的mirroredQuery流程，实现一个
// 并发请求url的fetch的变种。当第一个请求返回时，直接取消其它的请求。

func main() {
	canceled := make(chan struct{})
	resps := make(chan string, len(os.Args[1:]))

	for _, arg := range os.Args[1:] {
		go func(url string) {
			resps <- fetch(url, canceled)
		}(arg)
	}
	resp := <-resps
	close(canceled)
	fmt.Println(resp)
}

func fetch(url string, canceled chan struct{}) string {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return ""
	}
	req.Cancel = canceled

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("%s", b)
}
