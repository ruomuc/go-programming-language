package main

import "fmt"

// 练习5.19：使用 panic和 recover 写一个函数，它没有 return
// 语句，但是能够返回一个非零的值。
func main() {
	fmt.Println(panicFn())
}

func panicFn() (res int) {
	defer func() {
		res = 3
		recover()
		// recover 后会直接 return
	}()
	panic("aaa")
}
