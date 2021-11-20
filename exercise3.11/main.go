package main

import (
	"bytes"
	"fmt"
	"strings"
)

// 练习3.11：增强 comma 函数的功能，让其
// 正确处理浮点数，以及带有可选正负号的数字。
func main() {
	s1 := "1234567890"
	s2 := "-1234567890"
	s3 := "+1234567890"
	s4 := "-1234567890.1234"
	fmt.Println(comma(s1))
	fmt.Println(comma(s2))
	fmt.Println(comma(s3))
	fmt.Println(comma(s4))
}

func comma(s string) string {
	var (
		intStr     string
		floatStr   string
		signalStr  string
		tempStrArr []string
	)

	// 符号位
	if strings.Contains(s, "+") || strings.Contains(s, "-") {
		signalStr = s[:1]
		s = s[1:]
	}

	// 整数位 和 小数位
	intStr = s
	if idx := strings.Index(s, "."); idx != -1 {
		intStr = s[:idx]
		floatStr = s[idx+1:]
	}

	resStrBuf := bytes.NewBufferString(signalStr)
	intStrBuf := bytes.NewBufferString(intStr)

	for intStrBuf.Len() > 0 {
		b := intStrBuf.Next(3)
		tempStrArr = append(tempStrArr, string(b))
	}
	resStrBuf.WriteString(strings.Join(tempStrArr, ","))
	tempStrArr = tempStrArr[:0]

	floatStrBuf := bytes.NewBufferString(floatStr)
	for floatStrBuf.Len() > 0 {
		b := floatStrBuf.Next(3)
		tempStrArr = append(tempStrArr, string(b))
	}
	if len(tempStrArr) > 0 {
		resStrBuf.WriteByte('.')
	}
	resStrBuf.WriteString(strings.Join(tempStrArr, ","))

	return resStrBuf.String()
}
