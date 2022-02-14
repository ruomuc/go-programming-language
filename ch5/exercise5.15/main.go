package main

import (
	"errors"
	"fmt"
	"log"
)

// 练习5.15：模仿sum写两个变长函数 max 和 min。当
// 不带任何参数调用这些函数时应该怎么应对？编写类似函数
// 的变种，要求至少需要一个参数。

func main() {
	var a = [...]int{1, 34, 56, 6, 13, 231, 512, 321, 51, 23, 5}
	nums := make([]int, 0)
	for _, n := range a {
		nums = append(nums, n)
	}

	var (
		res int
		err error
	)
	res, err = max(nums...)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("max of nums: ", res)
	res, err = min(nums...)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("min of nums: ", res)

	fmt.Println("mustParamMax: ", mustParamMax(1111, nums...))
	fmt.Println("mustParamMin: ", mustParamMin(-11, nums...))
}

func max(nums ...int) (int, error) {
	if len(nums) == 0 {
		return -1, errors.New("must have at least one parameter")
	}
	max := nums[0]
	for _, n := range nums {
		if n > max {
			max = n
		}
	}
	return max, nil
}

func min(nums ...int) (int, error) {
	if len(nums) == 0 {
		return -1, errors.New("must have at least one parameter")
	}
	min := nums[0]
	for _, n := range nums {
		if n < min {
			min = n
		}
	}
	return min, nil
}

// 必须携带一个参数的变种 max 和 min
func mustParamMax(num int, nums ...int) (max int) {
	max = num
	if len(nums) == 0 {
		return
	}

	for _, n := range nums {
		if n > max {
			max = n
		}
	}
	return
}

func mustParamMin(num int, nums ...int) (min int) {
	min = num
	if len(nums) == 0 {
		return
	}

	for _, n := range nums {
		if n < min {
			min = n
		}
	}
	return
}
