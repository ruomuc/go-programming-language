package main

import "math/rand"

// 练习7.3：为gopl.io/ch4/treesort中的*tree类型（见4.4节）
// 写一个String方法，用于展示其中的值序列
func main() {
	data := make([]int, 50)
	for i := range data {
		data[i] = rand.Int() % 50
	}
	Sort(data)
}
