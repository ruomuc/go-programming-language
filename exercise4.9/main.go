package main

import (
	"bufio"
	"fmt"
	"os"
)

// 4.9 编写一个程序wordfreq来汇总输入文本文件中农每个单词出现的次数。
// 在第一次调用scan之前，需要使用input.Split(bufio.ScanWords)来将
// 文本行按照单词分割而不是行分割。
func main() {
	wordfreq()
}

func wordfreq() {
	countMap := make(map[string]int)

	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("read file err:", err)
		os.Exit(1)
	}
	input := bufio.NewScanner(file)
	input.Split(bufio.ScanWords)

	for input.Scan() {
		countMap[input.Text()]++
	}

	for key, value := range countMap {
		fmt.Printf("%s = %d\n", key, value)
	}
}
