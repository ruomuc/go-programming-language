package main

import (
	"bytes"
	"fmt"
)

// 练习6.3：(*IntSet).UnionWith计算了两个集合的并集，使用 | 操作符
// 对每个字进行安慰 “或” 操作。实现交集、差集和对称差运算。
//（两个集合的对称差包含只在某个集合中存在的元素）
func main() {
	var x, y IntSet
	x.AddAll(1, 144, 9, 3)
	y.AddAll(1, 2, 9)
	x.UniqueWith(&y)
	fmt.Printf("交集：x elements: %s\n", x.String())

	x.AddAll(144, 3)
	x.DiffWith(&y)
	fmt.Printf("差集：x elements: %s\n", x.String())

	x.AddAll(1, 9)
	x.SymmetryDiffWith(&y)
	fmt.Printf("对称差集：x elements: %s\n", x.String())
}

type IntSet struct {
	words []uint64
}

// SymmetryDiffWith将会对 s 与 t 做对称差集（只存在于s中或只存在于t中），
// 并将结果存入s中
func (s *IntSet) SymmetryDiffWith(t *IntSet) {
	for i, tw := range t.words {
		if i > len(s.words) {
			continue
		} else {
			// 先做 按位与，再做 按位异或
			s.words[i] ^= tw
		}
	}
}

// DiffWith将会对 s 与 t 做差集（存在于s中且不存在于t中），并将结果存入s中
func (s *IntSet) DiffWith(t *IntSet) {
	for i, tw := range t.words {
		if i > len(s.words) {
			continue
		} else {
			// 先做 按位与，再做 按位异或
			s.words[i] = s.words[i] ^ (s.words[i] & tw)
		}
	}
}

// UniqueWith将会对 s 和 t 做交集，并将结果存入 s 中
func (s *IntSet) UniqueWith(t *IntSet) {
	ls, lt := len(s.words), len(t.words)

	if ls > lt {
		s.words = s.words[0:lt]
	}

	for i, tw := range t.words {
		if i > ls {
			continue
		} else {
			s.words[i] &= tw
		}
	}
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
