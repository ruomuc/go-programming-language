package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"os"
)

// 练习5.2: 写一个函数，用于统计HTML文档树内所有元素的个数，如p、div、span等。
func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	elementCountMap := make(map[string]int)
	countElement(elementCountMap, doc)
	for ele, c := range elementCountMap {
		fmt.Printf("ele name: %s, count: %d\n", ele, c)
	}
}

func countElement(count map[string]int, node *html.Node) {
	if node == nil {
		return
	}
	if node.Type == html.ElementNode {
		count[node.Data]++
	}
	countElement(count, node.FirstChild)
	countElement(count, node.NextSibling)
}
