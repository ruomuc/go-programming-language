package main

import (
	"fmt"
	"io"
	"io/ioutil"
)

// 练习7.2：实现一个满足如下签名的 CountingWriter 函数，输入一个 io.Writer，
// 输出一个 io.Writer，输出一个封装了输入值的新 Writer，以及一个指向 int64 的
// 指针，该指针对应的值是新的 Writer 写入的字节数。
func main() {
	w, c := CountingWriter(ioutil.Discard)
	fmt.Fprintf(w, "hello,workld!")
	fmt.Println(*c)
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	bc := &ByteCounter{w, 0}
	return bc, &bc.written
}

type ByteCounter struct {
	w       io.Writer
	written int64
}

func (c *ByteCounter) Write(p []byte) (int, error) {
	n, err := c.w.Write(p)
	c.written += int64(n)
	return n, err
}
