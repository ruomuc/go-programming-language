package main

import "testing"

var args = []string{"a", "b", "c", "d", "e", "f", "g"}

/**
goos: windows
goarch: amd64
pkg: exercise1.3
BenchmarkNormal-12    	 4807720	       246 ns/op	      12 B/op	       6 allocs/op
PASS
ok  	exercise1.3	1.715s
*/
func BenchmarkNormal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		normal(args)
	}
}

/**
goos: windows
goarch: amd64
pkg: exercise1.3
BenchmarkStrJoin-12    	11150996	       101 ns/op	      16 B/op	       1 allocs/op
PASS
ok  	exercise1.3	1.510s
*/
func BenchmarkStrJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strJoin(args)
	}
}
