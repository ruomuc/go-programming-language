package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"hash"
)

// 练习4.2：编写一个程序，用于在默认情况下输出其标准输入的
// SHA256 散列，但也支持一个输出 SHA384 或 SHA512 散列
// 的命令行标记
func main() {
	var protocol string
	var content string

	flag.StringVar(&protocol, "p", "256", "protocol")
	flag.StringVar(&content, "c", "256", "content")

	fmt.Println(encrypt(protocol, content))
}

func encrypt(protocol string, content string) string {
	var h hash.Hash
	switch protocol {
	case "384":
		h = sha512.New384()
	case "512":
		h = sha512.New()
	default:
		h = sha256.New()
	}
	h.Write([]byte(content))
	fmt.Println(h.Sum(nil))
	return string(h.Sum(nil))
}
