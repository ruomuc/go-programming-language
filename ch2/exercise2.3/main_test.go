package main

import (
	"testing"
)

/**
goos: windows
goarch: amd64
pkg: exercise2.3
BenchmarkPopCount-12            1000000000               0.294 ns/op
BenchmarkPopCount2-12           128567740                9.50 ns/op
PASS
ok      exercise2.3     2.734s
*/

var testNumber uint64 = 8

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(testNumber)
	}
}

// 这个的性能变差了
func BenchmarkPopCount2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount2(testNumber)
	}
}
