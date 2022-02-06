package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"os"
)

// 练习5.3: 写一个函数，用于输出HTML文档树中所有文本节点的内容。
// 但不包括<script>或<style>元素，因为这些内容在Web浏览器中是不可见的。
func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	elementCountMap := make(map[string]int)
	countTextElement(elementCountMap, doc)
	for ele, c := range elementCountMap {
		fmt.Printf("ele name: %s, count: %d\n", ele, c)
	}
}

func countTextElement(count map[string]int, node *html.Node) {
	if node == nil {
		return
	}
	if node.Type == html.TextNode && node.Data != "script" && node.Data != "style" {
		count[node.Data]++
	}
	countTextElement(count, node.FirstChild)
	countTextElement(count, node.NextSibling)
}
