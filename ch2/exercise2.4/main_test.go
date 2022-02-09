package main

import (
	"testing"
)

/**
goos: windows
goarch: amd64
pkg: exercise2.4
BenchmarkPopCount-12            1000000000               0.282 ns/op
BenchmarkPopCount2-12           358208098                3.41 ns/op
PASS
ok      exercise2.4     2.130s
*/

var testNumber uint64 = 8

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(testNumber)
	}
}

// 比循环好点，但是没有查表快
func BenchmarkPopCount2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount2(testNumber)
	}
}
