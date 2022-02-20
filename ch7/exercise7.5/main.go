package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strings"
)

// 练习7.5：io 包中的 LimitReader 函数接受 io.Reader r 和字节数 n，
// 返回一个 Reader，该返回值从 r 读取数据，但在读取 n 字节后报告文件结束。
// 请实现该函数。
// func LimitReader(r io.Reader, n int64) io.Reader

func main() {
	lr := LimitReader(strings.NewReader("qwert"), 1)
	b, err := ioutil.ReadAll(lr)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%s", b)
}

type limitReader struct {
	r          io.Reader
	limitBytes int64
}

func (lr *limitReader) Read(p []byte) (n int, err error) {
	if lr.limitBytes <= 0 {
		return 0, io.EOF
	}
	if int64(len(p)) > lr.limitBytes {
		p = p[:lr.limitBytes]
	}
	n, err = lr.r.Read(p)
	lr.limitBytes -= int64(n)
	return
}

func LimitReader(r io.Reader, n int64) io.Reader {
	lr := &limitReader{
		r:          r,
		limitBytes: n,
	}
	return lr
}
