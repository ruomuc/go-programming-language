package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

// 练习5.4：扩展visit函数，使之能够获得到其他种类的链接地址，
// 比如图片、脚本或样式表的链接。
func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	linkMap := make(map[string][]string)
	visit(linkMap, doc)
	for linkType, links := range linkMap {
		fmt.Printf("type: %s: \n", linkType)
		for _, l := range links {
			fmt.Println(l)
		}
	}
}

func visit(links map[string][]string, n *html.Node) {
	if n == nil {
		return
	}

	if n.Type == html.ElementNode {
		switch n.Data {
		case "a":
			for _, a := range n.Attr {
				if a.Key == "href" {
					links[n.Data] = append(links[n.Data], a.Val)
				}
			}
		case "img":
			for _, s := range n.Attr {
				if s.Key == "src" {
					links[n.Data] = append(links[n.Data], s.Val)
				}
			}
		case "script":
			for _, s := range n.Attr {
				if s.Key == "src" {
					links[n.Data] = append(links[n.Data], s.Val)
				}
			}
		case "link":
			for _, m := range n.Attr {
				if m.Key == "media" {
					links[n.Data] = append(links[n.Data], m.Val)
				}
			}
		}
	}
	visit(links, n.FirstChild)
	visit(links, n.NextSibling)
}
