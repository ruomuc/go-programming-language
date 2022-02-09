package main

import "fmt"

// 练习4.5：编写一个就地处理函数，用于去除[]string slice 中相邻的重复字符元素。
func main() {
	strs := []string{"a", "a", "a", "b", "b", "c"}
	fmt.Println(clear(strs)) // [a b c]
}

func clear(strList []string) []string {
	for i, j := 0, 1; i < len(strList)-1; i, j = i+1, j+1 {
		if strList[i] == strList[j] {
			copy(strList[j:], strList[j+1:])
			strList = strList[:len(strList)-1]
			i--
			j--
		}
	}
	return strList
}
