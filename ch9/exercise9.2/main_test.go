package main

import "testing"

var testNumber uint64 = 8

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(testNumber)
	}
}
