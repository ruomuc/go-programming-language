package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"os"
)

// 练习5.12：5.5节（gopl.io/ch5/outline2）的startElement
// 和endElement函数共享一个全局变量depth。把它们变为匿名函数
// 以共享outline函数的一个局部变量。

func main() {
	for _, url := range os.Args[1:] {
		if err := outline(url); err != nil {
			log.Println(err)
		}
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

	var depth int

	forEachNode(doc, func(n *html.Node) {
		if n.Type == html.ElementNode {
			fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
			depth++
		}
	}, func(n *html.Node) {
		if n.Type == html.ElementNode {
			depth--
			fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		}
	})

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
