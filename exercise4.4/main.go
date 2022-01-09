package main

import "fmt"

// 练习4.4：编写一个rotate，实现一次遍历就可以完成元素旋转。
func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(rotate(nums, len(nums)-1)) // [10 9 8 7 6 5 4 3 2 1]
}

func rotate(nums []int, r int) []int {
	n := len(nums)
	// 声明一个切片,用于存储旋转后的结果
	ans := make([]int, n, n)
	for i := 0; i < n; i++ {
		index := i + r
		if index >= n {
			index = index - n
		}
		ans[i] = nums[index] // 这个是向左旋转
		// ans[index] = nums[i] // 这个是向右旋转
	}
	return ans
}
