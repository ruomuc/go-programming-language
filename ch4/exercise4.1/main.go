package main

import (
	"crypto/sha256"
	"fmt"
)

// 练习4.1：编写一个函数，用于统计SHA256 散列中
// 不同的位数。（见 2.6.2 节的PopCount）
func main() {
	h := sha256.New()
	h.Write([]byte("count sha256"))
	fmt.Println(h.Sum(nil))
	fmt.Println(countDiffBite(h.Sum(nil)))
}

func countDiffBite(h []byte) int {
	var count int

	cp := make(map[byte]int)

	// 统计不同位的个数
	for _, b := range h {
		cp[b]++
	}

	// 统计大于2的数量
	for _, p := range cp {
		if p >= 2 {
			count++
		}
	}
	return count
}
