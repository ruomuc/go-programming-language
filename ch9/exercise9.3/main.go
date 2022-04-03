package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"memo/memo"
	"net/http"
	"time"
)

// 练习 9.3： 扩展Func类型和(*Memo).Get方法，支持调用方提供一个可选的done channel，
// 使其具备通 过该channel来取消整个操作的能力(§8.9)。一个被取消了的Func的调用结果不应该被缓存。
const timeout = 5 * time.Minute

func main() {
	m := memo.New(httpGetBody)
	d := make(chan struct{})
	go func() {
		time.Sleep(timeout)
		close(d)
	}()
	data, err := m.Get("https://www.baidu.com", d)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(data)
}

func httpGetBody(url string, done <-chan struct{}) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
