package main

import (
	"fmt"
	"strings"
)

// 练习5.16：写一个变长版本的 Strings.Join 函数。

func main() {
	var strArr = [...]string{"hello", "world", "!", "\n", "how", "are", "you", "?"}
	strs := make([]string, 0)
	for _, s := range strArr {
		strs = append(strs, s)
	}
	fmt.Println(stringJoin(strs...))
}

func stringJoin(strs ...string) string {
	if len(strs) == 0 {
		return ""
	}
	var builder strings.Builder
	for _, s := range strs {
		builder.WriteString(s)
	}
	return builder.String()
}
