package main

import "fmt"

// 练习4.3：重写reverse函数，使用数组指针代替slice。
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
