package main

import "fmt"

// 练习2.4：写一个用于统计位的 PopCount，它在其实际参数的64位上执行移位操作
// 每次判断最右边的位，进而实现统计功能。把它与快查表版本的性能进行对比。
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
	for i := 0; x != 0; x = x >> 1 {
		if x&1 == 1 {
			i++
		}
		res = i
	}
	return res
}
