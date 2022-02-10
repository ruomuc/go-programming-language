package main

import (
	"fmt"
	"log"
	"sort"
)

// 练习5.11：现在有“线性代数”（linear algebra）这门课程，
// 它的先决课程是微积分（calculus）。扩展 topoSort 以函数输出结果。

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
	"linear algebra":        {"calculus"},
}

func main() {
	res, err := topoSort(prereqs)
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < len(res); i++ {
		fmt.Printf("%d: %s\n", i+1, res[i])
	}
}

func topoSort(m map[string][]string) (order []string, err error) {
	seen := make(map[string]bool)

	var visitAll func(items []string) error
	visitAll = func(items []string) error {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				if err := visitAll(m[item]); err != nil {
					return err
				}
				order = append(order, item)
			} else {
				return fmt.Errorf("has cycle: %s", item)
			}
		}
		return nil
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	err = visitAll(keys)
	return
}
