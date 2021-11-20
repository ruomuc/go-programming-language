package main

import (
	"bytes"
	"fmt"
	"strings"
)

// 练习3.10：编写一个非递归的 comma 函数，
// 运用 bytes.Buffer，而不是简单的字符串拼接
func main() {
	var s = "1234567890"
	fmt.Println(comma(s))
}

func comma(s string) string {
	var res []string
	buf := bytes.NewBufferString(s)

	for buf.Len() > 0 {
		b := buf.Next(3)
		res = append(res, string(b))
	}

	return strings.Join(res, ",")
}
