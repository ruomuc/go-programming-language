package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
)

// 练习5.7：开发startElement和endElement函数并应用到
// 一个通用的HTML输出代码中。输出注释节点、文本节点和所有
// 元素属性（<a href='...'>）。当一个元素没有子节点时，
// 使用简短的形式，比如<img/> 而不是 <img></img>。写一个
// 测试程序保证输出可以正确解析（参考第11章）。

func main() {
	for _, url := range os.Args[1:] {
		outline(url)
	}
}

func outline(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	forEachNode(doc, startElement, endElement)

	return nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

var depth int

func startElement(n *html.Node) {
	switch n.Type {
	case html.ElementNode:
		fmt.Printf("%*s<%s", depth*2, "", n.Data)
		for _, a := range n.Attr {
			fmt.Printf(" %s=%q", a.Key, a.Val)
		}
		// 判断是否有子节点
		if n.FirstChild == nil {
			fmt.Println("/>")
		} else {
			fmt.Println(">")
		}
		depth++
	case html.CommentNode:
		fmt.Printf("%*s<!--%s-->\n", depth*2, "", n.Data)
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		if n.FirstChild != nil {
			fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		}
	}
}
