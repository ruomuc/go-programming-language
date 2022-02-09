package main

import (
	"fmt"
	"os"
	"strings"
)

// 1.1 修改 echo 程序输出 os.Args[0],即参数的名字
func main() {
	fmt.Println(os.Args[0])
	fmt.Println(strings.Join(os.Args, " "))
}
