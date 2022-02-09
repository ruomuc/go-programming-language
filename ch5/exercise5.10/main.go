package main

import (
	"fmt"
)

// 练习5.10：重写 topoSort 以使用 map 代替 slice 并去掉开头的排序。
// 结果不是唯一的，验证这个结果是合法的拓扑排序。

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"database":              {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	res := topoSort(prereqs)
	for i := 0; i < len(res); i++ {
		fmt.Printf("%d: \t%s\n", i+1, res[i])
	}
}

func topoSort(m map[string][]string) map[int]string {
	order := make(map[int]string)
	seen := make(map[string]bool)

	var visitAll func(items []string)
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order[len(order)] = item
			}
		}
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	visitAll(keys)
	return order
}
