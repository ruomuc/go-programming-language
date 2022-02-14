package main

import (
	"fmt"
	"sort"
)

// 练习5.14：使用广度优先遍历搜索一个不同的拓扑结构。
// 比如，你可以借鉴拓扑排序的例子（有向图）里的课程依
// 赖关系，计算机文件系统的分层结构（树形结构），或者从
// 当前城市的官网上下载公共汽车或者地铁线路图（无向图）。

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
	var keys []string
	for key := range prereqs {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	bfs(keys)
}

// 广度优先遍历
func bfs(keys []string) {
	var queue []string
	seen := make(map[string]bool)

	queue = append(queue, keys...)
	for len(queue) > 0 {
		// 出队
		cur := queue[0]
		// 判断是否遍历过
		if seen[cur] {
			queue = queue[1:]
			continue
		}
		fmt.Println(cur)
		seen[cur] = true
		children := prereqs[cur]
		// 孩子结点入队
		queue = append(queue, children...)
		queue = queue[1:]
	}
}
