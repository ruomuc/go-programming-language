package main

import (
	"fmt"
	"os"
)

// 1.2 修改 echo 程序，输出参数的索引和值，每行一个
func main() {
	for idx, val := range os.Args {
		fmt.Printf("index: %d, value: %s\n", idx, val)
	}
}
