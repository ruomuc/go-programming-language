package main

import (
	"bufio"
	"fmt"
	"os"
)

// 1.4 修改 dup2 程序，输出出现重复行的文件的名称。
func main() {
	var dupFiles []string

	files := os.Args[1:]
	counts := make(map[string]int)

	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			continue
		}
		countLine(f, counts)
		f.Close()

		for _, n := range counts {
			if n > 1 {
				dupFiles = append(dupFiles, file)
				break
			}
		}
	}
	for _, fileName := range dupFiles {
		fmt.Printf("Duplicate file：%s\n", fileName)
	}
}

func countLine(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if input.Text() == "exit" {
			break
		}
		counts[input.Text()]++
	}
}
