package main

import (
	"testing"
)

/**
goos: windows
goarch: amd64
pkg: exercise2.5
BenchmarkPopCount-12            1000000000               0.556 ns/op
BenchmarkPopCount2-12           546457543                2.24 ns/op
PASS
ok      exercise2.5     2.312s
*/

var testNumber uint64 = 8

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(testNumber)
	}
}

// 比循环和移位都好，但是还是没查表快。
func BenchmarkPopCount2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount2(testNumber)
	}
}
