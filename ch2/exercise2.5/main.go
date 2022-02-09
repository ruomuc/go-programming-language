package main

import "fmt"

// 练习2.5：使用 x&(x-1) 可以清除 x 最右边的非零位，
// 利用该特点写一个 PopCount，然后评价它的性能。
func main() {
	fmt.Println(PopCount(9))
	fmt.Println(PopCount2(9))
}

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCount2(x uint64) int {
	var res int
	for x != 0 {
		x = x & (x - 1)
		res++
	}
	return res
}
