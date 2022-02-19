package main

import (
	"bufio"
	"bytes"
	"fmt"
)

// 练习7.1：使用类似 ByteCounter的想法，实现单词和行
// 的计数器。实现时考虑使用 bufio.ScanWords。
func main() {
	s := "hello, world!\nhello, golang!"
	var wc WordCounter
	fmt.Fprintf(&wc, s)
	fmt.Println(wc)

	var lc LineCounter
	fmt.Fprintf(&lc, s)
	fmt.Println(lc)
}

type WordCounter int

func (wc *WordCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))

	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		*wc++
	}
	return len(p), nil
}

type LineCounter int

func (lc *LineCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))

	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		*lc++
	}
	return len(p), nil
}
