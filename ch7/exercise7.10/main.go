package main

import (
	"fmt"
	"sort"
)

// 练习7.10：sort.Interface 也可以用于其他用途。试写一个函数
// IsPalindrome(s sort.Interface)bool 来判断一个序列是否回文，
// 即序列反转后是否保持不变。可以假定对于下标分别为 i、j的元素，如果
// !s.Less(i,j) && !s.Less(j,i)，那么两个元素相等。

func main() {
	s := c("abccba")
	fmt.Println(IsPalindrome(s))
}

func IsPalindrome(s sort.Interface) bool {
	for i, j := 0, s.Len()-1; i < j; i, j = i+1, j-1 {
		if !s.Less(i, j) && !s.Less(j, i) {
			continue
		}
		return false
	}
	return true
}

type c []rune

func (c c) Len() int {
	return len(c)
}

func (c c) Less(i, j int) bool {
	return c[i] < c[j]
}

func (c c) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
