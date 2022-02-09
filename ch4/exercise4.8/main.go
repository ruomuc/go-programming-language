package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

// 4.8修改charcount的代码来统计字母、数字和其他unicode
// 分类中的字符数量，可以使用函数unicode.IsLetter等。
func main() {
	charCount()
}

func charCount() {
	countMap := make(map[string]int)

	in := bufio.NewReader(os.Stdin)
	for {
		r, _, err := in.ReadRune()
		fmt.Println(r, err)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcout: %v\n", err)
			os.Exit(1)
		}

		if unicode.IsLetter(r) {
			countMap["letter"]++
		} else if unicode.IsNumber(r) {
			countMap["number"]++
		} else if r != unicode.ReplacementChar {
			countMap["other"]++
		}
		fmt.Println(countMap)
	}
	for t, c := range countMap {
		switch t {
		case "letter":
			fmt.Println("letter count: ", c)
		case "number":
			fmt.Println("number count: ", c)
		case "other":
			fmt.Println("other count: ", c)
		default:
			fmt.Println("no match count type, err")
		}
	}
}
