package main

import (
	"fmt"
	"sort"
)

// 练习3.12：编写一个函数判断两个字符串是否同文异构，
// 也就是，它们都含有相同的字符但排列顺序不同。
func main() {
	s1 := "qwerasdf"
	s2 := "asdfqwer"
	s3 := "asdfqqwewer"
	fmt.Println(isomerism(s1, s2))
	fmt.Println(isomerism(s1, s3))
}

func isomerism(s1, s2 string) bool {
	b1 := []byte(s1)
	b2 := []byte(s2)

	sort.Slice(b1, func(i, j int) bool {
		return b1[i] < b1[j]
	})

	sort.Slice(b2, func(i, j int) bool {
		return b2[i] < b2[j]
	})
	return string(b1) == string(b2)
}
