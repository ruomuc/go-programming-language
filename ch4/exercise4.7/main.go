package main

import (
	"fmt"
)

// 练习4.7：修改函数reverse，来翻转一个UTF-8编码的字符串中的字符元素，
// 传入参数是该字符串对应的字节slice类型（[]byte）。你可以做到不需要重新分配内存就实现该功能吗？
func main() {
	strs := "asadqweq"
	fmt.Println(string(reverse([]byte(strs))))
}

func reverse(strs []byte) []byte {
	for i, j := 0, len(strs)-1; i < j; i, j = i+1, j-1 {
		strs[i], strs[j] = strs[j], strs[i]
	}
	return strs
}
