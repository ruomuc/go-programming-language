package main

import (
	"errors"
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"os"
)

// 练习5.7：开发startElement和endElement函数并应用到
// 一个通用的HTML输出代码中。输出注释节点、文本节点和所有
// 元素属性（<a href='...'>）。当一个元素没有子节点时，
// 使用简短的形式，比如<img/> 而不是 <img></img>。写一个
// 测试程序保证输出可以正确解析（参考第11章）。
func main() {
	url := os.Args[1]
	id := os.Args[2]

	n, err := outline(url, id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(n.Data, ":")
	for _, attr := range n.Attr {
		fmt.Printf("%s = %s\n", attr.Key, attr.Val)
	}
}

func outline(url, id string) (*html.Node, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	node := ElementByID(doc, id)
	if node == nil {
		return nil, errors.New("not found")
	}
	return node, nil
}

func ElementByID(node *html.Node, id string) *html.Node {
	return forEachNode(node, id, findElement, findElement)
}

type findElementFn func(n *html.Node, id string) bool

func forEachNode(node *html.Node, id string, pre, post findElementFn) *html.Node {
	if !pre(node, id) {
		return node
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		if n := forEachNode(c, id, pre, post); n != nil {
			return n
		}
	}

	if !post(node, id) {
		return node
	}

	return nil
}

// 返回值代表是否继续查找
func findElement(n *html.Node, id string) bool {
	if n.Type == html.ElementNode {
		for _, attr := range n.Attr {
			if attr.Key == "id" && attr.Val == id {
				return false
			}
		}
	}
	return true
}
