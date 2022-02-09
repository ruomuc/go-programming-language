package main

import "fmt"

// 练习4.3：重写reverse函数，使用数组指针代替slice。
// tips:  在 golang 中，建立了 arrPtr := &arr 这种类似地址关系后，* 允许不写。
func main() {
	nums := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	reverse(&nums)
	fmt.Println(nums) // [10 9 8 7 6 5 4 3 2 1]
}

func reverse(nums *[10]int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
