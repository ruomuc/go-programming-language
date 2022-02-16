package main

import (
	"bytes"
	"fmt"
)

// 练习6.1：实现这些附加方法：
// func(*IntSet) Len()int //返回元素个数
// func(*IntSet) Remove(x int) // 从集合中去除元素x
// func(*IntSet) Clear() // 删除所有元素
// func(*IntSet) Copy() *IntSet // 返回集合的副本
func main() {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	//
	//y.Add(9)
	//y.Add(42)
	//
	//x.UnionWith(&y)
	fmt.Printf("x len: %d, \nx elements: %s\n", x.Len(), x.String())
	x.Remove(9)
	fmt.Printf("x len: %d, \nx elements: %s\n", x.Len(), x.String())
	x.Clear()
	fmt.Printf("x len: %d, \nx elements: %s\n", x.Len(), x.String())
	x.Add(111)
	xCopy := x.Copy()
	x.Add(222)
	fmt.Printf("x len: %d, \nx elements: %s\n", x.Len(), x.String())
	fmt.Printf("xCopy len: %d, \nxCopy elements: %s\n", xCopy.Len(), xCopy.String())
}

type IntSet struct {
	words []uint64
}

// Len将返回元素个数
func (s *IntSet) Len() int {
	var c int
	for _, word := range s.words {
		if word == 0 {
			continue
		}

		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				c++
			}
		}
	}
	return c
}

// Remove从集合中去除元素x
func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	s.words[word] ^= 1 << bit
}

// Clear将删除集合中所有元素
func (s *IntSet) Clear() {
	s.words = s.words[0:0]
}

// Copy将返回集合的副本
func (s *IntSet) Copy() *IntSet {
	cIntSet := &IntSet{words: []uint64{}}
	for _, w := range s.words {
		cIntSet.words = append(cIntSet.words, w)
	}
	return cIntSet
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add添加非负数x到集合中
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith将会对s和t做并集并将结果存在s中
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tw := range t.words {
		if i < len(s.words) {
			s.words[i] |= tw
		} else {
			s.words = append(s.words, tw)
		}
	}
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
