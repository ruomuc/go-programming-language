package main

import (
	"strings"
)

// 1.3 尝试测量可能低效的程序 和 使用strings.Join 的程序在执行时间上的差异。
func main() {
}

func normal(args []string) {
	var s, sep string
	for _, arg := range args {
		s = sep + arg
		sep = " "
	}
	_ = s
	// fmt.Println(s)
}

func strJoin(args []string) {
	s := strings.Join(args, " ")
	_ = s
	// fmt.Println(s)
}
