package main

import (
	"bytes"
	"fmt"
)

// 练习6.2：定义一个变长方法(*IntSet).AddAll(...int)，它允许接受一串
// 整型值作为参数，比如 s.AddAll(1,2,3)。
func main() {
	var x IntSet
	x.AddAll(1, 144, 9)
	fmt.Printf("x elements: %s\n", x.String())
}

type IntSet struct {
	words []uint64
}

// AddAll一次添加多个非负数到集合中
func (s *IntSet) AddAll(nums ...int) {
	for _, num := range nums {
		s.Add(num)
	}
}

// Add添加非负数x到集合中
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
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
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}
