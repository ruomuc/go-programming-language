package main

import (
	"bytes"
	"fmt"
)

// 练习6.5：Inset使用的每个字的类型都是uint64，但是64位的计算在32位
// 平台上的效率不高。改写程序以使用 uint 类型，这是适应平台的无符号整型。
// 除以64的操作可以用一个常量代表 32 位或者 64 位。你或许可以使用一个讨巧
// 的表达式 32<<(^uint(0)>>63)来表示除数

// 当 uint 为 64 位时，64 个 1111111... 右移 63 位变成 1，32 再左移 1 位变成 64
// 当 uint 位 32 位时，32 个 1111111... 右移 63 位变成 0, 32 再左移 0 位 变成 32
const uintSize = 32 << (^uint(0) >> 63)

func main() {
	var x IntSet
	x.Add(1)
	x.Add(11)
	fmt.Printf("x elements: %s\n", x.String())
}

type IntSet struct {
	words []uint
}

// Add添加非负数x到集合中
func (s *IntSet) Add(x int) {
	word, bit := x/uintSize, uint(x%uintSize)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// String方法将以字符串“{1,2,3}”的形式返回IntSet中的数
func (s IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < uintSize; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", uintSize*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}
