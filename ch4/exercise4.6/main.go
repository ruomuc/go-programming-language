package main

import (
	"fmt"
	"unicode"
)

// 练习4.6：编写一个就地处理函数，用于将一个 UTF-8编码的
// 字节 slice 中所有相邻的 Unicode 空白字符（unicode.IsSpace）缩减为一个ASCII空白字符。
func main() {
	strs := []byte{228, 184, 173, 32, 32, 33, 33, 229, 155, 189} // 中<空格><空格>!!国
	fmt.Println(string(clear(strs)))                             // 中<空格>!!国
}

func clear(strList []byte) []byte {
	for i, j := 0, 1; i < len(strList)-1; i, j = i+1, j+1 {
		if strList[i] == strList[j] && unicode.IsSpace(rune(strList[i])) {
			copy(strList[j:], strList[j+1:])
			strList = strList[:len(strList)-1]
			i--
			j--
		}
	}
	return strList
}
