package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// 1.7 函数 io.Copy(dst,src) 从src读，并写入dst。使用它代替 ioutil.ReadAll
// 来复制响应内容到 os.Stdout, 这样不需要装下整个响应数据流的缓冲区。确保检查
// io.Copy 返回的错误结果
func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}
