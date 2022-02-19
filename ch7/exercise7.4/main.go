package main

import (
	"bytes"
	"fmt"
	"golang.org/x/net/html"
	"io"
	"log"
)

// 练习7.4：strings.NewReader函数输入一个字符串，返回一个从
// 字符串读取数据且满足 io.Reader 接口（也满足其他接口）的值。
// 请自己实现该函数，并且通过它来让 HTML 分析器（参考5.2节）
// 支持以字符串作为输入

var htmlText = `<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<title>菜鸟教程(runoob.com)</title>
</head>
<body>
<h1>我的第一个标题</h1>
<p>我的第一个段落。</p>
</body>
</html>`

func main() {
	doc, err := html.Parse(NewReader(htmlText))
	if err != nil {
		log.Fatalln(err)
	}

	content := ""
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode {
			var buf bytes.Buffer
			html.Render(&buf, n)
			content = buf.String()
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	fmt.Println(content)
}

func NewReader(s string) io.Reader {
	sr := StringReader(s)
	return &sr
}

type StringReader string

func (s *StringReader) Read(p []byte) (n int, err error) {
	n = copy(p, *s)
	return n, io.EOF
}
