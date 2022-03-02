package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"strings"
)

// 练习7.18：使用基于标记的解析API，写一个程序来读入一
// 个任意的XML文档，构造出一颗树来展现XML中的主要节点。
// 节点包括两种类型：CharData节点表示文本字符串，Element
// 节点表示元素及属性。每个元素节点包含它的子节点数组。
// 可参考如下类型定义：
// import "encoding/xml"
// type Node interface {} // CharData 或 *Element
// type CharData string
// type Element struct {
//      Type     xml.Name
//		  Attr	   []xml.Attr
//			Children []Node
// }

const testXml = `
<note ml-update="aware" ml-stage="preload">
<to>Tove</to>
<from>Jani</from>
<heading>Reminder</heading>
<body>Don't forget me this weekend!</body>
</note>
`

func main() {
	node, err := parse(xml.NewDecoder(strings.NewReader(testXml)))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", node)
}

type Node interface {
	fmt.Stringer
} // CharData 或 *Element

type CharData string

func (c CharData) String() string {
	return string(c)
}

type Element struct {
	Type     xml.Name
	Attr     []xml.Attr
	Children []Node
}

func (ele *Element) String() string {
	var attrs, children string
	for _, attr := range ele.Attr {
		attrs += fmt.Sprintf(" %s=%q", attr.Name.Local, attr.Value)
	}
	for _, child := range ele.Children {
		children += child.String()
	}
	return fmt.Sprintf("<%s%s>%s</%s>", ele.Type.Local, attrs, children, ele.Type.Local)
}

func parse(dec *xml.Decoder) (Node, error) {
	var stack []*Element
	for {
		tok, err := dec.Token()
		if err != nil {
			return nil, err
		}

		switch tok := tok.(type) {
		case xml.StartElement:
			e := &Element{tok.Name, tok.Attr, []Node{}}
			if len(stack) > 0 {
				p := stack[len(stack)-1]
				p.Children = append(p.Children, e)
			}
			stack = append(stack, e)
		case xml.EndElement:
			if len(stack) == 0 {
				return nil, fmt.Errorf("unexpected tag closing")
			} else if len(stack) == 1 {
				return stack[0], nil
			}
			stack = stack[:len(stack)-1]
		case xml.CharData:
			if len(stack) > 0 {
				p := stack[len(stack)-1]
				p.Children = append(p.Children, CharData(tok))
			}
		}
	}
}
