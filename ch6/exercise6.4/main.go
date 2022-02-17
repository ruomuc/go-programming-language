package main

import (
	"bytes"
	"fmt"
)

// 练习6.4：添加方法 Elems 返回包含集合元素的 slice，
// 这适合在 range 循环中使用。
func main() {
	var x IntSet
	x.AddAll(1, 144, 9, 3)
	for _, e := range x.Elems() {
		fmt.Println(e)
	}
}

type IntSet struct {
	words []uint64
}

func (s *IntSet) Elems() []int {
	var res []int
	for i, word := range s.words {
		if word == 0 {
			continue
		}

		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				res = append(res, 64*i+j)
			}
		}
	}
	return res
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
