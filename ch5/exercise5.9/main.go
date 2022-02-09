package main

import (
	"fmt"
	"strings"
)

// 练习5.9：写一个函数expand(s string, f func(string)string)string
// 该函数替换参数 s 中的每一个子字符串 ”$foo“ 为 f("foo")的返回值。
func main() {
	s := "wqe$foowqeqe$foo"
	res := expand(s, replace)
	fmt.Println(res)
}

func expand(s string, f func(string) string) string {
	return strings.Replace(s, "$foo", f("foo"), 100)
}

func replace(s string) string {
	return s + "next"
}
