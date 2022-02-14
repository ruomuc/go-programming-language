package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
)

// 练习5.17：写一个变长函数 ElementByTagName，已知一个HTML
// 节点树和零个或多个名字，返回所有符合给出名字的元素。下面有两个
// 示例调用：
// func ElementByTagName(doc *html.Node, name ...string) []*html.Node
// images := ElementByTagName)(doc, "img")
// headings := ElementByTagName(doc, "h1", "h2“, "h3”, "h4“)

func main() {
	resp, err := http.Get("https://www.zhihu.com")
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	nodes := ElementByTagName(doc, "a", "script")
	for _, n := range nodes {
		fmt.Printf("%s = %s\n", n.Data, n.Attr)
	}
}

func ElementByTagName(doc *html.Node, name ...string) []*html.Node {
	var nodes []*html.Node
	if doc == nil {
		return nodes
	}
	if doc.Type == html.ElementNode {
		for _, n := range name {
			if doc.Data == n {
				nodes = append(nodes, doc)
			}
		}
	}
	nodes = append(nodes, ElementByTagName(doc.FirstChild, name...)...)
	nodes = append(nodes, ElementByTagName(doc.NextSibling, name...)...)
	return nodes
}
